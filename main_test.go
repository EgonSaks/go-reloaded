package main

import (
	"io/ioutil"
	"os/exec"
	"testing"
)

// This test file tests the go-reloaded project against the test cases on audit page
func TestGoReloaded(t *testing.T) {
	inputFile, outputFile := "sample.txt", "result.txt"

	//	Each element of testCases contains a pair of strings, the first string of the
	//	pair is what is to be written to the input file, the second is what should be
	//	the contents of the output file
	testCases := [][]string{
		{
			"1E (hex) files were added",
			"30 files were added",
		},
		{
			"It has been 10 (bin) years",
			"It has been 2 years",
		},
		{
			"Ready, set, go (up) !",
			"Ready, set, GO!",
		},
		{
			"I should stop SHOUTING (low)",
			"I should stop shouting",
		},
		{
			"Welcome to the Brooklyn bridge (cap)",
			"Welcome to the Brooklyn Bridge",
		},
		{
			"This is so exciting (up, 2)",
			"This is SO EXCITING",
		},
		{
			"I was sitting over there ,and then BAMM !!",
			"I was sitting over there, and then BAMM!!",
		},
		{
			"I was thinking ... You were right",
			"I was thinking... You were right",
		},
		{
			"I am exactly how they describe me: ' awesome '",
			"I am exactly how they describe me: 'awesome'",
		},
		{
			"As Elton John said: ' I am the most well-known homosexual in the world '",
			"As Elton John said: 'I am the most well-known homosexual in the world'",
		},
		{
			"There it was. A amazing rock!",
			"There it was. An amazing rock!",
		},

		{
			"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			"If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},

		{
			"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			"I have to pack 5 outfits. Packed 26 just to be sure",
		},

		{
			"Don't be sad ,because sad backwards is das . And das not good",
			"Don't be sad, because sad backwards is das. And das not good",
		},

		{
			"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			"Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},
	}

	for _, testCase := range testCases {
		// Write the string to be processed to the input file
		if err := ioutil.WriteFile(inputFile, []byte(testCase[0]), 0o664); err != nil {
			panic(err)
		}

		// Attempt to run the main project file with the input and output file
		// names as arguments
		if err := exec.Command("go", "run", ".", inputFile, outputFile).Run(); err != nil {
			t.Fatal(err)
		}

		// Read the contents of the output file, checking if it is equal to the
		// expected output
		if result, err := ioutil.ReadFile(outputFile); err != nil {
			panic(err)
		} else if string(result) != testCase[1] {
			t.Errorf("\nTest fails when given the test case:\n\t\"%s\","+
				"\n%s should contain:\n\t\"%s\",\n%s contains:\n\t\"%s\"\n\n",
				testCase[0], outputFile, testCase[1], outputFile, string(result))
		}
	}
}

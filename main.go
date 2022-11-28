package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readFile() string {
	if len(os.Args) == 1 {
		fmt.Println("Make sure you have the input file added to arguments ex: 'go run . <input>.txt <output>.txt'")
		os.Exit(1)
	} else if len(os.Args) == 2 {
		fmt.Println("Make sure you have the output file added to arguments ex: 'go run . <input>.txt <output>.txt'")
		os.Exit(2)
	} else if len(os.Args) > 3 {
		fmt.Println("Too many arguments, should be 'go run . <input>.txt <output>.txt")
		os.Exit(4)
	} else {

		inputFile := os.Args[1]
		file, err := ioutil.ReadFile(inputFile)
		if err != nil {
			fmt.Printf("Could not read the file due to this %s error \n", err)
		}
		return string(file)

	}
	return "run: go run . <input>.txt <output>.txt"
}

// Convert the word(s) before to (up), (low), (cap), (hex),
// (bin) and in case of `(low)`, `(up)`, `(cap)` -> `(low, <number>)`
func converter(s string) string {
	listOfWords := strings.Fields(s)

	for i, word := range listOfWords {
		switch word {

		case "(up)":
			listOfWords[i-1] = strings.ToUpper(listOfWords[i-1])
			listOfWords = append(listOfWords[:i], listOfWords[i+1:]...)
		case "(low)":
			listOfWords[i-1] = strings.ToLower(listOfWords[i-1])
			listOfWords = append(listOfWords[:i], listOfWords[i+1:]...)
		case "(cap)":
			listOfWords[i-1] = strings.Title(listOfWords[i-1])
			listOfWords = append(listOfWords[:i], listOfWords[i+1:]...)
		case "(hex)":
			listOfWords[i-1] = hexConverter(listOfWords[i-1])
			listOfWords = append(listOfWords[:i], listOfWords[i+1:]...)
		case "(bin)":
			listOfWords[i-1] = binConverter(listOfWords[i-1])
			listOfWords = append(listOfWords[:i], listOfWords[i+1:]...)

		case "(up,":
			number := strings.Trim(string(listOfWords[i+1]), listOfWords[i+1][1:])
			count, _ := strconv.Atoi(number)
			for n := 1; n <= count; n++ {
				listOfWords[i-n] = strings.ToUpper(listOfWords[i-n])
			}
			listOfWords = append(listOfWords[:i], listOfWords[i+2:]...)
		case "(low,":
			number := strings.Trim(string(listOfWords[i+1]), listOfWords[i+1][1:])
			count, _ := strconv.Atoi(number)
			for n := 1; n <= count; n++ {
				listOfWords[i-n] = strings.ToLower(listOfWords[i-n])
			}
			listOfWords = append(listOfWords[:i], listOfWords[i+2:]...)
		case "(cap,":
			number := strings.Trim(string(listOfWords[i+1]), listOfWords[i+1][1:])
			count, _ := strconv.Atoi(number)
			for n := 1; n <= count; n++ {
				listOfWords[i-n] = strings.Title(listOfWords[i-n])
			}
			listOfWords = append(listOfWords[:i], listOfWords[i+2:]...)
		}
	}
	return strings.Join(listOfWords, " ")
}

// Replace the word before with the decimal version of the word
func hexConverter(s string) string {
	decimal_num, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(int(decimal_num))
}

// Replace the word before with the decimal version of the word
func binConverter(s string) string {
	binary, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(int(binary))
}

// Put every instance of the punctuations `.`, `,`, `!`, `?`, `:` and `;`
// close to the previous word and with space apart from the next one.
func punctuationMarks(s string) string {
	listOfWords := strings.Fields(s)
	punctuationMarks := []string{",", ".", "!", "?", ":", ";"}

	// Check the beginning of the string
	for i, word := range listOfWords {
		for _, punctuation := range punctuationMarks {
			// If on the slices first index is punctuation and it's at the beginning of the, means at the start of the sentence
			if (string(word[0]) == punctuation) && (listOfWords[0] == listOfWords[i]) {
				listOfWords[i] = listOfWords[1]
			}
		}
	}

	// check punctuation on the middle of the string
	for i, word := range listOfWords {
		for _, punctuation := range punctuationMarks {
			if (string(word[0]) == punctuation) && (string(word[len(word)-1]) != punctuation) {
				listOfWords[i-1] = listOfWords[i-1] + punctuation
				// Start the new word from index 1, so you leave the punctuation from position 0
				listOfWords[i] = word[1:]
			}
		}
	}

	// Check punctuation at the end of the slice
	for i, word := range listOfWords {
		for _, punctuation := range punctuationMarks {
			if (string(word[0]) == punctuation) && (listOfWords[len(listOfWords)-1] == listOfWords[i]) {
				listOfWords[i-1] = listOfWords[i-1] + word
				listOfWords = listOfWords[:len(listOfWords)-1]
			}
		}
	}

	// Check punctuations inside the slice and align them next to the word when needed
	for i, word := range listOfWords {
		for _, punctuation := range punctuationMarks {
			if (string(word[0]) == punctuation) && (string(word[len(word)-1]) == punctuation) && (listOfWords[i] != listOfWords[len(listOfWords)-1]) {
				listOfWords[i-1] = listOfWords[i-1] + word
				// remove empty index
				listOfWords = append(listOfWords[:i], listOfWords[i+1:]...)
			}
		}
	}
	return strings.Join(listOfWords, " ")
}

// Find pairs of ' and place them to the right and left of the word in the middle of them, without any spaces.
func apostrophe(s string) string {
	str := ""

	var flag bool
	for i, word := range s {
		if (word == 39 || word == 96) && s[i-1] == ' ' {
			if flag {
				str = str[:len(str)-1]
				str = str + string(word)
				flag = false
			} else {
				str = str + string(word)
				flag = true
			}
		} else if i > 1 && (s[i-2] == 39 || s[i-2] == 96) && s[i-1] == ' ' {
			if flag {
				str = str[:len(str)-1]
				str = str + string(word)
			} else {
				str = str + string(word)
			}
		} else {
			str = str + string(word)
		}
	}
	return str
}

// Turn every instance of `a` into `an` if the next word begins with a vowel or a `h`.
func articles(s string) string {
	slice := strings.Split(s, " ")

	for i := 0; i < len(slice); i++ {
		if slice[i] == slice[len(slice)-1] {
			break
		}
		if (slice[i] == "a" || slice[i] == "A") &&
			(slice[i+1][0] == 'a' || slice[i+1][0] == 'e' || slice[i+1][0] == 'y' || slice[i+1][0] == 'u' || slice[i+1][0] == 'i' || slice[i+1][0] == 'o' || slice[i+1][0] == 'h') {
			slice[i] = slice[i] + "n"
		}
	}
	return strings.Join(slice, " ")
}

// write an output to new file based on given name from command line argument
func writeNewFile(s string) {
	outputFile := os.Args[2]
	output := []byte(s)
	os.WriteFile(outputFile, output, 0o664)
}

func main() {
	read := readFile()
	punctuations := punctuationMarks(converter(read))
	apostrophe := apostrophe(punctuations)
	articles := articles(apostrophe)
	writeNewFile(articles)
	fmt.Println(articles)
}

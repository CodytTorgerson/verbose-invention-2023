package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data, err := os.ReadFile("calb2.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fileContent := string(data)
	lines := strings.Split(fileContent, "\n")
	var score = 0
	for _, line := range lines {
		var combinedArray = append(digitFinder(line), numbersAsWordsFinder(line)...)
		if len(combinedArray) == 0 {
			continue
		}
		sort.Slice(combinedArray, func(i, j int) bool {
			return combinedArray[i].index < combinedArray[j].index
		})

		var firstDigit = combinedArray[0].number
		var secondDigit = combinedArray[len(combinedArray)-1].number
		var finalNumberAsString = firstDigit + secondDigit
		var finalNumber, _ = strconv.Atoi(finalNumberAsString)
		score += finalNumber
	}
	fmt.Println("total score: ", score)
}

func digitFinder(line string) []extractedNumberAsString {
	var listOfNumbers []extractedNumberAsString
	for index, char := range line {
		if unicode.IsDigit(char) {
			listOfNumbers = append(listOfNumbers, extractedNumberAsString{index: index, number: string(char)})
		}
	}
	return listOfNumbers
}

func numbersAsWordsFinder(line string) []extractedNumberAsString {
	var listOfNumbers []extractedNumberAsString
	var numbersAsWords = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9}
	for key, value := range numbersAsWords {
		indexOfString := strings.Index(line, key)
		if indexOfString != -1 {
			listOfNumbers = append(listOfNumbers, extractedNumberAsString{index: indexOfString, number: strconv.Itoa(value)})
		}

	}
	// this will find relevant duplicates
	for key, value := range numbersAsWords {
		indexOfString := strings.LastIndex(line, key)
		if indexOfString != -1 {
			listOfNumbers = append(listOfNumbers, extractedNumberAsString{index: indexOfString, number: strconv.Itoa(value)})
		}
	}

	return listOfNumbers
}

type extractedNumberAsString struct {
	index  int
	number string
}

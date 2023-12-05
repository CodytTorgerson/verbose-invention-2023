package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data, err := os.ReadFile("calibration.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fileContent := string(data)
	lines := strings.Split(fileContent, "\n")
	var score = 0
	for _, line := range lines {
		var combinedDigits = []string{digitCheckerForward(line), digitCheckerBackwards(line)}
		var combinedDigitsString = strings.Join(combinedDigits, "")
		var score2Add, err = strconv.Atoi(combinedDigitsString)
		if err != nil {
			fmt.Println("Error converting string to integer: ", err)
			continue
		}
		score += int(score2Add)
	}
	fmt.Println("total score: ", score)
}

func digitCheckerForward(line string) string {
	for _, char := range line {
		if unicode.IsDigit(char) {
			return string(char)
		}
	}
	return ""
}

func digitCheckerBackwards(line string) string {
	var reversedString = reverseString(line)

	for _, char := range reversedString {
		if unicode.IsDigit(char) {
			return string(char)
		}
	}
	return ""
}

func reverseString(line string) string {
	rs := []rune(line)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func numbersAsWordsChecker(line string) (int, bool) {
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
		if strings.Contains(line, key) {
			return value, true
		}
	}
	return -1, false
}

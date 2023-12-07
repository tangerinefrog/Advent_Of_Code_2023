package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func DayOnePartOne() {
	var filename = "input_day_one"
	lines := GetFileLines(filename)

	sum := 0
	for i := 0; i < len(lines); i++ {
		runes := []rune(lines[i])
		var numberString []rune
		for j := 0; j < len(runes); j++ {
			if unicode.IsDigit(runes[j]) {
				numberString = append(numberString, runes[j])
			}
		}
		if len(numberString) > 0 {
			var resultStr string

			if len(numberString) == 1 {
				resultStr = string([]rune{numberString[0], numberString[0]})
			} else {
				firstAndLast := []rune{numberString[0], numberString[len(numberString)-1]}
				resultStr = string(firstAndLast)
			}

			resultNum, _ := strconv.Atoi(resultStr)

			sum += resultNum
		}
	}

	fmt.Printf("Day 1 - Part 1 result: %d\n", sum)
}

func DayOnePartTwo() {
	var filename = "input_day_one"
	lines := GetFileLines(filename)

	sum := 0

	for lineIndex := 0; lineIndex < len(lines); lineIndex++ {
		firstNumber := FindFirstDigit(lines[lineIndex])
		lastNumber := FindLastDigit(lines[lineIndex])

		sum += firstNumber*10 + lastNumber
	}

	fmt.Printf("Day 1 - Part 2 result: %d\n", sum)
}

func FindFirstDigit(line string) int {

	strlen := len(line)
	strRunes := []rune(line)

	for i := 0; i < strlen; i++ {
		if unicode.IsDigit(strRunes[i]) {
			resultNum, _ := strconv.Atoi(string(strRunes[i]))
			return resultNum
		}

		for j := 0; j < len(digits); j++ {

			if strings.HasPrefix(line, digits[j]) {
				return j + 1
			}
		}
		line = line[1:]
	}

	return 0
}

func FindLastDigit(line string) int {

	strlen := len(line)
	strRunes := []rune(line)
	result := 0

	for i := 0; i < strlen; i++ {

		if unicode.IsDigit(strRunes[i]) {
			result, _ = strconv.Atoi(string(strRunes[i]))
		}

		for j := 0; j < len(digits); j++ {

			if strings.HasPrefix(line, digits[j]) {
				result = j + 1
			}
		}
		if len(line) > 0 {
			line = line[1:]
		}
	}

	return result
}

func GetFileLines(filename string) []string {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	content := strings.TrimRight(string(fileBytes), "\n")
	return strings.Split(content, "\n")
}

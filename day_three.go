package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type LineNumber struct {
	Num        int
	StartIndex int
	Length     int
}

type Position struct {
	row int
	col int
}

func DayThreePartOne() {
	filename := "input_day_three"
	lines := GetFileLines(filename)

	lineLen := len(lines[0])
	lineCount := len(lines)
	sum := 0

	for i, line := range lines {
		nums := GetLineNumbers(line)
		for _, num := range nums {
			if HasAdjacentSymbols(num, &lines, i, lineLen, lineCount) {
				sum += num.Num
			}
		}
	}

	fmt.Printf("Day 3 - Part 1 result: %d\n", sum)
}

func DayThreePartTwo() {
	filename := "input_day_three"
	lines := GetFileLines(filename)

	lineCount := len(lines)

	sum := 0

	gearPositions := GetGears(lines)
	lineNums := make(map[int][]LineNumber)
	for i, line := range lines {
		lineNums[i] = GetLineNumbers(line)
	}

	for _, pos := range gearPositions {
		ratio := GetGearRatio(pos, lineNums, lineCount)
		if ratio > 0{
			sum += ratio
		}
	}

	fmt.Printf("Day 3 - Part 2 result: %d\n", sum)
}

func GetLineNumbers(line string) (nums []LineNumber) {
	nums = make([]LineNumber, 0)

	var buf = make([]rune, 0)
	for i := 0; i < len(line); i++ {
		c := rune(line[i])

		if unicode.IsDigit(c) {
			buf = append(buf, c)
			if i == len(line) - 1 {
				num, _ := strconv.Atoi(string(buf))
				nums = append(nums, LineNumber{Num: num, StartIndex: i - len(buf), Length: len(buf)})
			}
		} else if len(buf) > 0 {
			num, _ := strconv.Atoi(string(buf))
			nums = append(nums, LineNumber{Num: num, StartIndex: i - len(buf), Length: len(buf)})
			buf = make([]rune, 0)
		}
	}

	return
}

func GetGears(lines []string) (positions []Position) {
	for i, line := range lines {
		for j, c := range line {
			if c == '*' {
				positions = append(positions, Position{row: i, col: j})
			}
		}
	}

	return
}

func GetGearRatio(gearPos Position, lineNums map[int][]LineNumber, lineCount int) int {
	gearNums := make([]int, 0)
	
	for row := gearPos.row - 1; row <= gearPos.row + 1; row++ {
		if row >= 0 && row < lineCount {
			nums := lineNums[row]
			for _, num := range nums {
				if(num.StartIndex <= gearPos.col + 1 && num.StartIndex + num.Length - 1 >= gearPos.col - 1){
					gearNums = append(gearNums, num.Num)
				}
			}
		}
	}
	if len(gearNums) == 2 {
		return gearNums[0] * gearNums[1]
	}

	return -1
}

func HasAdjacentSymbols(num LineNumber, lines *[]string, curRow int, lineLen int, lineCount int) bool {
	leftMargin := num.StartIndex - 1
	rightMargin := num.StartIndex + num.Length
	if leftMargin < 0 {
		leftMargin = 0
	}
	if rightMargin > lineLen {
		rightMargin = lineLen - 1
	}

	for runeInd := leftMargin; runeInd <= rightMargin; runeInd++ {
		for row := curRow - 1; row <= curRow + 1; row++ {
			if row >= 0 && row < lineCount {
				if IsSymbol(rune((*lines)[row][runeInd])) {
					return true
				}
			}
		}
	}

	return false
}

func IsSymbol(c rune) bool {
	return c != '.' && !unicode.IsDigit(c)
}

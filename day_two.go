package main

import (
	"fmt"
	"strconv"
	"strings"
)

var restrictions = map[string]int{"red": 12, "green": 13, "blue": 14}

func DayTwoPartOne() {
	var filename = "input_day_two"

	sum := 0
	lines := GetFileLines(filename)
	for i := 0; i < len(lines); i++ {
		gameSplit := strings.Split(lines[i], ": ")
		cubeGroups := strings.Split(gameSplit[1], "; ")
		isPossible := true

		for j := 0; j < len(cubeGroups); j++ {
			if IsGroupImpossible(cubeGroups[j]) {
				isPossible = false
				continue
			}
		}

		if isPossible {
			gameIdString := strings.Split(gameSplit[0], " ")[1]
			gameId, _ := strconv.Atoi(gameIdString)
			sum += gameId
		}
	}

	fmt.Printf("Day 2 - Part 1 result: %d\n", sum)
}

func IsGroupImpossible(group string) bool {
	draws := strings.Split(group, ", ")
	for i := 0; i < len(draws); i++ {
		count_color := strings.Split(draws[i], " ")
		count, _ := strconv.Atoi(count_color[0])
		color := count_color[1]

		if restrictions[color] < count {
			return true
		}
	}
	return false
}
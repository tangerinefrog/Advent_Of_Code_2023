package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Card struct {
	Points int
	Count  int
}

func DayFourPartOne() {
	lines := GetFileLines("input_day_4")

	sum := 0
	for _, line := range lines {
		wn, gn := ParseGameNumbers(line)
		sum += GetGamePoints(wn, gn)
	}

	fmt.Printf("Day 4 - Part 1 result: %d\n", sum)
}

func DayFourPartTwo() {
	lines := GetFileLines("input_day_4")

	sum := 0
	cards := make([]Card, 0)
	for _, line := range lines {
		card := Card{}
		wn, gn := ParseGameNumbers(line)
		card.Points = GetGameNextCardsCount(wn, gn)
		card.Count = 1
		cards = append(cards, card)
	}

	CountCardCopies(&cards)

	for _, card := range cards {
		sum += card.Count
	}

	fmt.Printf("Day 4 - Part 2 result: %d\n", sum)
}

func ParseGameNumbers(line string) (winNums []int, gameNums []int) {
	gameSection := strings.Split(line, ": ")[1]
	gameRows := strings.Split(gameSection, " | ")

	winNums = GetNumbers(gameRows[0])
	gameNums = GetNumbers(gameRows[1])

	return
}

func GetGamePoints(winNums []int, gameNums []int) (points int) {
	points = 0
	for _, winNum := range winNums {
		for _, gameNum := range gameNums {
			if winNum == gameNum {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
	}

	return
}

func GetGameNextCardsCount(winNums []int, gameNums []int) (points int) {
	points = 0
	for _, winNum := range winNums {
		for _, gameNum := range gameNums {
			if winNum == gameNum {
				points++
			}
		}
	}

	return
}

func CountCardCopies(cards *[]Card) {
	for cardInd, card := range *cards {
		for i := 0; i < card.Count; i++ {
			for j := 0; j < card.Points; j++ {
				if cardInd+j+1 < len(*cards) {
					(*cards)[cardInd+j+1].Count++
				}
			}
		}
	}
}

func GetNumbers(line string) (result []int) {
	buf := make([]rune, 0)

	for i, c := range line {
		if unicode.IsDigit(c) {
			buf = append(buf, c)
			if i == len(line)-1 {
				num, _ := strconv.Atoi(string(buf))
				result = append(result, num)
			}
		} else if len(buf) > 0 {
			num, _ := strconv.Atoi(string(buf))
			result = append(result, num)
			buf = make([]rune, 0)
		}
	}

	return
}

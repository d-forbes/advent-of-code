package day04

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Day04() {
	file, err := os.ReadFile("day04/day04Input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")

	cards := new(Storage)
	start := time.Now()
	fmt.Println("Day 04 Part A Result:", partA(lines, cards), "in", time.Since(start))

	start1 := time.Now()
	fmt.Println("Day 04 Part B Result:", partB(cards, len(lines)), "in", time.Since(start1))

}

type Card struct {
	number int
	won    int
	copies int
}

type Storage struct {
	cards []Card
}

func (storage *Storage) addCard(card Card) []Card {
	storage.cards = append(storage.cards, card)

	return storage.cards
}
func partA(lines []string, storage *Storage) int {
	re := regexp.MustCompile(`Card\s+(\d+):\s+(.*)\s+\|\s+(.*)\n?`)
	total := 0

	for _, line := range lines {
		found := re.FindAllStringSubmatch(line, -1)
		if found == nil {
			continue
		}

		var cardNum int
		if v, err := strconv.Atoi(found[0][1]); err == nil {
			cardNum = v
		}

		card := Card{number: cardNum, won: 0, copies: 1}
		numbers := strings.TrimSpace(found[0][3])
		var winning []string

		for _, v := range strings.Split(found[0][2], " ") {
			if v := strings.TrimSpace(v); v != "" {
				winning = append(winning, v)
			}
		}

		t := 0
		for _, winner := range winning {
			for _, num := range strings.Split(numbers, " ") {
				if num == winner {
					t++
				}
			}
		}

		if t > 0 {
			total += int(math.Pow(float64(2), float64(t-1)))
			card.won = t
		}

		storage.addCard(card)
	}

	return total
}

func partB(wonCards *Storage, lines int) int {
	total := 0
	for _, card := range wonCards.cards {
		for j := card.number; j < card.number+card.won+1 && j-1 < lines; j++ {
			wonCards.cards[j-1].copies += card.copies
		}
		total += card.copies
	}

	return total
}

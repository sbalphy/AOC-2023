package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func evaluateHand(hand string) int {
	counts := make([]int, 5)
	var jokers int
	for i := 0; i < len(hand); i++ {
		if hand[i] == 'J' {
			jokers++
			continue
		}
		for j := i; j < len(hand); j++ {
			if hand[j] == hand[i] {
				counts[i]++
			}
		}
	}
	slices.Sort(counts)
	high_set := counts[len(counts)-1] + jokers
	if high_set > 3 {
		return high_set + 1
	}
	if high_set == 3 {
		// hehe
		if counts[len(counts)-3] == 2 || (counts[len(counts)-2] == 2 && jokers == 1) {
			return 4
		}
		return 3
	}
	if high_set == 2 {
		if counts[len(counts)-2] == 2 {
			return 2
		}
		return 1
	}
	return 0
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	hand_to_bid := make(map[string]int)
	var hands []string
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		hand := input[0]
		bid, _ := strconv.Atoi(input[1])
		hand_to_bid[hand] = bid
		hands = append(hands, hand)
	}
	sort.Slice(hands, func(i, j int) bool {
		hand1type := evaluateHand(hands[i])
		hand2type := evaluateHand(hands[j])
		if hand1type != hand2type {
			return hand1type < hand2type
		}
		order := map[byte]int{
			'J': 0,
			'2': 1,
			'3': 2,
			'4': 3,
			'5': 4,
			'6': 5,
			'7': 6,
			'8': 7,
			'9': 8,
			'T': 9,
			'Q': 10,
			'K': 11,
			'A': 12,
		}
		for k := 0; k < len(hands[i]); k++ {
			if order[hands[i][k]] != order[hands[j][k]] {
				return order[hands[i][k]] < order[hands[j][k]]
			}
		}
		return false
	})
	var sum int
	for rank := 1; rank <= len(hands); rank++ {
		sum += rank * hand_to_bid[hands[rank-1]]
		fmt.Printf("%dth hand %v bid %d, product %d, sum %d \n", rank, hands[rank-1], hand_to_bid[hands[rank-1]], rank*hand_to_bid[hands[rank-1]], sum)
	}
	fmt.Printf("final sum: %d \n", sum)
}

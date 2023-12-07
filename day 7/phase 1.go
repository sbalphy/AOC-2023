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
	for i := 0; i < len(hand); i++ {
		for j := i; j < len(hand); j++ {
			if hand[j] == hand[i] {
				counts[i]++
			}
		}
	}
	slices.Sort(counts)
	if counts[len(counts)-1] > 3 {
		return counts[len(counts)-1] + 1
	}
	if counts[len(counts)-1] == 3 {
		// hehe
		if counts[len(counts)-3] == 2 {
			return 4
		}
		return 3
	}
	if counts[len(counts)-1] == 2 {
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
			'2': 0,
			'3': 1,
			'4': 2,
			'5': 3,
			'6': 4,
			'7': 5,
			'8': 6,
			'9': 7,
			'T': 8,
			'J': 9,
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

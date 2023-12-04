package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	sum := 205
	r_cleaner, _ := regexp.Compile("Card +[0-9]+:((?: |[0-9])+)\\|((?: |[0-9])+)")
	var line int
	var cardcount []int
	// :/
	for i := 0; i < 205; i++ {
		cardcount = append(cardcount, 1)
	}

	for scanner.Scan() {
		input := scanner.Text()
		cleaned_input := r_cleaner.FindStringSubmatch(input)
		winning_numbers := strings.Split(cleaned_input[1], " ")
		owned_numbers := strings.Split(cleaned_input[2], " ")
		var copies int
		fmt.Printf("Winners: ")
		winning_set := make(map[int]bool)
		for i := 0; i < len(winning_numbers); i++ {
			value, _ := strconv.Atoi(winning_numbers[i])
			winning_set[value] = true
		}
		winning_set[0] = false
		for i := 0; i < len(owned_numbers); i++ {
			value, _ := strconv.Atoi(owned_numbers[i])
			if winning_set[value] {
				fmt.Printf("%d ", value)
				copies++
			}
		}
		for i := 0; i < cardcount[line]; i++ {
			for j := 1; j <= copies; j++ {
				cardcount[line+j]++
			}
		}
		sum += copies
		fmt.Printf("Line %d has %d winning numbers. Current sum: %d \n", line, copies, sum)
		line++
	}
	var altsum int
	for i := 0; i < 205; i++ {
		altsum += cardcount[i]
	}
	fmt.Printf("final sum %d, altsum %d \n", sum, altsum)
}

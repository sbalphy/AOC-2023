package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	var sum int
	r_cleaner, _ := regexp.Compile("Card +[0-9]+:((?: |[0-9])+)\\|((?: |[0-9])+)")
	var line int
	for scanner.Scan() {
		input := scanner.Text()
		cleaned_input := r_cleaner.FindStringSubmatch(input)
		winning_numbers := strings.Split(cleaned_input[1], " ")
		owned_numbers := strings.Split(cleaned_input[2], " ")
		var power int
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
				power++
			}
		}
		if power > 0 {
			sum += int(math.Pow(float64(2), float64(power)-1))
		}
		fmt.Printf("Line %d has %d winning numbers. Current sum: %d \n", line, power, sum)
		line++
	}
	fmt.Printf("final sum %d \n", sum)
}

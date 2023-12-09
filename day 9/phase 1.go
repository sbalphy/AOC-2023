package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isZero(sequence []int) bool {
	for _, value := range sequence {
		if value != 0 {
			return false
		}
	}
	return true
}
func extrapolate(sequence []int) int {
	if isZero(sequence) {
		return 0
	}
	var change []int
	for i := 1; i < len(sequence); i++ {
		change = append(change, sequence[i]-sequence[i-1])
	}
	return sequence[len(sequence)-1] + extrapolate(change)
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	var sequences [][]int
	for scanner.Scan() {
		sequence_text := strings.Split(scanner.Text(), " ")
		sequence := make([]int, len(sequence_text))
		for i, value := range sequence_text {
			sequence[i], _ = strconv.Atoi(value)
		}
		sequences = append(sequences, sequence)
	}
	var sum int
	for i, sequence := range sequences {
		extrapolated_value := extrapolate(sequence)
		sum += extrapolated_value
		fmt.Printf("Sequence %d (%d ... %d) extrapolated to %d. Current sum: %d \n", i, sequence[0], sequence[len(sequence)-1], extrapolated_value, sum)
	}
	fmt.Printf("Final sum: %d \n", sum)
}

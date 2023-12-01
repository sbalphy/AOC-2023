package main

import (
	"fmt"
	"regexp"
)

func main() {
	// tem jeitos piores de mexer com strings ascii :)
	LUT := map[string]int{
		"zero":  0,
		"0":     0,
		"one":   1,
		"1":     1,
		"two":   2,
		"2":     2,
		"three": 3,
		"3":     3,
		"four":  4,
		"4":     4,
		"five":  5,
		"5":     5,
		"six":   6,
		"6":     6,
		"seven": 7,
		"7":     7,
		"eight": 8,
		"8":     8,
		"nine":  9,
		"9":     9,
	}

	// expressoes regulares, a primeira acha a primeira occorrencia e a segunda acha a ultima
	// usa comportamento non-greedy pra decidir qual eh a primeira e qual eh a ultima
	first_r, _ := regexp.Compile(".*?(zero|0|one|1|two|2|three|3|four|4|five|5|six|6|seven|7|eight|8|nine|9).*")
	last_r, _ := regexp.Compile(".*(zero|0|one|1|two|2|three|3|four|4|five|5|six|6|seven|7|eight|8|nine|9).*?")
	var sum int
	for {
		var input string
		fmt.Scanln(&input)
		if input == "" {
			break
		}
		first := LUT[first_r.FindStringSubmatch(input)[1]]
		last := LUT[last_r.FindStringSubmatch(input)[1]]
		sum += (first * 10) + last
		fmt.Printf("String: %s, first int: %d, last int: %d, current sum: %d \n", input, first, last, sum)
	}
	fmt.Printf("final sum: %d \n", sum)
}

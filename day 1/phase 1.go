package main

import "fmt"

func main() {
	var sum int
	for {
		var input string
		fmt.Scanln(&input)
		if input == "" {
			break
		}
		first := -1
		last := -1
		for i := 0; i < len(input); i++ {
			// com certeza tem jeitos melhores de mexer com strings ASCII
			// isInteger
			if '0' <= input[i] && input[i] <= '9' {
				if first == -1 {
					// converte para o int com base na tabela ascii
					first = int(input[i]) - 48
				}
				last = int(input[i]) - 48
			}
		}
		sum += (first * 10) + last
		fmt.Printf("String: %s, first int: %d, last int: %d, current sum: %d \n", input, first, last, sum)
	}
	fmt.Printf("final sum: %d \n", sum)
}

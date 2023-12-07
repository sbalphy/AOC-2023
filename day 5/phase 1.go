package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func read_block(scanner *bufio.Scanner, mapping map[int][2]int) {
	// come o header do bloco
	scanner.Scan()
	for {
		scanner.Scan()
		input := scanner.Text()
		if input == " " || input == "" {
			break
		}
		vals_string := strings.Split(input, " ")
		dest, _ := strconv.Atoi(vals_string[0])
		src, _ := strconv.Atoi(vals_string[1])
		length, _ := strconv.Atoi(vals_string[2])
		mapping[dest] = [2]int{src, length}
	}
	fmt.Println(mapping)
}

func compose(map_one map[int][2]int, map_two map[int][2]int) map[int][2]int {

}
func main() {
	f, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(f)
	r_cleaner, _ := regexp.Compile("[0-9]+")
	scanner.Scan()
	seeds_string := r_cleaner.FindAllString(scanner.Text(), -1)
	// lemos cada bloco separadamente
	var seeds []int
	for i := 0; i < len(seeds_string); i++ {
		val, _ := strconv.Atoi(seeds_string[i])
		seeds = append(seeds, val)
	}
	// come o espaco entre as seeds iniciais e o segundo bloco
	scanner.Scan()
	// definimos e preenchemos os outros blocos
	soil_to_seed := make(map[int][2]int)
	read_block(scanner, soil_to_seed)
	fertilizer_to_soil := make(map[int][2]int)
	read_block(scanner, fertilizer_to_soil)
	water_to_fertilizer := make(map[int][2]int)
	read_block(scanner, water_to_fertilizer)
	light_to_water := make(map[int][2]int)
	read_block(scanner, light_to_water)
	temperature_to_light := make(map[int][2]int)
	read_block(scanner, temperature_to_light)
	humidity_to_temperature := make(map[int][2]int)
	read_block(scanner, humidity_to_temperature)
	location_to_humidity := make(map[int][2]int)
	read_block(scanner, location_to_humidity)

}

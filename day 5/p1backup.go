package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type sortedmap struct {
	keys    []int
	mapping map[int][2]int
}

func newSortedMap(mapping map[int][2]int) sortedmap {
	smap := sortedmap{mapping: mapping}
	return smap
}

func fill_intervals(mapping map[int][2]int) []int {
	var keys []int
	for k := range mapping {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var start int
	initial_size := len(keys)
	for i := 0; i < initial_size; i++ {
		if start < keys[i] {
			keys = append(keys, start)
			mapping[start] = [2]int{start, keys[i] - start}
		}
		start = keys[i] + mapping[keys[i]][1]
	}
	return keys
}
func read_block(scanner *bufio.Scanner, mapping map[int][2]int) []int {
	// come o header do bloco
	scanner.Scan()
	for {
		scanner.Scan()
		input := scanner.Text()
		if input == " " || input == "" {
			break
		}
		fmt.Println(input)
		vals_string := strings.Split(input, " ")
		dest, _ := strconv.Atoi(vals_string[0])
		src, _ := strconv.Atoi(vals_string[1])
		length, _ := strconv.Atoi(vals_string[2])
		mapping[dest] = [2]int{src, length}
	}
	return fill_intervals(mapping)
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
	var soil_to_seed sortedmap
	soil_to_seed = newSortedMap(make(map[int][2]int))
	soil_to_seed.keys = read_block(scanner, soil_to_seed.mapping)
	fmt.Println(soil_to_seed)
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

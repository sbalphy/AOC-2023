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
	// regex to find game number and extract list of cube reveals
	r_initial_cleanup, _ := regexp.Compile("Game ([0-9]+): (.*)")
	// regex to identify number of cubes
	r_red, _ := regexp.Compile("([0-9]+) red")
	r_green, _ := regexp.Compile("([0-9]+) green")
	r_blue, _ := regexp.Compile("([0-9]+) blue")
	// max values for each cube
	max_red := 12
	max_green := 13
	max_blue := 14

	var sum int
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println(input)
		game, _ := strconv.Atoi(r_initial_cleanup.FindStringSubmatch(input)[1])
		valid := true
		reveals := strings.Split(r_initial_cleanup.FindStringSubmatch(input)[2], ";")
		for i := 0; i < len(reveals); i++ {
			// como eu faco um try except em go
			var red int
			red_matches := r_red.FindStringSubmatch(reveals[i])
			if len(red_matches) == 2 {
				red, _ = strconv.Atoi(red_matches[1])
			}
			var green int
			green_matches := r_green.FindStringSubmatch(reveals[i])
			if len(green_matches) == 2 {
				green, _ = strconv.Atoi(green_matches[1])
			}
			var blue int
			blue_matches := r_blue.FindStringSubmatch(reveals[i])
			if len(blue_matches) == 2 {
				blue, _ = strconv.Atoi(blue_matches[1])
			}
			if red > max_red || green > max_green || blue > max_blue {
				valid = false
				fmt.Printf("Validity for game %d failed at reveal %d \n", game, i)
				break
			}
		}
		if valid {
			sum += game
		}
		fmt.Printf("Game %d, validity %t, current sum %d \n", game, valid, sum)
	}
	fmt.Printf("final sum: %d \n", sum)
}

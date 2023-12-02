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

	var sum int
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println(input)
		game, _ := strconv.Atoi(r_initial_cleanup.FindStringSubmatch(input)[1])
		// max values for each cube
		max_red := 0
		max_green := 0
		max_blue := 0
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
			if red > max_red {
				max_red = red
			}
			if green > max_green {
				max_green = green
			}
			if blue > max_blue {
				max_blue = blue
			}
		}
		sum += max_red * max_green * max_blue
		fmt.Printf("The minimum values for game %d are R:%d, G:%d, B:%d. Power:%d, sum:%d \n", game, max_red, max_green, max_blue, max_red*max_green*max_blue, sum)
	}
	fmt.Printf("final sum: %d \n", sum)
}

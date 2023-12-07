package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	r_cleaner, _ := regexp.Compile("[0-9]+")
	scanner.Scan()
	time_string := r_cleaner.FindAllString(scanner.Text(), -1)
	scanner.Scan()
	dist_string := r_cleaner.FindAllString(scanner.Text(), -1)
	var races [][2]int
	for i := 0; i < len(time_string); i++ {
		time, _ := strconv.Atoi(time_string[i])
		dist, _ := strconv.Atoi(dist_string[i])
		val := [2]int{time, dist}
		races = append(races, val)
	}
	product := 1
	for i := 0; i < len(races); i++ {
		total_time := races[i][0]
		record_dist := races[i][1]
		var ways_to_beat int
		for held_time := 1; held_time < total_time; held_time++ {
			dist := held_time * (total_time - held_time)
			if dist > record_dist {
				ways_to_beat++
			}
		}
		product *= ways_to_beat
	}
	fmt.Println(product)
}

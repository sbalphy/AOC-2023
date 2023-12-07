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
	r_cleaner, _ := regexp.Compile("[0-9]+")
	scanner.Scan()
	time_str := strings.Join(r_cleaner.FindAllString(scanner.Text(), -1), "")
	scanner.Scan()
	dist_str := strings.Join(r_cleaner.FindAllString(scanner.Text(), -1), "")
	time, _ := strconv.ParseUint(time_str, 10, 64)
	dist, _ := strconv.ParseUint(dist_str, 10, 64)
	fmt.Println(time, dist)
	// -x2 + time * x - dist = 0
	del := time*time - 4*dist
	fmt.Println(del)
	root := math.Sqrt(float64(del))
	fmt.Printf("root is %f \n", root)
	x1 := float64(time)/2 + root/2
	x2 := float64(time)/2 - root/2
	fmt.Printf("%f - %f = %f) \n", x1, x2, x1-x2)

}

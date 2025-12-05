package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var lowerRange []int
var higherRange []int
var nums []int

var freshCount int = 0

func main() {
	file, err := os.Open("day5/part1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) > 0 {
			if strings.Contains(fields[0], "-") {
				split := strings.Split(fields[0], "-")
				lower, _ := strconv.Atoi(split[0])
				high, _ := strconv.Atoi(split[1])
				lowerRange = append(lowerRange, lower)
				higherRange = append(higherRange, high)
			} else {
				num, _ := strconv.Atoi(fields[0])
				nums = append(nums, num)
			}
		}
	}

	checkIfFresh()
	fmt.Println(freshCount)
}

func checkIfFresh() {
	for _, num := range nums {
		for i := 0; i < len(higherRange); i++ {
			if num >= lowerRange[i] && num <= higherRange[i] {
				freshCount++
				break
			}
		}
	}
}

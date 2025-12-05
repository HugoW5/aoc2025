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

var freshIDs int = 0
var overlaps int = 0

func main() {
	file, err := os.Open("day5/part2/input.txt")
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
			}
		}
	}

	findOverlappingIntervals() // merge intervals
	findFreshIDs()
	fmt.Println(freshIDs - overlaps)
}

func findOverlappingIntervals() {
	for i := 0; i < len(lowerRange); i++ {
		for j := i + 1; j < len(lowerRange); j++ {

			a1, a2 := lowerRange[i], higherRange[i]
			b1, b2 := lowerRange[j], higherRange[j]

			if a1 <= b2 && b1 <= a2 {
				fmt.Printf("%v-%v & %v-%v overlap\n", a1, a2, b1, b2)
				overlaps++
			}
		}
	}
}

func findFreshIDs() {
	for rangeIndex, low := range lowerRange {
		freshIDs += higherRange[rangeIndex] - low
	}
}

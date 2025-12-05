package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var banks []string
var joltage int = 0

func main() {
	file, err := os.Open("day3/part2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		banks = append(banks, fields...)
	}
	findHighestJoltage()

	fmt.Println(joltage)
}

func findHighestJoltage() {
	for _, bank := range banks {
		joltage += findLargestJoltageCombination(bank)
	}
}

func findLargestJoltageCombination(num string) int {
	toRemove := len(num) - 12
	stack := make([]byte, 0, len(num))

	for i := 0; i < len(num); i++ {
		for len(stack) > 0 && toRemove > 0 && stack[len(stack)-1] < num[i] {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, num[i])
	}
	tmpNum, _ := strconv.Atoi(string(stack[:12]))
	return tmpNum
}

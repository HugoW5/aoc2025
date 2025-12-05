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

func main() {
	file, err := os.Open("day3/part1/input.txt")
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
}

func findHighestJoltage() {
	joltage := 0

	for i := 0; i < len(banks); i++ {
		highestJoltage := 0
		for j := 0; j < len(banks[i])-1; j++ {
			num1 := string(banks[i][j])
			for k := j + 1; k < len(banks[i]); k++ {
				num2 := string(banks[i][k])
				num, _ := strconv.Atoi(num1 + num2)
				if num > highestJoltage {
					highestJoltage = num
				}
			}
		}
		joltage += highestJoltage
	}
	fmt.Println(joltage)
}

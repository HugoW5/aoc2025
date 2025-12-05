package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var ranges [][]int
var total int = 0

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Split(text, ",")
		for _, field := range fields {
			span := strings.Split(field, "-")
			min, _ := strconv.Atoi(span[0])
			max, _ := strconv.Atoi(span[1])
			ranges = append(ranges, []int{min, max})
		}
	}

	findInvalidIDs()
	fmt.Println(total)
}

func findInvalidIDs() {
	for _, span := range ranges {
		// lower span
		n := span[0]
		for n <= span[1] {
			if checkForPattern(strconv.Itoa(n)) {
				total += n
				fmt.Println(n)
			}
			n++
		}
	}
}
func checkForPattern(n string) bool {
	if len(n)%2 != 0 {
		return false
	}
	sameCharCount := 0
	for i := 0; i < len(n); i++ {
		if n[0] == n[i] {
			sameCharCount++
		}
	}
	if sameCharCount == len(n) {
		return true
	}

	if n[len(n)/2:] == n[:len(n)/2] {
		return true
	}

	return false
}

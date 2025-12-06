package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var nums [][]string = make([][]string, 1000)
var total int = 0

func main() {
	file, _ := os.Open("day6/part1/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for i, v := range strings.Fields(scanner.Text()) {
			nums[i] = append(nums[i], v)
		}
	}
	//Println(nums)
	calculateTotal()
	fmt.Println(total)
}

func calculateTotal() {
	for _, n := range nums {
		tmpTotal := 0
		for i := 0; i < len(n)-1; i++ {
			num, _ := strconv.Atoi(n[i])
			if n[len(n)-1] == "+" {
				tmpTotal += num
			} else {
				if tmpTotal == 0 {
					tmpTotal = num
				} else {
					tmpTotal *= num
				}
			}
		}
		fmt.Println(tmpTotal)
		total += tmpTotal
	}
}

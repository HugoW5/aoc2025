package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nums [][]string = make([][]string, 4)
var total int = 0

func main() {
	file, _ := os.Open("day6/part2/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for i, v := range strings.Fields(scanner.Text()) {
			nums[i] = append(nums[i], v)
		}
	}
	calculateTotal()
	fmt.Println(total)
}

func calculateTotal() {
	for i := 0; i < len(nums); i++ {
		// numbers := []string{}
		for j := 0; j < len(nums[i]); j++ {
			for k := len(nums[i][j]) - 1; k >= 0; k-- {
				fmt.Println(string(nums[i][j][k]))
			}
		}
		fmt.Println("-----")
	}
}

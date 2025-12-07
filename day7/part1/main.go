package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var rows []string
var beamsIndex [][]int = make([][]int, 142)
var splits int = 0

func main() {
	file, _ := os.Open("day7/part1/input.txt")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	i := strings.Index(rows[0], "S")
	rows[1] = replaceAtIndex(rows[1], '|', i)
	beamsIndex[1] = append(beamsIndex[1], i)

	splitBeam()
	print()
	fmt.Println(splits)
}

func splitBeam() {
	for !strings.Contains(rows[len(rows)-1], "|") {
		for i := 0; i < len(rows)-1; i++ {
			for _, beam := range beamsIndex[i] {
				if rows[i+1][beam] == '.' {
					rows[i+1] = replaceAtIndex(rows[i+1], '|', beam)
					beamsIndex[i+1] = append(beamsIndex[i+1], beam)
				} else if rows[i+1][beam] == '^' {
					beamsIndex[i+1] = append(beamsIndex[i+1], beam+1)
					beamsIndex[i+1] = append(beamsIndex[i+1], beam-1)

					rows[i+1] = replaceAtIndex(rows[i+1], '|', beam+1)
					rows[i+1] = replaceAtIndex(rows[i+1], '|', beam-1)
					splits++
				}
			}
		}
	}
}

func replaceAtIndex(in string, r rune, index int) string {
	out := []rune(in)
	out[index] = r
	return string(out)
}

func print() {
	for _, r := range rows {
		fmt.Println(r)
	}
}

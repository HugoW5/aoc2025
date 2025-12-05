package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var grid []string
var forklifts int = 0

func main() {
	file, err := os.Open("day4/part2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		grid = append(grid, fields...)
	}

	for i := 0; i < 100; i++ {
		calculateForkliftPositions()
	}

	fmt.Println(forklifts)
}

func calculateForkliftPositions() {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if string(grid[i][j]) == "@" {
				checkAdjustmentPositions(i, j)
			}
		}
	}
}

func checkAdjustmentPositions(i, j int) {
	adjustmentPositions := []struct {
		X int
		Y int
	}{
		{
			X: -1,
			Y: -1,
		},
		{
			X: 0,
			Y: -1,
		},
		{
			X: 1,
			Y: -1, ////123
		},
		{
			X: -1,
			Y: 0,
		},
		{
			X: 1,
			Y: 0, // 45
		},
		{
			X: -1,
			Y: 1,
		},
		{
			X: 0,
			Y: 1,
		},
		{
			X: 1,
			Y: 1, // 678
		},
	}
	paperRolls := 0
	for _, pos := range adjustmentPositions {
		tmpI := (i + pos.Y)
		tmpJ := (j + pos.X)
		if tmpI >= 0 && tmpI <= len(grid)-1 {
			if tmpJ >= 0 && tmpJ <= len(grid[0])-1 {
				if string(grid[tmpI][tmpJ]) == "@" {
					paperRolls++
				}
			}
		}
	}
	if paperRolls < 4 {
		forklifts++
		tmpStr := replaceAtIndex(grid[i], '.', j)
		grid[i] = tmpStr
	}

}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

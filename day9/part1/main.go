package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var tilesX []int
var tilesY []int

func main() {
	file, _ := os.Open("day9/part1/input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		corrods := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(corrods[0])
		y, _ := strconv.Atoi(corrods[1])
		tilesX = append(tilesX, x)
		tilesY = append(tilesY, y)
	}

	largest := findLargestSquare()
	fmt.Println(int(largest))
}

func findLargestSquare() float64 {
	largest := 0.0
	for i := 0; i < len(tilesX); i++ {
		for j := 0; j < len(tilesY); j++ {
			distanceX := math.Abs(float64(tilesX[i]-tilesX[j]) - 1)
			distanceY := math.Abs(float64(tilesY[i]-tilesY[j]) - 1)

			if (distanceX * distanceY) > largest {
				largest = (distanceX * distanceY)
			}
		}
	}
	return largest
}

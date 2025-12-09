package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type JunctionBox struct {
	ID int
	X  int
	Y  int
	Z  int
}

var junctionBoxes []JunctionBox

func main() {
	file, _ := os.Open("day8/part1/input.txt")
	scanner := bufio.NewScanner(file)
	id := 0
	for scanner.Scan() {
		for _, line := range strings.Fields(scanner.Text()) {
			cordinates := strings.Split(line, ",")
			tmpX, _ := strconv.Atoi(cordinates[0])
			tmpY, _ := strconv.Atoi(cordinates[1])
			tmpZ, _ := strconv.Atoi(cordinates[2])
			junctionBoxes = append(junctionBoxes, JunctionBox{
				ID: id,
				X:  tmpX,
				Y:  tmpY,
				Z:  tmpZ,
			})
			id++
		}
	}

	fmt.Printf("%+v\n", findClosestBox(junctionBoxes[0]))

	fmt.Println(junctionBoxes)
}

func findClosestBox(b1 JunctionBox) JunctionBox {
	tmpBoxIndex := 0
	var distance float64 = 1000000
	for i, b2 := range junctionBoxes {
		if b2.ID != b1.ID {
			d := calculateDistance(b1, b2)
			fmt.Println(d)
			if d < distance {
				distance = d
				tmpBoxIndex = i
			}
		}
	}
	return junctionBoxes[tmpBoxIndex]
}

func calculateDistance(j1, j2 JunctionBox) float64 {
	x := math.Pow(float64(j2.X)-float64(j1.X), 2)
	y := math.Pow(float64(j2.Y)-float64(j1.Y), 2)
	z := math.Pow(float64(j2.Z)-float64(j1.Z), 2)

	return math.Sqrt(x + y + z)
}

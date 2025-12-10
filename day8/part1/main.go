package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
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
var connections [][]int = make([][]int, 20)

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

	connectBoxes()
	arr := mapCircuitBoxes()
	fmt.Println("-----------")
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func mapCircuitBoxes() []int {
	var circuts []int = make([]int, 20)
	for i := 0; i < len(connections); i++ {
		for _, v := range connections[i] {
			if slices.Contains(connections[v], i) {
				break
			}
			circuts[i] += len(connections[v])
		}
		fmt.Println(i, connections[i])
	}
	return circuts
}

func connectBoxes() {
	for _, box := range junctionBoxes {
		closest := findClosestBox(box)
		connections[closest.ID] = append(connections[closest.ID], box.ID)
	}
}

func findClosestBox(b1 JunctionBox) JunctionBox {
	tmpBoxIndex := 0
	var distance float64 = 1000000
	for i, b2 := range junctionBoxes {
		if b2.ID != b1.ID {
			d := calculateDistance(b1, b2)
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

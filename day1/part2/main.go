package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var rotations []string
var dial int = 50
var timesDialAtZero = 0

func main() {

	file, err := os.Open("day1/part2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		rotations = append(rotations, fields...)
	}
	calculateDial()
	fmt.Println(timesDialAtZero)
}

func calculateDial() {

	for _, v := range rotations {
		direction := v[0]
		turnAmount, _ := strconv.Atoi(v[1:])
		if string(direction) == "L" {
			turnDialLeft(turnAmount)
		}
		if string(direction) == "R" {
			turnDialRight(turnAmount)
		}
	}
}

func turnDialRight(deg int) {
	for i := 0; i < deg; i++ {
		dial = (dial + 1) % 100
		if dial == 0 {
			timesDialAtZero++
		}
	}
}

func turnDialLeft(deg int) {
	for i := 0; i < deg; i++ {
		dial = (dial - 1 + 100) % 100
		if dial == 0 {
			timesDialAtZero++
		}
	}
}

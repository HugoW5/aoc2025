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
var timesDialAtZero int

func main() {
	file, err := os.Open("day1/part1/input.txt")
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
		dir := v[0]
		turnAmount, err := strconv.Atoi(v[1:])
		if err != nil {
			log.Fatalf("bad rotation %q: %v", v, err)
		}
		turnAmount = turnAmount % 100

		if string(dir) == "R" {
			turnDialRight(turnAmount)
		} else if string(dir) == "L" {
			turnDialLeft(turnAmount)
		}
		if dial == 0 {
			timesDialAtZero++
		}
	}
}

func turnDialRight(deg int) {
	dial = (dial + deg) % 100
}

func turnDialLeft(deg int) {
	dial = (dial - deg%100 + 100) % 100
}

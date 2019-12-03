package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type Instruction struct {
	direction string
	distance int
}

type Position struct {
	x int
	y int
	steps int
}

func main() {
	rawInput, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal("Something went wrong while reading the input file.")
		panic(err)
	}
	input := strings.ReplaceAll(string(rawInput), "\r", "")
	redWireRawInstructions, blueWireRawInstructions := strings.Split(strings.Split(input, "\n")[0], ","), strings.Split(strings.Split(input, "\n")[1], ",")

	redWireInstructions := populateWireInstructions(redWireRawInstructions)
	blueWireInstructions := populateWireInstructions(blueWireRawInstructions)

	var redPositions []Position
	redPos := Position{0, 0, 0}
	for _, r := range redWireInstructions {
		for i := 1; i <= r.distance; i++ {
			if r.direction == "R" {
				redPos.x++
			} else if r.direction == "L" {
				redPos.x--
			} else if r.direction == "U" {
				redPos.y++
			} else if r.direction == "D" {
				redPos.y--
			}
			redPos.steps++
			redPositions = append(redPositions, redPos)
		}
	}
	var bluePositions []Position
	bluePos := Position{0, 0, 0}
	for _, b := range blueWireInstructions {
		for i := 1; i <= b.distance; i++ {
			if b.direction == "R" {
				bluePos.x++
			} else if b.direction == "L" {
				bluePos.x--
			} else if b.direction == "U" {
				bluePos.y++
			} else if b.direction == "D" {
				bluePos.y--
			}
			bluePos.steps++
			bluePositions = append(bluePositions, bluePos)
		}
	}

	centralPort := Position{0, 0, 0}
	minDistance := math.MaxInt64
	minSteps := math.MaxInt64
	var x int
	var y int
	for _, rp := range redPositions {
		for _, bp := range bluePositions {
			if rp.x == bp.x && rp.y == bp.y {
				if dist := calculateManhattanDistance(rp, centralPort); dist < minDistance || rp.steps + bp.steps < minSteps {
					x = rp.x
					y = rp.y
					minDistance = dist
					minSteps = rp.steps + bp.steps
				}
			}
		}
	}
	fmt.Printf("Fastest wire intersection is: %d, %d\n", x, y)
	fmt.Printf("Distance from Central Port to closest wire intersection is: %d in %d total steps\n", minDistance, minSteps)
}

func populateWireInstructions(rawInstructions []string) []Instruction {
	wireInstructions := make([]Instruction, len(rawInstructions))
	currX := 0
	currY := 0
	maxX := 0
	maxY := 0
	for i, ins := range rawInstructions {
		direction := string([]rune(ins)[0])
		distance, err := strconv.ParseInt(string([]rune(ins)[1:]), 10, 64)
		if err != nil {
			panic("Error parsing distance")
		}

		if direction == "R" {
			currX += int(distance)
			if currX > maxX {
				maxX = currX
			}
		} else if direction == "L" {
			oldCurrX := currX
			currX -= int(distance)
			if currX < 0 {
				maxX += int(distance) - maxX
				currX = (oldCurrX - currX) - int(distance)
			}
		} else if direction == "U" {
			currY += int(distance)
			if currY > maxY {
				maxY = currY
			}
		} else if direction == "D" {
			oldCurrY := currY
			currY -= int(distance)
			if currY < 0 {
				maxY += int(distance) - maxY
				currY = (oldCurrY - currY) - int(distance)
			}
		}

		wireInstructions[i] = Instruction{
			direction: direction,
			distance:  int(distance),
		}
	}
	return wireInstructions
}

func calculateManhattanDistance(x Position, y Position) int {
	return int(math.Abs(float64(x.x - y.x)) + math.Abs(float64(x.y - y.y)))
}
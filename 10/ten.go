package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

const ASTEROID = "#"
const NOTHING = "."

func main() {
	m, height, width := readFile("input")
	posX := 0
	posY := 0
	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if m[y][x] == ASTEROID {
				c := findAsteroidsInSight(x, y, m, height, width)
				if c >= count {
					count = c
					posX = x
					posY = y
				}
			}
		}
	}
	fmt.Printf("%d Asteroids seen from position [%d, %d]", count, posX, posY)
}

func findAsteroidsInSight(x int, y int, m map[int]map[int]string, height int, width int) int {
	angles := map[float64]bool{}
	count := 0
	for y2 := 0; y2 < height; y2++ {
		for x2 := 0; x2 < width; x2++ {
			fmt.Printf("Checking pos1[%d, %d] and pos2[%d, %d]\n", x, y, x2, y2)
			if m[y2][x2] == ASTEROID && x != x2 && y != y2 {
				fmt.Printf("Found Asteroid at pos2[%d, %d]\n", x2, y2)
				if _, ok := angles[getAngle(x, y, x2, y2)]; !ok {
					count++
				}
			}
		}
	}
	return count
}

func getAngle(x1 int, y1 int, x2 int, y2 int) float64 {
	dX := x2 - x1
	dY := y1 - y2
	r := math.Atan2(float64(dY), float64(dX))
	fmt.Println(r)
	return r
}

func readFile(file string) (map[int]map[int]string, int, int) {
	input, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Something went wrong while reading the input file.")
		panic(err)
	}
	m := map[int]map[int]string{}
	for i, r := range strings.Split(string(input), "\n") {
		m[i] = map[int]string{}
		for j, c := range r {
			m[i][j] = string(c)
		}
	}
	height := len(m)
	width := len(m[0])
	return m, height, width
}

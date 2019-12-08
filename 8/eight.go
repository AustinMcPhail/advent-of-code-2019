package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const IMAGE_WIDTH = 25 // pixels
const IMAGE_HEIGHT = 6 // pixels
//const IMAGE_WIDTH = 3 // pixels
//const IMAGE_HEIGHT = 2 // pixels

const BLACK = 0
const WHITE = 1
const TRANSPARENT = 2
const BLACK_CHAR = " "
const WHITE_CHAR = "█"
const TRANSPARENT_CHAR = " "

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal("Something went wrong while reading the input file.")
		panic(err)
	}
	intStrings := strings.Split(string(input), "")
	imageData := make([]int, 0)
	for _, s := range intStrings {
		current, _ := strconv.Atoi(s)
		imageData = append(imageData, current)
	}
	imagePixels := map[int][IMAGE_HEIGHT][IMAGE_WIDTH]int{}
	layer := 0
	for len(imageData) != 0 {
		l := [IMAGE_HEIGHT][IMAGE_WIDTH]int{}
		for i := 0; i < IMAGE_HEIGHT; i++ {
			for j := 0; j < IMAGE_WIDTH; j++ {
				l[i][j] = imageData[0]
				imageData = removeAt(imageData, 0)
			}
		}
		imagePixels[layer] = l
		layer++
	}

	image := [IMAGE_HEIGHT][IMAGE_WIDTH]int{}
	for i := 0; i < IMAGE_HEIGHT; i++ {
		for j := 0; j < IMAGE_WIDTH; j++ {
			image[i][j] = -1
		}
	}

	for l := len(imagePixels) - 1; l > 0; l-- {
		for i := 0; i < IMAGE_HEIGHT; i++ {
			for j := 0; j < IMAGE_WIDTH; j++ {
				curr := image[i][j]
				if curr == -1 {
					image[i][j] = imagePixels[l][i][j]
				} else {
					switch curr {
					case BLACK:
						if imagePixels[l][i][j] != TRANSPARENT {
							image[i][j] = imagePixels[l][i][j]
						}
						break
					case WHITE:
						if imagePixels[l][i][j] != TRANSPARENT {
							image[i][j] = imagePixels[l][i][j]
						}
						break
					case TRANSPARENT:
						image[i][j] = imagePixels[l][i][j]
						break
					}
				}
				//fmt.Print(image[i][j])
			}
			//fmt.Println()
		}
		//fmt.Println()
	}

	for i := 0; i < IMAGE_HEIGHT; i++ {
		for j := 0; j < IMAGE_WIDTH; j++ {
			switch image[i][j] {
			case BLACK:
				fmt.Print(BLACK_CHAR + " ")
				break
			case WHITE:
				fmt.Print(WHITE_CHAR + " ")
				break
			case TRANSPARENT:
				fmt.Print(TRANSPARENT_CHAR + " ")
				break
			}
		}
		fmt.Println()
	}

}

func findOccurrences(target int, arr []int) int {
	count := 0
	for _, e := range arr {
		if e == target {
			count++
		}
	}
	return count
}

func removeAt(elements []int, index int) []int {
	e := append([]int(nil), elements...)
	return append(e[:index], e[index+1:]...)
}

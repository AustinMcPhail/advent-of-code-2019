package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	rawInput, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal("Something went wrong while reading the input file.")
		panic(err)
	}
	input := strings.Split(strings.ReplaceAll(string(rawInput), "\r", ""), "\n")

	fmt.Println(countCriteriaMatchesInRangePart1(input[0], input[1]))
	fmt.Println(countCriteriaMatchesInRangePart2(input[0], input[1]))
}

func countCriteriaMatchesInRangePart1(lower string, upper string) int {
	l, err := strconv.Atoi(lower)
	if err != nil {
		panic(err)
	}
	u, err := strconv.Atoi(upper)

	matchArr := make([]int, 0)
	for i := l; i < u; i++ {
		if validNumberOrder(i) && hasAdjacentTwins(i) {
			matchArr = append(matchArr, i)
		}
	}
	return len(matchArr)
}

func countCriteriaMatchesInRangePart2(lower string, upper string) int {
	l, err := strconv.Atoi(lower)
	if err != nil {
		panic(err)
	}
	u, err := strconv.Atoi(upper)

	matchArr := make([]int, 0)
	for i := l; i < u; i++ {
		if validNumberOrder(i) && hasAdjacentTwins2(i) {
			matchArr = append(matchArr, i)
		}
	}
	return len(matchArr)
}

func hasAdjacentTwins(i int) bool {
	sNum := strconv.Itoa(i)
	return sNum[0] == sNum[1] ||
		sNum[1] == sNum[2] ||
		sNum[2] == sNum[3] ||
		sNum[3] == sNum[4] ||
		sNum[4] == sNum[5]
}

func hasAdjacentTwins2(i int) bool {
	sNum := strconv.Itoa(i)
	for _, c := range sNum {
		foundTwins := strings.Contains(sNum, string(c) + "" + string(c))
		foundMore := strings.Contains(sNum, string(c) + "" + string(c) + "" + string(c))

		if foundTwins && !foundMore {
			return true
		}
	}
	return false
}

func validNumberOrder(i int) bool {
	sNum := strconv.Itoa(i)
	iNum := make([]int, len(sNum))
	iNum[0], _ = strconv.Atoi(string(rune(sNum[0])))
	iNum[1], _ = strconv.Atoi(string(rune(sNum[1])))
	iNum[2], _ = strconv.Atoi(string(rune(sNum[2])))
	iNum[3], _ = strconv.Atoi(string(rune(sNum[3])))
	iNum[4], _ = strconv.Atoi(string(rune(sNum[4])))
	iNum[5], _ = strconv.Atoi(string(rune(sNum[5])))

	return iNum[0] <= iNum[1] &&
		iNum[1] <= iNum[2] &&
		iNum[2] <= iNum[3] &&
		iNum[3] <= iNum[4] &&
		iNum[4] <= iNum[5]
}



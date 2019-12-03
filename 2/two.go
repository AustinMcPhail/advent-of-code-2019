package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const TARGET int = 19690720

type Result struct {
	value int
	noun int
	verb int
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal("Something went wrong while reading the input file.")
		panic(err)
	}
	intStrings := strings.Split(string(input), ",")
	var intCode []int
	for _, i := range intStrings {
		current, err := strconv.Atoi(i)
		if err != nil {
			panic(fmt.Sprintf("Something went wrong whilst converting a string {%v} to int", i))
		}
		intCode = append(intCode, current)
	}
	r := make(chan Result)
	go findNounAndVerb(TARGET, intCode, r)
	res := <-r
	fmt.Printf("100 * %d (noun) + %d (verb) = %d", res.noun, res.verb, 100 * res.noun + res.verb)
}

func findNounAndVerb(target int, intCode []int, r chan Result) {
	c := make(chan Result)
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			code := make([]int, len(intCode))
			copy(code, intCode)
			go evaluateIntCode(noun, verb, code, c)
		}
	}
	for res := range c {
		if res.value == target {
			r <- res
			break
		}
	}
}

func evaluateIntCode(noun int, verb int, intCode []int, c chan Result) {
	intCode[1] = noun
	intCode[2] = verb
	done := false
	for i := 0; i < len(intCode) && !done; i += 4 {
		switch intCode[i] {
		case 1:
			xPos := intCode[i + 1]
			yPos := intCode[i + 2]
			resPos := intCode[i + 3]
			intCode[resPos] = intCode[xPos] + intCode[yPos]
			break
		case 2:
			xPos := intCode[i + 1]
			yPos := intCode[i + 2]
			resPos := intCode[i + 3]
			intCode[resPos] = intCode[xPos] * intCode[yPos]
			break
		case 99:
			done = true
			break
		default:
			panic("Unknown opcode")
			break
		}
	}

	result := intCode[0]
	c <- Result{
		value: result,
		noun:  noun,
		verb:  verb,
	}
}
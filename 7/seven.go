package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Amp struct {
	prev *Amp
	input int
	intCode map[int]int
	output int
	next *Amp
}

const START = 0
var ELEMENTS_PART1 = []int{0, 1, 2, 3, 4}
var ELEMENTS_PART2 = []int{5, 6, 7, 8, 9}
func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal("Something went wrong while reading the input file.")
		panic(err)
	}

	combos := getAllCombos(ELEMENTS_PART1, make([]int, 0))
	fmt.Println(getThrust(combos, input))

	combos = getAllCombos(ELEMENTS_PART2, make([]int, 0))
	fmt.Println(getThrustWithFeedback(combos, input))
}

func getThrust(combos [][]int, input []byte) int {
	highest := 0
	for _, combo := range combos {
		aOut := evaluateIntCode([]int{combo[0], START}, input)
		bOut := evaluateIntCode([]int{combo[1], aOut}, input)
		cOut := evaluateIntCode([]int{combo[2], bOut}, input)
		dOut := evaluateIntCode([]int{combo[3], cOut}, input)
		eOut := evaluateIntCode([]int{combo[4], dOut}, input)

		if eOut > highest {
			highest = eOut
		}
	}
	return highest
}

func getThrustWithFeedback(combos [][]int, input []byte) int {
	highest := 0
	var aOut int
	var eOut int
	for i, combo := range combos {
		if i == 0 {
			aOut = evaluateIntCode([]int{combo[0], START}, input)
		} else {
			aOut = evaluateIntCode([]int{combo[0], eOut}, input)
		}
		bOut := evaluateIntCode([]int{combo[1], aOut}, input)
		cOut := evaluateIntCode([]int{combo[2], bOut}, input)
		dOut := evaluateIntCode([]int{combo[3], cOut}, input)
		eOut := evaluateIntCode([]int{combo[4], dOut}, input)

		if eOut > highest {
			highest = eOut
		}
	}
	return highest
}

func removeAt (elements []int, index int) []int {
	e := append([]int(nil), elements...)
	return append(e[:index], e[index+1:]...)
}

func getAllCombos(elements []int, prefix []int) [][]int {
	length := len(elements)
	if length == 0 {
		return [][]int{prefix}
	}
	combos := make([][]int, 0)
	for i := 0; i < len(elements); i++ {
		p := append(prefix, elements[i])
		e := removeAt(elements, i)
		combos = append(combos, getAllCombos(e, p)...)
	}
	return combos
}

func evaluateIntCode(input []int, rawIntCode []byte) int {
	intStrings := strings.Split(string(rawIntCode), ",")
	intCode := make(map[int]int)
	for i, s := range intStrings {
		current, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Something went wrong whilst converting a string {%v} to int", i))
		}
		intCode[i] = current
	}

	var output int
	done := false
	increment := 4
	for i := 0; i < len(intCode) && !done; i += increment {
		var opcode int
		firstParamMode := 0
		secondParamMode := 0
		switch len(strconv.Itoa(intCode[i])) {
		case 1, 2:
			opcode = intCode[i]
			break
		case 3:
			opcode, _ = strconv.Atoi(strconv.Itoa(intCode[i])[len(strconv.Itoa(intCode[i]))-2:])
			firstParamMode, _ = strconv.Atoi(string(strconv.Itoa(intCode[i])[0]))
			break
		case 4:
			opcode, _ = strconv.Atoi(strconv.Itoa(intCode[i])[len(strconv.Itoa(intCode[i]))-2:])
			firstParamMode, _ = strconv.Atoi(string(strconv.Itoa(intCode[i])[1]))
			secondParamMode, _ = strconv.Atoi(string(strconv.Itoa(intCode[i])[0]))
			break
		case 5:
			opcode, _ = strconv.Atoi(strconv.Itoa(intCode[i])[len(strconv.Itoa(intCode[i]))-2:])
			firstParamMode, _ = strconv.Atoi(string(strconv.Itoa(intCode[i])[2]))
			secondParamMode, _ = strconv.Atoi(string(strconv.Itoa(intCode[i])[1]))
			break
		}
		//fmt.Printf("POSITION: %d; VALUE: %d\n", i, intCode[i])
		//fmt.Printf("OPCODE: %d\n",opcode)
		switch opcode {
		case 1:
			var x int
			var y int
			if firstParamMode == 0 {
				x = intCode[intCode[i + 1]]
			} else {
				x = intCode[i + 1]
			}
			if secondParamMode == 0 {
				y = intCode[intCode[i + 2]]
			} else {
				y = intCode[i + 2]
			}
			resPos := intCode[i + 3]
			intCode[resPos] = x + y
			increment = 4
			break
		case 2:
			var x int
			var y int
			if firstParamMode == 0 {
				x = intCode[intCode[i + 1]]
			} else {
				x = intCode[i + 1]
			}
			if secondParamMode == 0 {
				y = intCode[intCode[i + 2]]
			} else {
				y = intCode[i + 2]
			}
			resPos := intCode[i + 3]
			intCode[resPos] = x * y
			increment = 4
			break
		case 3: // Get value and write to intCode[pos]
			pos := intCode[i + 1]
			if len(input) == 0 {
				panic("Something went wrong...")
			}
			intCode[pos] = input[0]
			input = removeAt(input, 0)
			increment = 2
			break
		case 4: // Output value from intCode[pos]
			var pos int
			if firstParamMode == 0 {
				pos = intCode[i + 1]
			} else {
				pos = i + 1
			}
			output = intCode[pos]
			input = append(input, output)
			increment = 2
			break
		case 5:
			var x int
			var y int
			if firstParamMode == 0 {
				x = intCode[intCode[i + 1]]
			} else {
				x = intCode[i + 1]
			}
			if secondParamMode == 0 {
				y = intCode[intCode[i + 2]]
			} else {
				y = intCode[i + 2]
			}
			if x != 0 {
				//fmt.Printf("true, i = %d\n", y)
				i = y
				increment = 0
			} else {
				increment = 3
			}
			break
		case 6:
			var x int
			var y int
			if firstParamMode == 0 {
				x = intCode[intCode[i + 1]]
			} else {
				x = intCode[i + 1]
			}
			if secondParamMode == 0 {
				y = intCode[intCode[i + 2]]
			} else {
				y = intCode[i + 2]
			}
			if x == 0 {
				//fmt.Printf("true, i = %d\n", y)
				i = y
				increment = 0
			} else {
				increment = 3
			}
			break
		case 7:
			var x int
			var y int
			if firstParamMode == 0 {
				x = intCode[intCode[i + 1]]
			} else {
				x = intCode[i + 1]
			}
			if secondParamMode == 0 {
				y = intCode[intCode[i + 2]]
			} else {
				y = intCode[i + 2]
			}
			if x < y {
				//fmt.Printf("true, intCode[%d] = 1\n", intCode[i + 3])
				intCode[intCode[i + 3]] = 1
			} else {
				//fmt.Printf("false, intCode[%d] = 0\n", intCode[i + 3])
				intCode[intCode[i + 3]] = 0
			}
			increment = 4
			break
		case 8:
			var x int
			var y int
			if firstParamMode == 0 {
				x = intCode[intCode[i + 1]]
			} else {
				x = intCode[i + 1]
			}
			if secondParamMode == 0 {
				y = intCode[intCode[i + 2]]
			} else {
				y = intCode[i + 2]
			}
			if x == y {
				//fmt.Printf("true, intCode[%d] = 1\n", intCode[i + 3])
				intCode[intCode[i + 3]] = 1
			} else {
				//fmt.Printf("false, intCode[%d] = 0\n", intCode[i + 3])
				intCode[intCode[i + 3]] = 0
			}
			increment = 4
			break
		case 99:
			done = true
			break
		default:
			panic("Unknown opcode: " + fmt.Sprintf(strconv.Itoa(opcode)))
			break
		}
	}
	return output
}
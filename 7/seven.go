package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"sync"
)

type Amp struct {
	prev    *Amp
	input   int
	intCode map[int]int
	output  int
	next    *Amp
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

	//combos := getAllCombos(ELEMENTS_PART1, make([]int, 0))
	//fmt.Println(getThrust(combos, input))

	combos := getAllCombos(ELEMENTS_PART2, make([]int, 0))
	fmt.Println(getThrustWithFeedback(combos, input))
}

//func getThrust(combos [][]int, input []byte) int {
//	highest := 0
//	for _, combo := range combos {
//		aOut := evaluateIntCode([]int{combo[0], START}, input)
//		bOut := evaluateIntCode([]int{combo[1], aOut}, input)
//		cOut := evaluateIntCode([]int{combo[2], bOut}, input)
//		dOut := evaluateIntCode([]int{combo[3], cOut}, input)
//		eOut := evaluateIntCode([]int{combo[4], dOut}, input)
//
//		if eOut > highest {
//			highest = eOut
//		}
//	}
//	return highest
//}

func getThrustWithFeedback(combos [][]int, input []byte) int {
	highest := 0
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	d := make(chan int)
	e := make(chan int)
	var wg sync.WaitGroup
	for i, combo := range combos {
		go func() {
			a <- combo[0]
			if i == 0 {
				a <- 0
			}
			b <- combo[1]
			c <- combo[2]
			d <- combo[3]
			e <- combo[4]
		}()

		wg.Add(1)
		go evaluateIntCode(input, a, b, wg, "A")
		wg.Add(1)
		go evaluateIntCode(input, b, c, wg, "B")
		wg.Add(1)
		go evaluateIntCode(input, c, d, wg, "C")
		wg.Add(1)
		go evaluateIntCode(input, d, e, wg, "D")
		wg.Add(1)
		go evaluateIntCode(input, e, a, wg, "E")

		wg.Wait()

		if eOut := <- e; eOut > highest {
			highest = eOut
		}
		fmt.Println("Done")
	}
	return highest
}

func removeAt(elements []int, index int) []int {
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

func evaluateIntCode(rawIntCode []byte, inChan chan int, outChan chan int, wg sync.WaitGroup, label string) int {
	defer wg.Done()
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

		//if label == "E" {
		//	fmt.Printf("Opcode %s - %d\n", label, opcode)
		//	fmt.Printf("Pos %s - %d\n", label, i)
		//	fmt.Printf("Intcode %s - %v\n", label, intCode)
		//}
		switch opcode {
		case 1:
			var x int
			var y int
			if firstParamMode == 0 {
				x = intCode[intCode[i+1]]
			} else {
				x = intCode[i+1]
			}
			if secondParamMode == 0 {
				y = intCode[intCode[i+2]]
			} else {
				y = intCode[i+2]
			}
			resPos := intCode[i+3]
			intCode[resPos] = x + y
			increment = 4
			break
		case 2:
			var x int
			var y int
			if firstParamMode == 0 {
				x = intCode[intCode[i+1]]
			} else {
				x = intCode[i+1]
			}
			if secondParamMode == 0 {
				y = intCode[intCode[i+2]]
			} else {
				y = intCode[i+2]
			}
			resPos := intCode[i+3]
			intCode[resPos] = x * y
			increment = 4
			break
		case 3: // Get value and write to intCode[pos]
			in := <-inChan
			fmt.Printf("%s - Receive: %d\n", label, in)
			pos := intCode[intCode[i+1]]
			//if label == "E" {
			//	fmt.Printf("%s - Receive: %d; Writing to intcode[%d]\n", label, in, pos)
			//}
			intCode[pos] = in
			increment = 2
			break
		case 4: // Output value from intCode[pos]
			var pos int
			if firstParamMode == 0 {
				pos = intCode[intCode[i+1]]
			} else {
				pos = intCode[i + 1]
			}
			output = intCode[pos]
			fmt.Printf("%s - Send: %d\n", label, output)
			outChan <- output
			increment = 2
			break
		case 5:
			var x int
			var y int
			if firstParamMode == 0 {
				x = intCode[intCode[i+1]]
			} else {
				x = intCode[i+1]
			}
			if secondParamMode == 0 {
				y = intCode[intCode[i+2]]
			} else {
				y = intCode[i+2]
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
				x = intCode[intCode[i+1]]
			} else {
				x = intCode[i+1]
			}
			if secondParamMode == 0 {
				y = intCode[intCode[i+2]]
			} else {
				y = intCode[i+2]
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
				x = intCode[intCode[i+1]]
			} else {
				x = intCode[i+1]
			}
			if secondParamMode == 0 {
				y = intCode[intCode[i+2]]
			} else {
				y = intCode[i+2]
			}
			if x < y {
				//fmt.Printf("true, intCode[%d] = 1\n", intCode[i + 3])
				intCode[intCode[i+3]] = 1
			} else {
				//fmt.Printf("false, intCode[%d] = 0\n", intCode[i + 3])
				intCode[intCode[i+3]] = 0
			}
			increment = 4
			break
		case 8:
			var x int
			var y int
			if firstParamMode == 0 {
				x = intCode[intCode[i+1]]
			} else {
				x = intCode[i+1]
			}
			if secondParamMode == 0 {
				y = intCode[intCode[i+2]]
			} else {
				y = intCode[i+2]
			}
			if x == y {
				//fmt.Printf("true, intCode[%d] = 1\n", intCode[i + 3])
				intCode[intCode[i+3]] = 1
			} else {
				//fmt.Printf("false, intCode[%d] = 0\n", intCode[i + 3])
				intCode[intCode[i+3]] = 0
			}
			increment = 4
			break
		case 99:
			done = true
			fmt.Printf("%s Done routine\n", label)
			break
		default:
			panic(fmt.Sprintf("%s - Unknown opcode: %s", label, strconv.Itoa(opcode)))
			break
		}
	}
	return output
}

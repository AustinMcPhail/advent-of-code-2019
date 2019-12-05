package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal("Something went wrong while reading the input file.")
		panic(err)
	}
	intStrings := strings.Split(string(input), ",")
	intCode := make(map[int]int)
	for i, s := range intStrings {
		current, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Something went wrong whilst converting a string {%v} to int", i))
		}
		intCode[i] = current
	}

	evaluateIntCode(intCode)
}

func evaluateIntCode(intCode map[int]int) {
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
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Enter System ID: ")
			scanner.Scan()
			input, _ := strconv.Atoi(scanner.Text())
			intCode[pos] = input
			increment = 2
			break
		case 4: // Output value from intCode[pos]
			var pos int
			if firstParamMode == 0 {
				pos = intCode[i + 1]
			} else {
				pos = i + 1
			}
			fmt.Println(intCode[pos])
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
			fmt.Println("Done!")
			done = true
			break
		default:
			panic("Unknown opcode: " + fmt.Sprintf(strconv.Itoa(opcode)))
			break
		}
		//fmt.Println(intCode)
	}
}

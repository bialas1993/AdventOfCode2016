package main

import (
	"fmt"
	"strings"
	"strconv"
)

const (
	INPUT_FILE = "input"
)

type ActionExecutor func(args []string)

var actions = map[string]ActionExecutor {
	"rect": rect,
	"rotate": rotate,
}

func rect(args []string) {

	axis := args[1][0:1]
	col, _ := strconv.Atoi(args[1:2][2])
	beforePixel, _ := strconv.Atoi(args[3:4][0])

	switch args[0] {
		case "row":
			rotateRow(axis, col, beforePixel)
			break
		case "column":
			rotateCol(axis, col, beforePixel)
			break
	}
}

func rotate(args []string) {
	print("rotate: ")
	fmt.Println(args)
}


func rotateRow(axis string, col int, beforePixel int) {
	fmt.Print("rotateRow: ")
	fmt.Println(axis, col, beforePixel)
}

func rotateCol(axis string, col int, beforePixel int) {
	fmt.Print("rotateCol: ")
	fmt.Println(axis, col, beforePixel)
}


func actionDispatcher(command string) {
	parseCommand := strings.Split(command, " ")
	actions[parseCommand[0]](parseCommand[1:])
}

func solve(input string) int {

	for _, line := range strings.Split(input, "\n") {
		actionDispatcher(line)
	}

	return 0
}

func main() {

	var testInput string =
`###....
###....
.......`

	testInput = "rect 3x2\nrotate column x=1 by 1\nrotate row y=0 by 4\nrotate column x=1 by 1"

	/*testExpected := "easter"
	if testResult := solve(testInput, findMostCommonChar); testResult != testExpected {
		panic(fmt.Sprintf("Test result is not correct: %s != %s ", testResult, testExpected))
	}*/



	fmt.Printf("[Task1] Displayed pixels: %d", solve(testInput))
}
package main

import (
	"fmt"
	"strings"
	"strconv"
	"runtime"
	"path"
	"io/ioutil"
)

const (
	INPUT_FILE = "input"
	WIDTH = 50
	HEIGHT = 6
)

type ActionExecutor func(args []string)

var screen = make([][]string, HEIGHT)

var actions = map[string]ActionExecutor {
	"rect": rect,
	"rotate": rotate,
}

func initScreen () {
	for y := 0; y < HEIGHT; y++ {
		var vect = make([]string, WIDTH)
		screen[y] = append(vect)
		for x := 0; x < WIDTH; x++ {
			screen[y][x] = "."
		}
	}
}


func calcPixels() int {
	pixels := 0
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if screen[y][x] == "#" {
				pixels++
			}
		}
	}
	return pixels
}

func rect(args []string) {
	print("rect: ")
	x, _ := strconv.Atoi(args[0][0:1])
	y, _ := strconv.Atoi(args[0][2:3])

	fmt.Println(x,y)

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			screen[i][j] = "#"
		}
	}
}

func rotate(args []string) {
	print("rotate: ")
	axis := args[1][0:1]

	col, _ := strconv.Atoi(args[1][2:])
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


func rotateRow(axis string, col int, beforePixel int) {
	fmt.Print("rotateRow: ")
	fmt.Println(axis, col, beforePixel)

	if beforePixel >= WIDTH {
		beforePixel = beforePixel % WIDTH
	}

	if beforePixel > 0 {
		move := make([]string, WIDTH - beforePixel)

		for x := 0; x < WIDTH; x++ {
			if x < WIDTH - beforePixel {
				move[x] = screen[col][x]

			} else {
				screen[col][WIDTH - x - 1] = screen[col][x]
			}
		}

		for y := range move {
			screen[col][y + beforePixel] = move[y]
		}
	}

}

func rotateCol(axis string, col int, beforePixel int) {
	fmt.Print("rotateCol: ")
	fmt.Println(axis, col, beforePixel)

	if beforePixel >= HEIGHT {
		beforePixel = beforePixel % HEIGHT
	}

	if beforePixel > 0 {
		move := make([]string, HEIGHT - beforePixel)
		for y := range screen {
			if y < HEIGHT - beforePixel {
				move[y] = screen[y][col]
			} else {
				screen[HEIGHT - y][col] = screen[y][col]
			}
		}

		for y := range move {
			screen[y + beforePixel][col] = move[y]
		}
	}
}


func actionDispatcher(command string) {
	parseCommand := strings.Split(command, " ")
	actions[parseCommand[0]](parseCommand[1:])
}

func solve(input string) int {

	for _, line := range strings.Split(input, "\n") {
		if len(line) > 0 {
			actionDispatcher(line)
		}
	}

	return 0
}

func main() {
	initScreen()

	/*testInput := "rect 3x2\nrotate column x=1 by 1\nrotate row y=0 by 4\nrotate column x=1 by 1"
	testInput = "rect 3x2\nrotate column x=1 by 12\nrotate row x=0 by 2"
	fmt.Printf("[Task1] Displayed pixels: %d", solve(testInput))*/
	/*testExpected := "easter"
	if testResult := solve(testInput, findMostCommonChar); testResult != testExpected {
		panic(fmt.Sprintf("Test result is not correct: %s != %s ", testResult, testExpected))
	}*/


	_, fileName, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(fileName), INPUT_FILE)
	inpBuff, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}


	solve(string(inpBuff))

	for _, line := range screen {
		fmt.Println(line)
	}



	fmt.Printf("[Task1] Displayed pixels: %d", calcPixels())
}
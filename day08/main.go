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
	split := strings.Split(args[0], "x")

	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			screen[i][j] = "#"
		}
	}
}

func rotate(args []string) {
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




func rotateCol(axis string, col int, beforePixel int) {
	if beforePixel > HEIGHT {
		beforePixel =  beforePixel % HEIGHT
	}

	if beforePixel > 0 {
		move := make([]string, HEIGHT - beforePixel)

		for y := 0; y < HEIGHT; y++ {
			if y < HEIGHT - beforePixel {
				move[y] = screen[y][col]
			} else {
				screen[y - HEIGHT + beforePixel][col] = screen[y][col]
			}
		}

		for y := 0; y < len(move); y++ {
			screen[y + beforePixel][col] = move[y]
		}
	}
}

func rotateRow(axis string, col int, beforePixel int) {
	if beforePixel > WIDTH {
		beforePixel =  beforePixel % WIDTH
	}

	if beforePixel > 0 {
		move := make([]string,  WIDTH - beforePixel)
		for y := 0; y < WIDTH; y++ {
			if y < WIDTH - beforePixel {
				move[y] = screen[col][y]
			} else {
				screen[col][y - WIDTH + beforePixel] = screen[col][y]
			}
		}

		for y := 0; y < len(move); y++ {
			screen[col][y + beforePixel] = move[y]
		}
	}
}

func main() {
	initScreen()

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
package main

import (
	"os/exec"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"time"
	"math/rand"
)

const (
	OUTPUT_SCREEN_BUFF = 2
)

var screenBuff = []map[int]string{}

type calc func(x int, y int) string
type drawElement func(x *int, y *int)

var drawFunctions = map[int]drawElement {
	0: drawElips,
	1: drawElips,
	2: drawElips,
	3: drawElips,
	4: drawElips,
	5: drawCross,
	6: drawCross,
	7: drawCross,
	8: drawCross,
}

func clearScreenBuff(height int, width int) {
	for i := 0; i < height; i++ {
		screenBuff = append(screenBuff, make(map[int]string, width))
		for j := 0; j < width; j++ {
			screenBuff[i][j] = " "
		}
	}
}

func drawBuff(height int, width int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Print(screenBuff[y][x])
		}
		println("")
	}
}



func setBoardToBuff() {
	for y := 1; y < 24; y++ {
		for x := 1; x < 48; x++ {
			if  y % 8  == 0  {
				screenBuff[y][x] = "#"
			}

			if x % 16 == 0 {
				screenBuff[y][x] = "#"
			}

		}
	}
}



func main() {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	outBuff, err := cmd.Output()

	out := strings.Fields(string(outBuff))
	height, _ := strconv.Atoi(out[0])
	width, _ := strconv.Atoi(out[1])
	width -= OUTPUT_SCREEN_BUFF
	height -= OUTPUT_SCREEN_BUFF + 2

	for  {
		var drawFunctionUses = [9]int {}

		fmt.Print("\033[1;1H\033[2J")

		var y, x, tempY int = 0, 0, 0

		clearScreenBuff(height, width)
		rand.Seed(time.Now().Unix())
		if rand.Intn(2) + 1 == 1 {
			drawFunctions[0] = drawCross
		}

		var drawes int = 9
		var draws int = 0


		for i := 0; i < drawes; i++ {
			if draws == drawes {
				break
			}

			rand.Seed(time.Now().UnixNano())
			idxFuncDraw := rand.Intn(drawes) + 0

			if drawFunctionUses[idxFuncDraw] == 1 {
				i--
				continue
			}

			//rand amount cross or circles | mod from itterations is better option
			if i % 3 == 0 && i > 0{
				x = 0
				tempY = y
			}

			y = tempY
			drawFunctions[idxFuncDraw](&x, &y)
			drawFunctionUses[idxFuncDraw] = 1
			draws++
		}

		setBoardToBuff()
		drawBuff(height, width)
		time.Sleep(500 * time.Millisecond)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func drawCross(startDrawX *int, startDrawY *int) {
	size := 16
	var tempX int
	for j := 0; j < size; j = j + 2 {
		tempX = *startDrawX
		for i := 0; i < size; i++ {
			char := " "
			if i + 1 == j || size - i == j {
				char = "#"
			}

			screenBuff[*startDrawY][tempX] = char
			tempX++
		}
		*startDrawY++
	}
	*startDrawX =  tempX
}

func drawElips(startDrawX *int, startDrawY *int) {
	r := 4
	var tempX int
	for x := -r ; x < r; x++ {
		tempX = *startDrawX
		for y := -r * 2; y < r * 2; y++ {
			elipsRes := ( float64(x * x) / .5 ) + (float64(y * y) / 2)

			char := " "
			if elipsRes >= 15 && elipsRes <= 20 {
				char = "*"
			}

			screenBuff[*startDrawY][tempX] = char
			tempX++
		}
		*startDrawY++
	}
	*startDrawX = tempX
}

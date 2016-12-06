package main

import (
	"strings"
	"math"
	"strconv"
	"reflect"
	"fmt"
)

type pos struct {
	x,y int
}

type dir struct {
	L,R int
}

func getField(v *dir, field string) int {
    r := reflect.ValueOf(v)
    f := reflect.Indirect(r).FieldByName(field)
    return int(f.Int())
}

var inp = "R1, L3, R5, R5, R5, L4, R5, R1, R2, L1, L1, R5, R1, L3, L5, L2, R4, L1, R4, R5, L3, R5, L1, R3, L5, R1, L2, R1, L5, L1, R1, R4, R1, L1, L3, R3, R5, L3, R4, L4, R5, L5, L1, L2, R4, R3, R3, L185, R3, R4, L5, L4, R48, R1, R2, L1, R1, L4, L4, R77, R5, L2, R192, R2, R5, L4, L5, L3, R2, L4, R1, L5, R5, R4, R1, R2, L3, R4, R4, L2, L4, L3, R5, R4, L2, L1, L3, R1, R5, R5, R2, L5, L2, L3, L4, R2, R1, L4, L1, R1, R5, R3, R3, R4, L1, L4, R1, L2, R3, L3, L2, L1, L2, L2, L1, L2, R3, R1, L4, R1, L1, L4, R1, L2, L5, R3, L5, L2, L2, L3, R1, L4, R1, R1, R2, L1, L4, L4, R2, R2, R2, R2, R5, R1, L1, L4, L5, R2, R4, L3, L5, R2, R3, L4, L1, R2, R3, R5, L2, L3, R3, R1, R3"
var visits = make([]pos, 0)
var startPos pos = pos{x: 0, y: 0}
var startDirection int = 0

var directionMapping = map[int]pos {
	0: {x: 0, y: 1},
	1: {x: 1, y: 0},
	2: {x: -1, y: 0},
	3: {x: 0, y: -1},
}

var turnsMapping = map[int]dir {
	0: {L: 2, R: 1},
	1: {L: 0, R: 3},
	2: {L: 3, R: 0},
	3: {L: 1, R: 2},
}

func main() {
	moves := strings.Split(inp, ", ")

	visits = append(visits, startPos)
	var currDirection int = startDirection
    	var currPosition pos = startPos

	for _, move := range moves  {
		turn := move[:1]
		steps, _ := strconv.Atoi(move[1:])

		d := turnsMapping[currDirection]
		currDirection = getField(&d, turn)
		nextStep := directionMapping[currDirection]
		for i := 1 ; i <= int(steps); i++ {
			currPosition = pos{x: currPosition.x + nextStep.x, y: currPosition.y + nextStep.y}

			if inSlice(currPosition) {
				fmt.Printf("Answer : %d", taxicab(startPos, currPosition))
				return
			}
			visits = append(visits, currPosition)
		}
	}
}

func taxicab(startPos pos, endPos pos) int {
	return int(math.Abs(float64(endPos.x - startPos.x)) + math.Abs(float64(endPos.y - startPos.y)))
}

func inSlice(p pos) bool {
	for _, z := range visits {
		if p == z {
			return true
		}
	}
	return false
}

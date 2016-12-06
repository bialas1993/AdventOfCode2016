package main

import (
	"fmt"
	"strings"
	"math"
	"reflect"
)

type pos struct {
	x,y int
}

type dir struct {
	L,R int
}

var inp = "L4, L3, R1, L4"
var visits = make([]pos, 0)
var startPos pos = pos{x: 0, y: 0}
var startDirection int = 0

var directionMap = map[int]pos {
	0: pos{x: 0, y: 1},
	1: pos{x: 1, y: 0},
	2: pos{x: -1, y: 0},
	3: pos{x: 0, y: -1},
}

func main() {

	visits = append(visits, startPos)

	moves := strings.Split(inp, ", ")
	fmt.Println(moves)

	//var currDirection int = startDirection
    	//var currposition pos = startPos

	for _, move := range moves  {
		direction := move[0]
		steps := int(move[1])

		//curr_dir = make_turn(curr_dir, turn)
		//next_step = dir_mapping[curr_dir]

		fmt.Println(reflect.TypeOf(steps).Kind())
		return

		fmt.Println("steps " + string(int(steps)))
		for i := 1 ; i <= int(steps); i++ {
			fmt.Println(i)
		}

		fmt.Print("\nDir: " + string(direction))
	}


	//for z := 0; z < 10; z++ {
	//	visits = append(visits, pos{x:z, y:z})
	//}
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

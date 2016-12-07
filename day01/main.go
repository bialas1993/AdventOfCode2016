package main

import (
	"strings"
	"math"
	"strconv"
	"reflect"
	"fmt"
	"io/ioutil"
	"runtime"
	"path"
)

const (
	INPUT_FILE = "input"
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
	_, fileName, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(fileName), INPUT_FILE)
	inpBuff, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	moves := strings.Split(string(inpBuff), ", ")

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

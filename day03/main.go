package main

import (
	"strings"
	"fmt"
	"runtime"
	"path"
	"io/ioutil"

	"strconv"
	"sort"
)

const (
	INPUT_FILE = "input"
)

var possibleTriangles int = 0

func main() {
	_, fileName, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(fileName), INPUT_FILE)
	inpBuff, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	trianglesSides := strings.Split(string(inpBuff), "\n")

	for _, input := range trianglesSides {
		sidesAsString := strings.Fields(input)
		sides := make([]int, 3)
		for i := 0; i < len(sidesAsString); i++ {
			side, _ := strconv.Atoi(sidesAsString[i])
			sides[i] = side
		}
		sort.Ints(sides)

		if (sides[0] + sides[1] > sides[2]) {
			possibleTriangles++
		}
	}

	fmt.Printf("Able to create a triangle: %d", possibleTriangles)
}
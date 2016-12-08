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

var bytes = [8]byte{30, 41, 20, 32, 30, 20, 32, 30}
var possibleTriangles int = 0


func convert( b []byte ) string {
    s := make([]string,len(b))
    for i := range b {
        s[i] = strconv.Itoa(int(b[i]))
    }
    return strings.Join(s,",")
}

func main() {
	_, fileName, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(fileName), INPUT_FILE)
	inpBuff, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	trianglesSides := strings.Split(string(inpBuff), convert(bytes[:]))

	for _, input := range trianglesSides {
		sidesAsString := strings.Split(strings.Trim(input, " "), " ")
		println(sidesAsString)

		sides := make([]int, 3)
		for i := 0; i < len(sidesAsString); i++ {
			side, _ := strconv.Atoi(sidesAsString[i])
			sides[i] = side
		}
		sort.Ints(sides)

		if (sides[0] + sides[1] >= sides[2]) {
			possibleTriangles++
		}
	}


	fmt.Printf("Able to create a triangle: %d", possibleTriangles)

}
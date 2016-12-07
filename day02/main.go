package main

import (
	"strings"
	"reflect"
	"path"
	"io/ioutil"
	"runtime"
	"bytes"
)

const (
	INPUT_FILE = "input"
)

type move struct {
	U,L,R,D int
}

var startPos = 5

var keyboardMapping = map[int] move {
	1: {U: 1, L: 1, R: 2, D: 4},
	2: {U: 2, L: 1, R: 3, D: 5},
	3: {U: 3, L: 2, R: 3, D: 6},
	4: {U: 1, L: 4, R: 5, D: 7},
	5: {U: 2, L: 4, R: 6, D: 8},
	6: {U: 3, L: 5, R: 6, D: 9},
	7: {U: 4, L: 7, R: 8, D: 7},
	8: {U: 5, L: 7, R: 9, D: 8},
	9: {U: 6, L: 8, R: 9, D: 9},
}

func main() {
	_, fileName, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(fileName), INPUT_FILE)
	inpBuff, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inpBuff), "\n")

	print("Toilet code: ")

	var code bytes.Buffer
	keyboardPos := startPos
	for _, line := range lines {

		for i := 0 ; i < len(line); i++ {
			lineMove := line[i:i+1]
			keyboardMove := keyboardMapping[keyboardPos]
			keyboardPos = getField(&keyboardMove, lineMove)
		}
		print(keyboardPos)
		code.WriteString(string(keyboardPos))
		//keyboardPos = 5
	}
}


func getField(v *move, field string) int {
    r := reflect.ValueOf(v)
    f := reflect.Indirect(r).FieldByName(field)
    return int(f.Int())
}
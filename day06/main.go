package main

import (
	"fmt"
	"strings"
	"runtime"
	"path"
	"io/ioutil"
)

const (
	INPUT_FILE = "input"
)

type symLambda func(string) string

func findMostCommonChar(line string) string {
	var character string
	var amount int = 0

	for _, char := range strings.Split(line, "") {
		c := strings.Count(line, char)

		if (c > amount) {
			amount = c
			character = char
		}
	}

	return character
}

func findRarestChar(line string) string {
	var character string
	var amount int = -1

	for _, char := range strings.Split(line, "") {
		c := strings.Count(line, char)
		if (amount == -1 || c < amount ) {
			amount = c
			character = char
		}
	}

	return character
}

func solve(input string, function symLambda) string {
	lines := strings.Split(input, "\n")
	parseLine := make(map[int]string)


	for _, line := range lines {
		for index, character := range strings.Split(line, "") {
			parseLine[index] += character
		}
	}

	var result = make([]string, len(parseLine))

	for index, line := range parseLine {
		result[index] = function(line)
	}

	return strings.Join(result, "")
}

func main() {

	var testInput string =
`eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`
	testExpected := "easter"
	if testResult := solve(testInput, findMostCommonChar); testResult != testExpected {
		panic(fmt.Sprintf("Test result is not correct: %s != %s ", testResult, testExpected))
	}

	testExpected = "advent"
	if testResult := solve(testInput, findRarestChar); testResult != testExpected {
		panic(fmt.Sprintf("Test result is not correct: %s != %s ", testResult, testExpected))
	}

	_, fileName, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(fileName), INPUT_FILE)
	inpBuff, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}


	fmt.Printf("[Task1] Repetition code decode: %s", solve(string(inpBuff), findMostCommonChar))
	println("")
	fmt.Printf("[Task2] Repetition code decode: %s", solve(string(inpBuff), findRarestChar))
}
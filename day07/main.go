package main

import (
	"fmt"
	"regexp"
	"strings"
	"runtime"
	"path"
	"io/ioutil"
)

const (
	INPUT_FILE = "input"
)

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func checkABBA(chars string) bool {
	if len(chars) >= 4 {
		for i := 0; i < len(chars) - 4; i++  {
			if strings.Count(chars, chars[0:1]) < len(chars) {
				if chars[i:i+2] == reverse(chars[i+2:i+4]) {
					return true
				}
			}
		}
	}
	return false
}

func solve(input string) int {
	var counter int = 0
	var re = regexp.MustCompile(`(\w+)`)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		matchAll := re.FindAllStringSubmatch(line, -1)

		var bracketsAbba = false
		for i := 1; i < len(matchAll); i += 2  {
			if checkABBA(matchAll[i][0]) {
				bracketsAbba = true
				break
			}
		}

		if !bracketsAbba {
			for i := 0; i < len(matchAll); i += 2  {
				if checkABBA(matchAll[i][0]) {
					counter++
					break
				}
			}
		}

	}
	return counter
}

func main() {
	testInput :=
`abba[mnop]qrst
abcd[bddb]xyyx
aaaa[qwer]tyui
ioxxoj[asdfgh]zxcvbn`

	 solve(testInput);
	//testExpected := 2
	//
	//testResult := 0
	//if testResult = solve(testInput); testResult != testExpected {
	//	panic(fmt.Sprintf("Test result is not correct: %d != %d ", testResult, testExpected))
	//}


	_, fileName, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(fileName), INPUT_FILE)
	inpBuff, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("IPs amount: %d", solve(string(inpBuff)))
}
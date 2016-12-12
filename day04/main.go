package main

import (
	"fmt"
	"regexp"
	"strings"
	"runtime"
	"path"
	"io/ioutil"
	"strconv"
	"sort"
)

const (
	INPUT_FILE = "input"
)

func removeDuplicatesUnordered(elements []string) []string {
    encountered := map[string]bool{}

    for v:= range elements {
	encountered[elements[v]] = true
    }

    result := []string{}
    for key := range encountered {
	result = append(result, key)
    }
    return result
}


func solve(input string) int {
var re = regexp.MustCompile(`([a-z\-]+)-(\d+)\[([a-z]+)\]`)
	var idsSum int = 0

	codeInput := strings.Split(input, "\n")
	for _, line := range codeInput {
		matchAll := re.FindAllStringSubmatch(line, -1)
		if len(matchAll) > 0 {
			match := matchAll[0]

			code := strings.Replace(match[1], "-", "", -1)
			id, _ := strconv.Atoi(match[2])
			checksum := strings.Split(match[3], "")
			checksumMatch := true

			cuteCodeMap := removeDuplicatesUnordered(strings.Split(code, ""))
			var chMap = make(map[string]int, len(cuteCodeMap))
			for _, character := range cuteCodeMap {
				chMap[character] = strings.Count(code, character)
			}

			if (!checksumMatch) {
				continue
			}

			n := map[int][]string{}
			var a []int
			for k, v := range chMap {
				n[v] = append(n[v], k)
			}

			for k := range n {
				a = append(a, k)
			}
			sort.Sort(sort.Reverse(sort.IntSlice(a)))

			outCode := ""
			for _, idx := range a {
				sort.Strings(n[idx])
				for _, char := range n[idx] {
					outCode += char
				}
			}

			if outCode[:5] == strings.Join(checksum, "") {
				idsSum += id
			}
		}
	}

	return idsSum
}

func main() {
	_, fileName, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(fileName), INPUT_FILE)
	inpBuff, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}


	testInput :=
`aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]`
	testExpected := 1514


	if testResult := solve(testInput); testResult != testExpected {
		panic(fmt.Sprintf("Test result is not correct: %s != %s ", testResult, testExpected))
	}


	fmt.Printf("Sum ids: %d", solve(string(inpBuff)))
}
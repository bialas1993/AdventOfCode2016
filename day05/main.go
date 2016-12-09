package main

import (
	"fmt"
	"crypto/md5"
	"strconv"
	"encoding/hex"
	"strings"
)

func solve_first(data string) {
	var i int = 0
	var passw string = ""

	for len(passw) < 8 {
		hasher := md5.New()
		hasher.Write([]byte((data + strconv.Itoa(i))))
		hash := hex.EncodeToString(hasher.Sum(nil)[:])
		i++

		if strings.HasPrefix(hash, "00000") {
			passw += hash[5:6]
		}
	}

	fmt.Printf("Result for first part: %s\n", passw)
}


func solve_second(data string) {
	var i, findCounter int = 0, 0
	solved := make([]string, 8)

	for findCounter < 8 {
		hasher := md5.New()
		hasher.Write([]byte((data + strconv.Itoa(i))))
		hash := hex.EncodeToString(hasher.Sum(nil)[:])
		i++

		if strings.HasPrefix(hash, "00000") {

			index, err := strconv.Atoi(hash[5:6])
			if err == nil && index < 8 {
				if solved[index] == "" {
					character :=  hash[6:7]
					solved[index] = character
					findCounter++
				}
			}
		}
	}

	fmt.Printf("Result for second part: %s\n", (strings.Join(solved, "")))
}

func main() {
	input := "ojvtpuvg"

	solve_first(input)
	solve_second(input)
}
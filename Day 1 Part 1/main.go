package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
func main() {
	var result int64 = 0
	data, err := (os.ReadFile("input.txt"))
	check(err)
	dataSlice := strings.Split(string(data), "\n")
	r := regexp.MustCompile("[0-9]")

	for i := 0; i < len(dataSlice); i++ {
		matches := r.FindAllString(dataSlice[i], -1)
		if matches == nil {
			continue
		}

		if len(matches) == 1 {
			extractedNumbers := []string{matches[0], matches[0]}
			joinedNumbers := strings.Join(extractedNumbers, "")
			newNumber, err := strconv.ParseInt(joinedNumbers, 10, 64)
			check(err)
			result += newNumber
		} else if len(matches) == 2 {
			joinedNumbers := strings.Join(matches, "")
			newNumber, err := strconv.ParseInt(joinedNumbers, 10, 64)
			check(err)
			result += newNumber
		} else {
			extractedNumbers := []string{matches[0], matches[len(matches)-1]}
			joinedNumber := strings.Join(extractedNumbers, "")
			newNumber, err := strconv.ParseInt(joinedNumber, 10, 64)
			check(err)
			result += newNumber
		}
	}

	fmt.Printf("Result: %d\n", result)
}

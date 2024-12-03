package day3

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func PartTwo(file string) {
	b, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	reForInstructions := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)

	instructions := reForInstructions.FindAllString(string(b), -1)

	reForNumbers := regexp.MustCompile(`\d+`)

	total := 0

	skipUntilDo := false

	for _, val := range instructions {

		if val == "do()" {
			skipUntilDo = false
			continue
		}

		if val == "don't()" || skipUntilDo {
			skipUntilDo = true
			continue
		}

		var intList []int = make([]int, 0)

		Map(reForNumbers.FindAllString(val, -1), &intList, func(element string) int {
			retval, _ := strconv.ParseInt(element, 10, 0)
			return int(retval)
		})

		subTotal := 1
		for _, val := range intList {

			subTotal *= val
		}

		total += subTotal
	}

	fmt.Printf("Total from mul instructions: %v", total)
}

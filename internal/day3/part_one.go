package day3

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Map[Source any, Dest any](src []Source, dst *[]Dest, mapFunc func(Source) Dest) {
	if cap(*dst) == 0 {
		(*dst) = make([]Dest, len(src))
	}

	for index, value := range src {
		(*dst)[index] = mapFunc(value)
	}
}

func PartOne(file string) {
	b, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	reForInstructions := regexp.MustCompile(`mul\(\d+,\d+\)`)

	instructions := reForInstructions.FindAllString(string(b), -1)

	reForNumbers := regexp.MustCompile(`\d+`)

	total := 0

	for _, val := range instructions {
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

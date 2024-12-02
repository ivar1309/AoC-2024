package day2

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type intList []int

func safe(report intList) bool {
	increasing := true
	decreasing := true
	withinRange := true

	var adjacentLengths intList = make(intList, 0)

	for index, value := range report {

		if index == 0 {
			continue
		}

		adjacentLengths = append(
			adjacentLengths,
			int(math.Abs(float64(value)-float64(report[index-1]))),
		)

		if value < report[index-1] {
			increasing = false
		}

		if value > report[index-1] {
			decreasing = false
		}
	}

	for _, value := range adjacentLengths {
		if value < 1 || value > 3 {
			withinRange = false
			break
		}
	}

	return withinRange && (increasing || decreasing)
}

func PartOne(file string) {
	var report intList = make(intList, 0)
	countSafe := 0

	b, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		report = make(intList, 0)

		fmt.Printf("line: %v\n", line)

		for _, number := range strings.Split(line, " ") {
			var a int
			fmt.Sscanf(number, "%d", &a)
			report = append(report, a)

		}

		if safe(report) {
			countSafe++
		}

		fmt.Printf("report: %v\n", report)
	}

	fmt.Printf("Total safe reports: %v", countSafe)
}

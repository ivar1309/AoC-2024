package day2

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

type report struct {
	list intList
	safe bool
}

func (r report) getRangeList() intList {
	var ranges intList = make(intList, 0)
	for idx, val := range r.list {
		if idx == 0 {
			continue
		}

		ranges = append(ranges, r.list[idx]-val)
	}

	return ranges
}

func outOfRange(value int) bool {
	return value < 1 || value > 3
}

func withinRange(report intList, level int) int {
	withinRangeFail := 0

	if level == 2 {
		return 0
	}

	for index, value := range report {

		if index == 0 {
			continue
		}

		if outOfRange(int(math.Abs(float64(value) - float64(report[index-1])))) {
			badLevelRemoved := slices.Concat(report[:index], report[index+1:])
			withinRangeFail += 1 + withinRange(badLevelRemoved, level+1)
		}
	}

	return withinRangeFail
}

func isIncreasing(report intList) float64 {
	cumSum := 0.0
	for index, value := range report {
		if index == 0 {
			continue
		}

		cumSum += float64(report[index-1]) / float64(value)
	}

	return cumSum / (float64(len(report)) - 1)
}

func increasing(report intList, level int) int {
	increasingFail := 0

	if level == 2 {
		return 0
	}

	for index, value := range report {

		if index == 0 {
			continue
		}

		if value < report[index-1] {
			badLevelRemoved := slices.Concat(report[:index], report[index+1:])
			increasingFail += 1 + increasing(badLevelRemoved, level+1)
		}
	}

	return increasingFail
}

func decreasing(report intList, level int) int {
	decreasingFail := 0

	if level == 2 {
		return 0
	}

	for index, value := range report {

		if index == 0 {
			continue
		}

		if value > report[index-1] {
			badLevelRemoved := slices.Concat(report[:index], report[index+1:])
			decreasingFail += 1 + decreasing(badLevelRemoved, level+1)
		}
	}

	return decreasingFail
}

func safe2(report intList) bool {
	monotonicityFail := 0
	if isinc := isIncreasing(report); isinc < 1 {
		fmt.Printf("%v\t\t%v (inc)\n", report, isinc)
		monotonicityFail = increasing(report, 0)
	} else {
		fmt.Printf("%v\t\t%v (dec)\n", report, isinc)
		monotonicityFail = decreasing(report, 0)
	}

	outOfRangeFails := withinRange(report, 0)

	if outOfRangeFails > 1 {
		return false
	}

	if (monotonicityFail) > 1 {
		return false
	}

	if (outOfRangeFails + monotonicityFail) > 1 {
		return false
	}

	return true
}

func PartTwo(file string) {
	var report intList = make(intList, 0)
	countSafe := 0

	b, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		report = make(intList, 0)

		for _, number := range strings.Split(line, " ") {
			var a int
			fmt.Sscanf(number, "%d", &a)
			report = append(report, a)

		}

		if safe2(report) {
			countSafe++
		}

	}

	fmt.Printf("Total safe reports: %v", countSafe)
}

package day1

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type intList []int

func Day1(file string) {
	var listA intList = make(intList, 0)
	var listB intList = make(intList, 0)

	b, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		var a, b int
		fmt.Sscanf(line, "%d   %d", &a, &b)
		listA = append(listA, a)
		listB = append(listB, b)
	}

	sort.Slice(listA, func(i, j int) bool {
		return listA[i] < listA[j]
	})

	sort.Slice(listB, func(i, j int) bool {
		return listB[i] < listB[j]
	})

	var total int = 0

	for idx := range listA {
		total += int(math.Abs(float64(listA[idx] - listB[idx])))
	}

	fmt.Printf("Total length: %v", total)
}

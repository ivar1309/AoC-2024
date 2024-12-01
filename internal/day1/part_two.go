package day1

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func PartTwo(file string) {
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

	var total int = 0
	memoize := make(map[int]int)

	for _, itemFromA := range listA {
		if val, ok := memoize[itemFromA]; ok {
			total += val * itemFromA
			continue
		}

		count := 0

		for _, itemFromB := range listB {
			if itemFromB == itemFromA {
				count += 1
			}
		}

		memoize[itemFromA] = count

		total += itemFromA * count
	}

	fmt.Printf("Total length: %v", total)
}

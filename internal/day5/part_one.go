package day5

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkIfAccepted(pages []string, rules map[string][]string) bool {
	for i, page := range pages {
		for j, page2 := range pages {
			if i == j {
				continue
			}

			if i > j && slices.Contains(rules[page], page2) {
				return false
			}
		}
	}

	return true
}

func PartOne(file string) {
	b, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	lines := strings.Split(string(b), "\r\n")

	rules := make(map[string][]string)
	updates := make([][]string, 0)

	for _, line := range lines {
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")

			if _, ok := rules[parts[0]]; !ok {
				rules[parts[0]] = make([]string, 0)
			}

			rules[parts[0]] = append(rules[parts[0]], parts[1])
		}

		if strings.Contains(line, ",") {
			updates = append(updates, strings.Split(line, ","))
		}
	}

	accepted := make([]bool, 0, len(updates))

	for _, pages := range updates {
		accepted = append(accepted, checkIfAccepted(pages, rules))
	}

	sumOfMiddles := 0

	for idx, update := range updates {
		numberOfUpdates := len(update)
		if accepted[idx] {
			subSum, _ := strconv.ParseInt(update[numberOfUpdates/2], 10, 0)
			sumOfMiddles += int(subSum)
			fmt.Printf("update: %v\n", update)
			fmt.Printf("update[numberOfUpdates/2]: %v\n", update[numberOfUpdates/2])
		}
	}

	fmt.Printf("sumOfMiddles: %v\n", sumOfMiddles)
}

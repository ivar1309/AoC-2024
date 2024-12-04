package day4

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func XmasMap(line string, mapFunc func(int) int) int {
	total := 0
	for idx := range line {
		total += mapFunc(idx)
	}

	return total
}

func getAllIndexes(i, N int) [][]int {
	indexes := make([][]int, 0, 4)

	// don't include horizontal if we are exceeding linelength
	if (i%N)+3 < N {
		indexes = append(indexes, []int{i, i + 1, i + 2, i + 3})

		// don't include right diagonal if word won't fit
		if (i + 3*(N+1)) < (N * N) {
			indexes = append(indexes, []int{i, i + N + 1, i + 2*(N+1), i + 3*(N+1)})
		}
	}

	// don't include vertical if exceeding number of lines
	if (i + 3*N) < (N * N) {
		indexes = append(indexes, []int{i, i + N, i + 2*N, i + 3*N})

		// don't include left diagonal if word won't fit
		if (i+3*(N-1))%N < N-3 {
			indexes = append(indexes, []int{i, i + N - 1, i + 2*(N-1), i + 3*(N-1)})
		}
	}

	return indexes
}

func PartOne(file string) {
	b, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	lineLength := strings.Index(string(b), "\r\n")

	oneLongLine := strings.Replace(string(b), "\r\n", "", -1)

	xmasOcurrences := XmasMap(oneLongLine, func(i int) int {

		count := 0
		for _, orientation := range getAllIndexes(i, lineLength) {
			word := ""

			for _, offset := range orientation {
				word += string(oneLongLine[offset])
			}

			if word == "XMAS" || word == "SAMX" {
				count += 1
			}

		}

		return count
	})

	fmt.Printf("XMAS found %v times!", xmasOcurrences)
}

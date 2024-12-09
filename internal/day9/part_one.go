package day9

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func swapSlice(s *[]string, a, b int) {
	(*s)[a], (*s)[b] = (*s)[b], (*s)[a]
}

func findRightMostNonEmpty(l []string, stop int) int {
	for i := len(l) - 1; i > 0; i-- {

		if i < stop {
			return -1
		}

		if l[i] != "." {
			return i
		}

	}

	return -1
}

func PartOne(file string) {
	b, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("b:\n%v\n", string(b))

	var layout = make([]string, 0)
	fileId := 0

	for idx, val := range strings.Split(string(b), "") {
		//fmt.Printf("idx: %v, val: %v\n", idx, val)
		block, _ := strconv.ParseInt(val, 10, 0)

		if idx%2 == 0 {
			for i := 0; i < int(block); i++ {
				layout = append(layout, fmt.Sprintf("%v", fileId))
			}

			fileId++
		} else {
			for i := 0; i < int(block); i++ {
				layout = append(layout, ".")
			}
		}
	}

	fmt.Printf("layout: %v\n", layout)

	for idx, val := range layout {
		if val != "." {
			continue
		}

		swapIdx := findRightMostNonEmpty(layout, idx)

		if swapIdx != -1 {
			swapSlice(&layout, idx, swapIdx)
		}
	}

	fmt.Printf("layout: %v\n", layout)

	checksum := 0

	for idx, val := range layout {
		if val == "." {
			break
		}

		intVal, _ := strconv.ParseInt(val, 10, 0)
		checksum += idx * int(intVal)
	}

	fmt.Printf("checksum: %v\n", checksum)
}

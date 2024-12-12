package day9

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func swapSubSlice(s *[]string, a, b, c, d int) {
	A := (*s)[a : b+1]
	B := (*s)[c : d+1]

	for i := 0; i < len(A); i++ {
		(*s)[a+i], (*s)[c+i] = B[i], A[i]
	}
	//(*s)[a], (*s)[b] = (*s)[b], (*s)[a]
}

func findRightMostNonEmptySlice(l []string, offset, stop int) (int, int) {
	start := 0
	end := 0
	prevCheck := ""
	check := ""

	for i := offset; i > 0; i-- {
		if l[i] == "." {
			continue
		} else {
			prevCheck = l[i]
			end = i
			break
		}
	}

	for i := end - 1; i > 0; i-- {
		//if i < stop {
		//	start = -1
		//	break
		//}

		check = l[i]

		if prevCheck != check {
			start = i + 1
			break
		}
	}

	return start, end
}

func findLeftMostEmptySlice(l []string, offset, size, stop int) (int, int) {
	start := 0
	end := -1
	prevCheck := ""
	check := ""

	if offset >= stop {
		return -1, -1
	}

	for i := offset; i < len(l); i++ {
		if l[i] != "." {
			continue
		} else {
			prevCheck = "."
			start = i
			break
		}
	}

	for i := start + 1; i < stop; i++ {

		check = l[i]

		if prevCheck != check {
			end = i - 1
			//fmt.Printf("end: %v\n", end)
			break
		}
	}

	if end == -1 {
		return -1, end
	}

	if end-start+1 < size {
		//fmt.Println("________________________")
		//fmt.Printf("start: %v\n", start)
		//fmt.Printf("end: %v\n", end)
		//fmt.Printf("offset: %v\n", offset)
		//fmt.Printf("size: %v\n", size)
		//fmt.Printf("stop: %v\n", stop)
		start, end = findLeftMostEmptySlice(l, end+1, size, stop)
	}

	return start, end
}

func PartTwo(file string) {
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

	fmt.Printf("layout: %v [%v]\n", layout, len(layout))

	var swapA, swapB, swapC, swapD int
	swapA = len(layout)

	//for _, val := range layout {
	for i := 0; i < len(layout); i++ {

		swapA, swapB = findRightMostNonEmptySlice(layout, swapA-1, 0)
		//fmt.Printf("swapA: %v\n", swapA)
		//fmt.Printf("swapB: %v\n", swapB)
		//fmt.Printf("swapC: %v\n", swapC)
		//fmt.Printf("swapD: %v\n", swapD)

		swapC, swapD = findLeftMostEmptySlice(layout, 0, swapB-swapA+1, swapA)

		if swapA != -1 && swapD != -1 {
			//fmt.Printf("layout slice: %v\n", layout[swapA:swapB+1])
			//fmt.Printf("layout slice: %v\n", layout[swapC:swapD+1])
			swapSubSlice(&layout, swapA, swapB, swapC, swapD)
		}

		//i += swapC
		//fmt.Printf("layout: %v\n", layout)
	}

	fmt.Printf("layout: %v [%v]\n", layout, len(layout))

	checksum := 0

	for idx, val := range layout {

		if val == "." {
			continue
		}

		intVal, _ := strconv.ParseInt(val, 10, 0)
		checksum += idx * int(intVal)
	}

	fmt.Printf("checksum: %v\n", checksum)
}

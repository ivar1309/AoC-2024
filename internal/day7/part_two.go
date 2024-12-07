package day7

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func decToTernSlice(dec int, size int) []int {
	tern := strconv.FormatInt(int64(dec), 3)

	length := len(tern)
	for i := 0; i < size-length; i++ {
		tern = "0" + tern
	}

	ternSlice := make([]int, 0)
	for _, v := range strings.Split(string(tern), "") {
		vInt, _ := strconv.ParseInt(v, 10, 0)
		ternSlice = append(ternSlice, int(vInt))
	}

	return ternSlice
}

func concat(a, b int) int {
	A := strconv.FormatInt(int64(a), 10)
	B := strconv.FormatInt(int64(b), 10)

	combined, _ := strconv.ParseInt(A+B, 10, 0)

	return int(combined)
}

func PartTwo(file string) {

	Operators = append(Operators, concat)

	b, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	calibrationSum := 0

	for _, line := range strings.Split(string(b), "\r\n") {
		parts := strings.Split(line, ": ")

		result, _ := strconv.ParseInt(parts[0], 10, 0)
		operands := stringToIntSLice(parts[1])

		combinations := int(math.Pow(3, float64(len(operands)-1)) - 1)

		for i := 0; i <= combinations; i++ {
			copyOfOperands := slices.Clone(operands)
			bin := decToTernSlice(i, len(copyOfOperands)-1)

			//for idx, opIdx := range bin {
			for j := 0; j < len(bin); j++ {
				opIdx := bin[j]

				frontTwo := copyOfOperands[:2]
				calculation := Operators[opIdx](frontTwo[0], frontTwo[1])
				copyOfOperands = slices.Concat([]int{calculation}, copyOfOperands[2:])

			}

			calculated := copyOfOperands[0]

			if calculated == int(result) {
				calibrationSum += calculated
				break
			}
		}
	}

	fmt.Printf("calibrationSum: %v\n", calibrationSum)
}

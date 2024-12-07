package day7

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func decToBinSlice(dec int, size int) []int {
	bin := strconv.FormatInt(int64(dec), 2)

	length := len(bin)
	for i := 0; i < size-length; i++ {
		bin = "0" + bin
	}

	binSlice := make([]int, 0)
	for _, v := range strings.Split(string(bin), "") {
		vInt, _ := strconv.ParseInt(v, 10, 0)
		binSlice = append(binSlice, int(vInt))
	}

	return binSlice
}

func stringToIntSLice(numbers string) []int {
	intSlice := make([]int, 0)

	for _, num := range strings.Split(numbers, " ") {
		intNum, _ := strconv.ParseInt(num, 10, 0)
		intSlice = append(intSlice, int(intNum))
	}

	return intSlice
}

func plus(a, b int) int {
	return a + b
}

func mult(a, b int) int {
	return a * b
}

var Operators = []func(int, int) int{plus, mult}

func PartOne(file string) {
	b, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	calibrationSum := 0

	for _, line := range strings.Split(string(b), "\r\n") {
		parts := strings.Split(line, ": ")

		result, _ := strconv.ParseInt(parts[0], 10, 0)
		operands := stringToIntSLice(parts[1])

		combinations := int(math.Pow(2, float64(len(operands)-1)) - 1)

		for i := 0; i <= combinations; i++ {
			calculated := operands[0]
			bin := decToBinSlice(i, len(operands)-1)

			for idx, opIdx := range bin {
				if opIdx == 1 && calculated == 0 {
					calculated = 1
				}

				if idx < len(operands)-1 {
					calculated = Operators[opIdx](calculated, operands[idx+1])
				}
			}

			if calculated == int(result) {
				calibrationSum += calculated
				break
			}
		}
	}

	fmt.Printf("calibrationSum: %v\n", calibrationSum)
}

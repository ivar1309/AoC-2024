package day4

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func PartTwo(file string) {
	b, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	numberOfColumns := strings.Index(string(b), "\r\n")
	oneLongLine := strings.Replace(string(b), "\r\n", "", -1)
	numberOfRows := len(oneLongLine) / numberOfColumns

	count := 0

	for i := 1; i < numberOfRows-1; i++ {
		for j := 1; j < numberOfColumns-1; j++ {
			middle := i*numberOfRows + j
			upperLeft := middle - numberOfColumns - 1
			upperRight := middle - numberOfColumns + 1
			lowerLeft := middle + numberOfColumns - 1
			lowerRight := middle + numberOfColumns + 1

			wordDown :=
				string(oneLongLine[upperLeft]) +
					string(oneLongLine[middle]) +
					string(oneLongLine[lowerRight])

			wordUp :=
				string(oneLongLine[lowerLeft]) +
					string(oneLongLine[middle]) +
					string(oneLongLine[upperRight])

			if (wordUp == "MAS" || wordUp == "SAM") && (wordDown == "MAS" || wordDown == "SAM") {
				count++
			}
		}
	}

	fmt.Printf("X-MAS found %v times!", count)
}

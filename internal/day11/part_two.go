package day11

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

	fmt.Printf("b:\n%v\n", string(b))

	stones := make([]Stone, 0)

	for _, val := range strings.Split(string(b), " ") {
		stones = append(stones, CreateNewStone(val))
	}

	for i := 0; i < 25; i++ {
		for j := len(stones) - 1; j >= 0; j-- {
			stones[j].Blink(&stones)
		}
	}

	fmt.Printf("stones: %v\n", len(stones))
}

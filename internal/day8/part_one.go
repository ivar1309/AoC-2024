package day8

import (
	"fmt"
	"log"
	"os"
)

func PartOne(file string) {
	b, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("b:\n%v\n", string(b))
}

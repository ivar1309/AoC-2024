package day11

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Stone struct {
	value  string
	number int
}

func CreateNewStone(val string) Stone {
	s := Stone{}
	s.value = val
	num, _ := strconv.ParseInt(val, 10, 0)
	s.number = int(num)

	return s
}

func (s *Stone) Blink(stoneList *[]Stone) {
	if s.value == "0" {
		s.Zero()
	} else if len(s.value)%2 == 0 {
		s.Split(stoneList)
	} else {
		s.Mult2024()
	}
}

func (s *Stone) Change(newVal string) {
	s.value = newVal
	num, _ := strconv.ParseInt(newVal, 10, 0)
	s.number = int(num)
}

func (s *Stone) Zero() {
	s.Change("1")
}

func (s *Stone) Mult2024() {
	newNum := s.number * 2024
	s.Change(fmt.Sprintf("%d", newNum))
}

func (s *Stone) Split(stoneList *[]Stone) {
	oldVal := s.value
	left := oldVal[:len(oldVal)/2]
	right := oldVal[len(oldVal)/2:]

	s.Change(left)

	firstNonZeroIndex := 0
	onlyZeroes := true
	for i, v := range right {
		if v != '0' {
			firstNonZeroIndex = i
			onlyZeroes = false
			break
		}
	}

	if onlyZeroes {
		right = "0"
	} else {
		right = strings.Join(slices.Delete(strings.Split(right, ""), 0, firstNonZeroIndex), "")
	}

	newStone := CreateNewStone(right)
	*stoneList = append(*stoneList, newStone)
}

func (s Stone) String() string {
	return s.value
}

func PartOne(file string) {
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

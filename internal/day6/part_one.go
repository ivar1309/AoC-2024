package day6

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

const (
	STAY  = 0
	UP    = -1
	DOWN  = 1
	LEFT  = -1
	RIGHT = 1
)

var Movement = []string{"^", ">", "v", "<"}

func GetMovement(dir int) (int, int) {
	switch dir {
	case 0:
		return 0, -1
	case 1:
		return 1, 0
	case 2:
		return 0, 1
	case 3:
		return -1, 0
	default:
		return 0, 0
	}
}

type Guard struct {
	Area      [][]string
	Direction int
	X         int
	Y         int
}

func (g *Guard) Move() bool {
	if g.collides() {
		g.Direction = (g.Direction + 1) % 4
	}

	g.Area[g.Y][g.X] = "X"
	X, Y := GetMovement(g.Direction)
	g.X += X
	g.Y += Y

	return g.hasLeftTheRoom(g.X, g.Y)
}

func (g *Guard) collides() bool {
	X, Y := GetMovement(g.Direction)
	x := g.X + X
	y := g.Y + Y

	if g.hasLeftTheRoom(x, y) {
		return false
	}

	return g.Area[y][x] == "#"
}

func (g *Guard) hasLeftTheRoom(x, y int) bool {
	return x < 0 || x == len(g.Area[0]) || y < 0 || y == len(g.Area)
}

func (g Guard) String() string {
	output := ""

	for _, row := range g.Area {
		output += strings.Join(row, "") + "\n"
	}

	return output
}

func PartOne(file string) {
	b, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	var mappedArea [][]string

	for _, line := range strings.Split(string(b), "\r\n") {
		mappedArea = append(mappedArea, strings.Split(line, ""))
	}

	var guard Guard

	for y, row := range mappedArea {
		for x, col := range row {
			if slices.Contains([]string{"^", "v", "<", ">"}, col) {
				guard = Guard{
					Area:      mappedArea,
					Direction: slices.Index(Movement, col),
					X:         x,
					Y:         y,
				}
			}
		}
	}

	fmt.Printf("guard:\n%v\n", guard)

	for {
		if guard.Move() {
			break
		}
	}

	fmt.Printf("guard:\n%v\n", guard)

	countX := 0

	for _, row := range guard.Area {
		for _, col := range row {
			if col == "X" {
				countX++
			}
		}
	}

	fmt.Printf("Number of locations that the guard visits: %v\n", countX)
}

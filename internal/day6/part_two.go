package day6

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Guard2 struct {
	Area      [][]string
	Direction int
	X         int
	Y         int
}

func (g *Guard2) Move() bool {
	if g.collides() {
		g.Direction = (g.Direction + 1) % 4
	}

	g.Area[g.Y][g.X] = "X"
	X, Y := GetMovement(g.Direction)
	g.X += X
	g.Y += Y

	return g.hasLeftTheRoom(g.X, g.Y)
}

func (g *Guard2) collides() bool {
	X, Y := GetMovement(g.Direction)
	x := g.X + X
	y := g.Y + Y

	if g.hasLeftTheRoom(x, y) {
		return false
	}

	return g.Area[y][x] == "#"
}

func (g *Guard2) hasLeftTheRoom(x, y int) bool {
	return x < 0 || x == len(g.Area[0]) || y < 0 || y == len(g.Area)
}

func (g Guard2) String() string {
	output := ""

	for _, row := range g.Area {
		output += strings.Join(row, "") + "\n"
	}

	return output
}

func PartTwo(file string) {
	b, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	var mappedArea [][]string

	for _, line := range strings.Split(string(b), "\r\n") {
		mappedArea = append(mappedArea, strings.Split(line, ""))
	}

	var guard Guard2

	for y, row := range mappedArea {
		for x, col := range row {
			if slices.Contains([]string{"^", "v", "<", ">"}, col) {
				guard = Guard2{
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

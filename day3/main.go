package main

import (
	_ "embed"
	"fmt"
)

//go:embed inputs.txt
var input string

func main() {
	sum := parseEngineSchematic(input)

	fmt.Printf("Part 1: %v", sum)
}

func parseEngineSchematic(schematic string) int {
	return 0
}

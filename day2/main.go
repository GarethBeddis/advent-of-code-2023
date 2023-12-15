package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var input string

func main() {
	cubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	result, pow := getPossibleGameCount(input, cubes)

	fmt.Printf("Part 1: %v\nPart 2: %v", result, pow)
}

func getPossibleGameCount(input string, cubes map[string]int) (int, int) {
	var count int
	var totalPow int

	for _, game := range strings.Split(input, "\n") {
		gameId, isGamePossible, pow := getGameState(game, cubes)
		if strings.TrimSpace(game) != "" && isGamePossible {
			count += gameId
		}
		totalPow += pow
	}

	return count, totalPow
}

func getGameState(game string, bag map[string]int) (gameId int, isGamePossible bool, powOfMinCubes int) {
	gameStr, bags, _ := strings.Cut(game, ":")
	gameNumber, err := strconv.Atoi(strings.Fields(gameStr)[1])
	if err != nil {
		panic(err)
	}

	cubes := strings.Split(bags, ";")
	isGamePossible = true
	cubeCounter := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, v := range cubes {
		cube := strings.Split(v, ",")
		for _, v := range cube {
			cubeFields := strings.Fields(v)
			color := cubeFields[1]
			amount, err := strconv.Atoi(cubeFields[0])
			if err != nil {
				panic(err)
			}

			if amount > cubeCounter[color] {
				cubeCounter[color] = amount
			}

			if cubeCounter[color] > bag[color] {
				isGamePossible = false
			}
		}
	}

	powOfMinCubes = cubeCounter["red"] * cubeCounter["green"] * cubeCounter["blue"]
	return gameNumber, isGamePossible, powOfMinCubes
}

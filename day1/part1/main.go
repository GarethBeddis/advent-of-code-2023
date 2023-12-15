package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inputFile, err := os.Open("../inputs.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	cv := processCalibrationDocument(inputFile)
	fmt.Printf("result: %v", cv)
}

func processCalibrationDocument(inputFile io.Reader) int {
	cv := getCalibrationValues(inputFile)
	return sumArray(cv)
}

func getCalibrationValues(r io.Reader) []int {
	fileScanner := bufio.NewScanner(r)

	fileScanner.Split(bufio.ScanLines)

	calibrationValues := []int{}

	for fileScanner.Scan() {
		numericLine := getNumericCharacters(fileScanner.Text())
		firstDigit := string(numericLine[0])
		lastDigit := string(numericLine[len(numericLine)-1])
		val, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			panic(err)
		}
		calibrationValues = append(calibrationValues, int(val))
	}

	return calibrationValues
}

func sumArray(a []int) int {
	sum := 0
	for _, v := range a {
		sum += int(v)
	}

	return sum
}

func getNumericCharacters(s string) string {
	return regexp.MustCompile("[^0-9]+").ReplaceAllString(s, "")
}

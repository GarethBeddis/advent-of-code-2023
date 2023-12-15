package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var digitMap = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

func main() {
	inputFile, err := os.Open("../inputs.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	result := processCalibrationDocument(inputFile)
	fmt.Print(result)
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
		line := fileScanner.Text()
		firstDigit := getFirstDigit(line)
		lastDigit := getLastdigit(line)
		val, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			panic(err)
		}
		calibrationValues = append(calibrationValues, int(val))
	}

	return calibrationValues
}

func getFirstDigit(s string) (d string) {
	chars := strings.Split(s, "")
	for i, c := range chars {
		if _, err := strconv.Atoi(c); err == nil {
			d = c
			break
		} else if isAlphabetical, ad := startsWithAlphabeticalDigit(s[i:]); isAlphabetical {
			d = digitMap[ad]
			break
		}
	}
	return
}

func startsWithAlphabeticalDigit(s string) (bool, string) {
	for i := range digitMap {
		if strings.HasPrefix(s, i) {
			return true, i
		}
	}
	return false, ""
}

func getLastdigit(s string) (d string) {
	chars := reverseArray(strings.Split(s, ""))

	for i, c := range chars {
		if _, err := strconv.Atoi(c); err == nil {
			d = c
			break
		} else if isAlphabetical, ad := endsWithAlphabeticalDigit(s[:len(s)-i]); isAlphabetical {
			d = digitMap[ad]
			break
		}
	}
	return
}

func endsWithAlphabeticalDigit(s string) (bool, string) {
	for i := range digitMap {
		if strings.HasSuffix(s, i) {
			return true, i
		}
	}
	return false, ""
}

func reverseArray(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func sumArray(a []int) int {
	sum := 0
	for _, v := range a {
		sum += int(v)
	}

	return sum
}

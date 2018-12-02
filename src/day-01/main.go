package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Atoi %s", err)
		}
		lines = append(lines, i)
	}
	return lines, scanner.Err()
}

func partOne() {
	lines, err := readLines("./input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	result := 0
	for _, v := range lines {
		result += v
	}
	fmt.Println(result)
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func partTwo() {
	lines, err := readLines("./input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	freqs := []int{}
	result := 0
	freqs = append(freqs, result)
	done := false
	i := 0
	for done != true {
		result += lines[i]
		if contains(freqs, result) {
			done = true
		} else {
			freqs = append(freqs, result)
			if len(lines)-1 == i {
				i = 0
			} else {
				i++
			}
		}
	}

	fmt.Println(result)
}

func main() {
	partOne()
	partTwo()
}

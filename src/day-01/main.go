package main

import (
	"bufio"
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

// PartOne solution
func PartOne() int {
	lines, err := readLines("./input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	result := 0
	for _, v := range lines {
		result += v
	}
	return result
}

// PartTwo solution
func PartTwo() int {
	lines, err := readLines("./input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	freqs := make(map[int]int)
	result := 0
	done := false
	i := 0
	for done != true {
		result += lines[i]
		_, ok := freqs[result]
		if ok {
			done = true
		} else {
			freqs[result]++
			if len(lines)-1 == i {
				i = 0
			} else {
				i++
			}
		}
	}

	return result
}

func main() {
	PartOne()
	PartTwo()
}

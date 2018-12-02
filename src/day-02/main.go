package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

type twoThree struct {
	two   int
	three int
}

// PartOne solution
func PartOne() int {
	lines, err := readLines("./input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	var result twoThree
	for _, line := range lines {
		m := make(map[string]int)
		characters := strings.Split(line, "")
		for _, c := range characters {
			_, ok := m[c]
			if ok {
				m[c]++
			} else {
				m[c] = 1
			}
		}
		two := false
		three := false
		for _, v := range m {
			if v == 2 {
				two = true
			} else if v == 3 {
				three = true
			}
		}
		if two {
			result.two++
		}
		if three {
			result.three++
		}
	}
	return result.two * result.three
}

func compareLines(lines []string) string {
	for i, lineI := range lines {
		ci := strings.Split(lineI, "")
		for j := i + 1; j < len(lines); j++ {
			var notSameLocations []int
			lineJ := lines[j]
			cj := strings.Split(lineJ, "")
			for x, v := range ci {
				if cj[x] != v {
					notSameLocations = append(notSameLocations, x)
				}
			}
			if len(notSameLocations) == 1 {
				x := notSameLocations[0]
				return strings.Join(append(ci[:x], ci[x+1:]...), "")
			}
		}
	}
	return ""
}

// PartTwo solution
func PartTwo() string {
	lines, err := readLines("./input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	result := compareLines(lines)
	return result
}

func main() {
	PartOne()
	PartTwo()
}

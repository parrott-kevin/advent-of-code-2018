package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func readLines() (map[string][]int, int) {
	path := "./input.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("open: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	claimed := make(map[string][]int)
	idMax := 0
	for scanner.Scan() {
		line := scanner.Text()
		splitStrings := regexp.MustCompile("#|@|,|:|x").Split(line, -1)
		var split []int
		for _, v := range splitStrings {
			if v != "" {
				t := strings.TrimSpace(v)
				i, err := strconv.Atoi(t)
				if err != nil {
					log.Fatalf("Atoi %s", err)
				}
				split = append(split, i)
			}
		}

		id := split[0]
		if id > idMax {
			idMax = id
		}
		xMin := split[1]
		xMax := xMin + split[3] - 1
		yMin := split[2]
		yMax := yMin + split[4] - 1
		for i := xMin; i <= xMax; i++ {
			for j := yMin; j <= yMax; j++ {
				key := strconv.Itoa(i) + "," + strconv.Itoa(j)
				claimed[key] = append(claimed[key], id)
			}
		}
	}

	if scanner.Err() != nil {
		log.Fatalf("scanner: %s", err)
	}

	return claimed, idMax
}

// PartOne solution
func PartOne() int {
	claimed, _ := readLines()
	result := 0
	for _, v := range claimed {
		if len(v) > 1 {
			result++
		}
	}
	return result
}

// PartTwo solution
func PartTwo() int {
	claimed, idMax := readLines()
	set := make(map[int]struct{})
	for _, v := range claimed {
		if len(v) > 1 {
			for _, id := range v {
				if _, ok := set[id]; !ok {
					set[id] = struct{}{}
				}
			}
		}
	}

	var keys []int
	for key := range set {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	result := -1
	for index, value := range keys {
		if index > 1 && value-keys[index-1] != 1 {
			result = value - 1
		}
	}
	if result == -1 && keys[0] == 1 {
		result = idMax
	}
	return result
}

func main() {
	PartOne()
	PartTwo()
}

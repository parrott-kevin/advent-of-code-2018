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

type rectangle struct {
	id   int
	xMin int
	xMax int
	yMin int
	yMax int
}

type grid struct {
	xMax       int
	yMax       int
	rectangles []rectangle
}

func readLines(path string) (grid, error) {
	file, err := os.Open(path)
	if err != nil {
		return grid{}, err
	}
	defer file.Close()

	canvas := grid{}
	var lines []rectangle
	scanner := bufio.NewScanner(file)
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
		xMin := split[1]
		xMax := xMin + split[3] - 1
		yMin := split[2]
		yMax := yMin + split[4] - 1
		if xMax > canvas.xMax {
			canvas.xMax = xMax
		}
		if yMax > canvas.yMax {
			canvas.yMax = yMax
		}
		r := rectangle{
			id,
			xMin,
			xMax,
			yMin,
			yMax,
		}
		lines = append(lines, r)
	}
	canvas.rectangles = lines
	return canvas, scanner.Err()
}

func claimedCanvas(canvas grid) map[string][]int {
	claimed := make(map[string][]int)
	for x := 0; x <= canvas.xMax; x++ {
		for y := 0; y <= canvas.yMax; y++ {
			for _, r := range canvas.rectangles {
				if x >= r.xMin && x <= r.xMax && y >= r.yMin && y <= r.yMax {
					key := strconv.Itoa(x) + "," + strconv.Itoa(y)
					claimed[key] = append(claimed[key], r.id)
				}
			}
		}
	}
	return claimed
}

// PartOne solution
func PartOne() int {
	canvas, err := readLines("./input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	claimed := claimedCanvas(canvas)

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
	canvas, err := readLines("./input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	claimed := claimedCanvas(canvas)

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
		result = canvas.rectangles[len(canvas.rectangles)-1].id
	}
	return result
}

func main() {
	PartOne()
	PartTwo()
}

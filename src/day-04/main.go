package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Record defines the properties of a guard's record
type Record struct {
	id   int
	date int
	time int
	text string
}

func readFile() []Record {
	path := "./input.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("open: %s", err)
	}
	defer file.Close()

	records := []Record{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// example line [1518-11-01 00:00] Guard #10 begins shift
		dateStr := strings.Replace(line[1:11], "-", "", -1)
		timeStr := strings.Replace(line[12:17], ":", "", -1)
		dateTimeStr := dateStr + timeStr

		date, err := strconv.Atoi(dateStr)
		if err != nil {
			log.Fatalf("Atoi %s", err)
		}

		time, err := strconv.Atoi(timeStr)
		if err != nil {
			log.Fatalf("Atoi %s", err)
		}

		dateTime, err := strconv.Atoi(dateTimeStr)
		if err != nil {
			log.Fatalf("Atoi %s", err)
		}

		r := Record{
			dateTime,
			date,
			time,
			line[19:],
		}
		records = append(records, r)
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].id < records[j].id
	})
	return records
}

// SleepTracker defines properties for the midnight hour
type SleepTracker struct {
	count   int
	minutes map[int]int
}

// HighScoreUser defines properties for longest sleeper
type HighScoreUser struct {
	id     int
	length int
}

// HighScoreMinute defines properties for most used minute
type HighScoreMinute struct {
	minute int
	length int
}

// FrequentMinute defines properties for most used minute
type FrequentMinute struct {
	id     int
	minute int
	length int
}

func countSheep() map[int]*SleepTracker {
	records := readFile()
	guard := 0
	sleepCount := 0
	sleepStart := 0

	sleepTime := map[int]*SleepTracker{}

	for _, record := range records {
		if strings.Contains(record.text, "Guard") {
			g, err := strconv.Atoi(record.text[7 : strings.Index(record.text, "begins")-1])
			guard = g
			if err != nil {
				log.Fatalf("Atoi %s", err)
			}
			sleepCount = 0
		}

		if record.text == "falls asleep" {
			sleepStart = record.time
		}

		if record.text == "wakes up" {
			sleepCount = record.time - sleepStart
			if item := sleepTime[guard]; item != nil {
				item.count += sleepCount
				for i := sleepStart; i < record.time; i++ {
					item.minutes[i]++
				}
			} else {
				minutes := make(map[int]int)
				for i := 0; i < 60; i++ {
					if i >= sleepStart && i < record.time {
						minutes[i] = 1
					} else {
						minutes[i] = 0
					}
				}
				sleepTime[guard] = &SleepTracker{
					sleepCount,
					minutes,
				}
			}
		}
	}
	return sleepTime
}

// PartOne function
func PartOne() int {
	highScoreUser := HighScoreUser{0, 0}
	sleepTime := countSheep()

	for k, v := range sleepTime {
		if v.count > highScoreUser.length {
			highScoreUser = HighScoreUser{k, v.count}
		}
	}

	highScoreMinute := HighScoreMinute{0, 0}
	for k, v := range sleepTime[highScoreUser.id].minutes {
		if v > highScoreMinute.length {
			highScoreMinute.minute = k
			highScoreMinute.length = v
		}
	}
	return highScoreUser.id * highScoreMinute.minute
}

// PartTwo solution
func PartTwo() int {
	sleepTime := countSheep()

	fm := FrequentMinute{0, 0, 0}
	for id, st := range sleepTime {
		for m, v := range st.minutes {
			if v > fm.length {
				fm = FrequentMinute{id, m, v}
			}
		}
	}

	return fm.id * fm.minute
}

func main() {
	PartOne()
	PartTwo()
}

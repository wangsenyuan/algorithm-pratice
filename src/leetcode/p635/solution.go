package main

import (
	"strings"
	"strconv"
	"sort"
	"fmt"
)

type Entry struct {
	year, month, day, hour, minute, second int
}

func (b Entry) Gt(a Entry) bool {
	if a.year < b.year {
		return true
	}

	if a.year == b.year && a.month < b.month {
		return true
	}

	if a.year == b.year && a.month == b.month && a.day < b.day {
		return true
	}

	if a.year == b.year && a.month == b.month && a.day == b.day && a.hour < b.hour {
		return true
	}

	if a.year == b.year && a.month == b.month && a.day == b.day && a.hour == b.hour && a.minute < b.minute {
		return true
	}

	if a.year == b.year && a.month == b.month && a.day == b.day && a.hour == b.hour && a.minute == b.minute && a.second < b.second {
		return true
	}

	return false
}

func (b Entry) Eq(a Entry) bool {
	return b.year == a.year && b.month == a.month && b.day == a.day && b.hour == a.hour && b.minute == a.minute && b.second == a.second
}

func MakeEntry(timestamp string) Entry {
	parts := strings.Split(timestamp, ":")
	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])
	hour, _ := strconv.Atoi(parts[3])
	minute, _ := strconv.Atoi(parts[4])
	second, _ := strconv.Atoi(parts[5])

	return Entry{year, month, day, hour, minute, second}
}

func GetPrevEntry(timestamp, gra string) Entry {
	times := make([]int, 6)
	parts := strings.Split(timestamp, ":")

	at := 6
	if gra == "Year" {
		at = 1
	} else if gra == "Month" {
		at = 2
	} else if gra == "Day" {
		at = 3
	} else if gra == "Hour" {
		at = 4
	} else if gra == "Minute" {
		at = 5
	} else if gra == "Second" {
		at = 6
	}

	for i := 0; i < at; i++ {
		times[i], _ = strconv.Atoi(parts[i])
	}
	return Entry{times[0], times[1], times[2], times[3], times[4], times[5]}
}

func GetNextEntry(timestamp, gra string) Entry {
	times := make([]int, 6)
	parts := strings.Split(timestamp, ":")

	at := 6
	if gra == "Year" {
		at = 1
	} else if gra == "Month" {
		at = 2
	} else if gra == "Day" {
		at = 3
	} else if gra == "Hour" {
		at = 4
	} else if gra == "Minute" {
		at = 5
	} else if gra == "Second" {
		at = 6
	}

	for i := 0; i < at; i++ {
		times[i], _ = strconv.Atoi(parts[i])
	}
	times[at-1] += 1

	return Entry{times[0], times[1], times[2], times[3], times[4], times[5]}
}

type Entries []Entry

func (entries Entries) Len() int {
	return len(entries)
}

func (entries Entries) Less(i, j int) bool {
	a, b := entries[i], entries[j]
	if a.year < b.year {
		return true
	}

	if a.year == b.year && a.month < b.month {
		return true
	}

	if a.year == b.year && a.month == b.month && a.day < b.day {
		return true
	}

	if a.year == b.year && a.month == b.month && a.day == b.day && a.hour < b.hour {
		return true
	}

	if a.year == b.year && a.month == b.month && a.day == b.day && a.hour == b.hour && a.minute < b.minute {
		return true
	}

	if a.year == b.year && a.month == b.month && a.day == b.day && a.hour == b.hour && a.minute == b.minute && a.second < b.second {
		return true
	}

	return false
}

func (entries Entries) Swap(i, j int) {
	entries[i], entries[j] = entries[j], entries[i]
}

type LogSystem struct {
	entries []Entry
	cache   map[Entry]int
}

func Constructor() LogSystem {
	return LogSystem{entries: make([]Entry, 0, 100), cache: make(map[Entry]int)}
}

func (this *LogSystem) Put(id int, timestamp string) {
	entry := MakeEntry(timestamp)
	this.entries = append(this.entries, entry)
	this.cache[entry] = id
}

func (this *LogSystem) Retrieve(s string, e string, gra string) []int {
	sort.Sort(Entries(this.entries))
	x, y := GetPrevEntry(s, gra), GetNextEntry(e, gra)
	i := sort.Search(len(this.entries), func(i int) bool {
		return this.entries[i].Gt(x) || this.entries[i].Eq(x)
	})
	j := sort.Search(len(this.entries), func(j int) bool {
		return this.entries[j].Gt(y) || this.entries[j].Eq(y)
	})

	var res = make([]int, j-i)

	for k := i; k < j; k++ {
		res[k-i] = this.cache[this.entries[k]]
	}

	return res
}

func main() {
	test2()
}

func test2() {
	logSystem := Constructor()
	logSystem.Put(1, "2017:01:01:23:59:59")
	logSystem.Put(2, "2017:01:02:23:59:59")
	fmt.Println(logSystem.Retrieve("2017:01:01:23:59:58", "2017:01:02:23:59:58", "Second"))
}

func test1() {
	logSystem := Constructor()
	logSystem.Put(1, "2017:01:01:23:59:59")
	logSystem.Put(2, "2017:01:01:22:59:59")
	logSystem.Put(3, "2016:01:01:00:00:00")
	fmt.Println(logSystem.Retrieve("2016:01:01:01:01:01", "2017:01:01:23:00:00", "Year"))
	fmt.Println(logSystem.Retrieve("2016:01:01:01:01:01", "2017:01:01:23:00:00", "Hour"))
}

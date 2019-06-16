package p937

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	n := len(logs)
	letters := make([]Log, n)
	digits := make([]Log, n)
	var a, b int
	for i := 0; i < n; i++ {
		log := logs[i]
		ss := strings.SplitAfterN(log, " ", 2)
		id := strings.TrimSpace(ss[0])
		left := strings.TrimSpace(ss[1])
		if left[0] >= 'a' && left[0] <= 'z' {
			letters[a] = Log{id, left}
			a++
		} else {
			digits[b] = Log{id, left}
			b++
		}
	}

	letters = letters[:a]

	sort.Sort(Logs(letters))

	digits = digits[:b]

	res := make([]string, n)
	var j int
	for i := 0; i < a; i++ {
		log := letters[i]
		res[j] = log.id + " " + log.left
		j++
	}

	for i := 0; i < b; i++ {
		log := digits[i]
		res[j] = log.id + " " + log.left
		j++
	}

	return res
}

type Log struct {
	id   string
	left string
}

type Logs []Log

func (logs Logs) Len() int {
	return len(logs)
}

func (logs Logs) Less(i, j int) bool {
	return logs[i].left < logs[j].left
}

func (logs Logs) Swap(i, j int) {
	logs[i], logs[j] = logs[j], logs[i]
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)
	events := make([][]int, n)
	for i := 0; i < n; i++ {
		events[i] = readNNums(scanner, 2)
	}

	fmt.Println(solve(n, events))
}

func solve(n int, events [][]int) int {
	evts := make(Events, n)
	for i := 0; i < n; i++ {
		evts[i] = Event{events[i][0], events[i][0] + events[i][1]}
	}
	sort.Sort(evts)

	var ans int
	last := -1
	for i := 0; i < n; i++ {
		if evts[i].start > last {
			ans++
			last = evts[i].end
		}
	}
	return ans
}

type Event struct {
	start, end int
}

type Events []Event

func (events Events) Len() int {
	return len(events)
}

func (events Events) Less(i, j int) bool {
	return events[i].end < events[j].end
}

func (events Events) Swap(i, j int) {
	events[i], events[j] = events[j], events[i]
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, X := readTwoNums(reader)
		S, E := make([]int, n), make([]int, n)
		for i := 0; i < n; i++ {
			S[i], E[i] = readTwoNums(reader)
		}

		fmt.Println(solve(n, X, S, E))
	}
}

func solve(n, X int, S, E []int) int {
	events := make([]Event, n)

	for i := 0; i < n; i++ {
		events[i] = Event{S[i], E[i], i}
	}
	sort.Sort(Events(events))

	C := make([]int, n)

	var res int

	for i := 0; i < n; i++ {
		var flag bool

		for j := 0; j <= i; j++ {
			if events[j].second >= events[i].first {
				C[j]++
				if C[j] == X {
					flag = true
				}
			}
		}

		if flag {
			res++
			for j := 0; j <= i; j++ {
				if events[j].second >= events[i].first {
					C[j]--
				}
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Event struct {
	first, second int
	index         int
}

type Events []Event

func (events Events) Len() int {
	return len(events)
}

func (events Events) Less(i, j int) bool {
	return events[i].second < events[j].second || events[i].second == events[j].second && events[i].first < events[j].first
}

func (events Events) Swap(i, j int) {
	events[i], events[j] = events[j], events[i]
}

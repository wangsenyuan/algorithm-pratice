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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, q := readTwoNums(scanner)
		A := readNNums(scanner, n)
		queries := make([][]int, q)
		for i := 0; i < q; i++ {
			queries[i] = readNNums(scanner, 3)
		}
		res := solve(n, A, queries)

		for _, ans := range res {
			fmt.Println(ans)
		}
	}
}

func solve(n int, A []int, queries [][]int) []int {
	events := make([]Event, 0, 2*n+len(queries))

	for i := 0; i < n-1; i++ {
		events = append(events, NewEvent(min(A[i], A[i+1]), 1, i))
		events = append(events, NewEvent(max(A[i], A[i+1]), 3, i))
	}

	for i := 0; i < len(queries); i++ {
		events = append(events, NewEvent(queries[i][2], 2, i))
	}

	sort.Sort(Events(events))

	arr := make([]int, n+1)

	update := func(pos int, v int) {
		pos++
		for pos <= n {
			arr[pos] += v
			pos += pos & -pos
		}
	}

	get := func(pos int) int {
		var res int
		pos++
		for pos > 0 {
			res += arr[pos]
			pos -= pos & -pos
		}
		return res
	}

	ans := make([]int, len(queries))

	for _, event := range events {
		if event.tp == 1 {
			update(event.i, 1)
		} else if event.tp == 2 {
			qr := queries[event.i]
			l, r := qr[0]-1, qr[1]-1
			ans[event.i] = get(r-1) - get(l-1)
		} else {
			update(event.i, -1)
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Event struct {
	y  int
	tp int
	i  int
}

func NewEvent(y int, tp int, i int) Event {
	return Event{y, tp, i}
}

type Events []Event

func (events Events) Len() int {
	return len(events)
}

func (events Events) Less(i, j int) bool {
	return events[i].y < events[j].y || (events[i].y == events[j].y && events[i].tp < events[j].tp)
}

func (events Events) Swap(i, j int) {
	events[i], events[j] = events[j], events[i]
}

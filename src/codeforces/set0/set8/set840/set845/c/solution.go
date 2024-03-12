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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	shows := make([][]int, n)
	for i := 0; i < n; i++ {
		shows[i] = readNNums(reader, 2)
	}

	res := solve(shows)

	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func solve(shows [][]int) bool {
	n := len(shows)
	events := make([]Event, 2*n)
	for i, cur := range shows {
		events[i*2] = Event{i, cur[0], 0}
		events[i*2+1] = Event{i, cur[1] + 1, 1}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].pos < events[j].pos || events[i].pos == events[j].pos && events[i].tp > events[j].tp
	})

	var open int
	for _, cur := range events {
		if cur.tp == 0 {
			open++
		} else {
			open--
		}
		if open > 2 {
			return false
		}
	}
	return true
}

type Event struct {
	id  int
	pos int
	tp  int
}

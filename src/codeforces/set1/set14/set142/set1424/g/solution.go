package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	E := make([][]int, n)
	for i := 0; i < n; i++ {
		E[i] = readNNums(reader, 2)
	}
	res := solve(E)
	fmt.Println(fmt.Sprintf("%d %d", res[0], res[1]))
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(records [][]int) []int {
	type Event struct {
		time int
		tp   int
	}

	n := len(records)

	events := make([]Event, 2*n)

	for i, cur := range records {
		events[2*i] = Event{cur[0], 0}
		events[2*i+1] = Event{cur[1] - 1, 1}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].time < events[j].time || events[i].time == events[j].time && events[i].tp < events[j].tp
	})

	var alive int
	var year int
	var most int
	for i := 0; i < len(events); i++ {
		cur := events[i]
		if cur.tp == 0 {
			alive++
		} else {
			alive--
		}
		if most < alive {
			year = cur.time
			most = alive
		}
	}

	return []int{year, most}
}

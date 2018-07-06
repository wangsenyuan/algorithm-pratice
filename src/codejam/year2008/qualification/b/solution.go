package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
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

func main() {
	// f, err := os.Open("./B-execise-practice.in")
	// f, err := os.Open("./B-small-practice.in")
	f, err := os.Open("./B-large-practice.in")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		T := readNum(scanner)
		NA, NB := readTwoNums(scanner)
		A := make([][]int, NA)
		for j := 0; j < NA; j++ {
			scanner.Scan()
			str := scanner.Text()
			A[j] = make([]int, 2)
			A[j][0] = ParseTime(str[:5])
			A[j][1] = ParseTime(str[6:])
		}
		B := make([][]int, NB)
		for j := 0; j < NB; j++ {
			scanner.Scan()
			str := scanner.Text()
			B[j] = make([]int, 2)
			B[j][0] = ParseTime(str[:5])
			B[j][1] = ParseTime(str[6:])
		}
		a, b := solve(T, NA, NB, A, B)
		fmt.Printf("Case #%d: %d %d\n", i, a, b)
	}
}

func ParseTime(str string) int {
	hs, ms := str[:2], str[3:]
	h, _ := strconv.Atoi(hs)
	m, _ := strconv.Atoi(ms)
	return h*60 + m
}

func solve(T int, NA, NB int, A [][]int, B [][]int) (int, int) {
	items := make(Items, NA+NB)

	for i := 0; i < NA; i++ {
		items[i] = Item{departure: A[i][0], arrival: A[i][1], from: 0}
	}
	for i := 0; i < NB; i++ {
		items[i+NA] = Item{departure: B[i][0], arrival: B[i][1], from: 1}
	}
	sort.Sort(items)

	// train from station b, arrive at a
	ss := make([]Station, 2)
	ss[0] = make(Station, 0, NB)
	ss[1] = make(Station, 0, NA)
	res := make([]int, 2)
	for i := 0; i < len(items); i++ {
		item := items[i]
		from := item.from
		to := 1 - from

		if ss[from].Len() > 0 && ss[from][0] <= item.departure {
			heap.Pop(&ss[from])
		} else {
			res[from]++
		}
		heap.Push(&ss[to], item.arrival+T)
	}

	return res[0], res[1]
}

type Item struct {
	departure, arrival int
	from               int
}

type Items []Item

func (ts Items) Len() int {
	return len(ts)
}

func (ts Items) Less(i, j int) bool {
	return ts[i].departure < ts[j].departure
}

func (ts Items) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

type Station []int

func (ts Station) Len() int {
	return len(ts)
}

func (ts Station) Less(i, j int) bool {
	return ts[i] < ts[j]
}

func (ts Station) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

func (ts *Station) Push(item interface{}) {
	train := item.(int)
	*ts = append(*ts, train)
}

func (ts *Station) Pop() interface{} {
	old := *ts
	n := len(old)
	last := old[n-1]
	*ts = old[:n-1]
	return last
}

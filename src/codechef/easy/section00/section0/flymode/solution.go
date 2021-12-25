package main

import (
	"bufio"
	"fmt"
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
	var n int
	fmt.Scanf("%d", &n)
	var hs = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &hs[i])
	}
	res := solve(n, hs)
	fmt.Println(res)
}

type event struct {
	value int
	open  bool
}

type orderEvents []event

func (this orderEvents) Len() int {
	return len(this)
}

func (this orderEvents) Less(i, j int) bool {
	return this[i].value < this[j].value
}

func (this orderEvents) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func solve(n int, hs []int) int {
	if n == 1 {
		return 1
	}
	events := make([]event, 2*(n-1))

	for i := 0; i < n-1; i++ {
		a, b := hs[i], hs[i+1]
		if a > b {
			a, b = b, a
		}
		events[2*i] = event{a, true}
		events[2*i+1] = event{b, false}
	}

	sort.Sort(orderEvents(events))

	res := 0
	cnt := 0
	for i := 0; i < len(events); i++ {
		if events[i].open {
			cnt++
		} else {
			cnt--
		}

		if cnt > res {
			res = cnt
		}
	}

	return res
}

func solve1(n int, hs []int) int {
	entry, exit := make([]int, n-1), make([]int, n-1)

	for i := 0; i < n-1; i++ {
		if hs[i] < hs[i+1] {
			entry[i] = hs[i]
			exit[i] = hs[i+1]
		} else {
			entry[i] = hs[i+1]
			exit[i] = hs[i]
		}
	}

	sort.Ints(entry)
	sort.Ints(exit)

	cnt := 1
	res := 1
	for i, j := 1, 0; i < n-1 && j < n-1; {
		if entry[i] < exit[j] {
			cnt++
			if cnt > res {
				res = cnt
			}
			i++
		} else {
			cnt--
			j++
		}
	}
	return res
}

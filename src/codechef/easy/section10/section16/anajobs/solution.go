package main

import (
	"bufio"
	"fmt"
	"os"
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
		line := readNNums(scanner, 3)
		N, K, J := line[0], line[1], line[2]
		events := make([][]int, N)
		for i := 0; i < N; i++ {
			events[i] = readNNums(scanner, 2)
		}
		fmt.Println(solve(N, K, J, events))
	}
}

func solve(N, K, J int, events [][]int) int {
	tableMap := make(map[int]*Table)

	for _, event := range events {
		t, sz := event[0], event[1]
		if entry, found := tableMap[t]; !found {
			entry = new(Table)
			tableMap[t] = entry
			entry.Update(sz)
		} else {
			entry.Update(sz)
		}
	}

	tables := make([]*Table, len(tableMap))
	var p int
	for _, t := range tableMap {
		tables[p] = t
		p++
	}

	check := func(ans int) bool {
		var num int

		for _, t := range tables {
			//concat sz_2 as most as can
			c2 := t.count2
			c1 := t.count1

			if ans == 1 && c2 > 0 {
				return false
			}

			var need int
			if c2*2 <= ans {
				need = (c2*2 + c1 + ans - 1) / ans
			} else {
				bit := ans & 1
				x := ans / 2
				y, r := c2/x, c2%x
				c1 = max(0, c1-y*bit)
				need = y + (r*2+c1+ans-1)/ans
			}

			num += (need + K - 1) / K
		}

		return num <= J
	}

	left, right := 1, 200000

	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Table struct {
	count1, count2 int
}

func (table *Table) Update(sz int) {
	if sz == 1 {
		table.count1++
	} else {
		table.count2++
	}
}

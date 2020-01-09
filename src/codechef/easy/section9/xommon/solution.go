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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)

	A := make([]uint64, n)

	var pos int

	scanner.Scan()

	for i := 0; i < n; i++ {
		pos = readUint64(scanner.Bytes(), pos, &A[i]) + 1
	}

	fmt.Println(solve(n, A))
}

func solve(n int, A []uint64) int {
	items := make([]Item, 0, n*n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			items = append(items, Item{i, j, A[j] ^ A[i]})
		}
	}

	sort.Sort(Items(items))

	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for _, item := range items {
		dp[item.second] = max(dp[item.second], dp[item.first]+1)
	}

	var res int

	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Item struct {
	first  int
	second int
	third  uint64
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	return items[i].third < items[j].third
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

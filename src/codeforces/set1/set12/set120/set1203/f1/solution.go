package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
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

func process(reader *bufio.Reader) bool {
	n, r := readTwoNums(reader)
	tasks := make([][]int, n)
	for i := range n {
		tasks[i] = readNNums(reader, 2)
	}
	return solve(r, tasks)
}

const inf = 1 << 60

func solve(r int, tasks [][]int) bool {
	n := len(tasks)

	var arr []pair

	for i := 0; i < n; i++ {
		if tasks[i][1] >= 0 {
			arr = append(arr, pair{tasks[i][0], tasks[i][1]})
		}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return a.first - b.first
	})

	for _, cur := range arr {
		if cur.first > r {
			return false
		}
		r += cur.second
	}

	arr = arr[:0]
	for i := 0; i < n; i++ {
		if tasks[i][1] < 0 {
			arr = append(arr, pair{tasks[i][0], tasks[i][1]})
		}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return (b.first + b.second) - (a.first + a.second)
	})

	for _, cur := range arr {
		if r < cur.first {
			return false
		}
		r += cur.second
	}

	return r >= 0
}

type pair struct {
	first  int
	second int
}

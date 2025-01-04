package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _ := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d ", cur))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) (res []int, pairs [][]int) {
	n := readNum(reader)
	pairs = make([][]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = readNNums(reader, 2)
	}
	res = solve(pairs)
	return
}

func solve(pairs [][]int) []int {
	n := len(pairs)
	f := func(check func([]int) bool, sorter func(a, b []int) int) []int {
		var arr [][]int
		for i := range n {
			if check(pairs[i]) {
				arr = append(arr, []int{pairs[i][0], pairs[i][1], i})
			}
		}
		slices.SortFunc(arr, sorter)
		res := make([]int, len(arr))
		for i, cur := range arr {
			res[i] = cur[2] + 1
		}
		return res
	}

	a := f(func(cur []int) bool {
		return cur[0] > cur[1]
	}, func(a, b []int) int {
		return a[0] - b[0]
	})

	b := f(func(cur []int) bool {
		return cur[0] < cur[1]
	}, func(a, b []int) int {
		return b[0] - a[0]
	})

	if len(a) >= len(b) {
		return a
	}
	return b
}

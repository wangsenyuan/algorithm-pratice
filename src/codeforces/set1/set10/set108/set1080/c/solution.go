package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		a := readNNums(reader, 4)
		b := readNNums(reader, 4)
		res := solve(n, m, a, b)
		buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
	}
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

func solve(n int, m int, a []int, b []int) []int {
	res := find([]int{1, 1, n, m})

	if b[0] <= a[2] && a[0] <= b[2] && b[1] <= a[3] && a[1] <= b[3] {
		if b[0] <= a[0] && b[1] <= a[1] && a[2] <= b[2] && a[3] <= b[3] {
			// 完全被包围了
			tb := find(b)
			res[0] -= tb[0]
			res[1] -= tb[1]
			res[1] += area(b)
		} else {
			c := []int{max(a[0], b[0]), max(a[1], b[1]), min(a[2], b[2]), min(a[3], b[3])}
			ta := find(a)
			tc := find(c)
			ta[0] -= tc[0]
			ta[1] -= tc[1]
			res[0] -= ta[0]
			res[1] -= ta[1]
			res[0] += area(a) - area(c)
			tb := find(b)
			res[0] -= tb[0]
			res[1] -= tb[1]
			res[1] += area(b)
		}
	} else {
		ta := find(a)
		tb := find(b)
		res[0] -= ta[0]
		res[1] -= ta[1]
		res[0] += area(a)
		res[0] -= tb[0]
		res[1] -= tb[1]
		res[1] += area(b)
	}

	return res
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

func area(x []int) int {
	dx := x[2] - x[0] + 1
	dy := x[3] - x[1] + 1
	return dx * dy
}

func find(x []int) []int {
	dx := x[2] - x[0] + 1
	dy := x[3] - x[1] + 1
	s := dx * dy
	if dx&1 == 0 || dy&1 == 0 {
		// 有一个维度是偶数，那么黑/白个一半
		return []int{s / 2, s / 2}
	}
	first := (x[0] & 1) ^ (x[1] & 1)
	if first == 0 {
		return []int{s/2 + 1, s / 2}
	}
	return []int{s / 2, s/2 + 1}
}

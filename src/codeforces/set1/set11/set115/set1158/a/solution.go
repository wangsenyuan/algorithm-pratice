package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	return solve(a, b)
}

func solve(a []int, b []int) int {
	x := slices.Max(a)
	y := slices.Min(b)
	if y < x {
		return -1
	}
	// y >= x
	sort.Ints(a)
	sort.Ints(b)
	n := len(a)
	m := len(b)
	p := m
	var res int
	for i := n - 1; i >= 0; i-- {
		// 每个女孩子至少要a[i]个糖果
		res += m * a[i]
		// 必须保证i送出去了一个a[i]
		if p == 0 {
			// 女孩子的条件都满足了，男孩子全部都送a[i]
			continue
		}
		if p == m {
			for p > 1 {
				res += b[p-1] - a[i]
				p--
			}
			if b[0] == a[i] {
				p--
			}
		} else {
			// p == 1
			res += b[p-1] - a[i]
			p--
		}
	}

	return res
}

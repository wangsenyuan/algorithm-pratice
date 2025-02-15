package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	buf.WriteTo(os.Stdout)
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
	n, m, k := readThreeNums(reader)
	h := readNNums(reader, n)
	x := readNNums(reader, n)
	return solve(h, x, m, k)
}

type pair struct {
	first  int
	second int
}

func solve(H []int, X []int, m int, k int) int {
	n := len(H)
	// open := make([]int, n)
	// end := make([]int, n)
	// 最多2*n个位置
	on := make([]int, 2*n)
	off := make([]int, 2*n)
	open := make([]int, n)
	end := make([]int, n)

	check := func(op int) bool {
		// 对于每个敌人计算一个能够把它消灭的区间
		clear(on)
		clear(off)

		var arr []int
		for i := range n {
			h, x := H[i], X[i]
			// 假设放在p处
			//( m - (p - x)) * op >= h
			// m + x - p >= h / op
			// p <= m + x - h / op
			end[i] = m + x - (h+op-1)/op
			// (m - (x - p)) * op >= h
			// m - x + p = h / op
			// p = h / op + x - m
			open[i] = -m + x + (h+op-1)/op
			if open[i] <= x && x <= end[i] {
				// 只保留有效的状态
				arr = append(arr, open[i])
				arr = append(arr, end[i])
			}
		}
		sort.Ints(arr)
		var it int
		for i := 1; i <= len(arr); i++ {
			if i == len(arr) || arr[i] > arr[i-1] {
				arr[it] = arr[i-1]
				it++
			}
		}
		arr = arr[:it]
		for i := range n {
			if open[i] <= X[i] && X[i] <= end[i] {
				j := sort.SearchInts(arr, open[i])
				on[j]++
				j = sort.SearchInts(arr, end[i])
				off[j]++
			}
		}
		var active int
		for i := range 2 * n {
			active += on[i]
			if active >= k {
				return true
			}
			active -= off[i]
		}

		return false
	}

	// 什么时候不行呢？
	// m太小的时候，是不行的， 始终干不掉k个
	l, r := 1, 1<<30
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if r == 1<<30 {
		return -1
	}
	return r
}

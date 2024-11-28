package main

import (
	"bufio"
	"fmt"
	"os"
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
	w, l := readTwoNums(reader)
	a := readNNums(reader, w-1)
	return solve(w, l, a)
}

func solve(w int, k int, a []int) int {
	n := len(a)
	// n == w - 1
	arr := make([]int, n+1)

	prev := make([]int, w+1)
	next := make([]int, w+1)

	remove := func(p int) {
		l, r := prev[p], next[p]
		next[l] = r
		prev[r] = l
	}

	use := func(p int, v int) int {
		// 能否把v个frog，安全放置在p的前面
		for v > 0 && p > 0 {
			if arr[p]+v <= a[p-1] {
				arr[p] += v
				v = 0
				return p
			}
			v -= a[p-1] - arr[p]
			arr[p] = a[p-1]
			q := prev[p]
			remove(p)
			p = q
		}

		return p
	}

	check := func(frogs int) bool {
		clear(arr)
		// fronts能否到达对岸
		for i := 1; i <= w; i++ {
			prev[i] = i - 1
			next[i] = i + 1
		}
		// 有一部分必须留在岸上
		if use(k, frogs) == 0 {
			return false
		}

		for i := 1; i+k < w; i++ {
			if use(i+k, arr[i]) <= i {
				return false
			}
		}
		return true
	}

	var sum int

	for _, x := range a {
		sum += x
	}

	l, r := 0, sum+1

	for l < r {
		mid := (l + r) / 2
		if !check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r - 1
}

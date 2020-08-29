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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	P := readNNums(reader, n)
	Q := readNNums(reader, n)

	res := solve(n, P, Q)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", res[i])
	}
}

func solve(n int, P []int, Q []int) []int {
	arr := make([]int, 4*n)
	add := make([]int, 4*n)

	update := func(p int, x int) {
		add[p] += x
		arr[p] += x
	}

	push := func(p int) {
		if add[p] != 0 {
			update(p*2+1, add[p])
			update(p*2+2, add[p])
			add[p] = 0
		}
	}

	var updateRange func(i int, l, r int, tl, tr int, x int)

	updateRange = func(i int, l, r int, tl, tr int, x int) {
		if tl >= r || tr <= l {
			return
		}
		if tl >= l && tr <= r {
			update(i, x)
			return
		}
		push(i)

		mid := (tl + tr) / 2

		updateRange(2*i+1, l, r, tl, mid, x)
		updateRange(2*i+2, l, r, mid, tr, x)

		arr[i] = min(arr[2*i+1], arr[2*i+2])
	}

	rev := make([]int, n+1)
	for i := 0; i < n; i++ {
		rev[P[i]] = i
	}

	decrease := func(x int) bool {
		updateRange(0, 0, rev[x]+1, 0, n, -1)
		if arr[0] >= 0 {
			return true
		}
		updateRange(0, 0, rev[x]+1, 0, n, 1)
		return false
	}

	ans := make([]int, n)
	x := n

	for i := 0; i < n; i++ {
		for decrease(x) {
			x--
		}
		ans[i] = x
		updateRange(0, 0, Q[i], 0, n, 1)
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)
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

func solve(a []int) int {
	n := len(a)

	// 如果存在一个区间[l...r]包含了数字 [1...num-1] 且这个num也在这个区间里
	// l和r这个区间可以用seg-tree 解决，但怎么知道这个区间里，包含了[1...num]呢？
	//
	ok := make([]bool, n+3)
	tr := NewTree(n + 2)
	last := make([]int, n+3)

	for i := 0; i < n; i++ {
		num := a[i]
		if num == 1 {
			ok[2] = true
		} else {
			ok[1] = true
			j := tr.get(1, num)
			if j > last[num] {
				ok[num] = true
			}
			last[num] = i + 1
		}
		tr.update(num, i+1)
	}

	for num := 2; num <= n+1; num++ {
		j := tr.get(1, num)
		if j > last[num] {
			ok[num] = true
		}
	}
	mex := 1
	for ok[mex] {
		mex++
	}

	return mex
}

type tree struct {
	arr []int
	n   int
}

func NewTree(n int) *tree {
	arr := make([]int, 2*n)
	return &tree{arr, n}
}

func (t *tree) update(i int, v int) {
	i += t.n
	t.arr[i] = v
	for i > 1 {
		t.arr[i>>1] = min(t.arr[i], t.arr[i^1])
		i >>= 1
	}
}

func (t *tree) get(l int, r int) int {
	l += t.n
	r += t.n
	res := 1 << 30
	for l < r {
		if l&1 == 1 {
			res = min(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

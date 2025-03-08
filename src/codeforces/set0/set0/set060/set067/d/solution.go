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
	n := readNum(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	return solve(a, b)
}

func solve(a []int, b []int) int {
	n := len(a)
	pos := make([]int, n)
	for i := range n {
		a[i]--
		b[i]--
		pos[b[i]] = i
	}
	dp := NewSegTree(n)

	for _, i := range a {
		j := pos[i]
		tmp := dp.Query(j+1, n)
		dp.Update(j, tmp+1)
	}

	return dp.Query(0, n)
}

type SegTree struct {
	arr  []int
	size int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	size := n
	return &SegTree{arr, size}
}

func (st *SegTree) Update(pos int, val int) {
	arr := st.arr
	n := st.size
	arr[pos+n] = val
	pos += n
	for i := pos; i > 1; i >>= 1 {
		arr[i>>1] = max(arr[i], arr[i^1])
	}
}

func (st *SegTree) Query(left, right int) int {
	arr := st.arr
	n := st.size
	left += n
	right += n
	var res int
	for left < right {
		if left&1 == 1 {
			res = max(res, arr[left])
			left++
		}
		if right&1 == 1 {
			right--
			res = max(res, arr[right])
		}
		left >>= 1
		right >>= 1
	}
	return res
}

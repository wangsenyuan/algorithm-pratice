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
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(n, m, a)
}

func solve(n int, m int, a []int) int {
	// f(sub_arr) = (a[i] + ... + a[j]) % m
	// = (pref[j] - pref[i-1]) % m

	cnt := NewSegTree(m)
	pref := NewSegTree(m)

	var sum int
	var res int

	cnt.Add(0, 1)
	pref.Add(0, 0)

	for _, x := range a {
		sum = (sum + x) % m
		// 前面比sum小的
		res += cnt.Query(0, sum+1)*sum - pref.Query(0, sum+1)
		// 前面比sum大的
		res += -pref.Query(sum+1, m) + cnt.Query(sum+1, m)*(sum+m)
		cnt.Add(sum, 1)
		pref.Add(sum, sum)
	}

	return res
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (seg *SegTree) Add(p int, v int) {
	p += seg.sz
	seg.arr[p] += v
	for p > 1 {
		seg.arr[p>>1] = seg.arr[p] + seg.arr[p^1]
		p >>= 1
	}
}

func (seg *SegTree) Query(l int, r int) int {
	l += seg.sz
	r += seg.sz
	var res int
	for l < r {
		if l&1 == 1 {
			res += seg.arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += seg.arr[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func bruteForce(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	var res int
	for i := 0; i < n; i++ {
		var sum int
		for j := i; j < n; j++ {
			sum += a[j]
			sum %= m
			res += sum
		}
	}
	return res
}

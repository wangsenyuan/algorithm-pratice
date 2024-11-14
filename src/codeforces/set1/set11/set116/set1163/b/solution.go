package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

const inf = 1 << 30

func solve(arr []int) int {
	x := slices.Max(arr)
	freq := make([]int, x+1)
	n := len(arr)

	xs := NewMultiSet(n + 1)

	add := func(num int) {
		if freq[num] > 0 {
			xs.Add(freq[num], -1)
		}
		freq[num]++
		xs.Add(freq[num], 1)
	}

	check := func(x int) bool {
		// num := arr[x]
		// 最大的freq，也应该是所有的freq
		a := xs.val[0]
		b := xs.Get(a)
		return a*b == x
	}

	var res int
	for i := 0; i < n; i++ {
		if i == 0 || check(i) {
			res = i + 1
		}
		add(arr[i])

		if res < i+1 {
			// 说明删除当前的num没有用
			x := freq[arr[i]]
			y := xs.Get(x)
			z := xs.Get(x + 1)
			if z == 1 && x*y+x+1 == i+1 {
				res = i + 1
			}
		}

		if res < i+1 {
			x := freq[arr[i]]
			y := xs.Get(x)
			// 删掉1个情况
			z := xs.Get(1)
			if z == 1 && x*y+1 == i+1 {
				res = i + 1
			}
		}
	}

	return res
}

type MultiSet struct {
	cnt []int
	val []int
	sz  int
}

func NewMultiSet(n int) *MultiSet {
	cnt := make([]int, 4*n)
	val := make([]int, 4*n)

	sz := n
	return &MultiSet{cnt, val, sz}
}

func (tr *MultiSet) Add(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.cnt[i] += v
			if tr.cnt[i] > 0 {
				tr.val[i] = l
			} else {
				tr.val[i] = -1
			}
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		tr.val[i] = max(tr.val[2*i+1], tr.val[2*i+2])
	}

	loop(0, 0, tr.sz-1)
}

func (tr *MultiSet) Get(p int) int {
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if l == r {
			return tr.cnt[i]
		}
		mid := (l + r) / 2
		if p <= mid {
			return loop(2*i+1, l, mid)
		}
		return loop(2*i+2, mid+1, r)
	}

	return loop(0, 0, tr.sz-1)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	l := readNNums(reader, n)
	r := readNNums(reader, n)

	res := solve(n, l, r)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	s := fmt.Sprintf("%v", res)

	fmt.Println(s[1 : len(s)-1])
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, l []int, r []int) []int {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		if l[i] > i {
			return nil
		}
		if r[n-1-i] > i {
			return nil
		}
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		return l[id[i]]+r[id[i]] < l[id[j]]+r[id[j]] || l[id[i]]+r[id[i]] == l[id[j]]+r[id[j]] && id[i] < id[j]
	})

	tr := NewSegTree(n, 0, plus)

	res := make([]int, n)
	cur := n
	for i := 0; i < n; {
		expect := l[id[i]] + r[id[i]]
		// 比a[id[j]]大的数有这么多， 只有j个数
		if expect > i {
			return nil
		}
		j := i
		for i < n && l[id[i]]+r[id[i]] == expect {
			res[id[i]] = cur
			if tr.Get(0, id[i]) != l[id[i]] {
				return nil
			}
			i++
		}

		for k := i - 1; k >= j; k-- {
			if tr.Get(id[k], n) != r[id[k]] {
				return nil
			}
		}

		for j < i {
			tr.Update(id[j], 1)
			j++
		}

		cur--
	}

	return res
}

func plus(a, b int) int {
	return a + b
}

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

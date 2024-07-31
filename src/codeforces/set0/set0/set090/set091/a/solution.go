package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s1 := readString(reader)
	s2 := readString(reader)
	res := solve(s1, s2)

	fmt.Println(res)
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(s1 string, s2 string) int {

	n := len(s1)

	trs := make([]*tree, 26)
	for i := 0; i < 26; i++ {
		trs[i] = NewTree(n)
	}

	for i := 0; i < n; i++ {
		x := int(s1[i] - 'a')
		trs[x].set(i, i)
	}

	res := 1

	for i, j := 0, 0; i < len(s2); i++ {
		x := int(s2[i] - 'a')
		k := trs[x].get(j, n)
		if k == inf {
			// must repeat
			res++
			k = trs[x].get(0, j+1)
		}
		if k == inf {
			// not found
			return -1
		}
		// advance
		j = (k + 1) % n
		if j == 0 && i+1 < len(s2) {
			res++
		}
	}

	return res
}

const inf = 1 << 30

type tree struct {
	arr []int
	n   int
}

func NewTree(n int) *tree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = inf
	}
	return &tree{arr, n}
}

func (tr *tree) get(l int, r int) int {
	l += tr.n
	r += tr.n
	res := inf
	for l < r {
		if l&1 == 1 {
			res = min(res, tr.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, tr.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func (tr *tree) set(p int, v int) {
	p += tr.n
	tr.arr[p] = v
	for p > 1 {
		tr.arr[p>>1] = min(tr.arr[p], tr.arr[p^1])
		p >>= 1
	}
}

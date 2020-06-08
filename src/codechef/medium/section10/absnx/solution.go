package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		fmt.Println(solve(n, A))
	}

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

func solve(n int, A []int) int64 {
	ans := find(n, A)
	reverse(A)
	ans += find(n, A)

	var cnt int64 = 1

	for i := 1; i <= n; i++ {
		if i == n || A[i] != A[i-1] {
			ans -= cnt * (cnt + 1) / 2
			cnt = 1
		} else {
			cnt++
		}
	}

	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func find(n int, A []int) int64 {
	tree := NewFenwick(n)

	st := make([]int, n)
	var p int

	li := make([][]int, n)
	for i := 0; i < n; i++ {
		li[i] = make([]int, 0, 10)
	}

	for i := 0; i < n; i++ {
		for p > 0 && A[st[p-1]] <= A[i] {
			p--
		}

		if p == 0 {
			tree.Update(i, 1)
		} else {
			li[st[p-1]] = append(li[st[p-1]], i)
		}
		st[p] = i
		p++
	}

	nxt := make([]int, n)
	p = 0

	for i := n - 1; i >= 0; i-- {
		for p > 0 && A[st[p-1]] >= A[i] {
			p--
		}
		if p == 0 {
			nxt[i] = n
		} else {
			nxt[i] = st[p-1]
		}
		st[p] = i
		p++
	}

	var res int64

	for i := 0; i < n; i++ {
		for _, j := range li[i] {
			tree.Update(j, 1)
		}
		res += tree.GetRange(i, nxt[i]-1)
		tree.Update(i, -1)
	}
	return res
}

type Fenwick struct {
	arr []int
	n   int
}

func NewFenwick(n int) Fenwick {
	arr := make([]int, n+1)
	return Fenwick{arr, n}
}

func (tree *Fenwick) Update(p int, v int) {
	arr := tree.arr
	n := tree.n

	p++

	for p <= n {
		arr[p] += v
		p += p & -p
	}
}

func (tree Fenwick) Get(p int) int64 {
	arr := tree.arr
	p++
	var res int64
	for p > 0 {
		res += int64(arr[p])
		p -= p & -p
	}
	return res
}

func (tree Fenwick) GetRange(l, r int) int64 {
	res := tree.Get(r)
	if l > 0 {
		res -= tree.Get(l - 1)
	}
	return res
}

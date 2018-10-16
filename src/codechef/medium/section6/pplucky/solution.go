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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)

		scanner.Scan()
		s := scanner.Text()
		fmt.Println(solve(n, s))
	}
}

func solve(n int, s string) int64 {
	st := make([]int, n)
	var p int

	L := make([]int, n)
	R := make([]int, n)
	for i := 0; i < n; i++ {
		R[i] = -1
	}
	sg := NewSegTree(n)

	for i := 0; i < n; i++ {
		if s[i] == '4' {
			st[p] = i
			p++
		} else if p > 0 {
			pos := st[p-1]
			p--
			L[pos] = sg.Query(pos, i) + 1
			R[i] = pos
			sg.Update(pos, L[pos])
		}
	}

	var res int64

	bit := NewBIT(n)

	C := make([]int64, n)

	for i := 0; i < n; i++ {
		if s[i] == '4' {
			C[i] = bit.Query(L[i] - 1)
		} else if R[i] >= 0 {
			res += int64(R[i]) + 1 - 2*C[R[i]]
			bit.Update(L[R[i]], 1)
		}
	}

	return res
}

type BIT []int64

func NewBIT(n int) BIT {
	arr := make([]int64, n+1)
	return BIT(arr)
}

func (bit BIT) Update(pos int, val int64) {
	n := len(bit) - 1
	pos++
	for pos <= n {
		bit[pos] += val
		pos += pos & -pos
	}
}

func (bit BIT) Query(pos int) int64 {
	var res int64
	pos++
	for pos > 0 {
		res += bit[pos]
		pos -= pos & -pos
	}
	return res
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

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(A []int) int {
	n := len(A)
	// 先找出 A[i] 作为最大值的范围 [l...r],
	// 然后找出 [l...i-1]中最大值的贡献，[i+1, r]的贡献
	// 考虑 s[l...i-1]
	// f(l...i-1) = A[x] * (rx - lx) (A[x]是最大值，且[lx..rx]是它的范围 )
	//  有rx <= i - 1,因为 A[i] >= A[x]
	//  l <= lx, 因为l是A[i]的左边距，说明存在 i == 0 || A[i-1] > A[i], A[x] < A[i]
	// 所以只需要计算 pref_sum[i-1] - pref_sum[l-1] 即可
	L := make([]int, n)
	R := make([]int, n)

	stack := make([]int, n)
	var p int

	for i := 0; i < n; i++ {
		for p > 0 && A[stack[p-1]] <= A[i] {
			p--
		}
		if p > 0 {
			L[i] = stack[p-1]
		} else {
			L[i] = -1
		}
		stack[p] = i
		p++
	}

	res := make(map[int]int)

	p = 0

	for i := n - 1; i >= 0; i-- {
		for p > 0 && A[stack[p-1]] < A[i] {
			p--
		}
		for p > 0 && A[stack[p-1]] == A[i] {
			res[0]++
			p--
		}

		if p > 0 {
			R[i] = stack[p-1]
		} else {
			R[i] = n
		}
		stack[p] = i
		p++
	}

	for i := 0; i < n; i++ {
		l := L[i]
		if l >= 0 && l < i {
			res[A[l]-A[i]]++
		}
		r := R[i]
		if r < n && i < r {
			res[A[r]-A[i]]++
		}
	}
	return len(res)
}

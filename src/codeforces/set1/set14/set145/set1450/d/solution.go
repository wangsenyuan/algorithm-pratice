package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(res)
		buf.WriteByte('\n')
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

func solve(A []int) string {
	n := len(A)
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = '0'
	}

	if min_of(A) != 1 {
		return string(res)
	}

	res[0] = '1'

	cnt := make([]int, n+2)
	for i := range A {
		A[i]--
		cnt[A[i]]++
		if cnt[A[i]] > 1 {
			res[0] = '0'
		}
	}
	l, r := 0, n-1
	for i := n - 1; i > 0; i-- {
		res[i] = '1'
		nxt := n - i - 1
		cnt[nxt]--
		if cnt[nxt] == 0 && (A[l] == nxt || A[r] == nxt) && cnt[nxt+1] > 0 {
			if A[l] == nxt {
				l++
			} else {
				r--
			}
			continue
		}
		break
	}

	return string(res)
}

func min_of(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		res = min(res, arr[i])
	}
	return res
}

func solve1(A []int) string {
	n := len(A)

	L := make([]int, n)

	stack := make([]int, n)
	var p int
	for i := 0; i < n; i++ {
		for p > 0 && A[stack[p-1]] >= A[i] {
			p--
		}
		if p == 0 {
			L[i] = -1
		} else {
			L[i] = stack[p-1]
		}
		stack[p] = i
		p++
	}

	R := make([]int, n)
	p = 0

	for i := n - 1; i >= 0; i-- {
		for p > 0 && A[stack[p-1]] >= A[i] {
			p--
		}
		if p == 0 {
			R[i] = n
		} else {
			R[i] = stack[p-1]
		}
		stack[p] = i
		p++
	}

	type Pair struct {
		first  int
		second int
	}

	B := make([]Pair, n)
	for i := 0; i < n; i++ {
		B[i] = Pair{A[i], i}
	}
	sort.Slice(B, func(i, j int) bool {
		return B[i].first < B[j].first
	})
	cnt := make([]int, n+2)
	for i, x := 0, 1; x <= n; x++ {
		var ln int
		for i < n && B[i].first == x {
			id := B[i].second
			ln = max(ln, R[id]-L[id]-1)
			i++
		}
		if ln == 0 {
			break
		}
		cnt[1]++
		cnt[ln+1]--
	}

	for i := 1; i <= n; i++ {
		cnt[i] += cnt[i-1]
	}
	res := make([]byte, n)
	for i := 1; i <= n; i++ {
		if cnt[i] == (n - i + 1) {
			res[i-1] = '1'
		} else {
			res[i-1] = '0'
		}
	}
	return string(res)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

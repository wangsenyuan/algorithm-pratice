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
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		S := make([][]int, m)
		for i := 0; i < m; i++ {
			S[i] = readNNums(reader, 2)
		}
		res := solve(A, B, S)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(A []int, B []int, S [][]int) bool {
	n := len(A)
	m := len(S)
	sum := make([]int64, n+1)

	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + int64(A[i]-B[i])
	}

	N := n + 1

	evt := make([][]int, N)

	addEvent := func(p int, i int) {
		if len(evt[p]) == 0 {
			evt[p] = make([]int, 0, 1)
		}
		evt[p] = append(evt[p], i)
	}

	for i, s := range S {
		l, r := s[0], s[1]
		addEvent(l-1, i)
		addEvent(r, i)
	}

	cnt := make([]int, m)

	st := make([]int, m)
	var p int

	push := func(x int) {
		st[p] = x
		p++
	}

	pop := func() int {
		res := st[p-1]
		p--
		return res
	}

	set := make([]int, N)
	for i := 0; i < N; i++ {
		set[i] = i
	}

	for i := 0; i <= n; i++ {
		if sum[i] == 0 {
			for _, j := range evt[i] {
				cnt[j]++
				if cnt[j] == 2 {
					push(j)
				}
			}
		}
		if i < n && sum[i] == 0 && sum[i+1] == sum[i] {
			set[i] = i + 1
		}
	}

	var find func(i int) int

	find = func(i int) int {
		if set[i] != i {
			set[i] = find(set[i])
		}
		return set[i]
	}

	for p > 0 {
		v := pop()
		i := find(S[v][0] - 1)
		for i < S[v][1] {
			set[i] = find(i + 1)
			for _, x := range evt[i+1] {
				if sum[i+1] != 0 {
					cnt[x]++
					if cnt[x] == 2 {
						push(x)
					}
				}
			}
			i = find(i)
		}
	}

	return find(0) == n
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

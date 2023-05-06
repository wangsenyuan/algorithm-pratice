package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, q, mod := readThreeNums(reader)
	A := readNNums(reader, n)
	B := readNNums(reader, n)
	Q := make([]string, q)
	for i := 0; i < q; i++ {
		Q[i] = readString(reader)
	}
	res := solve(A, B, mod, Q)
	var buf bytes.Buffer

	for _, ok := range res {
		if ok {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func add(a, b int, m int) int {
	a += b
	if a >= m {
		a -= m
	}
	return a
}

func sub(a, b int, m int) int {
	return add(a, m-b, m)
}

func solve(A []int, B []int, mod int, Q []string) []bool {
	n := len(A)
	F := make([]int, n+1)
	F[0] = 1
	F[1] = 1
	for i := 2; i <= n; i++ {
		F[i] = add(F[i-1], F[i-2], mod)
	}
	C := make([]int, n)
	for i := 0; i < n; i++ {
		C[i] = sub(A[i], B[i], mod)
	}

	D := make([]int, n)
	D[0] = C[0]
	if n > 1 {
		D[1] = sub(C[1], C[0], mod)
	}
	for i := 2; i < n; i++ {
		D[i] = sub(C[i], add(C[i-1], C[i-2], mod), mod)
	}
	var notZero int
	for i := 0; i < n; i++ {
		if D[i] != 0 {
			notZero++
		}
	}

	upd := func(ind int, delta int) {
		if D[ind] != 0 {
			notZero--
		}
		D[ind] = add(D[ind], delta, mod)
		if D[ind] != 0 {
			notZero++
		}
	}

	res := make([]bool, len(Q))

	for i, cur := range Q {
		buf := []byte(cur[2:])
		var l, r int
		pos := readInt(buf, 0, &l)
		readInt(buf, pos+1, &r)
		if cur[0] == 'A' {
			upd(l-1, 1)
			if r != n {
				upd(r, mod-F[r-l+1])
			}
			if r < n-1 {
				upd(r+1, mod-F[r-l])
			}
		} else {
			upd(l-1, mod-1)
			if r != n {
				upd(r, F[r-l+1])
			}
			if r < n-1 {
				upd(r+1, F[r-l])
			}
		}
		res[i] = notZero == 0
	}

	return res
}

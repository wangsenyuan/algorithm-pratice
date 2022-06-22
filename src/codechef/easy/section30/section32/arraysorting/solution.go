package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		P := readNNums(reader, n)
		res := solve(P)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, cur := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
		}
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(P []int) [][]int {
	n := len(P)
	var res [][]int

	for i := n; i > 0; {
		if P[i-1] == i {
			i--
			continue
		}
		// P[i-1] != i, and P[i-1] < i
		for j := 1; j < i; j++ {
			if P[j-1] == P[i-1]+1 {
				res = append(res, []int{j, i})
				P[j-1], P[i-1] = P[i-1], P[j-1]
			}
		}
	}

	return res
}

func solve1(P []int) [][]int {
	n := len(P)
	var L int
	for L+1 < n && P[L] < P[L+1] {
		L++
	}

	var res [][]int

	process := func(L, R int) {

		for r := R; r > L; {
			if P[L] < P[r] {
				r--
				continue
			}
			i := search(L+1, func(i int) bool {
				return P[i] > P[r]
			})
			res = append(res, []int{i + 1, r + 1})
			P[i], P[r] = P[r], P[i]
		}
	}

	for L < n-1 {
		R := L + 1
		for R+1 < n && P[R] < P[R+1] {
			R++
		}
		// P[R] > P[R+1] or R == n - 1
		process(L, R)
		L = R
	}

	return res
}

func search(n int, fn func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if fn(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

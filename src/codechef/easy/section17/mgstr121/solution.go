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

		S, _ := reader.ReadString('\n')

		solver := NewSolver(S[:n])
		q := readNum(reader)

		for q > 0 {
			q--
			L, R, K := readThreeNums(reader)
			res := solver.solve(L, R, K)
			if len(res) == 0 {
				buf.WriteString("NO\n")
			} else {
				buf.WriteString("YES\n")
				buf.WriteString(res)
				buf.WriteByte('\n')
			}
		}

	}
	fmt.Print(buf.String())
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

type Solver struct {
	cnt [][]int
}

func NewSolver(s string) *Solver {
	n := len(s)
	cnt := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		cnt[i] = make([]int, 26)
	}

	for i := 0; i < n; i++ {
		copy(cnt[i+1], cnt[i])
		cnt[i+1][int(s[i]-'A')]++
	}

	return &Solver{cnt}
}

func (s *Solver) solve(L, R, K int) string {
	if K > R-L+1 {
		return ""
	}

	cnt := s.cnt
	for i := 0; i < 26; i++ {
		if cnt[R][i]-cnt[L-1][i] > K {
			return ""
		}
	}

	// K >= max(L[i])
	buf := make([]byte, R-L+1)
	var pos int
	var max int

	for i := 25; i >= 0; i-- {
		tmp := cnt[R][i] - cnt[L-1][i]
		if tmp > max {
			max = tmp
		}
		for tmp > 0 {
			buf[pos] = byte('A' + i)
			pos++
			tmp--
		}
	}

	if max == K {
		return string(buf)
	}

	reverse(buf[pos-K:])

	return string(buf)
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

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
		S := make([][]int, n)
		for i := 0; i < n; i++ {
			S[i] = readNNums(reader, 2)
		}
		res := solve(S)

		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d %d %d\n", res[i][0], res[i][1], res[i][2]))
		}
	}

	fmt.Print(buf.String())
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

func solve(S [][]int) [][]int {

	n := len(S)

	has := make([][]bool, n)

	for i := 0; i < n; i++ {
		has[i] = make([]bool, n)
	}

	for _, s := range S {
		a, b := s[0], s[1]
		a--
		b--
		has[a][b] = true
	}

	var dfs func(l, r int)

	var res [][]int

	dfs = func(l, r int) {
		for d := l; d <= r; d++ {
			ok := true
			if d-1 >= l {
				ok = ok && has[l-1][d-2]
			}
			if d+1 <= r {
				ok = ok && has[d][r-1]
			}
			if ok {
				res = append(res, []int{l, r, d})
				if d-1 >= l {
					dfs(l, d-1)
				}
				if d+1 <= r {
					dfs(d+1, r)
				}
				break
			}
		}
	}

	dfs(1, n)
	return res
}

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

		line, _ := reader.ReadBytes('\n')

		var pos int
		var K int
		pos = readInt(line, 0, &K) + 1

		V := make([]int, K)

		for i := 0; i < K; i++ {
			pos = readInt(line, pos, &V[i]) + 1
		}

		L := readString(reader)

		S := readNum(reader)

		res := solve(K, V, L, S)

		if len(res) == 0 {
			res = "IMPOSSIBLE"
		}

		buf.WriteString(fmt.Sprintln(res))
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

func solve(K int, V []int, L string, S int) string {
	n := len(L)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, S+1)
	}
	can := make([]bool, S+1)
	can[0] = true
	dp[n][0] = true

	for i := n - 1; i >= 0; i-- {
		x := int(L[i] - 'a')
		v := V[x]
		for j := S - v; j >= 0; j-- {
			if can[j] {
				can[j+v] = true
				dp[i][j+v] = true
			}
		}
	}

	if !can[S] {
		return ""
	}

	var buf bytes.Buffer
	var p int
	for S > 0 && p < n {
		ii := -1
		for i := p; i < n; i++ {
			if dp[i][S] && (ii < 0 || L[i] < L[ii]) {
				ii = i
			}
		}
		buf.WriteByte(L[ii])
		S -= V[int(L[ii]-'a')]
		p = ii + 1
	}

	return buf.String()
}

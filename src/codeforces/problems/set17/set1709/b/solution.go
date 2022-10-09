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

	n, q := readTwoNums(reader)

	A := readNNums(reader, n)
	quests := make([][]int, q)

	for i := 0; i < q; i++ {
		quests[i] = readNNums(reader, 2)
	}

	res := solve(A, quests)

	for i := 0; i < q; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}

	fmt.Print(buf.String())
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

func solve(A []int, quests [][]int) []int64 {
	n := len(A)
	dp := make([]int64, n+1)
	for i := 1; i < n; i++ {
		dp[i+1] = dp[i] + int64(max(0, A[i-1]-A[i]))
	}

	fp := make([]int64, n+1)
	for i := n - 2; i >= 0; i-- {
		fp[i] = fp[i+1] + int64(max(0, A[i+1]-A[i]))
	}

	res := make([]int64, len(quests))

	for i, quest := range quests {
		s, t := quest[0], quest[1]
		if s == t {
			res[i] = 0
			continue
		}
		if s < t {
			res[i] = dp[t] - dp[s]
		} else {
			res[i] = fp[t-1] - fp[s-1]
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

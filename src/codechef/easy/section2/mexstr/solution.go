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
		s, _ := reader.ReadString('\n')
		res := solve(len(s)-1, s)
		buf.WriteString(res)
		buf.WriteByte('\n')
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

func solve(n int, s string) string {
	nr := make([]Pair, n)
	nr[n-1] = Pair{n, n}
	for i := n - 2; i >= 0; i-- {
		if s[i+1] == '0' {
			nr[i] = Pair{i + 1, nr[i+1].second}
		} else {
			nr[i] = Pair{nr[i+1].first, i + 1}
		}
	}

	if nr[0].first == n && s[0] != '0' {
		// all 1
		return "0"
	}
	if nr[0].second == n && s[0] != '1' {
		// all 0
		return "1"
	}
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		dp[i] = min(dp[nr[i].first], dp[nr[i].second]) + 1
	}
	var pivot int
	if s[0] != '1' {
		pivot = nr[0].second
	}

	nb := dp[pivot] + 1
	ans := make([]int, nb)
	ans[0] = 1
	for j := 1; j < nb; j++ {
		nz := nr[pivot].first
		if nz == n {
			pivot = nz
			continue
		}
		k := min(dp[nr[nz].first], dp[nr[nz].second])
		if k == nb-j-1 {
			ans[j] = 1
			pivot = nr[pivot].second
		} else {
			ans[j] = 0
			pivot = nz
		}
	}

	var buf bytes.Buffer

	for i := 0; i < len(ans); i++ {
		buf.WriteByte(byte('0' + ans[i]))
	}
	return buf.String()
}

type Pair struct {
	first, second int
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

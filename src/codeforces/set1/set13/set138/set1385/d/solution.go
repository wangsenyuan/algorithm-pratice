package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		s := readString(reader)[:n]
		res := solve(s)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const inf = 1 << 50

func solve(s string) int {
	n := len(s)

	pref := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		pref[i] = make([]int, 26)
	}

	for i := 0; i < n; i++ {
		copy(pref[i+1], pref[i])
		pref[i+1][int(s[i]-'a')]++
	}

	get := func(i int, w int, x int) int {
		ln := 1 << w
		cnt := pref[i+ln][x] - pref[i][x]
		return ln - cnt
	}

	k := bits.Len(uint(n))

	dp := make([][]int, n+1)
	fp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 26)
		fp[i] = make([]int, 26)
	}

	for i := 0; i < n; i++ {
		for x := 0; x < 26; x++ {
			dp[i][x] = 1
		}
		dp[i][int(s[i]-'a')] = 0
	}

	for w := 1; w < k; w++ {
		for i := 0; i+(1<<w) <= n; i += 1 << w {
			for x := 0; x < 25; x++ {
				u := get(i, w-1, x) + dp[i+(1<<(w-1))][x+1]
				v := dp[i][x+1] + get(i+1<<(w-1), w-1, x)
				fp[i][x] = min(u, v)
			}
		}
		dp, fp = fp, dp
	}

	return dp[0][0]
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		s, _ := reader.ReadString('\n')
		t, _ := reader.ReadString('\n')
		fmt.Println(solve(s[:n], t[:n]))
	}
}

const N = 2005
const INF = 1 << 30
var mem [][]int
var suf1 [][]int
var suf2 [][]int

func init() {
	mem = make([][]int, N)
	for i := 0; i < N; i++ {
		mem[i] = make([]int, N)
	}
	suf1 = make([][]int, 26)
	suf2 = make([][]int, 26)
	for i := 0; i < 26; i++ {
		suf1[i] = make([]int, N)
		suf2[i] = make([]int, N)
	}
}

func reset(n int) {
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			mem[i][j] = -1
		}
	}
	for i := 0; i < 26; i++ {
		for j := 0; j <= n + 1; j++ {
			suf1[i][j] = 0
			suf2[i][j] = 0
		}
	}
}


func solve(s, t string) int {
	n := len(s)
	reset(n)
	for i := n; i >= 1; i-- {
		for j := 0; j < 26; j++ {
			suf1[j][i] = suf1[j][i+1]
			suf2[j][i] = suf2[j][i+1]
		}
		suf1[ind(s[i-1])][i]++
		suf2[ind(t[i-1])][i]++
	}
	var dp func(i, j int) int

	dp = func(i, j int) int {
		if j == 0 {
			return 0
		}
		// i >= j
		if mem[i][j] >= 0 {
			return mem[i][j]
		}
		var ans = INF
		if i > 0 {
			ans = 1 + dp(i-1, j)
			if s[i-1] == t[j-1] {
				ans = min(ans, dp(i - 1, j - 1))
			}
		}
		k := ind(t[j-1])
		if suf1[k][i+1] > suf2[k][j+1] {
			ans = min(ans, dp(i, j - 1))
		}
		mem[i][j] = ans
		return ans
	}

	ans := dp(n, n)

	if ans >= INF {
		return -1
	}
	return ans
}

func ind(x byte) int {
	return int(x - 'a')
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
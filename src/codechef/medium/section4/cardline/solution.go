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
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(A, B)

		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const MAX_N = 2005

var dp [MAX_N][MAX_N]int

func init() {
	// dp[i][j] = maximum nodes u can take by breaking a cycle of length i to paths of lengths j+1 or j+2
	// dp[i][j] = max(j + dp[i-j - 1, j], j+1 + dp[i-j-2, j]
	for i := 2; i < MAX_N; i++ {
		for j := 0; j < i; j++ {
			if i-j-1 >= 0 && j+dp[i-j-1][j] > dp[i][j] {
				dp[i][j] = j + dp[i-j-1][j]
			}
			if i-j-2 >= 0 && j+1+dp[i-j-2][j] > dp[i][j] {
				dp[i][j] = j + 1 + dp[i-j-2][j]
			}
		}
	}
}

func solve(A []int, B []int) int {
	n := len(A)
	pa := make([]int, n+1)
	pb := make([]int, n+1)

	for i := 0; i < n; i++ {
		pa[A[i]] = i
		pb[B[i]] = i
	}

	cycles := make([]int, 0, n)
	var best int
	seen := make([]bool, n)
	for i := 0; i < n; i++ {
		if seen[i] {
			continue
		}
		p := i
		a := A[i]
		var sz int
		for {
			if seen[p] {
				break
			}
			seen[p] = true
			sz++
			p = pb[a]
			a = A[p]
		}
		cycles = append(cycles, sz)
		if sz == 1 {
			best++
		}
	}

	for l := 1; l <= n; l++ {
		var tot int
		for _, i := range cycles {
			tot += dp[i][l]
		}
		if best < tot {
			best = tot
		}
	}

	return best
}

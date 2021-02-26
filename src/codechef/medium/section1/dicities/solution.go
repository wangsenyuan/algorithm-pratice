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
		n, k := readTwoNums(reader)
		w := make([][]int, n)
		for i := 0; i < n; i++ {
			w[i] = readNNums(reader, n)
		}
		res := solve(n, k, w)
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

const INF = 1 << 30

func solve(N, K int, W [][]int) int {
	w := make([][]int, N)
	for i := 0; i < N; i++ {
		w[i] = make([]int, N)
	}

	//reverse NO, 现在要排除的是后K个节点
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			w[i][j] = W[N-1-i][N-1-j]
		}
	}

	tot := 1 << uint(N)

	cs := make([]int, tot)

	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			rem := (tot - 1) ^ (1 << uint(i)) ^ (1 << uint(j))

			for mask := rem; mask >= 0; mask = (mask - 1) & rem {
				cs[mask^(1<<uint(i))^(1<<uint(j))] += w[i][j]
				if mask == 0 {
					break
				}
			}
		}
	}

	extra := N - K
	dp := make([]int, 1<<uint(N))

	terminalMask := 1 << uint(N-1)

	for i := K - 2; i >= 0; i-- {
		maskLeft := 1 << uint(extra+i)
		maskRight := terminalMask
		terminalMask ^= maskLeft
		var stop int
		if i == 0 {
			stop = 1<<uint(extra) - 1
		}

		for mask := 1<<uint(extra) - 1; mask >= stop; mask-- {
			wt := cs[mask^terminalMask]
			tmp := INF
			for subMask := mask; ; subMask = (subMask - 1) & mask {
				wtLeft := cs[maskLeft^subMask]
				wtRight := cs[maskRight^subMask^mask]
				cut := wt - wtLeft - wtRight
				tmp = min(tmp, cut+dp[mask^subMask])
				if subMask == 0 {
					break
				}
			}
			dp[mask] = tmp
		}
	}

	return dp[1<<uint(extra)-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	g := make([]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g[u] |= 1 << v
		g[v] |= 1 << u
	}
	N := 1 << n

	dp := make([][]int, N)
	for i := range N {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[1<<i][i] = 1
	}

	for mask := 1; mask < N; mask++ {
		f := lowestSetBit(mask)
		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 1 && i != f {
				state := mask ^ (1 << i)
				for j := 0; j < n; j++ {
					if (g[i]>>j)&1 == 1 {
						dp[mask][i] += dp[state][j]
					}
				}
			}
		}
	}

	var ans int

	for state := 1; state < N; state++ {
		cnt := bits.OnesCount(uint(state))
		if cnt < 3 {
			continue
		}
		f := lowestSetBit(state)
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 && (g[f]>>i)&1 == 1 {
				ans += dp[state][i]
			}
		}
	}
	return ans / 2
}

func lowestSetBit(num int) int {
	for i := 0; i < 30; i++ {
		if (num>>i)&1 == 1 {
			return i
		}
	}
	return -1
}

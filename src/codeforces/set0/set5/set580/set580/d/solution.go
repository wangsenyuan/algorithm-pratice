package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	a := readNNums(reader, n)
	restrictions := make([][]int, k)
	for i := 0; i < k; i++ {
		restrictions[i] = readNNums(reader, 3)
	}

	res := solve(m, a, restrictions)

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

func solve(m int, a []int, restrictions [][]int) int {
	n := len(a)
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}

	for _, cur := range restrictions {
		x, y, c := cur[0]-1, cur[1]-1, cur[2]
		grid[x][y] = c
	}

	T := 1 << n
	dp := make([][]int, T)
	for i := 0; i < T; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[1<<i][i] = a[i]
	}

	for state := 1; state < T; state++ {
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				for j := 0; j < n; j++ {
					if (state>>j)&1 == 0 {
						next := state | 1<<j
						dp[next][j] = max(dp[next][j], dp[state][i]+grid[i][j]+a[j])
					}
				}
			}
		}
	}

	var res int
	for state := (1 << m) - 1; state < T; state++ {
		cnt := bits.OnesCount(uint(state))
		if cnt == m {
			res = max(res, slices.Max(dp[state]))
		}
	}

	return res
}

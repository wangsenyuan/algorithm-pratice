package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	res := solve(n)

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

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

type state struct {
	ways int
	free bool
}

func solve(n int) int {
	n *= 2
	dp := make([][]state, n+1)
	for i := range n + 1 {
		dp[i] = make([]state, n+1)
	}

	dp[0][0] = state{0, true}

	for l := 1; l <= n; l++ {
		for bal := 0; bal <= n; bal++ {
			var ways int
			var free bool
			if bal > 0 {
				ways = add(ways, dp[l-1][bal-1].ways)
				free = free || dp[l-1][bal-1].free
			}
			if bal+1 <= l-1 {
				ways = add(ways, dp[l-1][bal+1].ways)
				free = free || dp[l-1][bal+1].free
			}

			if free {
				// 有一个子节点可以用来连接当前节点
				dp[l][bal] = state{add(ways, 1), false}
			} else {
				dp[l][bal] = state{ways, true}
			}
		}
	}
	return dp[n][0].ways
}

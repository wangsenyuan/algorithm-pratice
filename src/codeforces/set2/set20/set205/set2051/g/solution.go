package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	s := make([]string, k)
	for i := range k {
		s[i] = readString(reader)
	}
	return solve(n, s)
}

const inf = 1e9

func solve(n int, qs []string) int {
	// n <= 20
	m := len(qs)
	events := make([]int, m)
	for i, x := range qs {
		j := strings.IndexByte(x, ' ')
		id, _ := strconv.Atoi(x[:j])
		if x[j+1] == '-' {
			id *= -1
		}
		events[i] = id
	}

	getSafeDist := func(u int, v int) int {
		res := 1
		cur := 1
		for _, id := range events {
			i := abs(id) - 1
			if i == u {
				if id > 0 {
					// u向v靠近了
					cur--
				}
			} else if i == v {
				if id < 0 {
					// v远离了u
					cur++
				}
			}
			// 保证cur >= 1
			if cur == 0 {
				res++
				cur++
			}
		}
		return res
	}

	// dist[i][j] = 如果i在j的前面，两个最短需要的安全距离
	dist := make([][]int, n)
	for i := range n {
		dist[i] = make([]int, n)
		for j := range n {
			if i != j {
				dist[i][j] = getSafeDist(i, j)
			}
		}
	}
	N := 1 << n
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = inf
		}
	}

	far := make([]int, n)

	for _, id := range events {
		if id > 0 {
			// id往前移动了多少
			far[id-1]++
		}
	}

	for i := 0; i < n; i++ {
		dp[1<<i][i] = 1 + far[i]
	}

	for state := 1; state < N; state++ {
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				// dp[state][i]
				for j := 0; j < n; j++ {
					if (state>>j)&1 == 0 {
						// i的起始位置
						tmp := dp[state][i] - far[i] + dist[i][j] + far[j]
						dp[state|(1<<j)][j] = min(dp[state|(1<<j)][j], tmp)
					}
				}
			}
		}
	}

	var ans int = inf

	for i := 0; i < n; i++ {
		ans = min(ans, dp[N-1][i])
	}

	return ans
}

func abs(num int) int {
	return max(num, -num)
}

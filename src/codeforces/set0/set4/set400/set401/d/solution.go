package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	res := solve(n, m)
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

const N = 20

var F [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = i * F[i-1]
	}
}

func solve(num int, m int) int {
	var arr []int
	freq := make([]int, 10)
	for i := num; i > 0; i /= 10 {
		x := i % 10
		freq[x]++
		arr = append(arr, x)
	}

	n := len(arr)
	// n <= 18
	tot := 1 << n

	dp := make([][]int, tot)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		// 不能以0开头
		if arr[i] > 0 {
			dp[1<<i][arr[i]%m] = 1
		}
	}

	for state := 0; state < tot; state++ {
		for x := 0; x < m; x++ {
			if dp[state][x] == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if (state>>j)&1 == 0 {
					next := state | (1 << j)
					dp[next][(x*10+arr[j])%m] += dp[state][x]
				}
			}
		}
	}

	res := dp[tot-1][0]

	for i := 0; i < 10; i++ {
		res /= F[freq[i]]
	}

	return res
}

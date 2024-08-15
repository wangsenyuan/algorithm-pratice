package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

const inf = 1 << 30

func solve(a []int) int {
	// n := len(a)

	n := len(a)

	dp := make([]int, n+1)

	stack := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i+1] = inf
		if a[i] == 0 {
			dp[i+1] = dp[i]
		}
		var cost int
		var top int
		for j := i; j >= 0; j-- {
			if a[j] == 1 {
				if top > 0 && stack[top-1] < 0 {
					cost += abs(j + 1 + stack[top-1])
					top--
				} else {
					stack[top] = j + 1
					top++
				}
			} else {
				if top > 0 && stack[top-1] > 0 {
					cost += stack[top-1] - (j + 1)
					top--
				} else {
					stack[top] = -(j + 1)
					top++
				}
			}
			if top == 0 {
				dp[i+1] = min(dp[i+1], dp[j]+cost)
			}
		}
	}

	return dp[n]
}

func abs(num int) int {
	return max(num, -num)
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)

	res := solve(n, k)

	if len(res) == 0 {
		fmt.Println(-1)
		return
	}

	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}

	fmt.Print(buf.String())
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

func solve(n int, k int) []int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = 1 + dp[i/2] + dp[i-i/2]
	}
	if dp[n] < k || k%2 == 0 {
		return nil
	}
	// k 不会很大， 差不多在lgn 的水平
	// F[n] >= k
	// 如果F[n] = k, 那么就是一个完全反序的排列
	arr := make([]int, n)

	var dfs func(l int, r int, x int) bool

	dfs = func(l int, r int, x int) bool {
		if x == 0 {
			return false
		}

		if x == 1 {
			for i := l; i < r; i++ {
				arr[i] = i
			}
			return true
		}
		if r-l == 1 {
			return false
		}
		if x > dp[r-l] {
			return false
		}

		if x == dp[r-l] {
			for i := l; i < r; i++ {
				arr[i] = r - 1 - (i - l)
			}
			return true
		}
		mid := (l + r) / 2
		y := max(1, min(dp[mid-l], x)-2)
		z := x - 1 - y
		ok := dfs(l, mid, y) && dfs(mid, r, z)
		if !ok {
			return false
		}

		if y > 1 || z > 1 {
			// 左边已经满足没有完全排好序，不需要再处理
			return true
		}
		// both 1
		for i := l; i < mid; i++ {
			j := i + (mid - l)
			arr[i], arr[j] = arr[j], arr[i]
		}
		return true
	}

	if !dfs(0, n, k) {
		return nil
	}
	for i := 0; i < n; i++ {
		arr[i]++
	}
	return arr
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

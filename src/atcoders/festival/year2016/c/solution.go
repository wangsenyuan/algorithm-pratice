package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	L := readNNums(reader, n)
	R := readNNums(reader, n)

	res := solve(L, R)

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

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % mod)
}

func solve(L []int, R []int) int {
	n := len(L)
	height := make([][]int, n)
	for i := 0; i < n; i++ {
		height[i] = make([]int, 2)
		if i == 0 || L[i] > L[i-1] {
			// must be a new height
			height[i][0] = L[i]
			height[i][1] = L[i]
			continue
		}
		// L[i] == L[i-1]
		height[i][0] = height[i-1][0]
		height[i][1] = 1
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 || R[i] > R[i+1] {
			// conflict with L
			if R[i] > height[i][0] || R[i] < height[i][1] {
				return 0
			}
			height[i][0] = R[i]
			height[i][1] = R[i]
			continue
		}
		// R[i] == R[i+1]
		height[i][0] = min(height[i][0], height[i+1][0])
		height[i][1] = max(height[i][1], 1)

		if height[i][0] < height[i][1] {
			return 0
		}
	}

	ans := 1
	for i := 0; i < n; i++ {
		ans = mul(ans, height[i][0]-height[i][1]+1)
	}
	return ans
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

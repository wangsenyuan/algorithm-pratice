package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	A := readNNums(reader, n)
	B := readNNums(reader, m)
	res := solve(A, B)
	fmt.Println(res)
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const MAX_X = 1e6 + 1

func solve(A []int, B []int) int {
	// len(A) <= 1000000
	// len(B) <= 1000000
	// A是distinct， B也是distinct
	// 假设进行了一次切分后，dp[i] = 此次切分后的以i结尾的符合条件的长度
	// 假设A的切分点使用新的切点，j
	// 在j前面的有可能会增加prefix,
	// 假设 dp[n-1] = x, dp[i] = y, and y = i + 1 and pos[A[0]] > pos[A[n-1] dp[i] += x
	// 在j后面的有可能会减少prefix， 假设 dp[i] = x, j > i - x, 那么 新的dp[i] = i - j
	// i < j + x
	pos := make([]int, MAX_X)
	for i := 0; i < MAX_X; i++ {
		pos[i]--
	}
	for i, num := range B {
		pos[num] = i
	}
	n := len(A)
	m := len(B)
	a := make([]int, 2*n+1)
	copy(a, A)
	copy(a[n:], A)
	a[2*n] = A[0]
	//b := extendArray(B)
	cost := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		if pos[a[i]] < 0 || pos[a[i+1]] < 0 {
			cost[i] = 2 * m
		} else {
			cost[i] = (pos[a[i+1]] - pos[a[i]] + m) % m
		}
	}
	var ans int
	for i := 0; i < n; i++ {
		if pos[A[i]] >= 0 {
			ans = 1
			break
		}
	}
	if ans == 0 {
		return 0
	}
	var cur int

	for l, r := 0, 0; l < 2*n; l++ {
		if r < l {
			r++
		}
		for r < 2*n && cur+cost[r] < m {
			cur += cost[r]
			r++
		}
		ans = max(ans, r-l+1)
		if r > l {
			cur -= cost[l]
		}
	}
	return min(ans, n, m)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(num ...int) int {
	res := num[0]
	for i := 1; i < len(num); i++ {
		if res > num[i] {
			res = num[i]
		}
	}
	return res
}

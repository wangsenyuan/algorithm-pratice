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
	s, f := readTwoNums(reader)
	res := solve(a, s, f)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

func solve(a []int, s int, f int) int {
	n := len(a)
	// 都于i来说 如果它是起点1,
	// 先试试diff数组，不行，再考虑segment tree
	// 假设在时刻 x = n时
	// n, 1, 2, 3.... n - 1
	// n - i >= s and n - i < f
	// i <= n - s
	// i > n - f
	//
	l, r := s, f-1
	// 符合条件的范围是
	var sum int
	for i := l; i <= r; i++ {
		sum += a[i%n]
	}
	best := sum
	ans := n

	for x := n - 1; x > 0; x-- {
		// 当x变小的时候，
		r = (r + 1) % n
		sum += a[r]
		sum -= a[l]
		l = (l + 1) % n
		if best <= sum {
			best = sum
			ans = x
		}
	}

	return ans
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

package main

import (
	"bufio"
	"fmt"
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
	n, k := readTwoNums(reader)
	return solve(n, k)
}
func solve(n int, k int) int {
	// a, b, c, d <= n >= 1
	// (a + b) - (c + d) = k
	// 那么在选定了 c+d后， a + b = c + d + k
	// 那么就是求 a+b = sum的计数
	// c+d 的计数 = n * n (不限制的情况下)
	// 但是 c + d + k <= 2 * n => c + d <= 2 * n - k
	// c + d + k >= 2
	// c >= 1, d >= 1
	// k 可以是负值
	cnt := make([]int, 2*n+2)
	for d := 1; d <= n; d++ {
		// c + d +k<= 2 * n >= 2
		c1 := min(2*n-k-d, n)
		c2 := max(2-k-d, 1)
		if c2 <= c1 {
			cnt[c2+d]++
			cnt[c1+d+1]--
		}
	}
	for i := 1; i <= 2*n; i++ {
		cnt[i] += cnt[i-1]
	}
	var res int
	// c + d 的计数方式
	for sum := 2; sum <= 2*n; sum++ {
		if sum+k >= 2 && sum+k <= 2*n {
			// sum + k = a + b
			// 1 + x = sum + k
			a1 := max(1, sum+k-n)
			a2 := min(n, sum+k-1)
			res += max(0, a2-a1+1) * cnt[sum]
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

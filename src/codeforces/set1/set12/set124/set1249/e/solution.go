package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, c := readTwoNums(reader)
	a := readNNums(reader, n-1)
	b := readNNums(reader, n-1)

	res := solve(n, c, a, b)

	s := fmt.Sprintf("%v", res)

	fmt.Print(s[1 : len(s)-1])
}

func solve(n int, c int, a []int, b []int) []int {
	// len(a) = n - 1
	ans := make([]int, n)

	dp := []int{0, c}

	var pa int
	var pb int

	for i := 1; i < n; i++ {
		pa += a[i-1]
		pb += b[i-1]
		// x从楼梯到达当前层
		x := pa + dp[0]
		y := pb + dp[1]
		ans[i] = min(x, y)

		dp[0] = min(dp[0], min(x, y)-pa)
		dp[1] = min(dp[1], min(y, x+c)-pb)
	}

	return ans
}

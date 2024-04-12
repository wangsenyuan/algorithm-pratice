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

const D = 62

func solve(a []int) int {
	// 最多18个1
	n := len(a)
	cnt := make([]int, n)

	var sum int

	var res int

	dp := make([]int, 2)
	dp[0]++

	for i, num := range a {
		cnt[i] = digits(num)
		sum += cnt[i]
		res += dp[sum&1]
		dp[sum&1]++

		var tmp, x int
		// x是最大值
		for k := i; k >= 0 && k >= i-D; k-- {
			tmp += cnt[k]
			x = max(x, cnt[k])
			if tmp%2 == 0 && tmp < 2*x {
				res--
			}
		}
	}
	// 以i为最大值的情况下，

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func digits(n int) int {
	var res int
	for n > 0 {
		res++
		n -= n & -n
	}
	return res
}

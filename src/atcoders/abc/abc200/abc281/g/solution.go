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

func solve(n int, mod int) int {
	add := func(a, b int) int {
		a += b
		if a >= mod {
			a -= mod
		}
		return a
	}

	sub := func(a, b int) int {
		return add(a, mod-b)
	}

	mul := func(num ...int) int {
		res := 1
		for _, x := range num {
			res = res * x % mod
		}
		return res
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[1][1] = 1

	C := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		C[i] = make([]int, n+1)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = add(C[i-1][j], C[i-1][j-1])
		}
	}

	N := n * (n - 1) / 2
	p2 := make([]int, N+10)
	p2[0] = 1
	for i := 1; i < len(p2); i++ {
		p2[i] = mul(2, p2[i-1])
	}

	for i := 1; i < n; i++ {
		// 还有 n - i 个节点没有处理
		for x := 1; x <= i; x++ {
			// dp[i][x] 表示到i为止有x个节点的距离是最远的
			y0 := sub(p2[x], 1)
			y := 1
			for j := 1; i+j < n; j++ {
				y = mul(y, y0)
				// 从中选择j个节点，作为新的一层
				// 每个有y0个选择去连接， 共有j个, pow(y0, j)
				// 这j个节点中间共有 j * (j - 1) / 2
				// 这些边可以连接，或者不连接
				tmp := mul(dp[i][x], C[n-1-i][j], y, p2[j*(j-1)/2])
				dp[i+j][j] = add(dp[i+j][j], tmp)
			}
		}

	}
	// 把n连接起来
	var res int
	for x := 1; x < n; x++ {
		res = add(res, mul(dp[n-1][x], sub(p2[x], 1)))
	}
	return res
}

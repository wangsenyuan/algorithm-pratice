package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)

	a := readNNums(reader, n)

	res := solve(a, k)

	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

const inf = 1 << 50

func solve(a []int, d int) []int {
	n := len(a)

	var arr []int

	var prod = 1

	for i := 0; i < n; i++ {
		if d%2 == 1 && a[i]%2 == 0 {
			continue
		}
		if d%5 != 0 && a[i]%5 == 0 {
			// d is not 0 and d is not 5
			continue
		}
		arr = append(arr, a[i])
		prod *= a[i]
		prod %= 10
	}

	if prod == d {
		return arr
	}

	sort.Ints(arr)
	n = len(arr)

	dp := make([][]int, n+1)
	fp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 10)
		fp[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			dp[i][j] = inf
			fp[i][j] = -1
		}
	}

	dp[0][1] = 1

	for i, num := range arr {
		// 如果当前数字放入集合
		// dp[i][j * num % 10] = dp[i][j]
		// 如果不放入集合 dp[i][j] = dp[i][j] * num
		for j := 0; j < 10; j++ {
			r := (j * num) % 10
			if dp[i][j] < dp[i+1][r] {
				dp[i+1][r] = dp[i][j]
				fp[i+1][r] = j
			}

			if dp[i][j]*num < dp[i+1][j] {
				dp[i+1][j] = dp[i][j] * num
				fp[i+1][j] = -1
			}
		}
	}

	if dp[n][d] >= inf {
		return nil
	}
	var res []int

	p, r := n, d
	for p > 0 {
		if fp[p][r] >= 0 {
			res = append(res, arr[p-1])
			r = fp[p][r]
		}

		p--
	}

	return res
}

func solve1(a []int, d int) []int {
	n := len(a)

	// 把乘法变成了log的加法
	dp := make([][]float64, n+1)
	fp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, 10)
		fp[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			dp[i][j] = -inf
			fp[i][j] = -1
		}
	}
	dp[0][1] = 0

	for i, num := range a {
		copy(dp[i+1], dp[i])
		tmp := math.Log2(float64(num))

		for r := 0; r < 10; r++ {
			nr := (r * num) % 10
			if dp[i][r]+tmp > dp[i+1][nr] {
				dp[i+1][nr] = dp[i][r] + tmp
				fp[i+1][nr] = r
			}
		}
	}

	if dp[n][d] < 0 {
		return nil
	}

	var res []int
	r := d
	for i := n; i > 0; i-- {
		if fp[i][r] >= 0 {
			res = append(res, a[i-1])
			r = fp[i][r]
		}
	}

	return res
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, l, k := readThreeNums(reader)
	a := readNNums(reader, n)

	res := solve(l, a, k)

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

const MOD = 1e9 + 7

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	a %= MOD
	b %= MOD
	return a * b % MOD
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func inv(a int) int {
	return pow(a, MOD-2)
}

func solve(l int, a []int, k int) int {
	n := len(a)

	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		return a[id[i]] < a[id[j]] || a[id[i]] == a[id[j]] && id[i] < id[j]
	})

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		dp[i][1] = 1
	}

	for j := 2; j <= k; j++ {
		var pt int
		var sum int
		for i := 0; i < n; i++ {
			for pt < n && a[id[pt]] <= a[id[i]] {
				sum = modAdd(sum, dp[id[pt]][j-1])
				pt++
			}
			dp[id[i]][j] = sum
		}
	}

	c := (l + n - 1) / n
	x := (l - 1) % n
	var res int

	for i := 0; i <= x; i++ {
		for j := 1; j <= min(c, k); j++ {
			res = modAdd(res, modMul(dp[i][j], c-j+1))
		}
	}

	for i := x + 1; i < n; i++ {
		for j := 1; j <= min(c-1, k); j++ {
			res = modAdd(res, modMul(dp[i][j], c-j))
		}
	}

	return res
}

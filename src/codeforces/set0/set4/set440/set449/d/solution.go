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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)

	res := solve(A)

	fmt.Println(res)
}

const H = 20

func solve(A []int) int {
	// how many and(arr) = 0
	// A[i] < 1e6
	T := 1 << H
	dp := make([]int, T)

	for _, a := range A {
		dp[a]++
	}

	// dp[i] = 所有i的超集的和

	for i := 0; i < H; i++ {
		for mask := T - 1; mask >= 0; mask-- {
			if (mask>>i)&1 == 0 {
				dp[mask] = add(dp[mask], dp[mask|(1<<i)])
			}
		}
	}
	// dp[i] 等于and后结果为i的子数组的长度，
	// fp[i] 等于有多少中选择得到and值为i的选择
	// res = fp[0] - fp[1] - fp[2] + fp[3] - fp[4] ...
	var res int

	for x := 0; x < T; x++ {
		var cnt int
		for i := 0; i < H; i++ {
			cnt += (x >> i) & 1
		}
		if cnt&1 == 1 {
			res = sub(res, pow(2, dp[x]))
		} else {
			res = add(res, pow(2, dp[x]))
		}
	}

	return res
}

func max_of(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}

const MOD = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func sub(a, b int) int {
	return add(a, MOD-b)
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

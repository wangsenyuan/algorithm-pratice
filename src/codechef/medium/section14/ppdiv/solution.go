package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
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
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		s, _ := reader.ReadBytes('\n')
		var N uint64
		readUint64(s, 0, &N)
		res := solve(N)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

const MOD = 1000000007

var inv6 = inv(6)

const H = 60

func solve(N uint64) int {
	dp := make([]uint64, H)
	dp[2] = solve2(N)
	for i := 3; i < H; i++ {
		dp[i] = solve3(uint64(i), N)
	}

	for i := H - 1; i >= 2; i-- {
		for j := i + i; j < H; j += i {
			dp[i] += MOD - dp[j]
			if dp[i] >= MOD {
				dp[i] -= MOD
			}
		}
	}
	res := N % MOD
	for i := 2; i < H; i++ {
		res += dp[i]
		res %= MOD
	}
	return int(res)
}

func solve2(N uint64) uint64 {
	if N < 4 {
		return 0
	}
	var cur uint64 = 2
	var ans uint64

	for cur <= N/cur {
		v := N / (cur * cur)
		lo := sqrt(N / v)
		ans += (N / (cur * cur) % MOD * (sumOfSquares(lo) + MOD - sumOfSquares(cur-1))) % MOD
		if ans >= MOD {
			ans -= MOD
		}
		cur = lo + 1
	}
	return ans
}

func sumOfSquares(n uint64) uint64 {
	n %= MOD
	return (((n * (n + 1)) % MOD * (2*n + 1)) % MOD * inv6) % MOD
}

func sqrt(N uint64) uint64 {
	x := uint64(math.Sqrt(float64(N)))
	for x*x > N {
		x--
	}
	for (x+1)*(x+1) <= N {
		x++
	}
	return x
}

func solve3(p uint64, N uint64) uint64 {
	var ans uint64
	for b := 2; ; b++ {
		tmp := pow(uint64(b), p, N)
		if tmp > N {
			break
		}
		ans += (N - N%tmp) % MOD
		if ans >= MOD {
			ans -= MOD
		}
	}
	return ans
}

func pow(a, p uint64, mx uint64) uint64 {
	var res uint64 = 1
	for p > 0 {
		p--
		res *= a
		if res > mx {
			res = mx + 1
			break
		}
	}
	return res
}

func inv(a uint64) uint64 {
	var res uint64 = 1
	for p := MOD - 2; p > 0; p >>= 1 {
		if p&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
	}
	return res
}

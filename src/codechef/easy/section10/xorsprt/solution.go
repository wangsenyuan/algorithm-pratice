package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := make([]int64, n)
		s, _ := reader.ReadBytes('\n')
		var pos int
		for i := 0; i < n; i++ {
			var num uint64
			pos = readUint64(s, pos, &num) + 1
			A[i] = int64(num)
		}
		res := solve(n, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

const MOD = 1000000007
const KK = 63

func solve(n int, A []int64) int {
	//dp := make([]map[int64]int64, KK+1)
	//for i := 0; i <= KK; i++ {
	//	dp[i] = make(map[int64]int64)
	//}
	//// dp[i][j] = count of ways get when num % (1 << i) = j
	//dp[0][0] = 1
	//fp := make([]int64, KK+1)
	//var sum int64
	//for i := 0; i < n; i++ {
	//	sum ^= A[i]
	//	for k := KK; k >= 0; k-- {
	//		cur := sum & ((1 << uint64(k)) - 1)
	//		fp[k] = dp[k][cur]
	//	}
	//	for k := KK; k > 0; k-- {
	//		cur := sum & ((1 << uint64(k)) - 1)
	//		dp[k][cur] += fp[k-1]
	//		dp[k][cur] %= MOD
	//	}
	//}
	//var res int64
	//
	//for k := 0; k <= KK; k++ {
	//	res += fp[k]
	//	res %= MOD
	//}

	var xor int64

	for i := 0; i < n; i++ {
		xor ^= A[i]
	}

	dp := make([]int64, KK)
	dp[0] = 1

	for i := 0; i < n-1; i++ {
		xor ^= A[i]

		for k := KK - 1; k > 0; k-- {
			if xor&((1<<uint64(k))-1) == 0 {
				if k == KK-1 {
					dp[k] = (dp[k] * 2) % MOD
				}
				dp[k] = (dp[k] + dp[k-1]) % MOD
			}
		}
	}
	var res int64

	for i := 0; i < KK; i++ {
		res += dp[i]
		res %= MOD
	}

	return int(res)
}

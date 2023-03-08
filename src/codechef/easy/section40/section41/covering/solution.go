package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	F := readNNums(reader, 1<<n)
	G := readNNums(reader, 1<<n)
	H := readNNums(reader, 1<<n)
	res := solve(n, F, G, H)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func addSos(dp []int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < 1<<n; j++ {
			if (j>>i)&1 == 1 {
				dp[j] = add(dp[j], dp[j^(1<<i)])
			}
		}
	}
}

func subSos(dp []int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < 1<<n; j++ {
			if (j>>i)&1 == 1 {
				dp[j] = sub(dp[j], dp[j^(1<<i)])
			}
		}
	}
}

func solve(n int, F, G, H []int) int {
	// R(D) = sum(F(A) * G(B) * H(C))
	// D 是 A | B | C 集合的子集
	// Sum(R(D))
	// 考虑D = 1, 只包含第0个元素的集合，
	// 那么A, B, C 中只要有一个包含第0个元素，即可
	// 假设确定了A，那么sum(F(A) * G(B) * H(C)) = F(A) * sum(G(B) * H(C))
	// (sum(F) - sum(A`))) * sum(G) * sum(H), A` = 不包含
	addSos(F, n)
	addSos(G, n)
	addSos(H, n)
	dp := make([]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		dp[i] = mul(F[i], mul(G[i], H[i]))
	}
	subSos(dp, n)

	for i := 0; i < n; i++ {
		for j := 0; j < 1<<n; j++ {
			if (j>>i)&1 == 0 {
				dp[j] = add(dp[j], dp[j^(1<<i)])
			}
		}
	}
	var res int

	for i := 0; i < 1<<n; i++ {
		res = add(res, dp[i])
	}

	return res
}

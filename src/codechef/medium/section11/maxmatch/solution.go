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

	tc, M := readTwoNums(reader)

	res := solve(M)

	for tc > 0 {
		tc--
		n := readNum(reader)
		fmt.Println(res[n])
	}
}

const N = 2010

func solve(M int) []int {
	MOD := int64(M)

	pw2 := make([]int64, N)
	pw2[0] = 1
	for i := 1; i < N; i++ {
		pw2[i] = (pw2[i-1] * 2) % MOD
	}
	dp := make([]int64, N)
	dp[1] = 0
	dp[2] = 1
	dp[3] = 7
	for i := 4; i < N; i++ {
		for r := 1; r < i; r++ {
			tmp := (pw2[(i-1)-r]*int64(r)*int64(i-r) + dp[i-r]) % MOD
			dp[i] = (dp[i] + tmp) % MOD
		}
	}

	fp := make([]int, N)
	fp[2] = 1
	fp[3] = 17
	for i := 4; i < N; i++ {
		var val int64
		for j := 1; j < i; j++ {
			tmp1 := int64(j) * int64(i-j) % MOD
			tmp := (tmp1 * tmp1) % MOD
			tmp = (tmp * pw2[i-1-j]) % MOD
			val = (val + tmp) % MOD
			val = (val + int64(fp[j])) % MOD
			val = (val + 2*tmp1*dp[i-j]) % MOD
		}
		fp[i] = int(val)
	}

	return fp
}

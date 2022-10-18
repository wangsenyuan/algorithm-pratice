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
		arr := readNNums(reader, 6)
		n, k, a, b, c, m := arr[0], arr[1], arr[2], arr[3], arr[4], arr[5]
		res := solve(n, k, a, b, c, m)
		buf.WriteString(fmt.Sprintln(fmt.Sprintf("%d", res)))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func modAdd(a, b int, mod int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func modMul(a, b int, mod int) int {
	r := int64(a) * int64(b)
	return int(r % int64(mod))
}

func pow(a, b int, mod int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a, mod)
		}
		a = modMul(a, a, mod)
		b >>= 1
	}
	return r
}

const MOD = 1000000007

const MAX_L = 1 << 20

func solve(n, k, a, b, c, m int) int {
	x := make([]int, MAX_L)
	x[1] = a % m
	vis := make([]int, MAX_L)
	vis[a] = 1

	var prefixLength, cycleLength int

	for i := 2; true; i++ {
		x[i] = int((int64(b)*int64(x[i-1]) + int64(c)) % int64(m))
		if vis[x[i]] > 0 {
			prefixLength = vis[x[i]] - 1
			cycleLength = i - vis[x[i]]
			break
		}
		vis[x[i]] = i
	}

	cycleCount := (n - prefixLength) / cycleLength
	suffixLength := (n - prefixLength) % cycleLength

	dp := make([]int, k+1)
	dp[0] = 1
	prefixDp := make([]int, k+1)
	suffixDp := make([]int, k+1)
	ndp := make([]int, k+1)
	cycleDp := make([]int, k+1)
	for i := 1; i <= prefixLength+cycleLength; i++ {
		for j := k; j >= 1; j-- {
			dp[j] = modAdd(dp[j], modMul(x[i], dp[j-1], MOD), MOD)
		}
		if i == prefixLength {
			copy(prefixDp, dp)
			clear(dp)
			dp[0] = 1
		}
		if i == prefixLength+suffixLength {
			copy(suffixDp, dp)
		}
		if i == n {
			copy(ndp, dp)
			break
		}
	}

	if n < prefixLength {
		return ndp[k]
	}

	copy(cycleDp, dp)

	if prefixLength != 0 {
		copy(dp, prefixDp)
	} else {
		clear(dp)
		dp[0] = 1
	}

	mulDp := func(a []int, b []int) {
		for i := k; i >= 1; i-- {
			var s int
			for j := 0; j <= i; j++ {
				s = modAdd(s, modMul(b[j], a[i-j], MOD), MOD)
			}
			a[i] = s
		}
	}

	for cycleCount > 0 {
		if cycleCount&1 == 1 {
			mulDp(dp, cycleDp)
		}
		mulDp(cycleDp, cycleDp)
		cycleCount >>= 1
	}

	if suffixLength != 0 {
		mulDp(dp, suffixDp)
	}

	return dp[k]
}

func clear(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = 0
	}
}

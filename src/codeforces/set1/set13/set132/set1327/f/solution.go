package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k, m := readThreeNums(reader)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 3)
	}
	res := solve(n, k, Q)
	fmt.Println(res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
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

func sub(a, b int) int {
	return add(a, mod-b)
}

func solve(n int, k int, Q [][]int) int {
	// 每一位可以单独处理
	var res = 1
	F := make([]int, n+2)
	F[0] = 1
	for d := 0; d < k; d++ {
		tmp := process(n, d, Q, F)
		res = mul(res, tmp)
	}

	return res
}

func process(n int, d int, Q [][]int, F []int) int {
	L := make([]int, n+1)
	diff := make([]int, n+2)
	for _, cur := range Q {
		if (cur[2]>>d)&1 == 0 {
			L[cur[1]] = max(L[cur[1]], cur[0])
		} else {
			diff[cur[0]]++
			diff[cur[1]+1]--
		}
	}
	var sum = 1
	var left int
	for i := 1; i <= n+1; i++ {
		for left < L[i-1] {
			sum = sub(sum, F[left])
			left++
		}
		diff[i] += diff[i-1]
		if diff[i] > 0 {
			F[i] = 0
			continue
		}
		F[i] = sum
		sum = mul(sum, 2)
	}

	return F[n+1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

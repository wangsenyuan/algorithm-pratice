package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	x, y := readTwoNums(reader)
	res := solve(x, y)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

const mod = 1e9 + 7

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

func solve(x, y int) int {
	if y%x != 0 {
		return 0
	}

	var fs []int

	checkAndAdd := func(f int) {
		if f%x != 0 {
			return
		}
		// f is multiple of x
		fs = append(fs, f)
	}

	for i := 1; i <= y/i; i++ {
		if y%i == 0 {
			checkAndAdd(i)
			if i*i != y {
				checkAndAdd(y / i)
			}
		}
	}

	sort.Ints(fs)

	n := len(fs)

	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		z := fs[i]
		dp[i] = pow(2, y/z-1)
		for j := i + 1; j < n; j++ {
			if fs[j]%z == 0 {
				dp[i] = sub(dp[i], dp[j])
			}
		}
	}

	return dp[0]
}

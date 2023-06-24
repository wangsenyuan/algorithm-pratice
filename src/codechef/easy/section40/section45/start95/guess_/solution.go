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
		res := solve(n)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

const MAX_N = 10000_000 + 10

var spf [MAX_N]int
var pct [MAX_N]int
var mul [MAX_N]int64
var div [MAX_N]int64
var ans [MAX_N]int64

const MOD = 1e9 + 7

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
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

func inverse(a int) int {
	return pow(a, MOD-2)
}

func init() {
	div[1] = 1
	for i := 2; i < MAX_N; i++ {
		if spf[i] == 0 {
			for j := i; j < MAX_N; j += i {
				spf[j] = i
			}
		}
		x := i / spf[i]
		pct[i] = 1 + pct[x]
		if spf[i] == spf[x] {
			div[i] = div[x] / mul[x] * (mul[x] + 1)
			mul[i] = mul[x] + 1
		} else {
			div[i] = div[x] * 2
			mul[i] = 2
		}
		ans[i] = 3*div[i] - 3
		ans[i] += int64(pct[i]) * (div[i] - 2)
	}

	div[1] = 0
	for i := 2; i < MAX_N; i++ {
		div[i] += div[i-1] - 1
		ans[i] += ans[i-1]
	}
}

func solve(n int) int {
	return modMul(int(ans[n]%MOD), inverse(int(2*div[n]%MOD)))
}

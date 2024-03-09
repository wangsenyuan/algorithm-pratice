package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	_, m := readTwoNums(reader)
	s := readString(reader)
	ops := make([]string, m)

	for i := 0; i < m; i++ {
		ops[i] = readString(reader)
	}

	res := solve(s, ops)

	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
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
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}

	return res
}

const N = 300010

var I [N]int

func init() {
	for i := 0; i < N; i++ {
		I[i] = pow(i, mod-2)
	}
}

func solve(s string, ops []string) []int {
	var zero bool
	res := 1
	buf := []byte(s)

	for i := 0; i < len(buf); i++ {
		if buf[i] == '?' {
			if i == 0 {
				zero = true
			} else {
				res = mul(res, i)
			}
		}
	}

	ans := make([]int, len(ops)+1)
	if zero {
		ans[0] = 0
	} else {
		ans[0] = res
	}

	for i := 0; i < len(ops); i++ {
		var j int
		pos := readInt([]byte(ops[i]), 0, &j)
		j--
		c := ops[i][pos+1]
		if buf[j] == '?' && c != '?' {
			if j == 0 {
				zero = false
			} else {
				res = mul(res, I[j])
			}
		} else if buf[j] != '?' && c == '?' {
			if j == 0 {
				zero = true
			} else {
				res = mul(res, j)
			}
		}
		buf[j] = c
		if !zero {
			ans[i+1] = res
		}
	}
	return ans
}

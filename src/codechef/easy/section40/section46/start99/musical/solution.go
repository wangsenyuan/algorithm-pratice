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
		n, m, k := readThreeNums(reader)
		C := readNNums(reader, n)
		L := readNNums(reader, k)
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = readNNums(reader, 2)
		}
		res := solve(m, C, L, A)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
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
	if len(bs) == 0 || bs[0] == '\n' {
		return readNum(reader)
	}
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

func solve(m int, C []int, L []int, A [][]int) []int {
	n := len(C)
	k := len(L)

	mark := make([]bool, n)
	for i := 0; i < k; i++ {
		if i > 0 {
			L[i] += L[i-1]
		}
		L[i] %= n
		mark[L[i]] = true
	}
	im := inverse(m)

	bad := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		bad[i] = make(map[int]bool)
		for j := 0; j < n; j++ {
			if mark[j] {
				bad[i][C[(i+j)%n]] = true
			}
		}
	}

	var tot int
	for i := 0; i < n; i++ {
		tmp := mul(sub(m, len(bad[i])), im)
		tot = add(tot, tmp)
	}

	ans := make([]int, n)

	for j, cur := range A {
		i, x := cur[0], cur[1]
		i--

		tmp := mul(sub(m, len(bad[i])), im)
		tot = sub(tot, tmp)

		if !bad[i][x] {
			tot = add(tot, 1)
		}
		ans[j] = tot
	}

	return ans
}

const MOD = 1e9 + 7

func mul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
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

func inverse(a int) int {
	return pow(a, MOD-2)
}

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

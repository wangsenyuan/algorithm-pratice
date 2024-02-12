package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)

		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const N = 100010

var factors [N][]int

func init() {
	for i := 1; i < N; i++ {
		factors[i] = getFactors(i)
	}
}

func getFactors(num int) []int {
	var res []int
	for x := 1; x <= num/x; x++ {
		if num%x == 0 {
			res = append(res, x)
			if num/x != x {
				res = append(res, num/x)
			}
		}
	}
	sort.Ints(res)
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

func sub(a, b int) int {
	return add(a, mod-b)
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

func solve(a []int) int {
	n := len(a)

	id := make([]int, n)

	for i := 0; i < n; i++ {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		// 在相等时先处理小的下标
		return a[id[i]] > a[id[j]] || a[id[i]] == a[id[j]] && id[i] < id[j]
	})
	var res int
	mark := make([]bool, n+1)
	left := n
	for _, i := range id {
		var cnt int
		for _, f := range factors[i+1] {
			if !mark[f] {
				cnt++
				mark[f] = true
				left--
			}
		}
		// cnt至少有一个
		tmp := pow(2, cnt)
		// 不包括一个都没有的情况
		tmp = sub(tmp, 1)
		// 未标记的情况，其他的可以随便标记或者不标记
		tmp = mul(tmp, pow(2, left))
		res = add(res, mul(tmp, a[i]))
	}

	return res
}

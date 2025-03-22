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
	for range tc {
		res := process(reader)
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
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

func inv(a int) int {
	return pow(a, mod-2)
}

const N = 2_000_00 + 10

var fact [N]int
var invFact [N]int

func init() {
	fact[0] = 1
	invFact[0] = 1
	for i := 1; i < N; i++ {
		fact[i] = mul(fact[i-1], i)
		invFact[i] = inv(fact[i])
	}
}

func C(n, k int) int {
	if k < 0 || n < k {
		return 0
	}
	return mul(fact[n], mul(invFact[k], invFact[n-k]))
}

func fill(h int, w int, k int) int {
	return mul(C(h, k), mul(C(w, k), fact[k]))
}

func solve(a []int) int {
	n := len(a)
	if a[n-1] != n && a[n-1] != -1 {
		return 0
	}
	a[n-1] = n
	res := 1

	var cur_side, cur_inserted int

	for i := 1; i <= len(a); i++ {
		v := a[i-1]
		if v < -1 || v > i {
			return 0
		}
		if v == -1 {
			continue
		}
		to_insert := v - cur_inserted
		if to_insert < 0 {
			return 0
		}

		up_h := cur_side - cur_inserted
		up_w := i - cur_side
		down_h := up_w
		down_w := i - cur_inserted
		var ways int
		for k := 0; k <= to_insert; k++ {
			x := fill(up_h, up_w, k)
			y := fill(down_h, down_w-k, to_insert-k)
			ways = add(ways, mul(x, y))
		}
		res = mul(res, ways)
		cur_inserted = v
		cur_side = i
	}
	return res
}

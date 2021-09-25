package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		res := solve(n, m)
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

const MOD = 998244353

func solve(n, m int) int {

	var tr int64
	var tm int64 = 1
	for k := 1; k <= (n + m - 2); k++ {
		v := max(0, k-m+1)
		u := min(k, n-1)
		ans := calc(int64(n), int64(m), int64(k), int64(u))
		if v > 0 {
			ans = sub(ans, calc(int64(n), int64(m), int64(k), int64(v-1)))
		}
		if v == 0 {
			ans = sub(ans, mult(int64(n), int64(m-k)))
		}

		if v <= k && u >= k {
			ans = sub(ans, mult(int64(n-k), int64(m)))
		}
		tr += ((ans * tm) % MOD)
		tr %= MOD
		tm *= 31
		tm %= MOD
	}
	return int(tr)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}

func calc(n, m, k, x int64) int64 {
	var tr int64
	tr = add(tr, mult(x+1, mult(n, m)))
	tr = sub(tr, mult(x+1, mult(n, k)))
	tr = add(mult(mult(mult(x, x+1), 499122177), n), tr)
	tr = add(mult(mult(mult(x, x+1), 499122177), k), tr)
	tr = sub(tr, mult(mult(mult(x, x+1), 499122177), m))
	tr = sub(tr, mult(mult(mult(x, x+1), 2*x+1), 166374059))
	tr = mult(tr, 2)
	return tr
}

func add(a, b int64) int64 {
	return (a + b) % MOD
}

func sub(a, b int64) int64 {
	return ((a-b)%MOD + MOD) % MOD
}

func mult(a, b int64) int64 {
	return (a * b) % MOD
}

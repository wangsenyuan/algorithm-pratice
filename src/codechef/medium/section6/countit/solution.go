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
		res := solve(n, m, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const MOD = 1000000007
const X = 412345

var fact []int64

func init() {
	fact = make([]int64, X)
	fact[0] = 1
	for i := 1; i < X; i++ {
		fact[i] = int64(i) * fact[i-1] % MOD
	}
}

func pow(a, b int) int {
	A := int64(a)
	var R int64 = 1
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(R)
}

func solve(n, m, k int) int {
	pn := make([]int, 2*(n+m)+10)
	pm := make([]int, 2*(n+m)+10)
	for i := 0; i < len(pn); i++ {
		pn[i] = pow(i, n)
		pm[i] = pow(i, m)
	}
	sum := make([]int64, 2*(n+m)+10)
	for i := 1; i < len(sum); i++ {
		val := int64(pn[i]-pn[i-1]) * int64(pm[i]-pm[i-1])
		val %= MOD
		if val < 0 {
			val += MOD
		}
		sum[i] = (sum[i-1] + val) % MOD
	}
	if k < len(sum) {
		return int(sum[k])
	}

	var gg int64 = 1
	for i := 1; i < len(sum); i++ {
		gg *= int64(k - i)
		gg %= MOD
	}
	var wow int64
	for i := 1; i < len(sum); i++ {
		val := (fact[i-1] * fact[2*(n+m)+9-i]) % MOD
		if (2*(n+m)+9-i)%2 == 1 {
			val *= -1
			val %= MOD
			val += MOD
			val %= MOD
		}
		val = int64(inverse(int(val)))
		val *= gg
		val %= MOD
		val *= int64(inverse(k - i))
		val %= MOD
		val *= sum[i]
		val %= MOD
		wow += val
		wow %= MOD
	}

	return int(wow)
}

func inverse(num int) int {
	return pow(num, MOD-2)
}

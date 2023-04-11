package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(arr []int) string {
		if len(arr) == 2 {
			fmt.Printf("1 %d %d\n", arr[0], arr[1])
		} else {
			fmt.Printf("2 %d %d %d\n", arr[0], arr[1], arr[2])
		}
		return readString(reader)
	}

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)

		res := solve(n, ask)

		var buf bytes.Buffer

		buf.WriteString("3")
		for i := 0; i < n; i++ {
			buf.WriteByte(' ')
			buf.WriteString(res[i])
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())

		readNum(reader)
	}
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

func solve(n int, ask func(arr []int) string) []string {
	// when ask len(arr) = 2, it gives gcd(A[arr[0]], A[arr[1]])
	//                   = 3, it gives lcm(A[arr[0]], A[arr[1]], A[arr[2]])

	askGcd := func(i int, j int) *big.Int {
		res := ask([]int{i, j})
		n := new(big.Int)
		n.SetString(res, 10)
		return n
	}

	askLcm := func(i int, j int, k int) *big.Int {
		res := ask([]int{i, j, k})
		n := new(big.Int)
		n.SetString(res, 10)
		return n
	}
	g := make([][]*big.Int, 4)
	// g[0], g[1], g[2]
	for i := 1; i <= 3; i++ {
		g[i] = make([]*big.Int, 5)
		for j := i + 1; j <= 4; j++ {
			g[i][j] = askGcd(i, j)
		}
	}
	l123 := askLcm(1, 2, 3)
	l124 := askLcm(1, 2, 4)
	l134 := askLcm(1, 3, 4)
	l234 := askLcm(2, 3, 4)
	p123 := divBigInt(mulBigInt(l123, g[1][2], g[2][3], g[1][3]), gcdBigInt(g[1][2], g[2][3]))
	p124 := divBigInt(mulBigInt(l124, g[1][2], g[2][4], g[1][4]), gcdBigInt(g[1][2], g[1][4]))
	p134 := divBigInt(mulBigInt(l134, g[1][3], g[3][4], g[1][4]), gcdBigInt(g[1][3], g[1][4]))
	p234 := divBigInt(mulBigInt(l234, g[2][3], g[3][4], g[2][4]), gcdBigInt(g[2][3], g[2][4]))

	allprod3 := mulBigInt(p123, p124, p134, p234)
	allprod := cbrt(allprod3)

	res := make([]string, n)
	a0 := divBigInt(allprod, p234)
	a1 := divBigInt(allprod, p134)
	res[0] = a0.String()
	res[1] = a1.String()
	res[2] = divBigInt(allprod, p124).String()
	res[3] = divBigInt(allprod, p123).String()

	a01 := mulBigInt(a0, a1)
	for i := 4; i < n; i++ {
		l := askLcm(1, 2, i+1)
		g1, g2 := askGcd(1, i+1), askGcd(2, i+1)
		tmp := mulBigInt(l, g1, g2, g[1][2])
		tmp = divBigInt(tmp, gcdBigInt(g1, g2))
		tmp = divBigInt(tmp, a01)
		res[i] = tmp.String()
	}

	return res
}

func mulBigInt(arr ...*big.Int) *big.Int {
	n := new(big.Int)
	n.SetInt64(1)
	for _, a := range arr {
		n = n.Mul(n, a)
	}
	return n
}

func divBigInt(a, b *big.Int) *big.Int {
	n := new(big.Int)
	return n.Div(a, b)
}

func gcdBigInt(a, b *big.Int) *big.Int {
	n := new(big.Int)
	return n.GCD(nil, nil, a, b)
}

func addBigInt(a, b *big.Int) *big.Int {
	n := new(big.Int)
	return n.Add(a, b)
}

func powBigInt(a *big.Int, b int) *big.Int {
	n := new(big.Int)
	return n.Exp(a, big.NewInt(int64(b)), nil)
}

func cbrt(x *big.Int) *big.Int {
	lo := big.NewInt(1)
	hi := big.NewInt(2)
	hi.SetString("1000000000000000000000000000002", 10)

	for lo.Cmp(hi) < 0 {
		mid := divBigInt(addBigInt(lo, hi), big.NewInt(2))
		tmp := powBigInt(mid, 3)
		if tmp.Cmp(x) < 0 {
			lo = addBigInt(mid, big.NewInt(1))
		} else {
			hi = mid
		}
	}
	return lo
}

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

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		l, r := readTwoNums(reader)
		res := solve(l, r)

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

const mod = 1000000007

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

func solve(l int, r int) int {
	res := calc(r)
	if l > 0 {
		res = add(res, mod-calc(l-1))
	}
	return res
}

const H = 62

func calc(num int) int {
	if num == 0 {
		return 0
	}
	// 0...num
	var res int

	for x := 0; x < H; x++ {
		lo := 1 << x
		hi := 1 << (x + 1)
		res = add(res, calcBase(lo, min(hi, num+1), x))
		if hi > num {
			break
		}
	}

	return res
}

func calcBase(lo int, hi int, base int) int {
	if base < 2 || lo >= hi {
		return 0
	}

	if base >= 59 {
		return calcBase1(lo, hi, base)
	}

	var res int

	i := getPowerOfBase(lo, base)
	cur := pow(base, i)

	for ; cur < hi; i++ {
		next := cur * base
		tmp := min(next, hi) - max(cur, lo)
		if tmp > 0 {
			res = add(res, mul(tmp%mod, i))
		}
		cur = next
	}

	return res
}

func calcBase1(lo int, hi int, base int) int {
	var res int
	i := getPowerOfBase(lo, base)
	cur := big.NewInt(0).Exp(big.NewInt(int64(base)), big.NewInt(int64(i)), nil)
	end := big.NewInt(int64(hi))

	start := big.NewInt(int64(lo))

	for cur.Cmp(end) < 0 {
		next := big.NewInt(0).Mul(cur, big.NewInt(int64(base)))

		x := next
		if end.Cmp(next) < 0 {
			x = end
		}

		y := cur
		if cur.Cmp(start) < 0 {
			y = start
		}

		diff := x.Sub(x, y).Int64()

		if diff > 0 {
			res = add(res, mul(int(diff%mod), i))
		}

		cur = next

		i++
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func getPowerOfBase(num int, base int) int {
	// base ** z <= num
	tmp := 1
	var pow int
	for tmp <= num/base && tmp*base <= num {
		tmp *= base
		pow++
	}
	return pow
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r *= a
		}
		a *= a
		b >>= 1
	}
	return r
}

func bruteForce(l int, r int) int {
	var res int

	for i := l; i <= r; i++ {
		res = add(res, calc1(i))
	}

	return res
}

func calc1(num int) int {
	if num < 4 {
		return 0
	}

	var y int
	for (1 << (y + 1)) <= num {
		y++
	}

	tmp := 1
	var z int
	for tmp*y <= num {
		tmp *= y
		z++
	}
	return z
}

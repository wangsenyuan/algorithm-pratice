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
		X := readNInt64s(reader, n)
		res := solve(X)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const MOD = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
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

const H = 61

func solve(X []int64) int {
	n := len(X)

	// d0 = 2  (0, 0, 0) 和  (0, 0, 1)
	//         (1 & 1) * (3 + 1) = 4
	// d1 = 12 (1, 1, 0) 和  (1, 1, 1)
	//         (2 & 2) = 2 * ( 3 + 2)  = 10
	// 确定j的时候， 需要知道 j 和 其他所有 sum(x[j] | x)
	// 如果x[j][d] = 1, 那么它的贡献是 n * pow(2, d)
	//  else  cnt[d] * pow(2, d)
	cnt := make([]int, H)

	for i := 0; i < H; i++ {
		for j := 0; j < n; j++ {
			if (X[j]>>i)&1 == 1 {
				cnt[i]++
			}
		}
	}
	var res int

	for j := 0; j < n; j++ {
		var tmp int
		pw := 1
		for i := 0; i < H; i++ {
			if (X[j]>>i)&1 == 1 {
				tmp = add(tmp, mul(n, pw))
			} else {
				// it is 0
				tmp = add(tmp, mul(cnt[i], pw))
			}
			pw = mul(pw, 2)
		}
		pw = 1
		for i := 0; i < H; i++ {
			if (X[j]>>i)&1 == 1 {
				res = add(res, mul(cnt[i], mul(pw, tmp)))
			}
			pw = mul(pw, 2)
		}
	}

	return res
}

func bruteForce(X []int64) int {
	var res int
	n := len(X)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				tmp := (X[i] & X[j]) * (X[j] | X[k])
				tmp %= MOD
				res = add(res, int(tmp))
			}
		}
	}

	return res
}

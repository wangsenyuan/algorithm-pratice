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
		A := readNInt64s(reader, n)
		B := readNInt64s(reader, n)
		res := solve(A, B)
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

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

const MOD = 1000000009

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func solve(start []int64, end []int64) int {
	n := len(start)
	dp := make([][]int, 52)
	for i := 0; i < 52; i++ {
		dp[i] = make([]int, 1<<n)
	}
	var calc func(lev int, mask int, t []int64) int
	calc = func(lev int, mask int, t []int64) int {
		if dp[lev][mask] < 0 {
			if lev == 0 {
				dp[lev][mask] = 1
			} else {
				var nmask int
				ma := int64(1 << (lev - 1))
				for i := 0; i < n; i++ {
					if (mask>>i)&1 == 0 && t[i]&ma > 0 {
						nmask |= 1 << i
					}
				}
				dp[lev][mask] = calc(lev-1, mask|nmask, t)
				for i := 0; i < n; i++ {
					if (nmask>>i)&1 == 1 {
						dp[lev][mask] = modAdd(dp[lev][mask], calc(lev-1, (mask|nmask)^(1<<i), t))
					}
					if (mask>>i)&1 == 1 {
						dp[lev][mask] = modAdd(dp[lev][mask], calc(lev-1, mask|nmask, t))
					}
				}
			}
		}

		return dp[lev][mask]
	}

	var nu int
	for i := 0; i < n; i++ {
		if start[i] > 0 {
			nu++
		}
	}
	t := make([]int64, n)
	var res int
	for mask := 0; mask < 1<<nu; mask++ {
		copy(t, end)
		sign := 1
		for i, k := 0, 0; i < n; i++ {
			if start[i] > 0 {
				if (mask>>k)&1 == 1 {
					t[i] = start[i] - 1
					sign *= -1
				}
				k++
			}
		}
		for i := 0; i < len(dp); i++ {
			for j := 0; j < len(dp[i]); j++ {
				dp[i][j] = -1
			}
		}
		val := calc(51, 0, t)
		if sign < 0 {
			res = modSub(res, val)
		} else {
			res = modAdd(res, val)
		}
	}

	return res
}

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
		s := readString(reader)
		res := solve(s)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const MOD = 1000000007

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

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func inv(a int) int {
	return pow(a, MOD-2)
}

func solve(s string) int {
	cnt := make([]int, 10)
	for i := 0; i < len(s); i++ {
		cnt[int(s[i]-'0')]++
	}
	var res int
	use := make([]int, 10)
	for a := 0; a <= 10; a++ {
		for b := a; b <= 10; b++ {
			for c := b; c <= 10; c++ {
				for d := c; d <= 10; d++ {
					for j := 0; j < 10; j++ {
						use[j] = 0
					}
					var sum int
					if a < 10 {
						use[a]++
						sum += a
					}
					if b < 10 {
						use[b]++
						sum += b
					}
					if c < 10 {
						use[c]++
						sum += c
					}
					if d < 10 {
						use[d]++
						sum += d
					}

					if sum == 0 || sum%9 > 0 {
						continue
					}

					tmp := 1
					for i := 0; i < 10; i++ {
						for j := 0; j < use[i]; j++ {
							tmp = modMul(tmp, modMul(cnt[i]-j, inv(j+1)))
						}
					}
					res = modAdd(res, tmp)
				}
			}
		}
	}
	return res
}

func solve1(s string) int {
	// 1, 2, 3, 4 digits
	// 9, 9 + 9, 3 * 9, 4 * 9
	dp := make([][]int, 5)
	for i := 0; i < 5; i++ {
		dp[i] = make([]int, i*10+1)
	}

	dp[0][0] = 1

	for i := 0; i < len(s); i++ {
		x := int(s[i] - '0')
		for k := 4; k > 0; k-- {
			for j := 0; j <= (k-1)*10; j++ {
				y := x + j
				dp[k][y] = modAdd(dp[k][y], dp[k-1][j])
			}
		}
	}

	var res int

	for k := 1; k <= 4; k++ {
		for num := 1; num <= k*10; num++ {
			var sum int
			tmp := num
			for tmp > 0 {
				sum = (sum + tmp%10) % 10
				tmp /= 10
			}
			if sum == 9 {
				res = modAdd(res, dp[k][num])
			}
		}
	}

	return res
}

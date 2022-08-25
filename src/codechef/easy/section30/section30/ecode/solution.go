package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line := readNInt64s(reader, 4)
	n := line[0]
	A := line[1]
	B := line[2]
	G := line[3]
	res := solve(n, A, B, G)

	fmt.Println(res)
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

func mul(x int64, y int64, mod int64) int64 {
	var res int64
	for y > 0 {
		if y&1 == 1 {
			res += x
			res %= mod
		}
		x = (x + x) % mod
		y >>= 1
	}
	return res
}

func solve(N int64, A int64, B int64, G int64) string {
	n := int(N)
	// n <= 12
	pw := make([]int64, n+1)
	pw[0] = 1
	for i := 1; i <= n; i++ {
		// 如果直接 pw[i-1] * A 会溢出
		pw[i] = mul(pw[i-1], A, B)
	}
	value := make([][]int64, 10)
	for i := 0; i < 10; i++ {
		value[i] = make([]int64, n+1)
		for j := 0; j <= n; j++ {
			value[i][j] = pw[j] * int64(i) % B
		}
	}
	buf := make([]int, n)

	if n <= 6 {
		var dfs func(p int, res int64) bool
		dfs = func(p int, res int64) bool {
			if p == n {
				return res == G
			}
			for i := 0; i < 10; i++ {
				buf[p] = i
				if dfs(p+1, (res+value[i][n-p-1])%B) {
					return true
				}
			}
			return false
		}

		dfs(0, 0)
	} else {
		m := n - 6

		mem := make(map[int64]int)

		var suffix func(p int, res int64)

		suffix = func(p int, res int64) {
			if p == m {
				if _, ok := mem[res]; !ok {
					var tmp int
					for i := m - 1; i >= 0; i-- {
						tmp = tmp*10 + buf[i]
					}
					mem[res] = tmp
				}
				return
			}
			for i := 0; i < 10; i++ {
				buf[p] = i
				suffix(p+1, (res+value[i][m-p-1])%B)
			}
		}
		suffix(0, 0)

		var prefix func(p int, res int64) bool

		prefix = func(p int, res int64) bool {
			if p == 6 {
				res = (G - res + B) % B
				if val, ok := mem[res]; ok {
					for i := 6; i < n; i++ {
						buf[i] = val % 10
						val /= 10
					}
					return true
				}
				return false
			}
			for i := 0; i < 10; i++ {
				buf[p] = i
				if prefix(p+1, (res+value[i][n-1-p])%B) {
					return true
				}
			}
			return false
		}

		prefix(0, 0)
	}

	var str bytes.Buffer

	for i := 0; i < n; i++ {
		str.WriteByte(byte('0' + buf[i]))
	}
	return str.String()
}

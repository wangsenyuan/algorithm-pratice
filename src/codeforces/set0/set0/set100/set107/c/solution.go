package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadBytes('\n')
	var n int
	pos := readInt(line, 0, &n)
	var y int64
	pos = readInt64(line, pos+1, &y)
	var m int
	readInt(line, pos+1, &m)
	P := make([][]int, m)
	for i := 0; i < m; i++ {
		P[i] = readNNums(reader, 2)
	}
	res := solve(n, y, P)
	if len(res) == 0 {
		fmt.Println("The times have changed")
	} else {
		var buf bytes.Buffer

		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
	}
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	sign := int64(1)
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := int64(0)
	for i < len(bytes) && (bytes[i] >= '0' && bytes[i] <= '9') {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func solve(n int, y int64, P [][]int) []int {
	y -= 2001
	G := make([]int, n)
	for _, p := range P {
		a, b := p[0], p[1]
		a--
		b--
		G[b] |= 1 << a
	}
	full := (1 << n) - 1
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	dp := make([]int64, (1<<n)+1)
	mark := make([]bool, 20)
	for i := 0; i < n; i++ {
		for ans[i] = 0; ; ans[i]++ {
			if !mark[ans[i]] {
				if ans[i] >= n {
					return nil
				}
				for j := 0; j < len(dp); j++ {
					dp[j] = 0
				}
				dp[0] = 1
				for s := 0; s < full; s++ {
					if dp[s] > 0 {
						cnt := popcount(s)
						for k := 0; k < n; k++ {
							if (s&(1<<k)) == 0 && ((ans[k] == -1 || ans[k] == cnt) && ((s & G[k]) == G[k])) {
								dp[s|(1<<k)] += dp[s]
							}
						}
					}
				}
				if y >= dp[full] {
					y -= dp[full]
				} else {
					mark[ans[i]] = true
					break
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		ans[i]++
	}
	return ans
}

func popcount(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}

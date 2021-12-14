package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		W := readNNums(reader, n)
		res := solve(W)

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

const N = 1025
const T = 1 << 18

var ok_digits [T]int
var mx [N][N]int

func init() {
	tot := 1 << 18
	for state := 0; state < tot; state++ {
		for j := 0; j < 9; j++ {
			ok_digits[state] = j
			k := (state >> uint(2*j)) % 4
			if k == 0 || k == 2 {
				break
			}
		}
	}

	tot = 1 << 10
	for i := 0; i < tot; i++ {
		for j := 0; j < tot; j++ {
			var k int
			for n := 0; n < 5; n++ {
				l := (i >> uint(2*n)) % 4
				m := (j >> uint(2*n)) % 4
				if l < m {
					l = m
				}
				k |= l << uint(2*n)
			}
			mx[i][j] = k
		}
	}
}

func solve(W []int) int64 {
	sort.Ints(W)
	dp := make([]int64, T)
	dp[0] = 1
	var res int64

	dig := make([]int, 10)

	for _, num := range W {
		var sz int
		for num > 0 {
			dig[sz] = num % 10
			sz++
			num /= 10
		}
		ok := true
		for i := 0; i < sz; i++ {
			if dig[i] < 4 {
				dig[i] = 0
			} else if dig[i] == 4 {
				dig[i] = 1
			} else if dig[i] < 7 {
				dig[i] = 2
			} else if dig[i] == 7 {
				dig[i] = 3
			} else {
				ok = false
				break
			}
		}
		if !ok {
			// invalid
			continue
		}
		var tmp int
		for i := 0; i < sz; i++ {
			tmp |= dig[i] << uint(2*i)
		}

		for i := 1<<uint(2*sz) - 1; i >= 0; i-- {
			if dp[i] == 0 {
				continue
			}
			mask := (mx[i/(1<<10)][tmp/(1<<10)] << 10) | mx[i%(1<<10)][tmp%(1<<10)]
			if ok_digits[mask] == sz {
				res += dp[i]
			}
			dp[mask] += dp[i]
		}
	}

	return res
}

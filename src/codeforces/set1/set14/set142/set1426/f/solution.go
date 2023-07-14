package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	readNum(reader)
	s := readString(reader)
	res := solve(s)
	fmt.Println(res)
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

const MOD = 1e9 + 7

func add(a int, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func mul(a int, b int) int {
	c := int64(a) * int64(b)
	c %= MOD
	return int(c)
}

func solve(s string) int {
	// abc
	n := len(s)
	pw := make([]int, n+1)
	pw[0] = 1
	for i := 1; i <= n; i++ {
		pw[i] = mul(pw[i-1], 3)
	}

	dp := make([]int, 16)
	fp := make([]int, 16)

	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < 16; j++ {
			fp[j] = 0
		}

		for j := 0; j <= 3; j++ {
			for k := 0; k <= 3; k++ {
				if dp[j*4+k] == 0 {
					continue
				}
				fp[j*4+k] = add(fp[j*4+k], dp[j*4+k])
				if j < 3 && (s[i] == '?' || int(s[i]-'a') == j) {
					nk := k
					if s[i] == '?' {
						nk++
					}
					fp[(j+1)*4+nk] = add(fp[(j+1)*4+nk], dp[j*4+k])
				}
			}
		}
		copy(dp, fp)
	}
	var cnt int
	for i := 0; i < n; i++ {
		if s[i] == '?' {
			cnt++
		}
	}
	var ans int
	for i := 0; i <= 3; i++ {
		if cnt >= i {
			ans = add(ans, mul(dp[3*4+i], pw[cnt-i]))
		}
	}
	return ans
}

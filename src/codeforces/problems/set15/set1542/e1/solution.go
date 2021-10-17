package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := 1

	for tc > 0 {
		tc--
		n, mod := readTwoNums(reader)
		res := solve(n, mod)
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

func solve(n int, mod int) int {
	f := make([][]int64, n+1)
	s := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int64, n*(n-1)/2+1)
		s[i] = make([]int64, n*(n-1)/2+1)
	}
	f[0][0] = 1
	s[0][0] = 1

	for i := 1; i < len(s[0]); i++ {
		s[0][i] = 1
	}

	MOD := int64(mod)

	for i := 1; i <= n; i++ {
		for j := 0; j <= n*(n-1)/2; j++ {
			if j-i >= 0 {
				f[i][j] = (s[i-1][j] - s[i-1][j-i] + MOD) % MOD
			} else {
				f[i][j] = s[i-1][j]
			}
			if j > 0 {
				s[i][j] = (s[i][j-1] + f[i][j]) % MOD
			} else {
				s[i][j] = f[i][j]
			}
		}
	}

	ans := make([]int64, n+1)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			for k := j + 1; k <= i; k++ {
				for u := 0; u <= (i-1)*(i-2)/2; u++ {
					t := u - (k - j) - 1
					if t < 0 {
						continue
					}
					ans[i] = (ans[i] + f[i-1][u]*s[i-1][t]%MOD) % MOD
				}
			}
		}
	}

	for i := 2; i <= n; i++ {
		ans[i] = (ans[i] + int64(i)*ans[i-1]%MOD) % MOD
	}

	return int(ans[n])
}

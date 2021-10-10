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
		n := readNum(reader)
		S := make([]string, n)
		for i := 0; i < n; i++ {
			S[i], _ = reader.ReadString('\n')
		}
		res := solve(n, S)
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

const MOD = 998244353

func solve(n int, S []string) int {
	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		if S[i][0] == '+' {
			readInt([]byte(S[i]), 2, &a[i+1])
		}
	}

	var ans int64

	f := make([][]int64, n+2)
	for i := 0; i <= n+1; i++ {
		f[i] = make([]int64, n+2)
	}

	for t := 1; t <= n; t++ {
		if a[t] == 0 {
			continue
		}
		for i := 0; i <= n+1; i++ {
			for j := 0; j <= n+1; j++ {
				f[i][j] = 0
			}
		}
		f[0][0] = 1
		for i := 1; i <= n; i++ {
			for j := 0; j <= i; j++ {
				if a[i] == 0 {
					// a minus
					if i <= t || j > 0 {
						f[i][max(j-1, 0)] = modAdd(f[i][max(j-1, 0)], f[i-1][j])
					}
				} else {
					if a[i] < a[t] || (a[i] == a[t] && i < t) {
						f[i][j+1] = modAdd(f[i][j+1], f[i-1][j])
					} else {
						f[i][j] = modAdd(f[i][j], f[i-1][j])
					}
				}

				if i != t {
					f[i][j] = modAdd(f[i][j], f[i-1][j])
				}
			}
		}

		for i := 0; i <= n; i++ {
			ans = modAdd(ans, f[n][i]*int64(a[t])%MOD)
		}
	}

	return int(ans)
}

func modAdd(a, b int64) int64 {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

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
		mat := make([]string, n)
		for i := 0; i < n; i++ {
			mat[i] = readString(reader)
		}
		res := solve(mat)
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

const inf = 1 << 30

func solve(mat []string) int {
	n := len(mat)
	f := make([]byte, n)
	d1 := make([]byte, n*2)
	d2 := make([]byte, n*2)

	var ans int

	for i, s := range mat {
		for j, c := range []byte(s) {
			f[j] ^= d1[i-j+n] ^ d2[i+j]
			if f[j] != c&1 {
				ans++
				f[j] ^= 1
				d1[i-j+n] ^= 1
				d2[i+j] ^= 1
			}
		}
	}

	return ans
}

func solve1(mat []string) int {
	n := len(mat)

	// 不大对，每个格子最多操作一次
	// 如果对于i,j 如果它的奇偶性和上面影响到它的数量
	// dp[i][j] 表示到i, j为止时,将其变为0时，必须发生变化的次数
	val := make([][]int, n)
	sum := make([][]int, n)
	for i := 0; i < n; i++ {
		val[i] = make([]int, n)
		sum[i] = make([]int, n)
	}

	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 1 {
				for k := max(0, j-1); k <= min(j+1, n-1); k++ {
					sum[i][j] ^= sum[i-1][k]
				}
			} else if i > 1 {
				if j == 0 {
					sum[i][j] ^= sum[i-2][j]
				} else {
					sum[i][j] ^= sum[i-1][j-1]
				}
				if j == n-1 {
					sum[i][j] ^= sum[i-2][j]
				} else {
					sum[i][j] ^= sum[i-1][j+1]
				}
				sum[i][j] ^= sum[i-2][j]
				sum[i][j] ^= val[i-1][j]
			}
			x := int(mat[i][j] - '0')
			if sum[i][j]^x != 0 {
				res++
				sum[i][j] ^= 1
				val[i][j] = 1
			}
		}
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
	return a + b - min(a, b)
}

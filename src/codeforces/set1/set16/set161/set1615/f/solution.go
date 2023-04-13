package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		s := readString(reader)[:n]
		t := readString(reader)[:n]
		res := solve(s, t)
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

const MOD = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func sub(a, b int) int {
	return add(a, MOD-b)
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

func solve(A, B string) int {
	n := len(A)

	A = flipString(A)
	B = flipString(B)
	pref := getDp(n, A, B)

	A = reverse(A)
	B = reverse(B)
	suf := getDp(n, A, B)

	var res int
	for i := 1; i < n; i++ {
		for j := -n; j <= n; j++ {
			tmp := mul(pref[i][j+2*n], suf[n-i][-j+2*n])
			tmp = mul(tmp, abs(j))
			res = add(res, tmp)
		}
	}

	return res
}

func getDp(n int, A, B string) [][]int {
	pref := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		pref[i] = make([]int, 4*n+1)
	}

	getValue := func(x byte) int {
		if x == '?' {
			return -1
		}
		return int(x - '0')
	}

	pref[0][2*n] = 1

	for i := 1; i <= n; i++ {
		for j := -i; j <= i; j++ {
			a := getValue(A[i-1])
			b := getValue(B[i-1])
			if a < 0 && b < 0 {
				for x := 0; x <= 1; x++ {
					for y := 0; y <= 1; y++ {
						nj := j - (x - y)
						pref[i][nj+2*n] = add(pref[i][nj+2*n], pref[i-1][j+2*n])
					}
				}
			} else if a < 0 {
				for x := 0; x <= 1; x++ {
					nj := j - (x - b)
					pref[i][nj+2*n] = add(pref[i][nj+2*n], pref[i-1][j+2*n])
				}
			} else if b < 0 {
				for y := 0; y <= 1; y++ {
					nj := j - (a - y)
					pref[i][nj+2*n] = add(pref[i][nj+2*n], pref[i-1][j+2*n])
				}
			} else {
				nj := j - (a - b)
				pref[i][nj+2*n] = add(pref[i][nj+2*n], pref[i-1][j+2*n])
			}
		}
	}

	return pref
}

func flipString(s string) string {
	buf := []byte(s)
	for i := 1; i < len(buf); i += 2 {
		if buf[i] == '1' {
			buf[i] = '0'
		} else if buf[i] == '0' {
			buf[i] = '1'
		}
	}
	return string(buf)
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

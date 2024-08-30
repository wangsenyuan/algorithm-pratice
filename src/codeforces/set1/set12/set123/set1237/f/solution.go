package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	existing := make([][]int, k)
	for i := 0; i < k; i++ {
		existing[i] = readNNums(reader, 4)
	}
	res := solve(n, m, existing)
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

const mod = 998244353

func mul(a, b int) int {
	return a * b % mod
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

const MX = 3610

var F [MX]int
var I [MX]int

func init() {
	F[0] = 1
	for i := 1; i < MX; i++ {
		F[i] = mul(i, F[i-1])
	}
	I[MX-1] = pow(F[MX-1], mod-2)
	for i := MX - 2; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}
}

func perm(n int, k int) int {
	return mul(F[n], I[n-k])
}

func solve(n int, m int, existing [][]int) int {

	badRow := make([]bool, n)
	badCol := make([]bool, m)

	for _, cur := range existing {
		badRow[cur[0]-1] = true
		badRow[cur[2]-1] = true
		badCol[cur[1]-1] = true
		badCol[cur[3]-1] = true
	}

	calc := func(ban []bool) ([]int, int) {
		n := len(ban)
		f := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, n/2+1)
			f[i][0] = 1
		}
		for i := 1; i < n; i++ {
			for j := 1; j <= (i+1)/2; j++ {
				f[i+1][j] = f[i][j]
				if !ban[i] && !ban[i-1] {
					f[i+1][j] = (f[i+1][j] + f[i-1][j-1]) % mod
				}
			}
		}
		empty := 0
		for _, b := range ban {
			if !b {
				empty++
			}
		}
		return f[n], empty
	}

	var ans int

	f, emptyR := calc(badRow)
	g, emptyC := calc(badCol)
	for i, v := range f { // i 个竖放
		for j, w := range g { // j 个横放
			if j > emptyR-i*2 || i > emptyC-j*2 {
				break
			}
			ans = (ans + v*w%mod*perm(emptyC-j*2, i)%mod*perm(emptyR-i*2, j)) % mod
		}
	}

	return ans
}

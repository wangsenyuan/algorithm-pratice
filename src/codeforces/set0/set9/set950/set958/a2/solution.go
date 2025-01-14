package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	_, _, _, _, res := process(reader)

	fmt.Println(res[0], res[3])
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func process(reader *bufio.Reader) (n int, m int, a []string, b []string, res []int) {
	n, m = readTwoNums(reader)
	a = make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	b = make([]string, m)
	for i := range m {
		b[i] = readString(reader)
	}
	res = solve(a, b)
	return
}

func solve(a []string, b []string) []int {
	n := len(a)
	m := len(b)

	bases := make([]Hash, m*m+1)
	bases[0] = NewHash(1)

	for i := 1; i <= m*m; i++ {
		bases[i] = bases[i-1].MulInt(26)
	}

	calc := func(s []string) [][]Hash {
		// dp[i][j] = 以(i, j)为右下角的，k*k的一块区域的hash
		h := len(s)
		w := len(s[0])
		dp := make([][]Hash, h)
		for i := range h {
			dp[i] = make([]Hash, w)
		}
		row := make([]Hash, h)
		for j := 0; j < w; j++ {
			if j >= m {
				for i := 0; i < h; i++ {
					x := int(s[i][j-m] - 'a')
					v := NewHash(x).Mul(bases[m-1])
					row[i] = row[i].Sub(v)
				}
			}
			for i := 0; i < h; i++ {
				row[i] = row[i].Mul(bases[1]).AddInt(int(s[i][j] - 'a'))
			}
			if j >= m-1 {
				var cur Hash

				for i := 0; i < h; i++ {
					if i >= m {
						cur = cur.Sub(row[i-m].Mul(bases[m*(m-1)]))
					}
					cur = cur.Mul(bases[m]).Add(row[i])
					if i >= m-1 {
						dp[i][j] = cur
					}
				}
			}
		}

		return dp
	}

	x := calc(a)
	y := calc(b)

	pos := make(map[Hash]int)

	for i := m - 1; i < m; i++ {
		for j := m - 1; j < n; j++ {
			cur := y[i][j]
			pos[cur] = i*n + j
		}
	}

	for i := m - 1; i < n; i++ {
		for j := m - 1; j < m; j++ {
			cur := x[i][j]
			if v, ok := pos[cur]; ok {
				i1, j1 := v/n, v%n
				return []int{i - m + 2, j - m + 2, i1 - m + 2, j1 - m + 2}
			}
		}
	}

	return nil
}

var MOD = [...]uint64{1000000007, 1000000009}

type Hash struct {
	h [2]uint64
}

func NewHash(v int) Hash {
	x := uint64(v)
	h := [2]uint64{x % MOD[0], x % MOD[1]}
	return Hash{h}
}

func (this Hash) Sub(that Hash) Hash {
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + MOD[i] - that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) Add(that Hash) Hash {
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) AddInt(v int) Hash {
	x := uint64(v)
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + x%MOD[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) Mul(that Hash) Hash {
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] * that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) MulInt(v int) Hash {
	x := uint64(v)
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] * x) % MOD[i]
	}
	return Hash{h}
}

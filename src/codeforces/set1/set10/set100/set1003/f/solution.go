package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
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

func solve(s string) int {
	n := len(s)

	words := strings.Split(s, " ")
	m := len(words)

	type Word struct {
		key Hash
		ln  int
	}

	ws := make([]Word, m)
	for i, w := range words {
		var key Hash
		for j := 0; j < len(w); j++ {
			key = key.MulInt(27).AddInt(int(w[j]-'a') + 1)
		}
		ws[i] = Word{key, len(w)}
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := m - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if ws[i] == ws[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
		}
	}

	var best int

	for i := 0; i < m; i++ {
		var sum int
		for j := 1; i+j <= m; j++ {
			sum += ws[i+j-1].ln
			cnt := 1
			for k := i + j; k+j <= m; k++ {
				if dp[i][k] >= j {
					cnt++
					k += j - 1
				}
			}
			if cnt > 1 {
				best = max(best, cnt*(sum-1))
			}
		}
	}

	return n - best
}

var MOD = [...]uint{1000000007, 1000000009, 1000000023}

type Hash struct {
	h [3]uint
}

func NewHash(x uint) Hash {
	h := [3]uint{uint(x) % MOD[0], uint(x) % MOD[1], uint(x) % MOD[2]}
	return Hash{h}
}

func (this Hash) Sub(that Hash) Hash {
	h := [3]uint{0, 0}
	for i := 0; i < 3; i++ {
		h[i] = (this.h[i] + MOD[i] - that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) Add(that Hash) Hash {
	h := [3]uint{0, 0}
	for i := 0; i < 3; i++ {
		h[i] = (this.h[i] + that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) AddInt(x int) Hash {
	h := [3]uint{0, 0}
	for i := 0; i < 3; i++ {
		h[i] = (this.h[i] + uint(x)%MOD[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) Mul(that Hash) Hash {
	h := [3]uint{0, 0}
	for i := 0; i < 3; i++ {
		h[i] = (this.h[i] * that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) MulInt(x int) Hash {
	h := [3]uint{0, 0}
	for i := 0; i < 3; i++ {
		h[i] = (this.h[i] * uint(x)) % MOD[i]
	}
	return Hash{h}
}

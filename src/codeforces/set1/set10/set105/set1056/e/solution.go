package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)
	t := readString(reader)
	res := solve(s, t)

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
func solve(s string, t string) int {
	cnt := make([]int, 2)
	for _, x := range s {
		cnt[int(x-'0')]++
	}

	bin := "01"
	// cnt[0] and cnt[1]
	if s[0] == '1' {
		cnt[0], cnt[1] = cnt[1], cnt[0]
		bin = "10"
	}

	n := len(t)
	// len(r0) * cnt[0] + len(r1) * cnt[1] = n
	pref := make([]Hash, n+1)
	pref[0] = NewHash(0)

	bases := make([]Hash, n+1)
	bases[0] = NewHash(1)

	B := Hash{[...]uint{10, 10}}

	for i, x := range []byte(t) {
		bases[i+1] = bases[i].Mul(B)
		v := int(x - '0')
		pref[i+1] = NewHash(v).Mul(bases[i]).Add(pref[i])
	}

	origin := NewHash(0)

	check := func(m int) bool {
		if m*cnt[0] >= n {
			return false
		}
		k := n - m*cnt[0]
		if k%cnt[1] != 0 {
			return false
		}
		k /= cnt[1]
		// 分别是长度
		r0 := pref[m]
		r1 := origin
		var p1 int
		var j int
		for i := 0; i < len(s); i++ {
			if s[i] == bin[0] {
				if j+m > n {
					return false
				}
				cur := pref[j+m].Sub(pref[j])
				if r0.Mul(bases[j]) != cur {
					return false
				}
				j += m
			} else {
				if j+k > n {
					return false
				}
				cur := pref[j+k].Sub(pref[j])
				if r1 == origin {
					r1 = cur
					p1 = j
					if r1 == r0.Mul(bases[j]) {
						// r0 = r1
						return false
					}
				} else if r1.Mul(bases[j-p1]) != cur {
					return false
				}
				j += k
			}
		}
		return true
	}

	var res int
	for i := 1; i < n; i++ {
		if check(i) {
			res++
		}
	}

	return res
}

var MOD = [...]uint{1000000007, 1000000009}

type Hash struct {
	h [2]uint
}

func NewHash(x int) Hash {
	h := [2]uint{uint(x) % MOD[0], uint(x) % MOD[1]}
	return Hash{h}
}

func (this Hash) Sub(that Hash) Hash {
	h := [2]uint{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + MOD[i] - that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) Add(that Hash) Hash {
	h := [2]uint{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) AddInt(x int) Hash {
	h := [2]uint{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + uint(x)%MOD[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) Mul(that Hash) Hash {
	h := [2]uint{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] * that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) MulInt(x int) Hash {
	h := [2]uint{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] * uint(x)) % MOD[i]
	}
	return Hash{h}
}

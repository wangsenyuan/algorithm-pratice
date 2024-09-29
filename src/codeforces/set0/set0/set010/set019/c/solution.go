package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)

	res := solve(a)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
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

func solve(a []int) []int {

	var B = Hash{[...]uint{10007, 1000001}}

	n := len(a)
	base := make([]Hash, n+1)
	base[0] = NewHash(1)

	for i := 1; i <= n; i++ {
		base[i] = base[i-1].Mul(B)
	}

	pref := make([]Hash, n+1)
	pref[0] = NewHash(0)
	var cur int

	get := func(l int, r int) Hash {
		res := pref[r+1]
		res = res.Sub(pref[l].Mul(base[r-l+1]))
		return res
	}

	check := func(l int, r int) bool {
		// 一半的长度是 r - l
		if l < cur || l-cur+1 < r-l {
			return false
		}
		x := get(l+1, r)
		y := get(l-(r-l)+1, l)
		return x == y
	}

	pos := make(map[int][]int)

	for i, x := range a {
		pref[i+1] = pref[i].Mul(B).AddInt(x)
		for _, j := range pos[x] {
			if check(j, i) {
				cur = j + 1
				break
			}
		}
		if pos[x] == nil {
			pos[x] = make([]int, 0, 1)
		}
		pos[x] = append(pos[x], i)
	}

	return a[cur:]
}

var MOD = [...]uint{1000000007, 1000000009}

type Hash struct {
	h [2]uint
}

func NewHash(x uint) Hash {
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

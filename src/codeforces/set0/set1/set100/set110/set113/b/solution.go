package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	a := readString(reader)
	b := readString(reader)
	c := readString(reader)
	res := solve(a, b, c)
	fmt.Println(res)
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(t, s1, s2 string) int {
	n := len(t)
	ht := CalcHash(t)
	h1 := CalcHash(s1)[len(s1)]
	h2 := CalcHash(s2)[len(s2)]

	bs := make([]Key, n+1)

	bs[0] = Key{1, 1}

	for i := 1; i <= n; i++ {
		bs[i] = bs[i-1].Add(0)
	}

	getKey := func(a int, b int) Key {
		// h[a] * bs(b - a) + diff = hash[b]
		tmp := ht[a].Mul2(bs[b-a])
		return ht[b].Sub(tmp)
	}

	valid := make(map[Key]bool)

	for i := 0; i < n; i++ {
		for j := i + max(len(s1), len(s2)); j <= n; j++ {
			pref := h1 == getKey(i, i+len(s1))
			suf := h2 == getKey(j-len(s2), j)
			if pref && suf {
				valid[getKey(i, j)] = true
			}
		}
	}
	return len(valid)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func CalcHash(s string) []Key {
	hs := make([]Key, len(s)+1)
	hs[0] = Key{0, 0}

	for i := 0; i < len(s); i++ {
		hs[i+1] = hs[i].Add(int(s[i] - 'a' + 1))
	}
	return hs
}

const M1 = 1000000007
const M2 = 1000000009

const BASE1 = 27
const BASE2 = 1337

type Key struct {
	first  int64
	second int64
}

func (this Key) Add(v int) Key {
	first := (this.first*BASE1%M1 + int64(v)) % M1
	second := (this.second*BASE2%M2 + int64(v)) % M2
	return Key{first, second}
}

func (this Key) Mul(a, b int64) Key {
	first := this.first * a % M1
	second := this.second * b % M2
	return Key{first, second}
}

func (this Key) Mul2(that Key) Key {
	first := this.first * that.first % M1
	second := this.second * that.second % M2
	return Key{first, second}
}

func (this Key) Sub(that Key) Key {
	first := (this.first + M1 - that.first) % M1
	second := (this.second + M2 - that.second) % M2
	return Key{first, second}
}

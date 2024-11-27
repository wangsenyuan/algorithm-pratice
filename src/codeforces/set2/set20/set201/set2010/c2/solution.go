package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)

	ok, res := solve(s)

	if !ok {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	fmt.Println(res)
}

func solve(s string) (bool, string) {
	n := len(s)
	next := make([]int, n)

	for i := 1; i < n; i++ {
		j := next[i-1]
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}

	m := next[n-1]

	if 2*m <= n {
		return false, ""
	}
	return true, s[:m]
}

func solve1(s string) (bool, string) {
	n := len(s)
	base := make([]Key, n+1)
	base[0] = Key{1, 1}
	for i := 1; i <= n; i++ {
		base[i] = base[i-1].Mul(B1, B2)
	}

	pref := make([]Key, n+1)
	pref[0] = Key{}

	for i := 1; i <= n; i++ {
		v := int(s[i-1] - 'a')
		// 这里会自动乘以系数
		pref[i] = pref[i-1].Add(v)
	}

	for l, r := 1, n-2; l <= r; l, r = l+1, r-1 {
		// t = s[l...r+1]
		// x = s[:l]
		// y = s[r+1:]
		u := pref[r+1]
		v := pref[n].Sub(pref[l].Mul2(base[n-l]))

		if u != v {
			continue
		}
		w := pref[r+1].Sub(pref[l].Mul2(base[r-l+1]))
		// w必须既是前缀，又是后缀
		a := pref[r-l+1]
		if w != a {
			continue
		}
		b := pref[n].Sub(pref[n-(r-l+1)].Mul2(base[r-l+1]))
		if w != b {
			continue
		}
		return true, s[:r+1]
	}

	return false, ""
}

const M = 1000000007

const B1 = 29
const B2 = 31

type Key struct {
	first  uint
	second uint
}

func (this Key) Add(v int) Key {
	first := (this.first*B1%M + uint(v)) % M
	second := (this.second*B2%M + uint(v)) % M
	return Key{first, second}
}

func (this Key) Mul(a, b int) Key {
	first := this.first * uint(a) % M
	second := this.second * uint(b) % M
	return Key{first, second}
}

func (this Key) Mul2(that Key) Key {
	first := this.first * that.first % M
	second := this.second * that.second % M
	return Key{first, second}
}

func (this Key) Sub(that Key) Key {
	first := (this.first + M - that.first) % M
	second := (this.second + M - that.second) % M
	return Key{first, second}
}

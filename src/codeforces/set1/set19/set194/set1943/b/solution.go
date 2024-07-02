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
		n, m := readTwoNums(reader)
		s := readString(reader)[:n]
		queries := make([][]int, m)
		for i := 0; i < m; i++ {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(s, queries)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		}
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

func solve(s string, queries [][]int) []int {
	n := len(s)
	// p1[i] = 和s[i]不同的最近的字符的位置
	p1 := make([]int, n)
	// dp[i] = 和s[i-1, i] 不同的位置
	p2 := make([]int, n)
	for i := 0; i < n; i++ {
		p1[i] = -1
		p2[i] = -1
	}

	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			p1[i] = i - 1
		} else {
			p1[i] = p1[i-1]
		}
		if i >= 2 && s[i] == s[i-2] {
			p2[i] = max(p2[i-2], p2[i-1])
		} else {
			p2[i] = i - 2
		}
	}

	base := make([]Key, n+1)
	base[0] = Key{1, 1}
	pref := make([]Key, n+1)
	pref[0] = Key{0, 0}

	for i := 1; i <= n; i++ {
		v := int(s[i-1]-'a') + 1
		base[i] = base[i-1].Mul(BASE1, BASE2)
		pref[i] = pref[i-1].Add(v)
	}

	suf := make([]Key, n+1)
	suf[n] = Key{0, 0}

	for i := n - 1; i >= 0; i-- {
		v := int(s[i]-'a') + 1
		suf[i] = suf[i+1].Add(v)
	}

	isPalidrome := func(l int, r int) bool {
		ln := r - l + 1
		a := pref[r+1].Sub(pref[l].Mul2(base[ln]))
		b := suf[l].Sub(suf[r+1].Mul2(base[ln]))
		return a == b
	}

	find := func(l int, r int) int {
		if p1[r] < l {
			// all a
			return 0
		}
		// 2 + 4 + ... + (r - l + 1)
		ln := r - l + 1
		m := (ln-2)/2 + 1
		res := (2 + 2*m) * m / 2
		if ln%2 == 0 && isPalidrome(l, r) {
			res -= ln
		}
		if p2[r] >= l && ln >= 3 {
			// 3， 5， 7。。。
			// 3 + 2 * k <= ln
			k := (ln - 3) / 2
			res += (3 + 3 + 2*k) * (k + 1) / 2
			if ln%2 == 1 && isPalidrome(l, r) {
				res -= ln
			}
		}
		// else ababab

		return res
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = find(cur[0]-1, cur[1]-1)
	}

	return ans
}

const M1 = 1000000007
const M2 = 1000000009

const BASE1 = 27
const BASE2 = 1337

type Key struct {
	first  int
	second int
}

func NewKey(s string) Key {
	var key Key
	for i := 0; i < len(s); i++ {
		key = key.Add(int(s[i]-'a') + 1)
	}
	return key
}

func (this Key) Add(v int) Key {
	first := (this.first*BASE1%M1 + v) % M1
	second := (this.second*BASE2%M2 + v) % M2
	return Key{first, second}
}

func (this Key) Mul(a, b int) Key {
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

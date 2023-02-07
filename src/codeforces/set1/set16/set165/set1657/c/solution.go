package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		readNum(reader)

		s := readString(reader)
		n, r := solve(s)

		buf.WriteString(fmt.Sprintf("%d %d\n", n, r))
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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
func solve(s string) (int, int) {
	n := len(s)
	var r int
	var res int

	for r < n {
		if s[r] == '(' {
			if r+1 == n {
				break
			}
			res++
			r += 2
		} else {
			i := r + 1
			for i < n && s[i] == '(' {
				i++
			}
			if i < n {
				// s[r] == ')'
				res++
				r = i + 1
			} else {
				break
			}
		}
	}
	return res, n - r
}
func solve1(s string) (int, int) {
	n := len(s)

	//if good, then remove
	// good = regular or palindrome
	// regular = open == 0
	// palindrome =?

	var open int
	var cnt int
	var r int

	base := NewKey(1)

	a := NewKey(0)
	b := NewKey(0)
	var l int
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			open++
			a = a.MulNum(3).AddNum(1)
			b = base.MulNum(1).AddKey(b)
		} else {
			open--
			a = a.MulNum(3).AddNum(2)
			b = base.MulNum(2).AddKey(b)
		}
		l++

		base = base.MulNum(3)

		if a == b && l > 1 || open == 0 && s[i-l+1] == '(' {
			cnt++
			r = i + 1
			base = NewKey(1)
			a = NewKey(0)
			b = NewKey(0)
			l = 0
			open = 0
		}
	}

	return cnt, n - r
}

const M1 = 1000000007
const M2 = 10000000009

type Key struct {
	first  int64
	second int64
}

func NewKey(x int) Key {
	X := int64(x)
	key := Key{X % M1, X % M2}
	return key
}

func (this Key) AddNum(x int) Key {
	X := int64(x)
	first := (this.first + X) % M1
	second := (this.second + X) % M2
	return Key{first, second}
}

func (this Key) MulNum(x int) Key {
	X := int64(x)
	first := this.first * X % M1
	second := this.second * X % M2
	return Key{first, second}
}

func (this Key) AddKey(that Key) Key {
	first := (this.first + that.first) % M1
	second := (this.second + that.second) % M2
	return Key{first, second}
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		ss := make([]string, n)
		for i := 0; i < n; i++ {
			ss[i] = readString(reader)
		}
		res := solve(ss)
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

const inf = 1 << 30

func solve(songes []string) int {
	n := len(songes)
	dp := make([][]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -inf
		}
	}
	for i := 0; i < n; i++ {
		dp[1<<i][i] = 1
	}

	ss := make([]Song, n)
	for i := 0; i < n; i++ {
		ss[i] = NewSong(i, songes[i])
	}

	check := func(i, j int) bool {
		return ss[i].genre == ss[j].genre || ss[i].writer == ss[j].writer
	}

	for state := 1; state < 1<<n; state++ {
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				for j := 0; j < n; j++ {
					if (state>>j)&1 == 0 && check(i, j) {
						dp[state|(1<<j)][j] = max(dp[state|(1<<j)][j], dp[state][i]+1)
					}
				}
			}
		}
	}
	var best int
	for state := 1; state < 1<<n; state++ {
		for i := 0; i < n; i++ {
			best = max(best, dp[state][i])
		}
	}
	return n - best
}

type Song struct {
	id     int
	writer Key
	genre  Key
}

func NewSong(id int, s string) Song {
	ss := strings.Split(s, " ")
	g := NewKey(ss[0])
	w := NewKey(ss[1])
	return Song{id, w, g}
}

const M1 = 1000000007
const M2 = 1000000009

const BASE1 = 27
const BASE2 = 1337

type Key struct {
	first  int64
	second int64
}

func NewKey(s string) Key {
	var key Key
	for i := 0; i < len(s); i++ {
		key = key.Add(int(s[i]-'a') + 1)
	}
	return key
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

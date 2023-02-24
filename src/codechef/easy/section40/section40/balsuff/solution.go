package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		s := readString(reader)[:n]
		res := solve(s, k)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(res)
			buf.WriteByte('\n')
		}
	}
	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

func solve(S string, K int) string {
	// max(freq[C]) - min(freq[C]) <= K for suffix
	n := len(S)
	// n <= 1e5
	freq := make([]int, 26)
	for i := 0; i < n; i++ {
		freq[int(S[i]-'a')]++
	}

	P := make([]Pair, 0, 26)
	for i := 0; i < 26; i++ {
		if freq[i] > 0 {
			P = append(P, Pair{freq[i], i})
		}
	}

	sort := func() {
		sort.Slice(P, func(i, j int) bool {
			return P[i].Less(P[j])
		})
	}

	sort()

	m := len(P)

	if m == 1 {
		return S
	}

	if P[0].first-P[m-1].first > K {
		return ""
	}

	check := func(i int) bool {
		// 如果减少一个P[i].second, 仍然能够满足条件
		a := P[0].first
		if i == 0 {
			a = max(a-1, P[1].first)
		}
		b := P[m-1].first
		if i == m-1 {
			b--
		} else {
			b = min(b, P[i].first-1)
		}

		return a-b <= K
	}

	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		x := -1
		for j := 0; j < m && P[j].first > 0; j++ {
			// 如果获取一个P[j].second, 后续的结果，仍然满足条件，那么就作为一个选项
			if check(j) {
				if x < 0 || P[x].second > P[j].second {
					x = j
				}
			}
		}
		if x < 0 {
			return ""
		}
		buf[i] = byte(P[x].second + 'a')
		P[x] = Pair{P[x].first - 1, P[x].second}
		sort()
	}

	return string(buf)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}

func (this Pair) Less(that Pair) bool {
	return this.first > that.first
}

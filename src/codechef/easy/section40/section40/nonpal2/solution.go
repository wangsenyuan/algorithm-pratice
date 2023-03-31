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
	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		s := readString(reader)[:n]
		res := solve(s)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
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

func solve(A string) string {
	// cnt[x] 表示x的频率
	// 从最多的开始排， 两个位置至少要大于等于2
	// cnt[x] > n / 3 那么就无法重排出无回文的字串
	// 现在cnt[x] <= n / 3
	// 是不是一定可以呢？
	freq := make([]int, 26)
	n := len(A)
	limit := (n + 2) / 3

	c1 := make([]int, 0, limit)
	c2 := make([]int, 0, limit)
	c3 := make([]int, 0, limit)

	for i := 1; i <= n; i++ {
		freq[int(A[i-1]-'a')]++
		if i%3 == 1 {
			c1 = append(c1, i)
		} else if i%3 == 2 {
			c2 = append(c2, i)
		} else {
			c3 = append(c3, i)
		}
	}
	var bad int

	type Pair struct {
		first  int
		second int
	}

	P := make([]Pair, 0, 26)

	for i, v := range freq {
		if v > limit {
			return ""
		}
		if v == limit {
			bad++
		}
		if v > 0 {
			P = append(P, Pair{v, i})
		}
	}

	afford := n % 3
	if afford == 0 {
		afford = 3
	}

	if bad > afford {
		return ""
	}

	sort.Slice(P, func(i, j int) bool {
		return P[i].first > P[j].first
	})

	buf := make([]byte, n)

	fillUp := func(c []int, x int) {
		for _, i := range c {
			buf[i-1] = byte(x + 'a')
		}
	}
	pos := make([]int, 3)

	for _, it := range P {
		if it.first == limit {
			if len(c1) == limit && pos[1] == 0 {
				fillUp(c1, it.second)
				pos[1] = limit
			} else if len(c2) == limit && pos[2] == 0 {
				fillUp(c2, it.second)
				pos[2] = limit
			} else {
				fillUp(c3, it.second)
				pos[0] = limit
			}
			continue
		} else {
			for i := 0; i < it.first; i++ {
				var j int
				if pos[0] < len(c3) {
					j = c3[pos[0]]
					pos[0]++
				} else if pos[2] < len(c2) {
					j = c2[pos[2]]
					pos[2]++
				} else {
					j = c1[pos[1]]
					pos[1]++
				}
				buf[j-1] = byte(it.second + 'a')
			}
		}
	}

	return string(buf)
}

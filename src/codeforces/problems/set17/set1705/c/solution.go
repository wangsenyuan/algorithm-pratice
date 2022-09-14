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
		n, m, k := readThreeNums(reader)
		s := readString(reader)
		s = s[:n]
		ops := make([][]int64, m)
		for i := 0; i < m; i++ {
			ops[i] = readNInt64s(reader, 2)
		}
		ques := make([]int64, k)
		for i := 0; i < k; i++ {
			ques[i] = readNInt64s(reader, 1)[0]
		}

		res := solve(n, s, ops, ques)

		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%c\n", x))
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

func solve(n int, s string, ops [][]int64, ques []int64) []byte {
	// 生成完整的字符串S肯定是不行的
	// 需要知道每个位置在原始s中的位置即可
	str := NewStr(s)

	for _, op := range ops {
		l, r := op[0], op[1]
		l--
		r--
		str = str.copy_paste(l, r)
	}

	res := make([]byte, len(ques))

	for i, k := range ques {
		res[i] = str.get_at(k)
	}

	return res
}

type Str struct {
	base     string
	delegate *Str
	length   int64
	left     int64
}

func NewStr(s string) *Str {
	return &Str{s, nil, int64(len(s)), 0}
}

func (s *Str) copy_paste(l, r int64) *Str {
	length := r - l + 1
	return &Str{s.base, s, s.length + length, l}
}

func (s *Str) get_at(k int64) byte {
	if k <= int64(len(s.base)) {
		return s.base[int(k-1)]
	}
	if k <= s.delegate.Len() {
		return s.delegate.get_at(k)
	}
	k -= s.delegate.Len()
	k += s.left
	return s.delegate.get_at(k)
}

func (s *Str) Len() int64 {
	if s == nil {
		return 0
	}
	return s.length
}

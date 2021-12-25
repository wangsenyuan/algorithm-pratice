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
		n := readNum(reader)

		ss, _ := reader.ReadString('\n')

		S := make([]string, n)

		var prev int
		var pos int
		for i := 0; i < n; i++ {
			for pos < len(ss) {
				if ss[pos] == ' ' || ss[pos] == '\n' {
					S[i] = ss[prev:pos]
					pos++
					prev = pos
					break
				}
				pos++
			}
		}
		res := solve(n, S)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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
func solve(n int, S []string) int {
	suffixes := make([]Pair, n)
	for i := 0; i < n; i++ {
		suffixes[i] = Pair{S[i][1:], i}
	}

	sort.Sort(Pairs(suffixes))
	compress := make([]int, n)
	var p int
	for i := 1; i <= n; i++ {
		compress[suffixes[i-1].second] = p
		if i == n || suffixes[i].first > suffixes[i-1].first {
			p++
		}
	}

	word := make([][]bool, 26)
	for i := 0; i < 26; i++ {
		word[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		x := int(S[i][0] - 'a')
		word[x][compress[i]] = true
	}

	var res int

	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			var a, b int
			for k := 0; k < n; k++ {
				if word[i][k] && !word[j][k] {
					a++
				} else if word[j][k] && !word[i][k] {
					b++
				}
			}
			res += a * b
		}
	}
	return res
}

type Pair struct {
	first  string
	second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first < this[j].first
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

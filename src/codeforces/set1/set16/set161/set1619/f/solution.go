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
		n, m, k := readThreeNums(reader)
		res := solve(n, m, k)
		for _, game := range res {
			for _, table := range game {
				buf.WriteString(fmt.Sprintf("%d", len(table)))
				for i := 0; i < len(table); i++ {
					buf.WriteString(fmt.Sprintf(" %d", table[i]))
				}
				buf.WriteByte('\n')
			}
		}
		if tc > 0 {
			buf.WriteByte('\n')
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

func solve(n int, m int, k int) [][][]int {
	// if n % m == 0
	// 随便分配即可， 因为b[i] = 0
	// else
	// x = n / m, r = n % m, 就有r个table，那么每轮game中，会有 r * (x + 1)个人 + 1
	// 所有，sum(b) = k * r * (x + 1)
	// avg = sum / n
	if n%m == 0 {
		return solve1(n, m, k)
	}
	ps := make(Persons, n)
	for i := 0; i < n; i++ {
		ps[i] = &Person{i + 1, 0}
	}
	// 每个人[1..n]需要参加avg次大桌会议
	res := make([][][]int, k)
	x, r := n/m, n%m

	for i := 0; i < k; i++ {
		// first r table
		res[i] = make([][]int, m)
		var p int
		sort.Sort(ps)
		for j := 0; j < r; j++ {
			res[i][j] = make([]int, x+1)
			for a := 0; a <= x; a++ {
				res[i][j][a] = ps[p].id
				ps[p].score++
				p++
			}
		}
		for j := r; j < m; j++ {
			res[i][j] = make([]int, x)
			for a := 0; a < x; a++ {
				res[i][j][a] = ps[p].id
				p++
			}
		}
	}

	return res
}

type Person struct {
	id    int
	score int
}

type Persons []*Person

func (pq Persons) Len() int {
	return len(pq)
}

func (pq Persons) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}

func (pq Persons) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func solve1(n int, m int, k int) [][][]int {
	res := make([][][]int, k)
	h := n / m
	for i := 0; i < k; i++ {
		res[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			// table[j]
			res[i][j] = make([]int, h)
			for a := 0; a < h; a++ {
				res[i][j][a] = j*h + a + 1
			}
		}
	}

	return res
}

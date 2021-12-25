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
		line := readNNums(reader, 5)
		res := solve(line[0], line[1], line[2], line[3], line[4])

		for i := 0; i < line[0]; i++ {
			buf.WriteString(res[i])
			buf.WriteByte('\n')
		}
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

func solve(n, m, r, g, b int) []string {
	tiles := make([]Pair, 3)
	tiles[0] = Pair{r, 'R'}
	tiles[1] = Pair{g, 'G'}
	tiles[2] = Pair{b, 'B'}

	sort.Sort(Pairs(tiles))

	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		res[i] = make([]byte, m)
	}

	fill := func(i, j int, x int) {
		res[i][j] = byte(x)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			d := (i + j) & 1
			if d == 0 && tiles[0].first > 0 {
				fill(i, j, tiles[0].second)
				(&tiles[0]).first--
			} else if d == 1 && tiles[1].first > 0 {
				fill(i, j, tiles[1].second)
				(&tiles[1]).first--
			} else {
				fill(i, j, tiles[2].second)
			}
		}
	}

	ans := make([]string, n)
	for i := 0; i < n; i++ {
		ans[i] = string(res[i])
	}

	return ans
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

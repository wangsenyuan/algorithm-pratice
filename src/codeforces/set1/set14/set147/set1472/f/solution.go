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
		readString(reader)
		n, m := readTwoNums(reader)
		blocks := make([][]int, m)
		for i := 0; i < m; i++ {
			blocks[i] = readNNums(reader, 2)
		}
		res := solve(n, blocks)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

func solve(n int, blocks [][]int) bool {
	cols := make(map[int]int)
	for _, cur := range blocks {
		x, y := cur[0], cur[1]
		x--
		y--
		cols[y] |= (1 << x)
	}

	type Pair struct {
		first  int
		second int
	}

	arr := make([]Pair, 0, len(cols)+1)

	for k, v := range cols {
		arr = append(arr, Pair{k, v})
	}

	arr = append(arr, Pair{n + 1, 3})

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].first < arr[j].first
	})

	has_last := false
	last_color := 0

	for _, cur := range arr {
		if cur.second != 3 && has_last {
			color := (cur.first + cur.second) % 2
			if last_color == color {
				return false
			}
			has_last = false
		} else if cur.second == 3 && has_last {
			return false
		} else if cur.second != 3 {
			last_color = (cur.first + cur.second) % 2
			has_last = true
		}
	}
	return true
}

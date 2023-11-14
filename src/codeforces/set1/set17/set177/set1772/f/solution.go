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

	n, m, k := readThreeNums(reader)

	pics := make([][]string, k+1)

	for i := 0; i <= k; i++ {
		readString(reader)
		pics[i] = make([]string, n)
		for j := 0; j < n; j++ {
			pics[i][j] = readString(reader)[:m]
		}
	}
	first, res := solve(pics)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n%d\n", first, len(res)))
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i][j]))
		}
		buf.WriteByte('\n')
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

func solve(pics [][]string) (int, [][]int) {
	// 相邻的两个的xor值，是相反的
	// 且相邻的两个，除了一位是不一样的，其他位都是一样的，且不一样的一位的四周是一样的
	// 然后组成一个图
	// 组成一个图 dp[state][i] => dp[next][j]
	n := len(pics)

	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{countAllowedOperations(pics[i]), i}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].first > ps[j].first
	})

	h := len(pics[0])
	w := len(pics[0][0])

	var ops [][]int

	for i := 0; i+1 < n; i++ {
		a := pics[ps[i].second]
		b := pics[ps[i+1].second]
		for r := 1; r+1 < h; r++ {
			for c := 1; c+1 < w; c++ {
				if a[r][c] != b[r][c] {
					ops = append(ops, []int{1, r + 1, c + 1})
				}
			}
		}
		ops = append(ops, []int{2, ps[i+1].second + 1})
	}

	return ps[0].second + 1, ops
}

type Pair struct {
	first  int
	second int
}

func countAllowedOperations(pic []string) int {
	h := len(pic)
	w := len(pic[0])

	var res int

	for i := 1; i+1 < h; i++ {
		for j := 1; j+1 < w; j++ {
			x := pic[i][j]
			if x != pic[i][j-1] && x != pic[i][j+1] && x != pic[i-1][j] && x != pic[i+1][j] {
				res++
			}
		}
	}

	return res
}

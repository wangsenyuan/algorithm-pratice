package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = readNNums(reader, m)
	}
	res := solve(a)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}

	fmt.Println(buf.String())
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

const inf = 1 << 60

var dd = []int{-1, 0, 1, 0, -1}

type pair struct {
	first  int
	second int
}

func min_of(a, b pair) pair {
	if a.first < b.first || a.first == b.first && a.second < b.second {
		return a
	}
	return b
}

func max_of(a, b pair) pair {
	if a.first > b.first || a.first == b.first && a.second > b.second {
		return a
	}
	return b
}

func solve(a [][]int) []int {
	n := len(a)
	m := len(a[0])

	// 好格子：自己是 1，或者存在一个小于自己的邻居
	ok := func(i, j int) bool {
		return a[i][j] == 1 ||
			j > 0 && a[i][j-1] < a[i][j] ||
			j+1 < m && a[i][j+1] < a[i][j] ||
			i > 0 && a[i-1][j] < a[i][j] ||
			i+1 < n && a[i+1][j] < a[i][j]
	}
	// 判断 (i,j)，以及 (i,j) 的邻居，是否都是好格子
	ok2 := func(i, j int) bool {
		return ok(i, j) &&
			(j == 0 || ok(i, j-1)) &&
			(j+1 == m || ok(i, j+1)) &&
			(i == 0 || ok(i-1, j)) &&
			(i+1 == n || ok(i+1, j))
	}

	// 收集坏格子
	type pair struct{ i, j int }
	badPos := []pair{}
	for i := range n {
		for j := range m {
			if !ok(i, j) {
				badPos = append(badPos, pair{i, j})
			}
		}
	}
	if len(badPos) == 0 {
		return []int{0}
	}

	ans := map[pair]struct{}{}
	// 除了交换 (bi,bj)，也可以通过交换 (bi,bj) 的邻居，使 (bi,bj) 变成一个好格子
	// 只需检查至多 5 个位置，因为 (bi,bj) 必须变成好格子
	bi, bj := badPos[0].i, badPos[0].j
	for _, p := range []pair{{bi, bj}, {bi, bj - 1}, {bi, bj + 1}, {bi - 1, bj}, {bi + 1, bj}} {
		if p.i < 0 || p.i == n || p.j < 0 || p.j == m {
			continue
		}
		for i := range n {
			for j := range m {
				// 交换其他所有点
				a[p.i][p.j], a[i][j] = a[i][j], a[p.i][p.j]
				// 交换离坏格子很远的点，必然是无效交换，所以先检查是否有坏格子仍然是坏格子
				for _, q := range badPos {
					if !ok(q.i, q.j) {
						goto next
					}
				}
				// 有效交换！进一步检查受到影响的 10 个点是否正常
				if ok2(p.i, p.j) && ok2(i, j) {
					// 注意去重
					ans[pair{min(p.i*m+p.j, i*m+j), max(p.i*m+p.j, i*m+j)}] = struct{}{}
				}
			next:
				a[p.i][p.j], a[i][j] = a[i][j], a[p.i][p.j]
			}
		}
	}

	if len(ans) > 0 {
		return []int{1, len(ans)}
	}
	return []int{2}
}

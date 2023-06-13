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

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m, k := readThreeNums(reader)
		participants := make([][]int, m)
		for i := 0; i < m; i++ {
			participants[i] = readNNums(reader, 2)
		}
		res := solve(n, k, participants)
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

func solve(n int, k int, participants [][]int) int {
	// 有n个题目，两个老师，每个老师可以选择连续的k个题目进行讲解
	// m个 participant, 每个人只对[li, ri]中间的题目感兴趣，
	// 如果老师讲解的题目是 [a..b] 则对于participant他听的讲解的个数 = max(0, min(ri, b) - max(a, li) + 1)
	// b - a + 1 == k
	// 在给定a和b的位置后，能否快速的计算出最大值？
	// 好像是可以的。
	// 但有一个条件是，两个老师课程不重叠，这个是否成立？
	type Player struct {
		l int
		r int
	}
	P := make([]Player, len(participants))
	for i := 0; i < len(participants); i++ {
		P[i] = Player{participants[i][0] - 1, participants[i][1]}
	}

	sort.Slice(P, func(i, j int) bool {
		return P[i].l+P[i].r < P[j].l+P[j].r
	})
	m := len(participants)
	suf := make([]int, m+1)

	for i := 0; i < n-k+1; i++ {
		var cur int
		for j := m - 1; j >= 0; j-- {
			cur += max(0, min(i+k, P[j].r)-max(i, P[j].l))
			suf[j] = max(suf[j], cur)
		}
	}
	ans := suf[0]

	for i := 0; i < n-k+1; i++ {
		var cur int
		for j := 0; j < m; j++ {
			cur += max(0, min(i+k, P[j].r)-max(i, P[j].l))
			ans = max(ans, cur+suf[j+1])
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}

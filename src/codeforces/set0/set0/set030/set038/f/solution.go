package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	winner, res := process(reader)
	fmt.Println(winner)
	fmt.Println(res[0], res[1])
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

func process(reader *bufio.Reader) (winner string, score []int) {
	n := readNum(reader)
	s := make([]string, n)
	for i := range n {
		s[i] = readString(reader)
	}
	winner, score = solve(s)
	return
}

type line struct {
	x   int
	sum int
	num int
	l   int
}

func solve(S []string) (winner string, score []int) {
	n := len(S)

	slices.SortFunc(S, func(a, b string) int {
		return len(b) - len(a)
	})

	var g [][]int
	var arr []line

	convert := func(s string) line {
		var x int
		var sum int
		for i := 0; i < len(s); i++ {
			y := int(s[i]-'a') + 1
			x = max(x, y)
			sum += y
		}
		var cnt int
		for i := 0; i < n; i++ {
			if strings.Contains(S[i], s) {
				cnt++
			}
		}

		return line{x, sum, cnt, len(s)}
	}

	pq := make(PQ, 0, n)

	id := make(map[string]int)
	add := func(x string) int {
		if v, ok := id[x]; ok {
			return v
		}
		m := len(arr)
		id[x] = m
		arr = append(arr, convert(x))
		g = append(g, []int{})

		heap.Push(&pq, x)
		return m
	}
	var pos int
	for pos < n && len(S[pos]) == len(S[0]) {
		add(S[pos])
		pos++
	}

	for pq.Len() > 0 {
		s := heap.Pop(&pq).(string)

		if len(s) > 1 {
			x := s[1:]
			i := add(x)
			g[i] = append(g[i], id[s])
			m := len(s)
			x = s[:m-1]
			i = add(x)
			g[i] = append(g[i], id[s])
		}
		for pos < n && len(S[pos]) == len(s)-1 {
			if _, ok := id[S[pos]]; !ok {
				add(S[pos])
			}
			pos++
		}
	}

	n = len(id)

	type state struct {
		win    bool
		first  int
		second int
	}

	comp := func(a state, b state) bool {
		// a是否优于状态b
		if a.win != b.win {
			return a.win
		}
		// a.win = b.win
		return a.first > b.first || a.first == b.first && a.second < b.second
	}

	// dp[i] = 在节点i时，获胜时的最大得分
	// 如果失败呢？
	// 仍然是最大的得分，且最小化对方的得分
	// dp[i] = 如果 dp[left[i]] 是一个失败位，可以选择
	dp := make([]state, n)
	for i := range n {
		cur := arr[i]
		s1 := cur.x*cur.sum + cur.num
		if len(g[i]) == 0 {
			dp[i] = state{true, s1, 0}
			continue
		}
		tmp := state{false, -1, 1 << 30}
		for _, j := range g[i] {
			if comp(dp[j], tmp) {
				tmp = dp[j]
			}
		}
		dp[i] = state{!tmp.win, tmp.second + s1, tmp.first}
	}
	var res state
	for i := n - 1; i >= 0; i-- {
		if arr[i].l == 1 && dp[i].win {
			// 从单个字符开始
			if comp(dp[i], res) {
				res = dp[i]
			}
		}
	}
	if res.win {
		return "First", []int{res.first, res.second}
	}

	for i := n - 1; i >= 0; i-- {
		if arr[i].l == 1 {
			// dp[i].first < dp[i].second
			if comp(dp[i], res) {
				res = dp[i]
			}
		}
	}

	return "Second", []int{res.first, res.second}
}

type PQ []string

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return len(pq[i]) > len(pq[j])
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x any) {
	s := x.(string)
	*pq = append(*pq, s)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	s := old[n-1]
	*pq = old[:n-1]
	return s
}

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
		X := readNNums(reader, n)
		Y := readNNums(reader, n)
		res := solve(X, Y)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(X []int, Y []int) int {
	// X[i] < X[i+]
	// X[i] < Y[i]
	// X & Y distinct
	// 1 2 3 4
	// 6 7 5 8
	// 先把food3取走，然后person3先到达， 取走food1
	// 然后person1到达 取走food2，
	// 然后person2到达， 取走food4
	// 看起来要按照Y排序
	// 如果person i 取走了 food j, (person j 后到达）
	// 那么 dp[i] = dp[j]+1
	n := len(X)

	type Event struct {
		id   int
		time int
		tp   int
	}

	events := make([]Event, 2*n)
	for i := 0; i < n; i++ {
		events[2*i] = Event{i, X[i], 0}
		events[2*i+1] = Event{i, Y[i], 1}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].time < events[j].time
	})

	next := make([]int, n)
	for i := 0; i < n; i++ {
		next[i] = -1
	}
	n2 := n*2 + 1
	// 假设person id[i] 的 food被取走，那么此时，person id[i] 将去获取现有的最小index的food
	// 该food属于person id[j]
	// dp[i] = dp[j] + 1
	// person id[i] 到达时，所有的food
	arr := make([]int, 2*n2)

	for i := 0; i < 2*n2; i++ {
		arr[i] = n
	}
	set := func(p int, v int) {
		p += n2
		arr[p] = v
		for p > 1 {
			arr[p>>1] = min(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l int, r int) int {
		res := n
		l += n2
		r += n2
		for l < r {
			if l&1 == 1 {
				res = min(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = min(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	for _, evt := range events {
		i := evt.id
		if evt.tp == 0 {
			// food arrival
			set(X[i], i)
		} else {
			// 如果person i 的食物被取走了, 他/她会取走剩余食物中index最小的那份
			set(X[i], n)
			next[i] = get(0, Y[i])
		}
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}

	var dfs func(i int) int

	dfs = func(i int) int {
		if dp[i] < 0 {
			if next[i] < n {
				dp[i] = 1 + dfs(next[i])
			} else {
				dp[i] = 0
			}
		}
		return dp[i]
	}

	var ans int

	for i := 0; i < n; i++ {
		ans = max(ans, dfs(i))
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	return a + b - min(a, b)
}

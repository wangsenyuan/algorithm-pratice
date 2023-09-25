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
	n := readNum(reader)
	rows := make([][]int, n)
	for i := 0; i < n; i++ {
		rows[i] = readNNums(reader, 2)
	}
	m := readNum(reader)
	Q := make([]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNum(reader)
	}

	res := solve(rows, Q)
	var buf bytes.Buffer
	for i := 0; i < m; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

const H = 30

func solve(rows [][]int, W []int) []int {
	// 对于w
	// 移动n行中的某些行，假设每行移动的距离为k[i]
	//  k[0] | k[1] | ... | k[n-1] = w
	// 目的是使得，移动后在某一列能够获得最大值
	// 能否每次都得到n？
	// 答案是否定的, 假设两行，它们距离很远
	// 它们没有办法移动足够远的距离，从而使的它们有重叠的部分
	// 这里有两个问题
	//  1. 如何保证移动的距离 or = w
	//  2. 使得它们有最多的一列
	//  3. 向左移动时，还需要保证不超出边界
	//  对于第2点，看起来时要往中间靠拢？
	// 其中提到，可以对同一行操作多次
	// 那么这个是不是相当于随便移动了？
	// 这个是因为，假设有个目的地，保证了2、3，此时如果 or < w,那么随便选一行，向右移动，再移动回来，就可以保证or = w
	// 这里额外的要求是 or & w = w
	//  所以，这里有个就是对于任何一行，它移动的k，= w的子码
	// 如果行移动了k距离，那么就是在 [l-k, r - k] 范围内有值
	// 即使只有一个w，也很难计算呐
	// 还有一个，就是加入只移动1个距离，那么对w的贡献是1，但是，其实可以一直按照1移动
	// 那这么看，那不就可以都可以移动到1号位了？也就是按照w的最低位进行移动
	// 似乎看到希望了
	// 对于给定的d，如何计算它的最大值呢？

	n := len(rows)
	events := make([]Pair, 3*n)
	// 现在是 n * lg * lg
	// 能不能再消掉个lg？
	var it int
	addEvent := func(a, b int) {
		events[it] = Pair{a, b}
		it++
	}

	check := func(d int) int {
		if d == 0 {
			return n
		}
		// 只需要所有的一直往east移动就可以了，
		// 然后在某个地方，寻找重叠区域
		// 当长度超过dist时，它可以cover住[0..dist]个区域
		// 如果 l % dist < r % dist 只能cover [l...r)区域
		// 如果 l % dist >= r % dist, 只能cover [0..r) 和 [l...)区域
		dist := 1 << d
		it = 0
		for i := 0; i < n; i++ {
			l, r := rows[i][0], rows[i][1]
			l--
			if r-l >= dist {
				addEvent(0, 1)
			} else {
				l %= dist
				r %= dist
				if l < r {
					addEvent(l, 1)
					addEvent(r, -1)
				} else {
					addEvent(0, 1)
					addEvent(r, -1)
					addEvent(l, 1)
				}
			}
		}

		sort.Slice(events[:it], func(i, j int) bool {
			return events[i].first < events[j].first ||
				events[i].first == events[j].first && events[i].second < events[j].second
		})

		var best int
		var cnt int

		for i := 0; i < it; i++ {
			cnt += events[i].second
			best = max(best, cnt)
		}

		return best
	}

	res := make([]int, H)
	// 高位的结果，貌似也可以应用到低位，考虑移动4位的情况下，答案是x
	// 那么移动2位，其实相当于移动了2次2位后，得到高位的结果
	for i := H - 1; i >= 0; i-- {
		res[i] = check(i)
	}

	ans := make([]int, len(W))

	for i, w := range W {
		var j int
		for w&1 == 0 {
			j++
			w >>= 1
		}
		ans[i] = res[j]
	}
	return ans
}

type Pair struct {
	first  int
	second int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

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
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(n, a, m)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, a []int, coins int) int {
	//想象最后的结果中
	// 大体可以分为两类
	// 有x个是从0出发到达的
	// 剩下的y个是从n+1出发到达的
	// 假设到达位置，要传送至哪里呢？取决于下一个更近的是哪一个
	// 但是第一个必须是从0出发的
	ps := make([]Pair, n)
	for i := 1; i <= n; i++ {
		ps[i-1] = Pair{min(i, n+1-i) + a[i-1], i + a[i-1]}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].first < ps[j].first
	})

	sum := make([]int, n)

	for i := 0; i < n; i++ {
		sum[i] = ps[i].first
		if i > 0 {
			sum[i] += sum[i-1]
		}
	}
	var res int
	for i := 0; i < n; i++ {
		tmp := coins - ps[i].second
		if tmp < 0 {
			continue
		}

		j := search(n, func(j int) bool {
			price := sum[j]
			if j > i {
				price -= ps[i].first
			}
			return price > tmp
		})
		if j > i {
			j--
		}
		res = max(res, j+1)
	}

	return res
}

func solve1(n int, a []int, coins int) int {
	//想象最后的结果中
	// 大体可以分为两类
	// 有x个是从0出发到达的
	// 剩下的y个是从n+1出发到达的
	// 假设到达位置，要传送至哪里呢？取决于下一个更近的是哪一个
	// 但是第一个必须是从0出发的
	ps := make([]Pair, n)
	for i := 1; i <= n; i++ {
		ps[i-1] = Pair{min(i, n+1-i) + a[i-1], i}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].first < ps[j].first
	})
	// 第一个位置有关系
	val := make([]Pair, n+1)
	pos := make([]int, n+1)
	for i := 0; i < n; i++ {
		pos[ps[i].second] = i
		val[ps[i].second] = ps[i]
	}

	sum := make([]int, n+2)

	set := func(i int, v int) {
		i++
		for i < len(sum) {
			sum[i] += v
			i += i & -i
		}
	}

	get := func(i int) int {
		i++
		var res int
		for i > 0 {
			res += sum[i]
			i -= i & -i
		}
		return res
	}

	for i := 0; i < n; i++ {
		set(i, ps[i].first)
	}

	var res int

	for i := 1; i <= n; i++ {
		if coins < i+a[i-1] {
			continue
		}
		// 如果先到达i
		tmp := coins - (i + a[i-1])
		set(pos[i], -val[i].first)

		j := search(n, func(j int) bool {
			return get(j) > tmp
		})

		if j <= pos[i] {
			res = max(res, j+1)
		} else {
			res = max(res, j)
		}

		set(pos[i], val[i].first)
	}

	return res
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}

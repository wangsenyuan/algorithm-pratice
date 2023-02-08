package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	E := make([][]int, m)

	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}

	P := make([]Pair, len(E))
	for i := 0; i < len(E); i++ {
		P[i] = Pair{int64(E[i][2]), i}
	}
	sort.Slice(P, func(i, j int) bool {
		return P[i].first < P[j].first
	})

	set := NewSet(n)

	ask2 := func(s string) int64 {
		return ask(s, E, P, set)
	}

	res := solve(n, m, ask2)

	output(res)
}

func output(res int64) {
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()

	fmt.Fprintf(f, "! %d\n", res)
}

type set struct {
	arr  []int
	size []int
}

func NewSet(n int) *set {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i]++
	}
	return &set{arr, cnt}
}

func (s *set) Find(x int) int {
	if s.arr[x] != x {
		s.arr[x] = s.Find(s.arr[x])
	}
	return s.arr[x]
}

func (s *set) Union(a, b int) bool {
	a = s.Find(a)
	b = s.Find(b)
	if a == b {
		return false
	}
	if s.size[a] < s.size[b] {
		a, b = b, a
	}
	s.size[a] += s.size[b]
	s.arr[b] = a
	return true
}

func (s *set) Reset() {
	for i := 0; i < len(s.arr); i++ {
		s.arr[i] = i
		s.size[i] = 1
	}
}

func ask(s string, E [][]int, P []Pair, set *set) int64 {
	set.Reset()

	var res int64
	for i := 0; i < len(P); i++ {
		j := P[i].second
		if s[j] == '1' {
			u, v := E[j][0], E[j][1]
			u--
			v--
			if set.Union(u, v) {
				res += P[i].first
			}
		}
	}
	return res
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
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
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

func solve(n int, m int, ask func(string) int64) int64 {
	// ask 返回的是最大，工作中的，full spanning tree 的value
	// 考虑全部打开设置为on，得到的结果结果是最大值
	// 关闭其中一个，如果结果没有变化，至少说明它不是最大值中的一条边，但也不一定是需要的那条边
	//    如果有变化，那么它就是属于其中，且其他的边(在同一个连通图里面) < 它， 它的值等于变化的部分
	// 但有这个信息，要怎么继续处理？还有一个个的开，似乎太慢了。
	// n个点，在全部链接在一起的情况下，需要n-1条边，
	//   如果分成x个森林的话，需要n-x条边
	// 可以先计算出x吗？
	// 可以通过m次查询知道每条边的长度，然后按照从低到高添加边进来

	buf := make([]byte, m)
	for i := 0; i < m; i++ {
		buf[i] = '0'
	}

	query := func(i int) int64 {
		buf[i] = '1'
		res := ask(string(buf))
		buf[i] = '0'
		return res
	}

	P := make([]Pair, m)
	for i := 0; i < m; i++ {
		l := query(i)
		P[i] = Pair{l, i}
	}

	sort.Slice(P, func(i, j int) bool {
		return P[i].first < P[j].first
	})

	var res int64

	buf[P[0].second] = '1'
	res += P[0].first

	for i := 1; i < m; i++ {
		tmp := query(i)
		if tmp == res {
			continue
		}
		res = tmp
		buf[i] = '1'
	}
	return res
}

type Pair struct {
	first  int64
	second int
}

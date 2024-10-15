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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		wires := make([][]int, n)
		for i := 0; i < n; i++ {
			wires[i] = readNNums(reader, 2)
		}
		res := solve(wires)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, cur := range res {
			buf.WriteString(fmt.Sprintf("%d %d %d\n", cur[0], cur[1], cur[2]))
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

func compact(wires [][]int) (map[int]int, []int) {
	pos := make(map[int]int)
	for _, wire := range wires {
		u, v := wire[0], wire[1]
		pos[u]++
		pos[v]++
	}
	var arr []int
	for k := range pos {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	n := len(arr)
	for i := 0; i < n; i++ {
		pos[arr[i]] = i
	}

	return pos, arr
}

func solve(wires [][]int) [][]int {
	pos, arr := compact(wires)

	n := len(arr)

	set := NewDSU(n)

	var extra []int

	for i, wire := range wires {
		u, v := wire[0], wire[1]
		u = pos[u]
		v = pos[v]
		if !set.Union(u, v) {
			extra = append(extra, i)
		}
	}

	// comps是每个集合的根节点，也是用来连接的节点
	var comps []int
	for i := 0; i < n; i++ {
		j := set.Find(i)
		if i == j {
			comps = append(comps, i)
		}
	}

	by_comps := make([][]int, len(comps))

	for _, i := range extra {
		u := pos[wires[i][0]]
		u = set.Find(u)
		u = sort.SearchInts(comps, u)
		by_comps[u] = append(by_comps[u], i)
	}

	ids := make([]int, len(comps))

	for i := 0; i < len(comps); i++ {
		ids[i] = i
	}

	sort.Slice(ids, func(i int, j int) bool {
		return len(by_comps[ids[i]]) > len(by_comps[ids[j]])
	})

	use := make([]int, len(wires))

	for i := 0; i < len(wires); i++ {
		use[i] = -1
	}

	// 和第一个去连接
	var ans [][]int
	extra = by_comps[ids[0]]
	for i := 1; i < len(ids) && len(extra) > 0; i++ {
		e_id := extra[len(extra)-1]
		u := wires[e_id][0]

		c_id := ids[i]
		b := comps[c_id]

		ans = append(ans, []int{e_id + 1, u, arr[b]})
		use[e_id] = len(ans) - 1

		extra = extra[:len(extra)-1]
		extra = append(extra, by_comps[c_id]...)
	}
	if len(extra) > 0 {
		// 肯定都连起来了，否则环不会有剩余
		return ans
	}
	// 现在已经没有环了
	set.Reset()

	deg := make([]int, n)
	line := make([]int, n)

	for i, wire := range wires {
		u, v := wire[0], wire[1]
		if use[i] < 0 {
			u = pos[u]
		} else {
			b := ans[use[i]][2]
			u = pos[b]
		}
		v = pos[v]
		set.Union(u, v)
		deg[u]++
		deg[v]++
		line[u] = i
		line[v] = i
	}
	// 然后用 deg = 1 的去连接就可以了
	// 不可能有 deg = 0的部分
	comps = comps[:0]
	for i := 0; i < n; i++ {
		if set.Find(i) == i {
			comps = append(comps, i)
		}
	}
	// 这里使用的是节点编号
	ends := make([]int, len(comps))

	for u := 0; u < n; u++ {
		if deg[u] == 1 {
			v := set.Find(u)
			if u == v {
				// 这种情况存在那种一条线, 连接两个节点的情况
				continue
			}
			v = sort.SearchInts(comps, v)
			// 只需要记录一个就可以了
			ends[v] = u
		}
	}

	// 依次连接就可以了, 不需要排序
	for i := 1; i < len(comps); i++ {
		u := ends[i-1]
		e_id := line[u]
		// now change back
		u = arr[u]
		b := arr[comps[i]]
		ans = append(ans, []int{e_id + 1, u, b})
	}

	return ans
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}

func (this *DSU) Reset() {
	for i := 0; i < len(this.arr); i++ {
		this.arr[i] = i
		this.cnt[i] = 1
	}
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d\n", len(cur)))
		for _, x := range cur {
			buf.WriteString(x)
			buf.WriteByte('\n')
		}
	}
	buf.WriteTo(os.Stdout)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func process(reader *bufio.Reader) [][]string {
	nums := readNNums(reader, 4)
	k := nums[0]
	n := nums[1]
	m := nums[2]
	q := nums[3]
	basics := make([]string, n)
	for i := range n {
		basics[i] = readString(reader)
	}
	comps := make([]string, m)
	for i := range m {
		comps[i] = readString(reader)
	}
	qs := make([]string, q)
	for i := range q {
		qs[i] = readString(reader)
	}
	return solve(k, basics, comps, qs)
}

type pair struct {
	first  int
	second int
}

func solve(k int, basics []string, composites []string, qs []string) [][]string {
	// 先要组成树
	// bfs
	id := make(map[string]int)
	for i, cur := range basics {
		id[cur] = i
	}
	n := len(basics)
	m := len(composites)
	g := make([][]pair, n)

	compNames := make([]string, m)
	expect := make([]int, m)
	need := make([][]int, m)

	for i := range m {
		need[i] = make([]int, n)
		cur := composites[i]
		j := strings.IndexByte(cur, ':')
		compNames[i] = cur[:j]
		cur = cur[j+2:]
		for len(cur) > 0 {
			j = strings.IndexByte(cur, ' ')
			arti := cur[:j]
			j++
			var cnt int
			for j < len(cur) && cur[j] != ',' {
				x := int(cur[j] - '0')
				cnt = cnt*10 + x
				j++
			}
			need[i][id[arti]] = cnt
			expect[i] |= 1 << id[arti]
			g[id[arti]] = append(g[id[arti]], pair{i, cnt})
			if j == len(cur) {
				break
			}
			cur = cur[j+2:]
		}
	}

	cnt := make([]int, n)
	flag := make([]int, m)
	res := make([]int, m)
	play := func(x int) []string {
		clear(res)
		clear(cnt)
		clear(flag)

		for _, cur := range qs {
			var pos int
			var num int
			for pos < len(cur) && cur[pos] != ' ' {
				num = num*10 + int(cur[pos]-'0')
				pos++
			}
			if num != x {
				continue
			}
			arti := cur[pos+1:]
			j := id[arti]
			cnt[j]++
			// 不确定这是个数组？
			for _, cur := range g[j] {
				if cur.second == cnt[j] {
					i := cur.first
					flag[i] |= 1 << j
					if flag[i] == expect[i] {
						// 完成了一个
						res[i]++
						// 消耗掉了，需要把这些消耗掉的给减去
						for u := 0; u < n; u++ {
							cnt[u] -= need[i][u]
							if cnt[u] < need[i][u] {
								flag[i] ^= 1 << u
							}
						}
					}
					break
				}
			}
		}
		var arr []string

		for i := 0; i < n; i++ {
			if cnt[i] > 0 {
				arr = append(arr, fmt.Sprintf("%s %d", basics[i], cnt[i]))
			}
		}

		for i := 0; i < m; i++ {
			if res[i] > 0 {
				arr = append(arr, fmt.Sprintf("%s %d", compNames[i], res[i]))
			}
		}

		sort.Strings(arr)

		return arr
	}

	var ans [][]string

	for x := 1; x <= k; x++ {
		ans = append(ans, play(x))
	}

	return ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _ := process(reader)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) (res []int, roads [][]int) {
	n, m := readTwoNums(reader)
	roads = make([][]int, m)
	for i := range m {
		roads[i] = readNNums(reader, 4)
	}
	res = solve(n, roads)
	return
}

func solve(n int, roads [][]int) []int {
	m := len(roads)
	g := NewGraph(n+1, 2*m)
	for i, cur := range roads {
		u, v := cur[0], cur[1]
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}
	a := make([]int, n+1)

	// l * g 越小的， 似乎限制越大。
	// 假设节点u，它一圈的边的l * g 知道的情况下，那么a[u]， 至多是它们的gcd
	// 因为 a[u]要能整除它们。
	// 如果有一个节点不满足条件，那么就没有答案
	// 两外，相邻的u，v如果u变大，那么v就要变小
	type node struct {
		id int
		x  int
		y  int
	}

	nodes := make([]node, 0, n)

	for u := 1; u <= n; u++ {
		var x int
		y := 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			e := g.val[i]
			y = lcm(y, roads[e][2])
			x = gcd(x, roads[e][3])
		}
		if x == 0 {
			a[u] = 1
		} else if x%y != 0 {
			return nil
		} else {
			// a[u]要能被所有的gcd整除，且能够整除所有的gcd * lcm
			// 也就是说 a[u] % y = 0, x % a[u] = 0
			// 所以y是a[u]最小值， x是a[u]的最大值
			nodes = append(nodes, node{u, x, y})
		}
	}

	slices.SortFunc(nodes, func(a, b node) int {
		return b.y - a.y
	})

	var dfs func(u int) bool

	dfs = func(u int) bool {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			e := g.val[i]
			l, m := roads[e][2], roads[e][3]
			if a[v] == 0 {
				if l*m%a[u] != 0 {
					return false
				}
				tmp := l * m / a[u]
				if gcd(tmp, a[u]) != l {
					return false
				}
				a[v] = tmp
				if !dfs(v) {
					// revert
					a[v] = 0
					return false
				}
			} else {
				if a[u]*a[v] != l*m {
					return false
				}
			}
		}
		return true
	}

	for _, cur := range nodes {
		u := cur.id
		if a[u] != 0 {
			continue
		}
		for y := cur.y; y <= 1e6 && y <= cur.x; y += cur.y {
			if cur.x%y != 0 {
				continue
			}
			a[u] = y
			if dfs(u) {
				break
			}
			a[u] = 0
		}
		if a[u] == 0 {
			return nil
		}
	}

	return a[1:]
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	c := gcd(a, b)
	return a / c * b
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

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
	n, m := readTwoNums(reader)
	bridges := make([][]int, m)
	for i := 0; i < m; i++ {
		bridges[i] = readNNums(reader, 2)
	}
	res := solve(n, bridges)
	if len(res) == 0 {
		fmt.Println("NO")
	} else {
		var buf bytes.Buffer
		buf.WriteString("YES\n")
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		fmt.Println(buf.String())
	}
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func solve(n int, E [][]int) []int {
	k := n / 3
	// (a, b), (b, c), a 和 b要用不同的颜色，b 和 c要用不同的颜色
	g := createGraph(n, E)

	color := make([]int, n)

	var dfs func(u int)

	dfs = func(u int) {
		for v := range g[u] {
			if color[v] == 0 {
				color[v] = -color[u]
				dfs(v)
			}
		}
	}

	for u := 0; u < n; u++ {
		if color[u] == 0 {
			color[u] = 1
			dfs(u)
		}
	}

	var left []int
	var right []int
	for i := 0; i < n; i++ {
		if color[i] < 0 {
			left = append(left, i)
		} else {
			right = append(right, i)
		}
	}

	if len(left)%3 == 0 {
		// good enough
		return solve0(n, k, left, right)
	}
	if len(left)%3 == 2 {
		left, right = right, left
	}
	// len(left) % 3 == 1
	// find one vertex have at least two anti-edges
	for i, u := range left {
		if len(g[u]) <= len(right)-2 {
			return solve1(n, k, g, left, right, i)
		}
	}
	return solve2(n, k, g, left, right)
}

func createGraph(n int, E [][]int) []map[int]bool {
	g := make([]map[int]bool, n)

	for i := 0; i < n; i++ {
		g[i] = make(map[int]bool)
	}

	add := func(u, v int) {
		g[u][v] = true
	}

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		add(u, v)
		add(v, u)
	}
	return g
}

func solve0(n int, k int, left, right []int) []int {
	res := make([]int, n)

	for i := 0; i < len(left); i += 3 {
		res[left[i]] = k
		res[left[i+1]] = k
		res[left[i+2]] = k
		k--
	}

	for i := 0; i < len(right); i += 3 {
		res[right[i]] = k
		res[right[i+1]] = k
		res[right[i+2]] = k
		k--
	}
	return res
}

func solve1(n int, k int, g []map[int]bool, left, right []int, x int) []int {
	// left[x] have (>=) 2 anti-edges
	first, second := -1, -1
	u := left[x]
	for i, v := range right {
		if g[v][u] {
			continue
		}
		if first < 0 {
			first = i
		} else {
			second = i
		}
		if second >= 0 {
			break
		}
	}
	a := right[first]
	b := right[second]

	copy(left[x:], left[x+1:])
	copy(right[second:], right[second+1:])
	copy(right[first:], right[first+1:])

	res := solve0(n, k-1, left[:len(left)-1], right[:len(right)-2])
	res[u] = k
	res[a] = k
	res[b] = k
	return res
}

func solve2(n int, k int, g []map[int]bool, left []int, right []int) []int {
	// len(left) % 3 == 1, len(right) % 3 == 2
	// need to find two vertices have at least two ant-
	first, second := -1, -1
	for i, u := range right {
		if len(g[u])+2 <= len(left) {
			if first < 0 {
				first = i
			} else {
				second = i
			}
		}
		if second >= 0 {
			break
		}
	}
	if second < 0 {
		return nil
	}

	var cand []int

	for i, v := range left {
		if !g[v][right[first]] {
			cand = append(cand, i)
		}
		if len(cand) == 2 {
			break
		}
	}
	for i, v := range left {
		if !g[v][right[second]] {
			cand = append(cand, i)
		}
		if len(cand) == 4 {
			break
		}
	}

	kee := make([]int, 4)
	for i := 0; i < 4; i++ {
		kee[i] = left[cand[i]]
	}

	sort.Ints(cand)

	for i := len(cand) - 1; i >= 0; i-- {
		j := cand[i]
		copy(left[j:], left[j+1:])
	}
	a := right[first]
	b := right[second]
	copy(right[second:], right[second+1:])
	copy(right[first:], right[first+1:])
	res := solve0(n, k-2, left[:len(left)-4], right[:len(right)-2])
	// a 和 kee[0] 相同
	res[kee[0]] = k
	res[kee[1]] = k
	res[a] = k
	res[kee[2]] = k - 1
	res[kee[3]] = k - 1
	res[b] = k - 1
	return res
}

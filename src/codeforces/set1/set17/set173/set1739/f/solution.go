package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	C := make([]int, n)
	words := make([]string, n)

	for i := 0; i < n; i++ {
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &C[i])
		pos++
		words[i] = normalize(string(s[pos:]))
	}

	res := solve(C, words)
	fmt.Println(res)
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

const INF = 1 << 30
const N = 10043
const K = 12

func solve(C []int, words []string) string {
	// 对于word, 是否easy
	//  要判断word[i] 是否和 word[i-1]是相邻的
	// 首先可以排除一部分word，a和b要相邻,b和c相邻，c要和a相邻, 组成一个环，长度 > 2
	// 剩下的就是是链式的

	var verts []*Vertex

	var trie [N][K]int
	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			trie[i][j] = -1
		}
	}
	newNode := func() int {
		verts = append(verts, NewVertex())
		return len(verts) - 1
	}

	find := func(x, y int) int {
		if trie[x][y] == -1 {
			trie[x][y] = newNode()
			verts[trie[x][y]].pch = y
			verts[trie[x][y]].pa = x
		}
		return trie[x][y]
	}

	var jump func(x int, y int) int

	var getLink func(x int) int

	getLink = func(x int) int {
		if x >= len(verts) {
			return 0
		}

		if verts[x].link != -1 {
			return verts[x].link
		}
		if x == 0 || verts[x].pa == 0 {
			verts[x].link = 0
		} else {
			// 找到父节点的suf-link, follow pch
			verts[x].link = jump(getLink(verts[x].pa), verts[x].pch)
		}
		return verts[x].link
	}

	jump = func(x int, y int) int {
		if x >= len(verts) {
			return 0
		}
		if verts[x].jump[y] != -1 {
			return verts[x].jump[y]
		}
		if trie[x][y] != -1 {
			verts[x].jump[y] = trie[x][y]
			return verts[x].jump[y]
		}
		if x == 0 {
			verts[x].jump[y] = 0
			return verts[x].jump[y]
		}
		verts[x].jump[y] = jump(getLink(x), y)
		return verts[x].jump[y]
	}

	add := func(s string, cost int) {
		cur := 0
		for i := 0; i < len(s); i++ {
			cur = find(cur, int(s[i]-'a'))
		}
		verts[cur].cost += cost
	}

	var calc func(x int) int
	calc = func(x int) int {
		if verts[x].ncost < 0 {
			verts[x].ncost = verts[x].cost
			y := getLink(x)
			if y != x {
				verts[x].ncost += calc(y)
			}
		}
		return verts[x].ncost
	}
	// for root
	newNode()

	for i, word := range words {
		tmp := check(word)
		if len(tmp) == 0 {
			continue
		}
		add(tmp, C[i])
		add(reverse(tmp), C[i])
	}

	dp := make([][]int, 1<<K)
	pdp := make([][]Pair, 1<<K)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(verts)+1)
		pdp[i] = make([]Pair, len(verts)+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -INF
		}
	}

	dp[0][0] = 0

	for state := 0; state < 1<<K; state++ {
		for j := 0; j <= len(verts); j++ {
			for z := 0; z < K; z++ {
				if (state>>z)&1 == 1 {
					continue
				}
				ns := jump(j, z)
				ex := calc(ns)
				nm := state | (1 << z)
				if dp[nm][ns] < dp[state][j]+ex {
					dp[nm][ns] = dp[state][j] + ex
					pdp[nm][ns] = Pair{z, j}
				}
			}
		}
	}
	var ans []byte
	curMask := (1 << K) - 1
	var curState int
	for j := 0; j <= len(verts); j++ {
		if dp[curMask][j] > dp[curMask][curState] {
			curState = j
		}
	}
	for curMask != 0 {
		cc := pdp[curMask][curState].first
		dd := pdp[curMask][curState].second
		ans = append(ans, byte('a'+cc))
		curMask ^= 1 << cc
		curState = dd
	}
	return string(ans)
}

type Pair struct {
	first  int
	second int
}

type Vertex struct {
	cost  int
	ncost int
	pa    int
	pch   int
	link  int
	jump  [K]int
}

func NewVertex() *Vertex {
	v := new(Vertex)
	v.pa = -1
	for i := 0; i < K; i++ {
		v.jump[i] = -1
	}
	v.pch = -1
	v.link = -1
	v.cost = 0
	// not cal yet
	v.ncost = -1
	return v
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}

func check(word string) string {
	adj := make([]map[int]bool, K)
	for i := 0; i < K; i++ {
		adj[i] = make(map[int]bool)
	}
	for i := 1; i < len(word); i++ {
		x := int(word[i-1] - 'a')
		y := int(word[i] - 'a')
		adj[x][y] = true
		adj[y][x] = true
	}
	var cnt int
	var x int = -1
	for i := 0; i < K; i++ {
		if len(adj[i]) == 0 {
			continue
		}
		if len(adj[i]) > 2 {
			return ""
		}
		if len(adj[i]) == 1 {
			x = i
		}
		cnt++
	}
	if x < 0 {
		return ""
	}
	// len(adj[x]) = 1
	var arr []int
	arr = append(arr, x)
	for len(adj[x]) > 0 {
		for y := range adj[x] {
			delete(adj[x], y)
			delete(adj[y], x)
			x = y
		}
		arr = append(arr, x)
	}
	if cnt != len(arr) {
		return ""
	}
	pos := make([]int, K)
	for i := 0; i < len(arr); i++ {
		pos[arr[i]] = i
	}

	for i := 0; i+1 < len(word); i++ {
		a := int(word[i] - 'a')
		b := int(word[i+1] - 'a')
		if abs(pos[a]-pos[b]) > 1 {
			return ""
		}
	}
	buf := make([]byte, len(arr))

	for i := 0; i < len(arr); i++ {
		buf[i] = byte('a' + arr[i])
	}

	return string(buf)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

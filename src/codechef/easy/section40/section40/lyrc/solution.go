package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	m := readNum(reader)
	P := make([]string, m)
	for i := 0; i < m; i++ {
		P[i] = readString(reader)
	}
	n := readNum(reader)
	S := make([]string, n)
	for i := 0; i < n; i++ {
		S[i] = readString(reader)
	}
	res := solve(P, S)

	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

// const MAX_P = 5001
const MAX_S = 50010
const H = 63

func getIndex(x byte) int {
	if x >= 'a' && x <= 'z' {
		return int(x - 'a')
	}
	if x >= 'A' && x <= 'Z' {
		return 26 + int(x-'A')
	}
	return 62
}

func solve(P []string, S []string) []int {
	m := len(P)
	//n := len(S)
	// P[i] count in S
	// 如果用kmp， M * （5000 + 50000 * 100) <= 500 * 50000 = 25 * 10e8
	// 显然是不行的。
	// 如果用hash可行吗？
	// 100 * （5000 + 50000 * 100）
	nodes := m * MAX_S
	ac := make([][]int, nodes)
	for i := 0; i < nodes; i++ {
		ac[i] = make([]int, H)
		for j := 0; j < H; j++ {
			ac[i][j] = -1
		}
	}
	fail := make([]int, nodes)
	deg := make([]int, nodes)
	cnt := make([]int, nodes)
	tail := make([]int, m)

	var sz = 1
	nextVertex := func() int {
		//res := sz
		sz++
		return sz - 1
	}

	insert := func(s string) int {
		var cur int
		for i := 0; i < len(s); i++ {
			id := getIndex(s[i])
			if ac[cur][id] < 0 {
				ac[cur][id] = nextVertex()
			}
			cur = ac[cur][id]
		}
		return cur
	}

	for i := 0; i < m; i++ {
		tail[i] = insert(P[i])
	}
	//fail to root is root
	fail[0] = 0
	//fail to root is pointing towards root hence indeg is 1
	deg[0]++
	for i := 0; i < H; i++ {
		if ac[0][i] != -1 {
			//make all valid children of root fail to root.
			fail[ac[0][i]] = 0
			deg[0]++
		} else {
			ac[0][i] = 0
		}
	}
	// make transitions
	que := make([]int, sz+10)
	var front, end int
	for i := 0; i < H; i++ {
		if ac[0][i] > 0 {
			que[end] = ac[0][i]
			end++
		}
	}

	for front < end {
		u := que[front]
		front++
		for i := 0; i < H; i++ {
			//if the current node exists
			if ac[u][i] > 0 {
				//if string fails at the current position, it should
				//next point towards where the parents' failure
				//function was pointing
				fail[ac[u][i]] = ac[fail[u]][i]
				deg[ac[fail[u]][i]]++
				que[end] = ac[u][i]
				end++
			} else {
				//make the child point towards the fail of parent
				ac[u][i] = ac[fail[u]][i]
			}
		}
	}
	// query
	for _, text := range S {
		var pos int
		//refer to insert function to understand what goes on over here
		//we are iterating through the trie
		for j := 0; j < len(text); j++ {
			pos = ac[pos][getIndex(text[j])]
			//increasing the frequency of the current node
			cnt[pos]++
		}
	}

	front = 0
	end = 0
	for i := 0; i <= sz; i++ {
		if deg[i] == 0 {
			que[end] = i
			end++
		}
	}
	for front < end {
		tmp := que[front]
		front++
		//if we reach to root we break.
		if tmp == 0 {
			break
		}
		//we are virtually removing the current node,
		//so we remove its pointer to the failure function by
		//reducing degree to its failure node.
		deg[fail[tmp]]--
		//adding the frequency of the current node to its
		//failure node
		cnt[fail[tmp]] += cnt[tmp]

		if deg[fail[tmp]] == 0 {
			que[end] = fail[tmp]
			end++
		}
	}

	res := make([]int, m)

	for i := 0; i < m; i++ {
		res[i] = cnt[tail[i]]
	}

	return res
}

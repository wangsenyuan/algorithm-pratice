package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	S := make([]string, n)
	for i := 0; i < n; i++ {
		S[i] = readString(reader)
	}
	m := readNum(reader)
	Q := make([]string, m)
	for i := 0; i < m; i++ {
		Q[i] = readString(reader)
	}
	res := solve(S, Q)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

const N = 1e5

func solve(S []string, Q []string) []int64 {
	trie := make([][]int, 1)
	pv := make([][]int, 1)

	addNode := func(n int) []int {
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = -1
		}
		return arr
	}

	trie[0] = addNode(26)
	pv[0] = addNode(20)

	n := len(S)
	tail := make([]int, n)
	ln := make([]int, n)

	for i := 0; i < n; i++ {
		var pos int
		for j := 0; j < len(S[i]); j++ {
			x := int(S[i][j] - 'a')
			if trie[pos][x] == -1 {
				trie[pos][x] = len(trie)
				trie = append(trie, addNode(26))
				cur := addNode(20)
				pv = append(pv, cur)
				p := pos
				for k := 0; k < 20; k++ {
					cur[k] = p
					if p != -1 {
						p = pv[p][k]
					}
				}
			}
			pos = trie[pos][x]
		}
		tail[i] = pos
		ln[i] = len(S[i])
	}

	//ss := make([]string, n)
	//copy(ss, S)

	que := make([][]int, 0, len(Q))

	for i := 0; i < len(Q); i++ {
		if Q[i][0] == '1' {
			var x, y, z int
			s := []byte(Q[i])
			pos := readInt(s, 2, &x)
			x--
			pos = readInt(s, pos+1, &y)
			readInt(s, pos+1, &z)
			que = append(que, []int{1, x, y, z})
		} else if Q[i][0] == '2' {
			var x, y int
			s := []byte(Q[i])
			pos := readInt(s, 2, &x)
			x--
			pos = readInt(s, pos+1, &y)
			z := Q[i][pos+1:]
			//ss = append(ss, z)
			pos = tail[x]
			goup := ln[x] - y
			for j := 0; j < 20; j++ {
				if (goup>>j)&1 == 1 {
					pos = pv[pos][j]
				}
			}

			for j := 0; j < len(z); j++ {
				c := int(z[j] - 'a')
				if trie[pos][c] == -1 {
					trie[pos][c] = len(trie)
					trie = append(trie, addNode(26))
					cur := addNode(20)
					pv = append(pv, cur)
					p := pos
					for k := 0; k < 20; k++ {
						cur[k] = p
						if p != -1 {
							p = pv[p][k]
						}
					}
				}
				pos = trie[pos][c]
			}

			tail = append(tail, pos)
			ln = append(ln, y+len(z))
			que = append(que, []int{2, n})
			n++
		} else {
			var x int
			s := []byte(Q[i])
			readInt(s, 2, &x)
			x--
			que = append(que, []int{3, x})
		}
	}

	//n = len(tail)
	in := make([]int, N)
	out := make([]int, N)

	var dfs func(v int)
	var time int

	dfs = func(v int) {
		in[v] = time
		time++
		for i := 0; i < 26; i++ {
			if trie[v][i] != -1 {
				dfs(trie[v][i])
			}
		}
		out[v] = time
	}

	dfs(0)

	tr := NewFenwick(N + 1)
	t := make([]int64, N)

	var ans []int64

	for i := 0; i < len(que); i++ {
		if que[i][0] == 1 {
			x, y, z := que[i][1], que[i][2], que[i][3]
			goup := ln[x] - y
			pos := tail[x]
			for j := 0; j < 20; j++ {
				if (goup>>j)&1 == 1 {
					pos = pv[pos][j]
				}
			}
			tr.Add(in[pos], z)
			tr.Add(out[pos], -z)
		} else if que[i][0] == 2 {
			x := que[i][1]
			t[x] = tr.Get(in[tail[x]])
		} else {
			x := que[i][1]
			ans = append(ans, tr.Get(in[tail[x]])-t[x])
		}
	}

	return ans
}

type Fenwick struct {
	node []int64
}

func NewFenwick(n int) *Fenwick {
	arr := make([]int64, n)
	return &Fenwick{arr}
}

func (t *Fenwick) Add(x int, v int) {
	for x < len(t.node) {
		t.node[x] += int64(v)
		x |= (x + 1)
	}
}

func (t *Fenwick) Get(x int) int64 {
	var res int64
	for x >= 0 {
		res += t.node[x]
		x = (x & (x + 1)) - 1
	}
	return res
}

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

	q := readNum(reader)
	Q := make([][]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNNums(reader, 2)
	}

	res := solve(n, S, Q)

	var buf bytes.Buffer

	for i := 0; i < q; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}

	fmt.Print(buf.String())
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
	for {
		s, _ := reader.ReadString('\n')
		if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
			continue
		}
		for i := 0; i < len(s); i++ {
			if s[i] == '\n' {
				return s[:i]
			}
		}
		return s
	}
}

func readNum(reader *bufio.Reader) (a int) {
	for {
		bs, _ := reader.ReadBytes('\n')
		if len(bs) == 0 || len(bs) == 1 && bs[0] == '\n' {
			continue
		}
		readInt(bs, 0, &a)
		return
	}
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
	for {
		x := 0
		bs, _ := reader.ReadBytes('\n')
		if len(bs) == 0 || len(bs) == 1 && bs[0] == '\n' {
			continue
		}
		for i := 0; i < n; i++ {
			for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
				x++
			}
			x = readInt(bs, x, &res[i])
		}
		return res
	}
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

func solve(n int, S []string, Q [][]int) []int {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(S[i])
		buf.WriteByte('a' + 26)
	}

	T := buf.String()

	sa := NewStateMachine(T)

	pos_in_sa := make([][]int, n)

	for i := 0; i < n; i++ {
		l := len(S[i])

		pos_in_sa[i] = make([]int, l)
		cur := 0
		for j := 0; j < l; j++ {
			x := int(S[i][j] - 'a')
			cur = sa.states[cur].next[x]
			pos_in_sa[i][j] = cur
		}
	}

	ans := make([]int, len(Q))

	for i, cur := range Q {
		p, d := cur[0], cur[1]
		p--
		if d > len(S[p]) {
			continue
		}
		ans[i] = sa.children_cnt[pos_in_sa[p][d-1]]
	}
	return ans
}

const MAX_N = 1000009

type State struct {
	len  int
	link int
	next []int
}

func NewState() *State {
	res := new(State)
	res.link = -1
	res.next = make([]int, 27)
	for i := 0; i < 27; i++ {
		res.next[i] = -1
	}
	res.len = 0
	return res
}

type StateMachine struct {
	states       []*State
	children_cnt []int
}

func NewStateMachine(S string) *StateMachine {
	n := len(S)
	states := make([]*State, 2*MAX_N)
	states[0] = NewState()
	sz := 1
	var last int

	nextState := func() int {
		res := NewState()
		states[sz] = res
		sz++
		return sz - 1
	}

	extend := func(c byte) {
		cur := nextState()
		states[cur].len = states[last].len + 1
		p := last
		x := int(c - 'a')
		for p >= 0 && states[p].next[x] < 0 {
			states[p].next[x] = cur
			p = states[p].link
		}

		if p < 0 {
			states[cur].link = 0
		} else {
			q := states[p].next[x]
			if states[p].len+1 == states[q].len {
				states[cur].link = q
			} else {
				clone := nextState()
				states[clone].len = states[p].len + 1
				copy(states[clone].next, states[q].next)
				states[clone].link = states[q].link
				for p >= 0 && states[p].next[x] == q {
					states[p].next[x] = clone
					p = states[p].link
				}
				states[q].link = clone
				states[cur].link = clone
			}
		}
		last = cur
	}

	for i := 0; i < n; i++ {
		extend(S[i])
	}

	children_cnt := make([]int, 2*n)

	var dfs func(u int)

	dfs = func(u int) {
		if children_cnt[u] > 0 {
			return
		}
		children_cnt[u]++
		for i := 0; i < 26; i++ {
			if states[u].next[i] < 0 {
				continue
			}
			dfs(states[u].next[i])
			children_cnt[u] += children_cnt[states[u].next[i]]
		}
	}

	dfs(0)

	return &StateMachine{states, children_cnt}
}

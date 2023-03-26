package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--

		n, k := readTwoNums(reader)
		P := readNNums(reader, n)
		A := readNNums(reader, n)
		res := solve(k, P, A)

		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const MOD = 1e9 + 7

func solve(m int, p []int, a []int) int {
	n := len(p)
	P := make([]int, n+1)
	copy(P[1:], p)
	A := make([]int, n+1)
	copy(A[1:], a)
	outedge := make([]int, n+1)
	indegree := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if P[i] >= i {
			outedge[i] = P[i]
		} else {
			outedge[i] = outedge[P[i]]
		}
		indegree[outedge[i]]++
	}
	que := make([]int, n)
	var front, end int
	push := func(v int) {
		que[end] = v
		end++
	}
	pop := func() int {
		u := que[front]
		front++
		return u
	}
	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			push(i)
		}
	}
	var ans int64 = 1
	M := int64(m)
	for front < end {
		u := pop()
		if A[u] == 0 {
			ans = (ans * M) % MOD
		}
		indegree[outedge[u]]--
		if indegree[outedge[u]] == 0 {
			push(outedge[u])
		}
	}
	cycleValues := make(map[int]int)
	for u := 1; u <= n; u++ {
		if indegree[u] == 0 {
			continue
		}
		cycleValues[A[u]]++
	}
	if len(cycleValues) > 2 {
		return 0
	}
	if len(cycleValues) == 2 {
		if _, ok := cycleValues[0]; !ok {
			return 0
		}
	} else {
		if _, ok := cycleValues[0]; ok {
			ans = (ans * M) % MOD
		}
	}
	return int(ans)
}

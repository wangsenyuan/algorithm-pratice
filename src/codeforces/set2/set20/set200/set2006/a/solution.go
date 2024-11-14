package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	s := readString(reader)
	return solve(n, edges, s)
}

func solve(n int, edges [][]int, s string) int {
	deg := make([]int, n)

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		deg[u]++
		deg[v]++
	}

	cnt := make([]int, 4)

	for i := 1; i < n; i++ {
		if s[i] == '?' {
			cnt[3]++
		}
		if deg[i] == 1 {
			if s[i] == '0' {
				cnt[0]++
			} else if s[i] == '1' {
				cnt[1]++
			} else {
				cnt[2]++
			}
		}
	}
	cnt[3] -= cnt[2]

	var score int
	if s[0] == '?' {

		if cnt[0] > cnt[1] {
			// 必须先处理s[0], 且放置为1
			score = cnt[0] + cnt[2]/2
		} else if cnt[0] < cnt[1] {
			score = cnt[1] + cnt[2]/2
		} else {
			if cnt[3]%2 == 1 {
				// 让对方先走，逼迫它去修改s[0]
				score = cnt[0] + (cnt[2]+1)/2
			} else {
				score = cnt[0] + cnt[2]/2
			}
		}
	} else {
		// cnt[3] no use
		if s[0] == '0' {
			score += cnt[1]
		} else {
			score += cnt[0]
		}
		score += (cnt[2] + 1) / 2
	}

	return score
}

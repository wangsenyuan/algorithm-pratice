package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, 3*n)
	sessions := make([][]int, n)
	for i := range n {
		sessions[i] = readNNums(reader, 3)
	}
	k := readNum(reader)
	return solve(n, a, sessions, k)
}

func solve(n int, a []int, sessions [][]int, k int) []int {
	pos := make([]int, 3*n+1)
	for i, x := range a {
		pos[x] = i
	}

	marked := make([]bool, 3*n+1)
	for _, session := range sessions {
		i := slices.Index(session, k)
		if i < 0 {
			// k 还没有出现
			for _, x := range session {
				marked[x] = true
			}
		} else {
			u, v, w := session[0], session[1], session[2]
			if u == k {
				u = w
			} else if v == k {
				v = w
			}
			// u, v are the parters
			if u > v {
				u, v = v, u
			}
			if min(pos[u], pos[v]) < pos[k] {
				// k is not a leader
				return case1(n, k)
			}
			marked[u] = true
			marked[v] = true
			// k is the leader
			// w is the k
			var res1 []int
			var res2 []int
			for j := 1; j <= 3*n; j++ {
				if j == k {
					continue
				}
				if !marked[j] || j > v {
					res2 = append(res2, j)
				} else {
					res1 = append(res1, j)
				}
			}

			sort.Ints(res1)
			sort.Ints(res2)
			return append(res1, res2...)
		}
	}

	return nil
}

func case1(n int, k int) []int {
	var res []int
	for i := 1; i <= 3*n; i++ {
		if i != k {
			res = append(res, i)
		}
	}
	return res
}

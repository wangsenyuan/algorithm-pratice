package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := make([]string, n)
	for i := 0; i < n; i++ {
		A[i] = readString(reader)
	}

	res := solve(n, A)

	if len(res) == 0 {
		fmt.Println(-1)
		return
	}

	fmt.Printf("%d %d %d\n", res[0], res[1], res[2])
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

func solve(n int, A []string) []int {
	// 感觉用dfs就可以了吧

	vis := make([]int, n)
	trace := make([]int, n)
	for i := 0; i < n; i++ {
		trace[i] = -1
	}

	var cycle []int

	var dfs func(u int)

	dfs = func(u int) {
		vis[u]++
		for v := 0; v < n; v++ {
			if len(cycle) > 0 {
				return
			}
			if A[u][v] == '1' {
				// a directed edge
				if vis[v] == 2 {
					continue
				}
				if vis[v] == 1 {
					// a cycle
					x := u
					for x != v {
						cycle = append(cycle, x)
						x = trace[x]
					}
					cycle = append(cycle, v)
					return
				}
				trace[v] = u
				dfs(v)
			}
		}
		vis[u]++
	}

	for i := 0; i < n && len(cycle) == 0; i++ {
		if vis[i] == 0 {
			dfs(i)
		}
	}

	if len(cycle) == 0 {
		return nil
	}

	reverse(cycle)

	for len(cycle) > 3 {
		u := cycle[0]
		v := cycle[2]
		if A[v][u] == '1' {
			// found one
			cycle = cycle[:3]
		} else {
			// A[u][v] = '1'
			cycle[1] = u
			cycle = cycle[1:]
		}
	}
	for i := 0; i < 3; i++ {
		cycle[i]++
	}
	return cycle
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

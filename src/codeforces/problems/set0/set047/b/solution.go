package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	conds := make([]string, 3)

	for i := 0; i < 3; i++ {
		conds[i] = readString(reader)
	}

	res := solve(conds)

	if len(res) == 0 {
		fmt.Println("Impossible")
	} else {
		fmt.Println(res)
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

func solve(conditions []string) string {
	conn := make([][]int, 3)
	for i := 0; i < 3; i++ {
		conn[i] = make([]int, 3)
	}

	for _, cond := range conditions {
		a := int(cond[0] - 'A')
		b := int(cond[2] - 'A')
		m := cond[1]
		if m == '>' {
			if conn[a][b] == -1 {
				return ""
			}
			conn[a][b] = 1
			conn[b][a] = -1
		} else {
			// m == '<'
			if conn[a][b] == 1 {
				return ""
			}
			conn[a][b] = -1
			conn[b][a] = 1
		}
	}

	// conn[a][b] == 1 => a > b
	vis := make([]int, 3)

	var dfs func(u int) bool

	dfs = func(u int) bool {
		if vis[u] == 1 {
			return false
		}
		if vis[u] == 2 {
			return true
		}
		vis[u]++
		for v := 0; v < 3; v++ {
			if u != v {
				if conn[u][v] == 1 {
					if !dfs(v) {
						return false
					}
				}
			}
		}
		vis[u]++
		return true
	}

	for u := 0; u < 3; u++ {
		if !dfs(u) {
			return ""
		}
	}

	if conn[0][1] == 1 && conn[0][2] == 1 {
		if conn[1][2] == 1 {
			return "CBA"
		}
		return "BCA"
	}

	if conn[1][0] == 1 && conn[1][2] == 1 {
		if conn[0][2] == 1 {
			return "CAB"
		}
		return "ACB"
	}

	if conn[0][1] == 1 {
		return "BAC"
	}

	return "ABC"
}

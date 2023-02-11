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
		n := readNum(reader)
		S := readString(reader)
		res := solve(n, S)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			for i := 0; i < len(res); i++ {
				buf.WriteString(fmt.Sprintf("%d %d\n", res[i][0], res[i][1]))
			}
		}
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

func solve(n int, S string) [][]int {
	//S[i] = 0, then deg[i] is even, else deg[i] is odd
	// and if (i, j) connected, then no edge from i + 1 go outside
	// 且它们组成一棵树
	// 所以有 sum(deg) = 2 * (n - 1)
	// 所有的叶子节点的deg = 1
	// 假设(i, j) 连起来了， 它们中间的节点要么通过i连接，要们通过j连接
	// 考虑一棵树的话，通过dfs的方式，可以生成一个满足条件的环，但问题是怎么找到root？
	// root是不是不重要？一棵树，从任何节点都可以dfs
	var cnt_odd int
	for i := 0; i < n; i++ {
		if S[i] == '1' {
			cnt_odd++
		}
	}
	if cnt_odd == 0 || cnt_odd&1 == 1 {
		return nil
	}

	inc := func(i int) int {
		return (i + 1) % n
	}

	var res [][]int

	for i := 1; i < n; i++ {
		if S[i-1] == '1' {
			j := inc(i)
			for j != i {
				k := j
				prev := i
				for k != i {
					res = append(res, []int{prev + 1, k + 1})
					prev = k
					k = inc(k)
					if S[prev] == '1' {
						break
					}
				}
				j = k
			}
			break
		}
	}

	return res
}

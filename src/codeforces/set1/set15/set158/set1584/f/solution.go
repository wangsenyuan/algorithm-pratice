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
		words := make([]string, n)
		for i := 0; i < n; i++ {
			words[i] = readString(reader)
		}
		res := solve(words)

		buf.WriteString(fmt.Sprintf("%d\n%s\n", len(res), res))
	}
	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

type Pair struct {
	first  int
	second int
}

var cc = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func solve(words []string) string {

	// 计算每个字母在各个字符串中的位置
	n := len(words)
	pos := make([][][]int, 123)

	for i := 0; i < 123; i++ {
		pos[i] = make([][]int, n)
	}

	for i, word := range words {
		pos[0][i] = []int{-1, -1}
		for j, x := range word {
			if pos[x][i] == nil {
				pos[x][i] = []int{j, j}
			}
			pos[x][i][1] = j
		}
	}

	from := make([][]Pair, 223)
	mem := make([][]int, 123)
	for i := 0; i < 123; i++ {
		mem[i] = make([]int, 1<<n)
		from[i] = make([]Pair, 1<<n)
		for j := 0; j < 1<<n; j++ {
			mem[i][j] = -1
		}
	}

	var dfs func(state int, u int) int

	dfs = func(state int, u int) (res int) {
		if mem[u][state] >= 0 {
			return mem[u][state]
		}
		var mask int
		var next int
		for _, c := range cc {
			var tmp int
			for i := 0; i < n; i++ {
				if pos[c][i] == nil {
					tmp = -1
					break
				}
				pos_u := pos[u][i][(state>>i)&1]
				if pos[c][i][0] > pos_u {
					// 最好使用这个位置
					continue
				}
				if pos[c][i][1] <= pos_u {
					tmp = -1
					break
				}
				tmp |= 1 << i
			}
			if tmp < 0 {
				continue
			}
			res_v := dfs(tmp, int(c))
			if res_v > res {
				res = res_v
				mask = tmp
				next = int(c)
			}
		}

		from[u][state] = Pair{mask, next}
		res++
		mem[u][state] = res
		return res
	}

	dfs(0, 0)

	var lcs []byte

	for p := from[0][0]; p.second > 0; p = from[p.second][p.first] {
		lcs = append(lcs, byte(p.second))
	}

	return string(lcs)
}

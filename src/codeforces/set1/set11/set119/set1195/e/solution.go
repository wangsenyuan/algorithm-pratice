package main

import (
	"bufio"
	"fmt"
	"os"
)

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
func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	line1 := readNNums(reader, 4)
	n, m, a, b := line1[0], line1[1], line1[2], line1[3]
	line2 := readNNums(reader, 4)
	g0, x, y, z := line2[0], line2[1], line2[2], line2[3]
	return solve(n, m, a, b, g0, x, y, z)
}

type pair struct {
	first  int
	second int
}

func solve(n int, m int, a int, b int, g0 int, x int, y int, z int) int {
	sz := n * m
	g := make([]int, sz)
	g[0] = g0
	for i := 1; i < sz; i++ {
		g[i] = (g[i-1]*x + y) % z
	}

	cols := make([][]pair, m)
	pos := make([][]int, m)
	for i := 0; i < m; i++ {
		cols[i] = make([]pair, n)
		// pos[i][0] for tail, pos[i][1] for head
		pos[i] = make([]int, 2)
	}

	for i := 0; i < a; i++ {
		for j := 0; j < m; j++ {
			val := g[i*m+j]
			for pos[j][1] > pos[j][0] && cols[j][pos[j][1]-1].first >= val {
				pos[j][1]--
			}
			cols[j][pos[j][1]] = pair{val, i}
			pos[j][1]++
		}
	}

	row := make([]pair, m)
	var res int

	for i := a; i <= n; i++ {
		var head, tail int
		for j := 0; j < m; j++ {
			if tail < head && row[tail].second == j-b {
				tail++
			}
			val := cols[j][pos[j][0]].first
			for head > tail && row[head-1].first >= val {
				head--
			}
			row[head] = pair{val, j}
			head++
			if j >= b-1 {
				res += row[tail].first
			}
		}
		if i < n {
			for j := 0; j < m; j++ {
				// 这个肯定是成立的
				if cols[j][pos[j][0]].second == i-a {
					pos[j][0]++
				}
				val := g[i*m+j]
				for pos[j][1] > pos[j][0] && cols[j][pos[j][1]-1].first >= val {
					pos[j][1]--
				}
				cols[j][pos[j][1]] = pair{val, i}
				pos[j][1]++
			}
		}
	}

	return res
}

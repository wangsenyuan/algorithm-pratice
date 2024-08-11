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
		n, k := readTwoNums(reader)
		chords := make([][]int, k)
		for i := 0; i < k; i++ {
			chords[i] = readNNums(reader, 2)
		}
		res := solve(n, chords)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(n int, chords [][]int) int {

	k := len(chords)

	m := n - k

	vis := make([]bool, 2*n)

	var all [][]int
	for _, cur := range chords {
		a, b := cur[0]-1, cur[1]-1
		vis[a] = true
		vis[b] = true
		a, b = min(a, b), max(a, b)
		all = append(all, []int{a, b})
	}

	var free []int
	for i := 0; i < 2*n; i++ {
		if !vis[i] {
			free = append(free, i)
		}
	}

	for i := 0; i < m; i++ {
		all = append(all, []int{free[i], free[i+m]})
	}

	var res int

	for i, cur := range all {
		a, b := cur[0], cur[1]
		for j, oth := range all {
			if i == j {
				continue
			}
			c, d := oth[0], oth[1]
			if d < a || b < c || a < c && d < b || c < a && b < d {
				continue
			}
			res++
		}
	}

	return res / 2
}

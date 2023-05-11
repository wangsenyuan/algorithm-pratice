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
		m, n := readTwoNums(reader)
		field := make([]string, m)
		for i := 0; i < m; i++ {
			field[i] = readString(reader)[:n]
		}
		cnt, res := solve(field)
		buf.WriteString(fmt.Sprintf("%d %d\n", cnt, res))
		if tc > 0 {
			readString(reader)
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

const MOD = 151109

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

const MAX = 76867

var tr [1 << 20]int
var xp [20][20][MAX]int
var touched [8858835]int

func init() {

	isOk := func(state int) bool {
		var nv int
		var prev int
		for state > 0 {
			if state&1 == 1 {
				nv += prev
				prev = 1
			} else {
				prev = 0
			}
			state >>= 1
		}
		return nv <= 1
	}
	var id int
	for state := 0; state < 1<<20; state++ {
		if isOk(state) {
			tr[state] = id
			id++
		}
	}

	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			for k := 0; k < MAX; k++ {
				xp[i][j][k] = -1
			}
		}
	}
}

func solve(field []string) (int, int) {
	var touchCnt int
	m := len(field)
	n := len(field[0])
	ALL := (1 << n) - 1

	var dfs func(x int, y int, state int) int

	dfs = func(x int, y int, state int) int {
		if y >= n {
			return dfs(x+1, 0, state)
		}
		if x >= m {
			return (1 << 9) + 0
		}
		if xp[x][y][tr[state]] >= 0 {
			return xp[x][y][tr[state]]
		}
		touched[touchCnt] = (((tr[state] << 5) + y) << 5) + x
		touchCnt++
		var v, w int
		if field[x][y] == '.' && (y == 0 || state&1 == 0) && (state>>(n-1))&1 == 0 {
			tmp := dfs(x, y+1, ((state<<1)|1)&ALL)
			v = 1 + (tmp & 511)
			w = tmp >> 9
		}
		tmp := dfs(x, y+1, (state<<1)&ALL)
		if tmp&511 > v {
			v = tmp & 511
			w = tmp >> 9
		} else if tmp&511 == v {
			w = add(w, tmp>>9)
		}
		xp[x][y][tr[state]] = (w << 9) + v
		return xp[x][y][tr[state]]
	}

	ans := dfs(0, 0, 0)

	for i := 0; i < touchCnt; i++ {
		xp[touched[i]&31][(touched[i]>>5)&31][touched[i]>>10] = -1
	}

	return ans & 511, ans >> 9
}

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
		n := readNum(reader)
		cmds := make([][]int, n)
		for i := 0; i < n; i++ {
			cmds[i] = readNNums(reader, 2)
		}
		res := solve(cmds)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

const INF = 1e9 + 10

func solve(cmds [][]int) int {
	// cmd[i] = [t, x], at time t move to x
	var res int
	var cur int
	for i := 0; i < len(cmds); {
		t, x := cmds[i][0], cmds[i][1]
		dir := sign(x - cur)
		took := abs(x - cur)
		if took == 0 {
			// at time t move to position x, 如果它已经在位置x了， 不需要移动
			res++
			i++
			continue
		}
		j := i
		until := t + took
		for ; i < len(cmds) && cmds[i][0] < until; i++ {
			if i+1 < len(cmds) {
				tmp_pos := cur + dir*(cmds[i][0]-t)
				tmp_pos2 := cur + dir*(cmds[i+1][0]-t)
				if dir > 0 {
					tmp_pos2 = min(tmp_pos2, x)
				} else {
					tmp_pos2 = max(tmp_pos2, x)
				}

				y := cmds[i][1]

				if y >= min(tmp_pos, tmp_pos2) && y <= max(tmp_pos, tmp_pos2) {
					res++
				}
			} else {
				if i == j {
					// it will stop at x anyway
					res++
				} else {
					// i > j
					tmp_pos := cur + dir*(cmds[i][0]-t)
					tmp_pos2 := x
					y := cmds[i][1]
					if y >= min(tmp_pos, tmp_pos2) && y <= max(tmp_pos, tmp_pos2) {
						res++
					}
				}
			}
		}

		cur = x
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	return a + b - min(a, b)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func sign(num int) int {
	if num == 0 {
		return 0
	}
	if num < 0 {
		return -1
	}
	return 1
}

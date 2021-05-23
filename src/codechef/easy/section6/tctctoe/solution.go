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

	board := make([]string, 3)

	for tc > 0 {
		tc--
		for i := 0; i < 3; i++ {
			board[i], _ = reader.ReadString('\n')
		}
		res := solve(board)
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

var states map[int]int

func init() {

	states = make(map[int]int)

	var dfs func(cur []byte, turn byte)

	dfs = func(cur []byte, turn byte) {
		v := encode(cur)
		states[v] = 0
		x := win(cur, 'X')
		if x {
			states[v] = 2
			return
		}
		o := win(cur, 'O')
		if o {
			states[v] = 1
			return
		}
		next := byte(int('X') + int('O') - int(turn))

		for i := 0; i < 9; i++ {
			if cur[i] == '_' {
				cur[i] = turn
				dfs(cur, next)
				cur[i] = '_'
			}
		}
	}

	buf := []byte("_________")

	dfs(buf, 'X')
}

func encode(s []byte) int {
	var res int
	for i := 0; i < 9; i++ {
		var v int
		if s[i] == 'X' {
			v = 2
		} else if s[i] == 'O' {
			v = 1
		}
		res = res*3 + v
	}
	return res
}

func solve(board []string) int {
	buf := make([]byte, 9)
	var cnt int
	for i := 0; i < 3; i++ {
		copy(buf[i*3:], board[i][:3])
		for j := 0; j < 3; j++ {
			if board[i][j] == '_' {
				cnt++
			}
		}
	}
	v := encode(buf)
	if w, found := states[v]; found {
		if w == 0 {
			if cnt == 0 {
				return 1
			}
			return 2
		}
		return 1
	}
	return 3
}

//
//func solve1(board []string) int {
//	var cntx, cnto int
//
//	for i := 0; i < 3; i++ {
//		for j := 0; j < 3; j++ {
//			if board[i][j] == 'X' {
//				cntx++
//			} else if board[i][j] == 'O' {
//				cnto++
//			}
//		}
//	}
//
//	if cnto > cntx || cntx-cnto > 1 {
//		return 3
//	}
//	xwin := win(board, 'X')
//	owin := win(board, 'O')
//
//	if xwin && owin {
//		return 3
//	}
//
//	if owin && cntx != cnto {
//		return 3
//	}
//
//	if xwin && cntx == cnto {
//		return 3
//	}
//
//	if xwin || owin || cntx+cnto == 9 {
//		return 1
//	}
//	return 2
//}

func win(board []byte, x byte) bool {
	for r := 0; r < 3; r++ {
		ok := true
		for c := 0; c < 3; c++ {
			if board[r*3+c] != x {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}

	for c := 0; c < 3; c++ {
		ok := true
		for r := 0; r < 3; r++ {
			if board[r*3+c] != x {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	ok1 := true
	ok2 := true
	for i := 0; i < 3; i++ {
		if board[i*3+i] != x {
			ok1 = false
		}
		if board[i*3+2-i] != x {
			ok2 = false
		}
	}

	return ok1 || ok2
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

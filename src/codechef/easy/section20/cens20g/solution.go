package main

import (
	"bufio"
	"bytes"
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var ans bytes.Buffer

	for tc > 0 {
		tc--
		S, _ := reader.ReadString('\n')
		S = S[:len(S)-1]
		x1, y1 := readTwoNums(reader)
		solver := NewSolver(S, x1, y1)

		Q := readNum(reader)
		for Q > 0 {
			Q--
			x2, y2 := readTwoNums(reader)
			res := solver.Query(x2, y2)
			if res {
				ans.WriteString(fmt.Sprintf("YES %d\n", abs(y2-y1)+abs(x2-x1)))
			} else {
				ans.WriteString("NO")
			}
			ans.WriteByte('\n')
		}
	}
	fmt.Println(ans.String())
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type Solver struct {
	x1, y1 int
	cnt    []int
}

func NewSolver(S string, x1, y1 int) Solver {
	cnt := make([]int, 4)
	for i := 0; i < len(S); i++ {
		if S[i] == 'R' {
			cnt[0]++
		} else if S[i] == 'L' {
			cnt[1]++
		} else if S[i] == 'U' {
			cnt[2]++
		} else {
			cnt[3]++
		}
	}

	return Solver{x1, y1, cnt}
}

func (solver Solver) Query(x2, y2 int) bool {
	x1, y1 := solver.x1, solver.y1
	cnt := solver.cnt
	if x2 >= x1 && cnt[0] < x2-x1 {
		return false
	}
	if x2 < x1 && cnt[1] < x1-x2 {
		return false
	}
	if y2 >= y1 && cnt[2] < y2-y1 {
		return false
	}

	if y2 < y1 && cnt[3] < y1-y2 {
		return false
	}
	return true
}

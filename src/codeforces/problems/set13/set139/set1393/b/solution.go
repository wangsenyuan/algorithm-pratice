package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	A := readNNums(reader, n)

	solver := NewSolver(n, A)

	q := readNum(reader)

	var buf bytes.Buffer

	for q > 0 {
		q--
		s, _ := reader.ReadBytes('\n')

		var x int
		readInt(s, 2, &x)
		var res bool
		if s[0] == '+' {
			res = solver.Add(x)
		} else {
			res = solver.Add(-x)
		}

		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

type Solver struct {
	planks map[int]int
	two    int
	four   int
}

func NewSolver(n int, A []int) Solver {
	planks := make(map[int]int)

	for i := 0; i < n; i++ {
		planks[A[i]]++
	}

	var two int
	var four int
	for _, v := range planks {
		two += v / 2
		four += v / 4
	}

	return Solver{planks, two, four}
}

func (solver *Solver) Add(x int) bool {
	sign := 1
	if x < 0 {
		x = -x
		sign = -1
	}
	y := solver.planks[x]

	two, four := solver.two, solver.four

	two -= y / 2
	four -= y / 4

	y += sign

	two += y / 2
	four += y / 4

	solver.two = two
	solver.four = four
	solver.planks[x] = y

	return four >= 1 && two >= 4
}

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

	for tc > 0 {
		tc--
		n, q := readTwoNums(reader)
		A := readNNums(reader, n)
		solver := solve(n, A)

		for q > 0 {
			q--
			l, r := readTwoNums(reader)
			fmt.Println(solver.Query(l, r))
		}
	}
}

type Solver struct {
	n    int
	zero []int
	msb  [][]int
	on   [][]int
}

const H = 30

func solve(n int, A []int) Solver {
	zero := make([]int, n+1)
	msb := make([][]int, n+1)
	on := make([][]int, n+1)
	msb[0] = make([]int, H)
	on[0] = make([]int, H)
	for i := 0; i < n; i++ {
		zero[i+1] = zero[i]
		x := A[i]
		if x == 0 {
			zero[i+1]++
		}
		msb[i+1] = make([]int, H)
		on[i+1] = make([]int, H)
		var j = -1
		for k := 0; k < H; k++ {
			on[i+1][k] = on[i][k]
			if (x>>uint(k))&1 == 1 {
				on[i+1][k]++
				j = k
			}
		}
		for k := 0; k < H; k++ {
			msb[i+1][k] = msb[i][k]
			if j == k {
				msb[i+1][k]++
			}
		}
	}

	return Solver{n, zero, msb, on}
}

func (solver Solver) Query(l, r int) int64 {
	var ans int64
	ans += int64(solver.zero[r]-solver.zero[l-1]) * int64(r-l+1)

	for k := 0; k < H; k++ {
		a := solver.msb[r][k] - solver.msb[l-1][k]
		b := solver.on[r][k] - solver.on[l-1][k]
		ans += int64(a) * int64(b-a)
	}

	return ans
}

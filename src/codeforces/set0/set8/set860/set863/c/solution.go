package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	k, a, b := readThreeNums(reader)
	A := make([][]int, 3)
	for i := 0; i < 3; i++ {
		A[i] = readNNums(reader, 3)
	}
	B := make([][]int, 3)
	for i := 0; i < 3; i++ {
		B[i] = readNNums(reader, 3)
	}

	res := solve(k, a, b, A, B)

	fmt.Printf("%d %d\n", res[0], res[1])
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

func solve(k int, a int, b int, A [][]int, B [][]int) []int {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			A[i][j]--
			B[i][j]--
		}
	}

	occ := make([][]int, 3)
	for i := 0; i < 3; i++ {
		occ[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			occ[i][j] = -1
		}
	}
	a--
	b--

	check := func(a, b int) Round {
		if a == b {
			return Round{0, 0}
		}
		if a == 0 && b == 2 || a == 1 && b == 0 || a == 2 && b == 1 {
			return Round{1, 0}
		}
		return Round{0, 1}
	}

	var arr []Round
	arr = append(arr, Round{0, 0})
	x, y := a, b
	for {
		cur := check(x, y)
		last := arr[len(arr)-1]
		arr = append(arr, last.Add(cur))
		occ[x][y] = len(arr) - 1
		u, v := A[x][y], B[x][y]
		if occ[u][v] >= 0 {
			// keep x, y
			break
		}
		x, y = u, v
	}
	if k < len(arr) {
		return arr[k].asSlice()
	}
	u, v := A[x][y], B[x][y]
	res := arr[occ[u][v]-1]
	k -= occ[u][v] - 1
	if u == x && v == y {
		tmp := check(u, v).Mul(k)
		return res.Add(tmp).asSlice()
	}

	ln := occ[x][y] - occ[u][v] + 1
	cnt, rem := k/ln, k%ln

	res1 := arr[occ[x][y]].Sub(arr[occ[u][v]-1]).Mul(cnt)

	res2 := arr[occ[u][v]-1+rem].Sub(arr[occ[u][v]-1])

	res = res.Add(res1)
	res = res.Add(res2)

	return res.asSlice()
}

type Round struct {
	alice int
	bob   int
}

func (a Round) Add(b Round) Round {
	return Round{a.alice + b.alice, a.bob + b.bob}
}

func (a Round) Sub(b Round) Round {
	return Round{a.alice - b.alice, a.bob - b.bob}
}

func (a Round) asSlice() []int {
	return []int{a.alice, a.bob}
}

func (a Round) Mul(k int) Round {
	return Round{a.alice * k, a.bob * k}
}

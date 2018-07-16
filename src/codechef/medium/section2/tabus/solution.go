package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	firstLine := readNNums(scanner, 3)
	N, T, M := firstLine[0], firstLine[1], firstLine[2]
	buses := make([][]int, M)
	for i := 0; i < M; i++ {
		buses[i] = readNNums(scanner, 4)
	}
	fmt.Printf("%d\n", solve(N, T, M, buses))
}

func solve(N int, T int, M int, buses [][]int) int {

	bs := make(Buses, M)
	for i := 0; i < M; i++ {
		bs[i] = Bus{buses[i][0] - 1, buses[i][1] - 1, buses[i][2], buses[i][3]}
	}

	sort.Sort(bs)

	S := make([][]int, N)
	for i := 0; i < N; i++ {
		S[i] = make([]int, 0, 10)
	}
	check := func(W int) bool {
		for i := 0; i < N; i++ {
			S[i] = S[i][0:0]
		}
		S[0] = append(S[0], 0)

		for i := 0; i < M; i++ {
			bus := bs[i]
			u, v, s, e := bus.u, bus.v, bus.s, bus.e

			if len(S[u]) == 0 || S[u][0] > s {
				// all end time from start station u, is after time s,
				continue
			}
			k := sort.Search(len(S[u]), func(k int) bool {
				return S[u][k] > s
			})
			if s-S[u][k-1] <= W {
				S[v] = append(S[v], e)
			}
		}

		return len(S[N-1]) > 0 && S[N-1][0] <= T
	}

	left, right := 0, 1000000010

	if !check(right) {
		return -1
	}

	for left < right {
		mid := left + (right-left)>>1
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Bus struct {
	u, v, s, e int
}

type Buses []Bus

func (this Buses) Len() int {
	return len(this)
}

func (this Buses) Less(i, j int) bool {
	return this[i].e < this[j].e
}

func (this Buses) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

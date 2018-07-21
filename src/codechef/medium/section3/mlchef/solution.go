package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
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

	n := readNum(scanner)

	H := make([]int, n)
	P := make([]int, n)

	for i := 0; i < n; i++ {
		H[i], P[i] = readTwoNums(scanner)
	}
	Q := readNum(scanner)
	input := make([][]int, Q)

	for i := 0; i < Q; i++ {
		scanner.Scan()
		bs := scanner.Bytes()
		var kind int
		pos := readInt(bs, 0, &kind)
		if kind == 1 {
			input[i] = make([]int, 3)
			input[i][0] = 1
			pos = readInt(bs, pos+1, &input[i][1])
			readInt(bs, pos+1, &input[i][2])
		} else {
			input[i] = make([]int, 2)
			input[i][0] = 2
			readInt(bs, pos+1, &input[i][1])
		}
	}

	// fmt.Printf("[debug] %d %v %v %d %v", n, H, P, Q, input)

	res := solve(n, H, P, Q, input)
	for _, ans := range res {
		fmt.Println(ans)
	}
}

func solve(N int, H []int, P []int, Q int, input [][]int) []int {
	start := make([]int, N+1)
	end := make([]int, N+1)
	X := make([]int, N+1)

	outs := make([][]int, N+1)

	for i := 1; i <= N; i++ {
		p := P[i-1]
		if outs[p] == nil {
			outs[p] = make([]int, 0, 3)
		}
		outs[p] = append(outs[p], i)
	}

	var dfs func(u int, time *int)

	dfs = func(u int, time *int) {
		if u > 0 {
			X[*time] = H[u-1]
		}
		start[u] = *time
		*time++
		for _, v := range outs[u] {
			dfs(v, time)
		}
		end[u] = *time
	}

	dfs(0, new(int))

	st := CreateSegTree(N+1, X)

	res := make([]int, Q)
	var j int

	for i := 0; i < Q; i++ {
		q := input[i]
		if q[0] == 1 {
			a, x := q[1], q[2]
			st.Poison(start[a]+1, end[a]-1, x)
		} else {
			a := q[1]
			res[j] = st.CountAlive(start[a]+1, end[a]-1)
			j++
		}
	}

	return res[:j]
}

type SegTree struct {
	poison []int
	alive  []int
	health []int
	size   int
}

func CreateSegTree(n int, X []int) SegTree {
	poison := make([]int, 4*n)
	alive := make([]int, 4*n)
	health := make([]int, 4*n)
	var loop func(i int, start, end int)

	loop = func(i int, start, end int) {
		if start == end {
			health[i] = X[start]
			alive[i] = 1
			return
		}
		mid := (start + end) >> 1
		loop(2*i+1, start, mid)
		loop(2*i+2, mid+1, end)
		alive[i] = alive[2*i+1] + alive[2*i+2]
		health[i] = min(health[2*i+1], health[2*i+2])
	}

	loop(0, 0, n-1)

	return SegTree{poison, alive, health, n}
}

func (st *SegTree) Poison(left, right int, vol int) {
	if left > right {
		return
	}
	poison := st.poison
	health := st.health
	alive := st.alive
	var loop func(i int, start, end int, vol int)

	loop = func(i int, start, end int, vol int) {
		if end < left || start > right {
			//out of range
			return
		}
		mid := (start + end) >> 1

		if left <= start && end <= right {
			poison[i] += vol
			health[i] -= vol
			if health[i] <= 0 {
				// need to progate
				if start == end {
					//leaf node
					health[i] = math.MaxInt32
					poison[i] = 0
					alive[i] = 0
				} else {
					health[2*i+1] -= poison[i]
					poison[2*i+1] += poison[i]
					health[2*i+2] -= poison[i]
					poison[2*i+2] += poison[i]
					poison[i] = 0
					loop(2*i+1, start, mid, 0)
					loop(2*i+2, mid+1, end, 0)
					health[i] = min(health[2*i+1], health[2*i+2])
					alive[i] = alive[2*i+1] + alive[2*i+2]
				}
			}
			return
		}
		health[2*i+1] -= poison[i]
		poison[2*i+1] += poison[i]
		health[2*i+2] -= poison[i]
		poison[2*i+2] += poison[i]
		loop(2*i+1, start, mid, vol)
		loop(2*i+2, mid+1, end, vol)
		health[i] = min(health[2*i+1], health[2*i+2])
		alive[i] = alive[2*i+1] + alive[2*i+2]
	}

	loop(0, 0, st.size-1, vol)
}

func (st SegTree) CountAlive(left, right int) int {
	if left > right {
		return 0
	}
	alive := st.alive
	var loop func(i int, start, end int) int

	loop = func(i int, start, end int) int {
		if end < left || start > right {
			return 0
		}
		if left <= start && end <= right {
			return alive[i]
		}
		mid := (start + end) >> 1
		return loop(2*i+1, start, mid) + loop(2*i+2, mid+1, end)
	}

	return loop(0, 0, st.size-1)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

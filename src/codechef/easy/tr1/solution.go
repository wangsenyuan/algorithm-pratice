package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = sign * tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readPair(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	bs := scanner.Bytes()
	x := readInt(bs, 0, &a)
	readInt(bs, x+1, &b)
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)
	for i := 0; i < t; i++ {
		n := readNum(scanner)
		roads := make([][]int, n-1)
		for j := 0; j < n-1; j++ {
			a, b := readPair(scanner)
			roads[j] = []int{a, b}
		}
		fmt.Println(solve(n, roads))
	}
}

const MOD = int64(1000000007)

func solve(n int, roads [][]int) int {
	tree := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		tree[i] = make([]int, 3)
	}

	for _, road := range roads {
		a, b := road[0], road[1]
		i := 0
		for i < 3 && tree[a][i] > 0 {
			i++
		}
		tree[a][i] = b
		i = 0
		for i < 3 && tree[b][i] > 0 {
			i++
		}
		tree[b][i] = a
	}

	im := make([]int64, n+1)
	var dfs func(v int, p int) int
	dfs = func(v int, p int) int {
		var res int64
		var cnt int
		for i := 0; i < 3 && tree[v][i] > 0; i++ {
			w := tree[v][i]
			if w == p {
				continue
			}
			tc := dfs(w, v)
			res += int64(cnt*tc) % MOD
			cnt += tc
		}
		im[v] = (res*2 + 2*int64(cnt) + 1) % MOD
		return cnt + 1
	}
	dfs(1, 0)
	var res int64
	for i := 1; i <= n; i++ {
		res = (res + int64(i)*im[i]) % MOD
	}
	return int(res)
}

func solve1(n int, roads [][]int) int {
	tree := make([][]int, n)
	for i := 0; i < n; i++ {
		tree[i] = make([]int, 3)
		tree[i][0] = -1
		tree[i][1] = -1
		tree[i][2] = -1
	}

	for _, road := range roads {
		a, b := road[0]-1, road[1]-1
		i := 0
		for i < 3 && tree[a][i] >= 0 {
			i++
		}
		tree[a][i] = b
		i = 0
		for i < 3 && tree[b][i] >= 0 {
			i++
		}
		tree[b][i] = a
	}

	var eulerTour func(v int, p int, depth int)

	level := make([]int, 2*n-1)
	euler := make([]int, 2*n-1)
	firstOccurrence := make([]int, n)
	for i := 0; i < n; i++ {
		firstOccurrence[i] = -1
	}
	idx := 0
	eulerTour = func(v int, p int, depth int) {
		euler[idx] = v
		level[idx] = depth
		idx++
		if firstOccurrence[v] < 0 {
			firstOccurrence[v] = idx - 1
		}
		for i := 0; i < 3 && tree[v][i] >= 0; i++ {
			w := tree[v][i]
			if w == p {
				continue
			}
			eulerTour(w, v, depth+1)
			euler[idx] = v
			level[idx] = depth
			idx++
		}
	}

	eulerTour(0, -1, 0)

	st := ConstructST(level, len(level))

	im := make([]int, n+1)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x, y := firstOccurrence[i], firstOccurrence[j]
			if x > y {
				y, x = x, y
			}
			z := euler[st.Query(x, y)] + 1
			im[z]++
		}
	}
	var ans int64
	for i := 1; i <= n; i++ {
		ans = (ans + int64(i*im[i])) % MOD
	}

	return int(ans)
}

type ST struct {
	tree []int
	arr  []int
	n    int
}

func ConstructST(arr []int, n int) ST {
	x := log2(n) + 1
	sz := 2*(1<<uint(x)) - 1
	tree := make([]int, sz)

	var process func(idx int, left int, right int)

	process = func(idx int, left int, right int) {
		if left == right {
			tree[idx] = left
			return
		}
		mid := (left + right) / 2
		process(2*idx+1, left, mid)
		process(2*idx+2, mid+1, right)

		if arr[tree[2*idx+1]] < arr[tree[2*idx+2]] {
			tree[idx] = tree[2*idx+1]
		} else {
			tree[idx] = tree[2*idx+2]
		}
	}

	process(0, 0, n-1)
	return ST{tree, arr, n}
}

func (st ST) Query(qs int, qe int) int {
	var process func(idx int, left int, right int) int
	process = func(idx int, left int, right int) int {
		if qs <= left && right <= qe {
			return st.tree[idx]
		} else if right < qs || left > qe {
			return -1
		}
		mid := (left + right) / 2
		q1 := process(2*idx+1, left, mid)
		q2 := process(2*idx+2, mid+1, right)
		if q1 == -1 {
			return q2
		}
		if q2 == -1 {
			return q1
		}
		if st.arr[q1] < st.arr[q2] {
			return q1
		}
		return q2
	}

	return process(0, 0, st.n-1)
}

func log2(n int) int {
	var ans int
	for n > 1 {
		n >>= 1
		ans++
	}
	return ans
}

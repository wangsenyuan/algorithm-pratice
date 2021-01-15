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
		n, w := readTwoNums(reader)
		A := readNNums(reader, n)
		solver := solve(n, w, A)

		q := readNum(reader)

		for q > 0 {
			q--
			p, c := readTwoNums(reader)
			res := solver.Update(p, c)
			buf.WriteString(fmt.Sprintf("%d\n", res))
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
	ans   int64
	W     int
	n     int
	A     []int
	tree1 *SegTree
	tree2 *SegTree
}

type SegTree struct {
	arr       []int
	sz        int
	initValue int
	fn        func(int, int) int
}

func (tree *SegTree) Update(pos int, v int) {
	pos += tree.sz
	tree.arr[pos] = v
	for pos > 0 {
		tree.arr[pos>>1] = tree.fn(tree.arr[pos], tree.arr[pos^1])
		pos >>= 1
	}
}

func (tree *SegTree) Find(l, r int) int {
	l += tree.sz
	r += tree.sz

	ans := tree.initValue

	for l < r {
		if l&1 == 1 {
			ans = tree.fn(ans, tree.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			ans = tree.fn(ans, tree.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return ans
}

func solve(n, w int, A []int) *Solver {
	//arr1 for max value
	arr1 := make([]int, 2*n)
	//arr2 for min value
	arr2 := make([]int, 2*n)

	for i := 0; i < len(arr1); i++ {
		arr1[i] = -1
		arr2[i] = n
	}

	tree1 := SegTree{arr1, n, -1, max}
	tree2 := SegTree{arr2, n, n, min}
	var ans int64
	var j int
	for i := 1; i <= n; i++ {
		if i == n || A[i] != A[i-1] {
			tree1.Update(i-1, i-1)
			tree2.Update(i-1, i-1)
			ans += int64(max(0, i-j-(w-1)))
			j = i
		}
	}

	return &Solver{ans, w, n, A, &tree1, &tree2}
}

func (solver *Solver) Update(p int, color int) int64 {
	p--

	if solver.A[p] == color {
		return solver.ans
	}

	ways := func(l int) int64 {
		return int64(max(0, l-(solver.W-1)))
	}
	ans := solver.ans
	n := solver.n

	if n == 1 {
		return ways(1)
	}

	if p == 0 {
		j := solver.tree2.Find(0, n)
		if j > p {
			// [0...j] are same
			ans = ans - ways(j+1) + ways(1) + ways(j)
		} else {
			// j == p
			jj := solver.tree2.Find(1, n)
			if color == solver.A[jj] {
				// now [0...jj] are same
				ans = ans - ways(1) - ways(jj) + ways(jj+1)
			}
			// else no change
		}
	} else if p == n-1 {
		i := solver.tree1.Find(0, p)
		// [i+1...p] are same
		if i+1 < p {
			ans = ans - ways(n-i) + ways(1) + ways(n-i-1)
		} else {
			// i + 1 == p
			ii := solver.tree1.Find(0, n-1)
			if color == solver.A[ii+1] {
				ans = ans - ways(1) - ways(n-1-ii) + ways(n-ii)
			}
			// else no change
		}
	} else {
		i := solver.tree1.Find(0, p)
		j := solver.tree2.Find(p, n)
		ii := solver.tree1.Find(0, i)
		jj := solver.tree2.Find(j+1, n)

		if i+1 == p && p == j {
			// previous p is a standalone point
			if solver.A[ii+1] == solver.A[jj] {
				if color == solver.A[jj] {
					ans = ans - ways(1) - ways(jj-p) - ways(i-ii) + ways(jj-ii)
				}
				// else no change
			} else {
				if color == solver.A[ii+1] {
					ans = ans - ways(1) - ways(i-ii) + ways(p-ii)
				} else if color == solver.A[jj] {
					ans = ans - ways(1) - ways(jj-j) + ways(jj-p+1)
				}
				// else no change
			}
		} else if i+1 == p {
			// p < j, after change, it will break
			if solver.A[i] != color {
				//just a new standalone point
				ans = ans - ways(j-i) + ways(1) + ways(j-p)
			} else {
				// p need to append to before
				ans = ans - ways(j-i) - ways(i-ii) + ways(j-p) + ways(p-ii)
			}
		} else if p == j {
			// [i+1....p] are same
			if solver.A[j+1] != color {
				ans = ans - ways(p-i) + ways(1) + ways(p-1-i)
			} else {
				ans = ans - ways(p-i) - ways(jj-p) + ways(p-1-i) + ways(jj-(p-1))
			}
		} else {
			// i + 1 < p && p < j
			ans = ans - ways(j-i) + ways(1) + ways(j-p) + ways(p-1-i)
		}
	}

	if p > 0 && solver.A[p-1] != solver.A[p] {
		solver.tree1.Update(p-1, -1)
		solver.tree2.Update(p-1, n)
	}

	if p == n-1 || solver.A[p] != solver.A[p+1] {
		solver.tree1.Update(p, -1)
		solver.tree2.Update(p, n)
	}

	solver.A[p] = color

	if p > 0 && solver.A[p-1] != solver.A[p] {
		solver.tree1.Update(p-1, p-1)
		solver.tree2.Update(p-1, p-1)
	}

	if p == n-1 || solver.A[p] != solver.A[p+1] {
		solver.tree1.Update(p, p)
		solver.tree2.Update(p, p)
	}

	solver.ans = ans

	return ans
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

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
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)

	pairs := make([][]int, n)

	for i := 0; i < n; i++ {
		pairs[i] = readNNums(scanner, 2)
	}

	fmt.Println(solve(n, pairs))
}

func solve(n int, pairs [][]int) int64 {
	sort.Sort(PS(pairs))

	fs := make([]int, n)
	for i := 0; i < n; i++ {
		fs[i] = pairs[i][1]
	}

	sort.Ints(fs)

	ps := make([]int, n)
	for i := 0; i < n; i++ {
		ps[i] = sort.SearchInts(fs, pairs[i][1])
	}

	bit := BIT(n)

	fn := func(read func(int) int, update func(int, int)) int64 {
		var ans int64

		for i := n - 1; i >= 0; i-- {
			ans += int64(read(ps[i]))
			update(ps[i], 1)
		}

		return ans
	}

	return bit(fn)
}

func BIT(n int) func(fn func(func(int) int, func(int, int)) int64) int64 {
	tree := make([]int, n+1)

	read := func(i int) int {
		var sum int
		i++
		for i > 0 {
			sum += tree[i]
			i -= i & (-i)
		}

		return sum
	}

	update := func(i int, val int) {
		i++
		for i <= n {
			tree[i] += val
			i += i & (-i)
		}
	}

	gn := func(fn func(func(int) int, func(int, int)) int64) int64 {
		return fn(read, update)
	}

	return gn
}

func solve1(n int, pairs [][]int) int64 {
	sort.Sort(PS(pairs))

	ps := make([]int, n)
	for i := 0; i < n; i++ {
		ps[i] = pairs[i][1]
	}

	tmp := make([]int, n)

	merge := func(left, mid, right int) int64 {
		var cnt int64

		i, j, k := left, mid, left

		for i < mid || j < right {
			if j == right || (i < mid && ps[i] <= ps[j]) {
				tmp[k] = ps[i]
				i++
			} else {
				tmp[k] = ps[j]
				cnt += int64(mid - i)
				j++
			}
			k++
		}
		copy(ps[left:right], tmp[left:right])
		return cnt
	}

	var mergeSort func(left, right int) int64

	mergeSort = func(left, right int) int64 {
		if left+1 >= right {
			return 0
		}

		mid := left + (right-left)/2
		ans := mergeSort(left, mid) + mergeSort(mid, right)
		return ans + merge(left, mid, right)
	}

	return mergeSort(0, n)
}

type PS [][]int

func (ps PS) Len() int {
	return len(ps)
}

func (ps PS) Less(i, j int) bool {
	return ps[i][0] < ps[j][0]
}

func (ps PS) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

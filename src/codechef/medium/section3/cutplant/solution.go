package main

import (
	"bufio"
	"fmt"
	"math"
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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		A := readNNums(scanner, n)
		B := readNNums(scanner, n)
		res := solve(n, A, B)
		fmt.Println(res)
	}
}

func solve(n int, A []int, B []int) int {
	for i := 0; i < n; i++ {
		if A[i] < B[i] {
			return -1
		}
	}

	H := make(Pairs, n)
	for i := 0; i < n; i++ {
		H[i] = Pair{B[i], i}
	}
	sort.Sort(H)

	arr1 := make([]int, 4*n)
	arr2 := make([]int, 4*n)

	var create func(i int, start, end int)
	create = func(i int, start, end int) {
		if start == end {
			arr1[i] = A[start]
			arr2[i] = B[start]
			return
		}
		mid := (start + end) >> 1
		create(2*i+1, start, mid)
		create(2*i+2, mid+1, end)
		arr1[i] = min(arr1[2*i+1], arr1[2*i+2])
		arr2[i] = max(arr2[2*i+1], arr2[2*i+2])
	}

	create(0, 0, n-1)

	var getMin func(i int, start, end int, left, right int) int

	getMin = func(i int, start, end int, left, right int) int {
		if left > end || right < start {
			return math.MaxInt32
		}
		if left <= start && end <= right {
			return arr1[i]
		}
		mid := (start + end) >> 1
		a := getMin(2*i+1, start, mid, left, right)
		b := getMin(2*i+2, mid+1, end, left, right)
		return min(a, b)
	}

	var getMax func(i int, start, end int, left, right int) int

	getMax = func(i int, start, end int, left, right int) int {
		if left > end || right < start {
			return 0
		}
		if left <= start && end <= right {
			return arr2[i]
		}
		mid := (start + end) >> 1
		a := getMax(2*i+1, start, mid, left, right)
		b := getMax(2*i+2, mid+1, end, left, right)
		return max(a, b)
	}

	var ans int
	lastCut := -1

	for i := 0; i < n; i++ {
		p := H[i]
		if A[p.second] == p.first {
			continue
		}
		if lastCut == -1 ||
			H[lastCut].first != p.first ||
			getMin(0, 0, n-1, H[lastCut].second+1, p.second) < p.first ||
			getMax(0, 0, n-1, H[lastCut].second+1, p.second) > p.first {
			ans++
		}
		lastCut = i
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first > this[j].first || (this[i].first == this[j].first && this[i].second < this[j].second)
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func solve1(n int, A []int, B []int) int {
	for i := 0; i < n; i++ {
		if A[i] < B[i] {
			return -1
		}
	}
	var ans int

	que := make([]int, n)
	var front, end int

	for i := 0; i < n; i++ {
		for end > front && que[end-1] < B[i] {
			end--
		}
		for front < end && que[front] > A[i] {
			front++
		}

		if A[i] != B[i] && (front == end || B[i] != que[end-1]) {
			ans++
			que[end] = B[i]
			end++
		}
	}

	return ans
}

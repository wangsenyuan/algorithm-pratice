package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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
		nums := readNNums(scanner, n)
		res := solve(n, nums)
		fmt.Println(res)
	}
}

func solve(n int, nums []int) int {
	xs := make([]int, 2*n)
	copy(xs, nums)
	copy(xs[n:], nums)

	// at most n
	arr := make([]int, n)

	check := func(index int) int {
		var end int
		arr[end] = xs[index]
		end++
		for i := 1; i < n; i++ {
			j := index + i
			// find the insert position for nums[j]
			k := sort.Search(end, func(k int) bool {
				return arr[k] >= xs[j]
			})
			if k == end {
				arr[end] = xs[j]
				end++
			} else {
				arr[k] = xs[j]
			}
		}
		return end
	}

	ps := make(Pairs, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{nums[i], i}
	}
	sort.Sort(ps)

	var best int
	for i := 0; i < 40 && i < n; i++ {
		tmp := check(ps[i].second)
		if tmp > best {
			best = tmp
		}
	}
	return best
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

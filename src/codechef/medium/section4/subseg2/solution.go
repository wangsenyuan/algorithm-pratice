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
	N, Q := readTwoNums(scanner)

	courses := make([][]int, N)
	for i := 0; i < N; i++ {
		courses[i] = readNNums(scanner, 2)
	}

	plans := make([][]int, Q)
	for i := 0; i < Q; i++ {
		plans[i] = readNNums(scanner, 2)
	}
	res := solve(N, Q, courses, plans)
	for _, ans := range res {
		fmt.Println(ans)
	}
}

func solve(N, Q int, courses [][]int, plans [][]int) []int {
	items := make(Items, N)
	for i := 0; i < N; i++ {
		items[i] = Item{courses[i][0], courses[i][1]}
	}
	sort.Sort(items)

	var j int
	for i := 1; i < N; i++ {
		a := items[j]
		b := items[i]
		if b.start > a.start {
			j++
			items[j] = items[i]
		}
	}
	j++
	items = items[:j]

	parent := make([][]int, j+1)

	for i := 0; i <= j; i++ {
		parent[i] = make([]int, 20)
	}

	for i := 0; i < 20; i++ {
		parent[j][i] = j
	}

	findParent := func(i int) int {
		cur := items[i]
		return sort.Search(j, func(k int) bool {
			return items[k].start > cur.end
		})
	}

	for i := j - 1; i >= 0; i-- {
		parent[i][0] = findParent(i)
	}
	for i := j - 1; i >= 0; i-- {
		for k := 1; k < 20; k++ {
			parent[i][k] = parent[parent[i][k-1]][k-1]
		}
	}

	res := make([]int, Q)

	for i := 0; i < Q; i++ {
		query := plans[i]
		start, end := query[0], query[1]

		a := sort.Search(j, func(a int) bool {
			return items[a].start >= start
		})
		var ans int
		if a < j && items[a].end <= end {
			ans = 1
			for k := 19; k >= 0; k-- {
				p := parent[a][k]
				if p < j && items[p].end <= end {
					ans += 1 << uint(k)
					a = p
				}
			}
		}

		res[i] = ans
	}

	return res
}

type Item struct {
	start, end int
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	if items[i].end < items[j].end {
		return true
	}
	if items[i].end == items[j].end && items[i].start > items[j].start {
		return true
	}
	return false
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

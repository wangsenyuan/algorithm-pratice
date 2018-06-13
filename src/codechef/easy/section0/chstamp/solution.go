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
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
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
	t := readNum(scanner)
	for i := 0; i < t; i++ {
		n, m := readTwoNums(scanner)
		stamps := readNNums(scanner, n)
		offers := make([][]int, m)
		for j := 0; j < m; j++ {
			offers[j] = readNNums(scanner, 3)
		}
		fmt.Println(solve(n, stamps, m, offers))
	}
}

func solve(n int, stamps []int, m int, offers [][]int) int64 {
	sort.Sort(oo(offers))

	maxType := make(map[int]int)

	getMax := func(x int) int {
		if px, found := maxType[x]; found {
			return px
		}
		maxType[x] = x
		return x
	}

	var find func(x int, set map[int]int) int

	find = func(x int, set map[int]int) int {
		if px, found := set[x]; found && px != x {
			set[x] = find(px, set)
		} else {
			set[x] = x
		}
		return set[x]
	}

	for i := m - 1; i >= 0; {
		set := make(map[int]int)
		j := i
		tps := make(map[int]bool)
		for j >= 0 && offers[j][0] == offers[i][0] {
			a, b := offers[j][1], offers[j][2]
			tps[a] = true
			tps[b] = true
			pa, pb := find(a, set), find(b, set)
			if pa != pb {
				set[pa] = pb
			}
			j--
		}
		max := make(map[int]int)
		for x := range tps {
			px := find(x, set)
			tmp := getMax(x)
			if tmp > max[px] {
				max[px] = tmp
			}
		}

		for x := range tps {
			px := find(x, set)
			maxType[x] = max[px]
		}
		i = j
	}

	var ans int64
	for i := 0; i < n; i++ {
		x := stamps[i]
		y := getMax(x)
		ans += int64(y)
	}
	return ans
}

type oo [][]int

func (this oo) Len() int {
	return len(this)
}

func (this oo) Less(i, j int) bool {
	return this[i][0] < this[j][0]
}

func (this oo) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

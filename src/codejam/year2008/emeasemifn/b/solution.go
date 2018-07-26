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

	for i := 1; i <= tc; i++ {
		n := readNum(scanner)
		C := make([]string, n)
		L := make([]int, n)
		R := make([]int, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			bs := scanner.Bytes()
			var pos int
			for bs[pos] != ' ' {
				pos++
			}
			C[j] = string(bs[:pos])
			pos = readInt(bs, pos+1, &L[j])
			readInt(bs, pos+1, &R[j])
		}
		res := solve(n, C, L, R)
		if res < 0 {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", i)
		} else {
			fmt.Printf("Case #%d: %d\n", i, res)
		}
	}
}

func solve1(n int, C []string, L []int, R []int) int {
	os := make(OfferSlice, n)
	for i := 0; i < n; i++ {
		os[i] = Offer{C[i], L[i], R[i]}
	}
	sort.Sort(os)

	colorMap := make(map[string]bool)
	colorArr := make([]string, 0, n)

	for i := 0; i < n; i++ {
		if colorMap[C[i]] {
			continue
		}
		colorArr = append(colorArr, C[i])
		colorMap[C[i]] = true
	}

	find := func(a, b, c string) int {
		arr := make(OfferSlice, 0, n)
		for i := 0; i < n; i++ {
			if a == os[i].color || b == os[i].color || c == os[i].color {
				arr = append(arr, os[i])
			}
		}

		var cur int
		var cnt int

		var i int

		for i < len(arr) {
			if cur == 10000 {
				return cnt
			}
			most := cur
			for i < len(arr) && arr[i].left <= cur+1 {
				if arr[i].right > most {
					most = arr[i].right
				}
				i++
			}

			if most == cur {
				// no extension
				return -1
			}
			cur = most
			cnt++
		}

		if cur < 10000 {
			return -1
		}
		return cnt
	}

	findAndUpdateRes := func(a, b, c string, res *int) {
		tmp := find(a, b, c)
		if tmp > 0 && (*res < 0 || *res > tmp) {
			*res = tmp
		}
	}

	best := -1
	for i := 0; i < len(colorArr); i++ {
		for j := i; j < len(colorArr); j++ {
			for k := j; k < len(colorArr); k++ {
				findAndUpdateRes(colorArr[i], colorArr[j], colorArr[k], &best)
			}
		}
	}
	return best
}

func solve(n int, C []string, L []int, R []int) int {
	os := make(OfferSlice, n)
	for i := 0; i < n; i++ {
		os[i] = Offer{C[i], L[i], R[i]}
	}
	sort.Sort(os)

	best := -1

	var dfs func(i int, cur int, cnt int)
	used := make(map[string]bool)
	dfs = func(i int, cur int, cnt int) {
		if best > 0 && cnt > best {
			return
		}

		if cur == 10000 {
			if best < 0 || best > cnt {
				best = cnt
			}
			return
		}
		if i == n {
			return
		}

		for j := i; j < n && os[j].left <= cur+1; j++ {
			if os[j].right < cur {
				continue
			}
			if used[os[j].color] {
				dfs(j+1, os[j].right, cnt+1)
			} else if len(used) < 3 {
				used[os[j].color] = true
				dfs(j+1, os[j].right, cnt+1)
				delete(used, os[j].color)
			}
		}
	}
	dfs(0, 0, 0)

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Offer struct {
	color       string
	left, right int
}

type OfferSlice []Offer

func (this OfferSlice) Len() int {
	return len(this)
}

func (this OfferSlice) Less(i, j int) bool {
	return this[i].left < this[j].left || (this[i].left == this[j].left && this[i].right > this[j].right)
}

func (this OfferSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

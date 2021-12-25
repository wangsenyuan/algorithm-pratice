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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		var n int
		scanner.Scan()
		readInt(scanner.Bytes(), 0, &n)
		ls := make([]int, n)
		scanner.Scan()
		bs := scanner.Bytes()
		for i, pos := 0, -1; i < n; i++ {
			pos = readInt(bs, pos+1, &ls[i])
		}

		var k int
		scanner.Scan()
		readInt(scanner.Bytes(), 0, &k)
		bands := make([][]int, k)
		for i := 0; i < k; i++ {
			var a, b int
			scanner.Scan()
			bs = scanner.Bytes()
			pos := readInt(bs, 0, &a)
			readInt(bs, pos+1, &b)
			bands[i] = []int{a, b}
		}
		res := solve(n, ls, k, bands)
		fmt.Println(res)
	}
}

func solve(n int, ls []int, k int, bands [][]int) int {
	sort.Ints(ls)

	for i := 0; i < n; i++ {
		ls[i] *= 7
	}

	for i := 0; i < k; i++ {
		bands[i][0] *= 11
		bands[i][1] *= 11
	}
	var res int
	used := make([]bool, k)

	for i := 0; i < n; i++ {
		x, right := -1, 0
		for j := 0; j < k; j++ {
			if used[j] {
				continue
			}
			if bands[j][0] <= ls[i] && ls[i] <= bands[j][1] {
				if x == -1 || bands[j][1] < right {
					x, right = j, bands[j][1]
				}
			}
		}
		if x != -1 {
			res++
			used[x] = true
		}
	}

	return res
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from

	for i < len(bytes) && bytes[i] == ' ' {
		i++
	}

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

	for {
		var n int
		scanner.Scan()
		readInt(scanner.Bytes(), 0, &n)
		if n == 0 {
			break
		}

		coords := make([][]int, n)
		for i := 0; i < n; i++ {
			var a, b int
			scanner.Scan()
			bs := scanner.Bytes()
			readInt(bs, readInt(bs, 0, &a), &b)
			coords[i] = []int{a, b}
		}
		res := solve(n, coords)
		fmt.Println(res)
	}
}

func solve(n int, coords [][]int) int64 {
	sort.Sort(BF(coords))

	segs := make([][]int, 0, len(coords))

	for i := 1; i < len(coords); i++ {
		a, b := coords[i-1], coords[i]
		if a[0] == b[0] {
			// vertical
			segs = append(segs, []int{a[1], b[1] - a[1]})
		}
	}

	if len(segs) == 0 {
		return 0
	}

	sort.Sort(BF(segs))
	var ans int64
	cnt := int64(1)
	for i := 1; i <= len(segs); i++ {
		if i < len(segs) && segs[i][0] == segs[i-1][0] && segs[i][1] == segs[i-1][1] {
			cnt++
		} else {
			ans += cnt * (cnt - 1) / 2
			cnt = 1
		}
	}

	return ans
}

type BF [][]int

func (this BF) Len() int {
	return len(this)
}

func (this BF) Less(i, j int) bool {
	return this[i][0] < this[j][0] || this[i][0] == this[j][0] && this[i][1] < this[j][1]
}

func (this BF) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

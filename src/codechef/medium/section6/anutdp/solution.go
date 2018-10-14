package main

import "sort"

func main() {

}

var parent map[Pair]int
var valid []int

const N = 100000000

func init() {
	valid = make([]int, 0, 100000)
	que := make([]int, 0, 100000)
	que = append(que, 0)
	que = append(que, 1)
	var pos int
	for pos+2 <= len(que) {
		a, b := que[pos], que[pos+1]
		pos += 2

		if b == 0 {
			valid = append(valid, a)
		}
		c, d := b, a+b
		if _, found := parent[Pair{c, d}]; !found && d <= N {
			parent[Pair{c, d}] = a
			que = append(que, c)
			que = append(que, d)
		}
		d = 0
		if _, found := parent[Pair{c, d}]; !found {
			parent[Pair{c, d}] = a
			que = append(que, c)
			que = append(que, d)
		}
	}
	sort.Ints(valid)
}

type Pair struct {
	first, second int
}

func solve(L, R, N int) (bool, string) {

}

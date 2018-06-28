package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, T int
	fmt.Scanf("%d %d", &N, &T)

	input := make([][]int, N)
	for i := 0; i < N; i++ {
		input[i] = make([]int, 3)
		fmt.Scanf("%d %d %d", &input[i][0], &input[i][1], &input[i][2])
	}

	fmt.Println(solve(N, T, input))
}

func solve(N int, T int, input [][]int) int {
	items := make(Items, N)

	for i := 0; i < N; i++ {
		items[i] = Item{input[i][0] - 1, input[i][1], input[i][2]}
	}
	sort.Sort(items)

	FULL := 1<<6 - 1
	var flag int

	ts := make([]int, 6)

	for i := 0; i < 6; i++ {
		ts[i] = -1
	}
	var total int
	for i := 0; i < N; i++ {
		it := items[i]
		k, c, q := it.k, it.c, it.q
		if ts[k] < 0 {
			ts[k] = c
			total += c
		} else if ts[k] > c {
			total -= ts[k]
			ts[k] = c
			total += c
		}

		flag |= 1 << uint(k)

		if flag == FULL && total <= T {
			return q
		}
	}
	return 0
}

type Item struct {
	k, c, q int
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	return items[i].q > items[j].q
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

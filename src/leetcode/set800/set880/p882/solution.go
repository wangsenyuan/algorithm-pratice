package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)

type Solution struct {
	total int
	rects []Rect
}

func Constructor(rects [][]int) Solution {
	var total int

	n := len(rects)

	rs := make([]Rect, n)

	area := func(rect []int) int {
		return (rect[2] - rect[0] + 1) * (rect[3] - rect[1] + 1)
	}

	for i := 0; i < n; i++ {
		tmp := area(rects[i])
		// current rect cover [0, tmp)
		rs[i] = Rect{rects[i][0], rects[i][1], rects[i][2], rects[i][3], total + tmp}
		total += tmp
	}

	return Solution{total, rs}
}

func (this *Solution) Pick() []int {
	cand := rand.Intn(this.total) + 1

	i := sort.Search(len(this.rects), func(i int) bool {
		return this.rects[i].end >= cand
	})

	rect := this.rects[i]

	x := rand.Intn(rect.c - rect.a + 1)

	y := rand.Intn(rect.d - rect.b + 1)
	return []int{rect.a + x, rect.b + y}
}

type Rect struct {
	a, b, c, d int
	end        int
}

func main() {
	rects := [][]int{{1, 1, 5, 5}}
	solution := Constructor(rects)
	cnt := make([][]int, 6)
	for i := 0; i < 6; i++ {
		cnt[i] = make([]int, 6)
	}
	for i := 0; i < 25; i++ {
		res := solution.Pick()
		if res != nil {
			x, y := res[0], res[1]
			if x < 1 || x > 5 || y < 1 || y > 5 {
				fmt.Fprintf(os.Stderr, "%d %d out of range\n", x, y)
			} else {
				cnt[x][y]++
			}
		}
	}
	fmt.Println(cnt)
}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)

type Solution struct {
	ps    []Pair
	total int
}

func Constructor(w []int) Solution {
	// v := make([]int, len(w))
	// copy(v, w)
	// sort.Ints(v)
	var start int
	n := len(w)
	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{start + w[i] - 1, i}
		start += w[i]
	}

	return Solution{ps, start}
}

type Pair struct {
	first, second int
}

func (this *Solution) PickIndex() int {
	cand := rand.Intn(this.total)
	ps := this.ps
	fmt.Fprintf(os.Stderr, "cand %d\n", cand)
	i := sort.Search(len(ps), func(i int) bool {
		return ps[i].first >= cand
	})
	return ps[i].second
}

func main() {
	solution := Constructor([]int{1, 2, 3})
	cnt := make([]int, 3)
	for i := 0; i < 100; i++ {
		cnt[solution.PickIndex()]++
	}
	// fmt.Println(solution.PickIndex())
	fmt.Println(cnt)
}

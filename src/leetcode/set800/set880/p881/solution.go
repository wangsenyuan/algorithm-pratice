package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	rows, cols int
	zeros      int
	mp         map[int]int
}

func Constructor(n_rows int, n_cols int) Solution {
	zeros := n_rows * n_cols
	mp := make(map[int]int)
	return Solution{n_rows, n_cols, zeros, mp}
}

func (this *Solution) Flip() []int {
	if this.zeros == 0 {
		return nil
	}
	cand := rand.Intn(this.zeros)
	this.zeros--
	res := cand
	if val, found := this.mp[cand]; found {
		res = val
	}
	if val, found := this.mp[this.zeros]; found {
		this.mp[cand] = val
	} else {
		this.mp[cand] = this.zeros
	}
	a, b := res/this.cols, res%this.cols
	return []int{a, b}
}

func (this *Solution) Reset() {
	this.zeros = this.rows * this.cols
	this.mp = make(map[int]int)
}

func main() {
	solution := Constructor(5, 5)
	cnt := make([][]int, 5)
	for i := 0; i < 5; i++ {
		cnt[i] = make([]int, 5)
	}

	for i := 0; i < 30; i++ {
		res := solution.Flip()
		if res != nil {
			x, y := res[0], res[1]
			cnt[x][y]++
		}
	}
	fmt.Println(cnt)
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, p int
	var k, pos int
	fmt.Scanf("%d %d %d", &n, &k, &p)

	frogs := make([]*Frog, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &pos)
		frogs[i] = &Frog{pos, i}
	}

	sort.Sort(Frogs(frogs))

	f := make([]int, n)
	idx := make([]int, n)

	for i := 0; i < n; i++ {
		if i > 0 && frogs[i].pos <= frogs[i - 1].pos+k {
			f[i] = f[i - 1]
		} else {
			f[i] = i
		}
		idx[frogs[i].idx] = i
	}

	var a, b int
	for i := 0; i < p; i++ {
		fmt.Scanf("%d %d", &a, &b)
		a--
		b--
		x, y := idx[a], idx[b]

		if f[y] == f[x] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

type Frog struct {
	pos int
	idx int
}

type Frogs []*Frog

func (f Frogs) Len() int {
	return len(f)
}

func (f Frogs) Less(i, j int) bool {
	return f[i].pos < f[j].pos
}

func (f Frogs) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

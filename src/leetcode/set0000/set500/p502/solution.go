package main

import (
	"sort"
	"fmt"
	"container/heap"
)

func main() {
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}
	fmt.Println(findMaximizedCapital(2, 0, profits, capital))
}

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	n := len(Profits)

	projects := make([]*Project, n)
	for i := 0; i < n; i++ {
		projects[i] = &Project{Profits[i], Capital[i]}
	}

	sort.Sort(Projects(projects))

	h := &IntHeap{}
	heap.Init(h)

	j := 0
	for i := 0; i < k; i++ {
		for j < n && projects[j].capital <= W {
			heap.Push(h, projects[j].profit)
			j++
		}
		if h.Len() == 0 {
			break
		}
		W += heap.Pop(h).(int)
	}
	return W
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0: n - 1]
	return x
}

type Project struct {
	profit  int
	capital int
}
type Projects []*Project

func (ps Projects) Len() int {
	return len(ps)
}

func (ps Projects) Less(i, j int) bool {
	return ps[i].capital < ps[j].capital
}

func (ps Projects) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

/*

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	n := len(Profits)

	projects := make([]*Project, n)
	for i := 0; i < n; i++ {
		projects[i] = &Project{Profits[i], Capital[i]}
	}

	sort.Sort(Projects(projects))

	pickProject := func(projects []*Project, w int) ([]*Project, int) {
		i := sort.Search(len(projects), func(i int) bool {
			return projects[i].capital > w
		});
		x, mxProfit := 0, 0
		for j := 0; j < i; j++ {
			if projects[j].profit > mxProfit {
				x = j
				mxProfit = projects[j].profit
			}
		}
		return dropAt(projects, x), mxProfit
	}

	for k > 0 && len(projects) > 0 {
		left, p := pickProject(projects, W)
		projects = left
		W += p
		k--
	}
	return W
}

func dropAt(project []*Project, i int) []*Project {
	dst := make([]*Project, len(project) - 1)
	copy(dst, project[:i])
	copy(dst[i:], project[i + 1:])
	return dst
}


*/

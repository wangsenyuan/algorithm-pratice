package p436

import "sort"

/**
 * Definition for an interval.
 */
type Interval struct {
	Start int
	End   int
}

func findRightInterval(intervals []Interval) []int {
	nodes := make([]*Node, len(intervals))

	for i, x := range intervals {
		nodes[i] = &Node{x.Start, x.End, i}
	}
	sort.Sort(Nodes(nodes))
	ans := make([]int, len(intervals))
	for i, v := range nodes {
		ans[v.index] = -1

		if i < len(nodes)-1 && v.end <= nodes[len(nodes)-1].start {
			a, b := 0, len(nodes)-1
			for a <= b {
				c := a + (b-a)/2
				if nodes[c].start < v.end {
					a = c + 1
				} else {
					b = c - 1
				}
			}
			ans[v.index] = nodes[a].index
		}
	}

	return ans
}

type Node struct {
	start, end, index int
}

/**
* sort interval by end
 */
type Nodes []*Node

//length of the array
func (a Nodes) Len() int {
	return len(a)
}

func (a Nodes) Less(i, j int) bool {
	return a[i].start < a[j].start
}

func (a Nodes) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

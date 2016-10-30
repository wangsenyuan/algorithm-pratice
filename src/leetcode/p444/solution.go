package main

import "fmt"

/**
[5,3,2,4,1]
[[5,3,2,4],[4,1],[1],[3],[2,4]
[1,1000000000]]
**/
func main() {
	org := []int{1, 2}
	seqs := [][]int{{1, 2}, {2, 1}}
	fmt.Println(sequenceReconstruction(org, seqs))
}

func sequenceReconstruction(org []int, seqs [][]int) bool {
	n := len(org)
	degree := make(map[int]int)
	graph := make(map[int][]int, n+1)
	words := make(map[int]bool)
	for _, seq := range seqs {
		for i := 0; i < len(seq); i++ {
			y := seq[i]
			words[y] = true
			if _, has := degree[y]; !has {
				degree[y] = 0
			}
			if i == 0 {
				continue
			}

			x := seq[i-1]
			graph[x] = append(graph[x], y)
			degree[y]++
		}
	}
	if len(words) != n {
		return false
	}

	var visit func(v int) bool
	visit = func(v int) bool {
		if !words[v] || degree[v] > 0 {
			return false
		}
		delete(words, v)
		for _, w := range graph[v] {
			degree[w]--
		}
		for _, w := range graph[v] {
			if degree[w] > 0 {
				continue
			}
			return visit(w)
		}
		return true
	}

	visit(org[0])
	return len(words) == 0
}

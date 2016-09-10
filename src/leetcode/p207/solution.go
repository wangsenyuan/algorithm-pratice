package main

import "fmt"

func main() {
	prerequisites := [][]int{[]int{1, 0}, []int{1, 0}}
	fmt.Println(canFinish(2, prerequisites))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)

	for _, pre := range prerequisites {
		i, j := pre[0], pre[1]
		if len(graph[i]) == 0 {
			graph[i] = []int{j}
		} else {
			graph[i] = append(graph[i], j)
		}
	}

	visited := make([]int, numCourses)

	var visit func(i int) bool
	visit = func(i int) bool {
		visited[i] = 1

		for _, j := range graph[i] {
			if visited[j] == 1 {
				return false
			}
			if visited[j] == 2 {
				continue
			}
			if !visit(j) {
				return false
			}
		}
		visited[i] = 2
		return true
	}

	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 && !visit(i) {
			return false
		}
	}
	return true
}

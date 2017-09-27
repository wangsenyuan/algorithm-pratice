package main

import "fmt"

func findRedundantDirectedConnection(edges [][]int) []int {
	var n int

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if a > n {
			n = a
		}
		if b > n {
			n = b
		}
	}

	parent := make([]int, n+1)

	var findRoot func(v int, roots []int, visited []bool) int

	findRoot = func(v int, roots []int, visited []bool) int {
		var w int
		if parent[v] == 0 {
			w = v
		} else if visited[v] {
			w = roots[v]
		} else {
			// to detect loop
			visited[v] = true
			w = findRoot(parent[v], roots, visited)
		}

		roots[v] = w
		return w
	}

	shouldRemove := func(x, y int) bool {
		for _, edge := range edges {
			a, b := edge[0], edge[1]
			if a == x && b == y {
				continue
			}
			parent[b] = a
		}
		root := 0
		for i := 1; i <= n; i++ {
			if parent[i] == 0 {
				if root != 0 {
					// multiple roots
					return false
				}

				root = i
			}
		}

		roots := make([]int, n+1)
		visited := make([]bool, n+1)
		for i := 1; i <= n; i++ {
			if !visited[i] {
				tmp := findRoot(i, roots, visited)
				if tmp != root {
					return false
				}
			}
		}
		return true
	}

	for _, edge := range edges {
		a, b := edge[0], edge[1]

		pb := parent[b]
		if pb == 0 {
			parent[b] = a
		} else if pb != a {
			if shouldRemove(a, b) {
				return edge
			}
			return []int{pb, b}
		}
	}

	//find loop
	slow, fast := 1, 1

	for slow != fast {
		slow = parent[slow]
		if slow == 0 {
			break
		}

		fast = parent[fast]
		if fast == 0 {
			break
		}
		fast = parent[fast]
		if fast == 0 {
			break
		}
	}

	if slow == 0 || fast == 0 {
		return nil
	}

	loop := make([]bool, n+1)

	for i := slow; loop[i] == false; i = parent[i] {
		loop[i] = true
	}

	for i := len(edges) - 1; i >= 0; i-- {
		a, b := edges[i][0], edges[i][1]
		if loop[a] && loop[b] {
			return edges[i]
		}
	}

	return nil
}

func main() {
	test1 := [][]int{{1, 2}, {1, 3}, {2, 3}}
	fmt.Println(findRedundantDirectedConnection(test1))

	test2 := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}
	fmt.Println(findRedundantDirectedConnection(test2))

	test3 := [][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}}
	fmt.Println(findRedundantDirectedConnection(test3))

	test4 := [][]int{{5, 2}, {5, 1}, {3, 1}, {3, 4}, {3, 5}}
	fmt.Println(findRedundantDirectedConnection(test4))

	test5 := [][]int{{3, 4}, {4, 1}, {1, 2}, {2, 3}, {5, 1}}
	fmt.Println(findRedundantDirectedConnection(test5))
}

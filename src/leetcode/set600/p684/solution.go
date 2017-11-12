package main

import "fmt"

func findRedundantConnection(edges [][]int) []int {
	root := make(map[int]int)

	var findRoot func(x int) int

	findRoot = func(x int) int {
		if y, found := root[x]; found {
			if x == y {
				return x
			}
			r := findRoot(y)
			root[x] = r
			return r
		} else {
			root[x] = x
			return x
		}
	}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		ra := findRoot(a)
		rb := findRoot(b)
		if ra == rb {
			return edge
		}
		root[ra] = rb
	}
	return nil
}

func main() {
	test1 := [][]int{{1, 2}, {1, 3}, {2, 3}}
	fmt.Println(findRedundantConnection(test1))
	test2 := [][]int{{1, 2}, {1, 3}, {3, 1}}
	fmt.Println(findRedundantConnection(test2))
	test3 := [][]int{{2, 3}, {5, 2}, {1, 5}, {4, 2}, {4, 1}}
	fmt.Println(findRedundantConnection(test3))

}

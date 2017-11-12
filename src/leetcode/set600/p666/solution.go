package main

import "fmt"

func pathSum(nums []int) int {
	tree := make([][]int, 5)
	for i := 0; i < 5; i++ {
		tree[i] = make([]int, 16)
		for j := 0; j < 16; j++ {
			tree[i][j] = -1
		}
	}

	putIntoTree := func(num int) {
		v := num % 10
		p := (num / 10) % 10
		d := (num / 100) % 10
		tree[d-1][p-1] = v
	}

	for _, num := range nums {
		putIntoTree(num)
	}

	var sum func(i, j, r int)
	var res int
	sum = func(i, j, r int) {
		r += tree[i][j]
		if i == 4 || (tree[i+1][2*j] < 0 && tree[i+1][2*j+1] < 0) {
			res += r
			return
		}

		if tree[i+1][2*j] >= 0 {
			sum(i+1, 2*j, r)
		}

		if tree[i+1][2*j+1] >= 0 {
			sum(i+1, 2*j+1, r)
		}
	}

	sum(0, 0, 0)

	return res
}

func main() {
	//fmt.Println(pathSum([]int{113, 215, 221}))
	//fmt.Println(pathSum([]int{113, 221}))
	//fmt.Println(pathSum([]int{113, 221, 333, 345}))
	//fmt.Println(pathSum([]int{111, 217, 221, 315, 415}))
	fmt.Println(pathSum([]int{115, 215, 224, 325, 336, 446, 458}))
}

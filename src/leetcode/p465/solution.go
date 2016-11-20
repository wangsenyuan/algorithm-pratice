package main

import "fmt"

func main() {
	transactions := [][]int{[]int{0, 1, 10}, []int{2, 0, 5}}
	fmt.Println(minTransfers(transactions))
}

func minTransfers(transactions [][]int) int {
	tx := make(map[int]map[int]int)
	for _, transaction := range transactions {
		x, y, z := transaction[0], transaction[1], transaction[2]
		if tx[x] == nil {
			tx[x] = make(map[int]int)
		}
		tx[x][y] = z
	}

	res := make(map[int]map[int]int)
	processed := make(map[int]bool)

	for x := range tx {
		if !processed[x] {
			transfer(tx, x, processed, res)
		}
	}

	ans := 0
	for x := range res {
		for _, z := range res[x] {
			if z > 0 {
				ans++
			}
		}
	}
	return ans
}

func transfer(tx map[int]map[int]int, x int, processed map[int]bool, res map[int]map[int]int) {
	if processed[x] {
		res = processLoop(res, x)
	}
	processed[x] = true

	for y, z := range tx[x] {
		if processed[y] {
			continue
		}
		res = processLine(res, x, y, z)
		transfer(tx, x, processed, res)
	}
}

func processLine(res map[int]map[int]int, x int, y int, z int) map[int]map[int]int {
	xs := res[x]
	if res[y] == nil {
		res[y] = make(map[int]int)
	}
	for a, b := range xs {
		if b == z {
			xs[a] = 0
			res[y][a] = z
		} else if b > z {
			xs[a] = b - z
			res[y][a] = z
		} else {
			xs[a] = b
			res[y][x] = z - b
		}
	}
	return res
}

func processLoop(res map[int]map[int]int, x int) map[int]map[int]int {
	xs := res[x]

	for a, b := range xs {
		c := res[a][x]
		if c == 0 {
			continue
		}

		if c > b {
			res[a][x] = c - b
			res[x][a] = 0
		} else if c < b {
			res[a][x] = 0
			res[x][a] = b - c
		} else {
			res[a][x] = 0
			res[x][a] = 0
		}
	}
	return res
}

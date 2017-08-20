package main

import (
	"sort"
	"fmt"
)

func Part(n int) string {
	var enum func(prod, left, start int)

	products := make([]int, 0, n)

	enum = func(prod, left, start int) {
		//fmt.Printf("enum %d %d\n", prod, left)
		products = append(products, prod*left)
		for i := start; 2*i <= left; i++ {
			enum(prod*i, left-i, i)
		}
	}

	enum(1, n, 1)

	sort.Ints(products)
	m := len(products)

	sum := 0

	j := 0
	i := 0
	for i < m {
		sum += products[j]
		for i < m && products[i] == products[j] {
			i++
		}
		if i < m {
			j++
			products[j] = products[i]
		}
	}

	m = j + 1

	//fmt.Println(products[:m])

	min, max := products[0], products[m-1]
	rg := max - min

	median := 0.0
	if m%2 == 1 {
		median = float64(products[m/2])
	} else {
		a, b := products[m/2], products[m/2-1]
		median = 0.5 * (float64(a) + float64(b))
	}

	avg := float64(sum) / float64(m)

	return fmt.Sprintf("Range: %d Average: %.2f Median: %.2f", rg, avg, median)
}

func main() {
	fmt.Println(Part(1))
	fmt.Println(Part(2))
	fmt.Println(Part(3))
	fmt.Println(Part(4))
	fmt.Println(Part(5))
	fmt.Println(Part(6))
	fmt.Println(Part(7))
	fmt.Println(Part(8))


}

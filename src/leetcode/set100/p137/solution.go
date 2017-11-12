package main

import "fmt"

func singleNumber(nums []int) int {
	xs := make([]int, 32)

	for _, num := range nums {
		for i := 0; i < 32; i++ {
			if (num & (1 << uint(i))) > 0 {
				xs[i]++
			}
		}
	}
	var res int

	for i := 0; i < 32; i++ {
		if xs[i]%3 == 0 {
			continue
		}
		res |= (1 << uint(i))
	}

	return int(int32(res))
}

func main() {
	nums := []int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}
	fmt.Println(singleNumber(nums))
	//fmt.Println(math.MaxInt32)
}

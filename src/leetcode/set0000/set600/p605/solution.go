package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))

}

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		if flowerbed[i] == 1 {
			continue
		}

		if i > 0 && flowerbed[i-1] == 1 {
			continue
		}

		if i < len(flowerbed)-1 && flowerbed[i+1] == 1 {
			continue
		}

		n--
		flowerbed[i] = 1
	}

	return n == 0
}

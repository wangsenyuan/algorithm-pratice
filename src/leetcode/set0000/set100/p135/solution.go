package main

import "fmt"

func main() {
	ratings := []int{4, 2, 3, 4, 1}
	fmt.Println(candy(ratings))
}

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	total, prev, countDown := 1, 1, 0

	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] {
			if countDown > 0 {
				total += countDown * (countDown + 1) / 2
				if countDown >= prev {
					total += countDown - prev + 1
				}
				countDown = 0
				prev = 1
			}
			if ratings[i] == ratings[i-1] {
				prev = 1
			} else {
				prev++
			}
			total += prev
		} else {
			countDown++
		}
	}
	if countDown > 0 {
		total += countDown * (countDown + 1) / 2
		if countDown >= prev {
			total += countDown - prev + 1
		}
	}
	return total
}

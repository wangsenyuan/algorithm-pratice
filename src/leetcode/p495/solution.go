package main

import "fmt"

func main() {
	//fmt.Println(findPosisonedDuration([]int{1, 4}, 2))
	fmt.Println(findPosisonedDuration([]int{1, 2}, 2))

}

func findPosisonedDuration(timeSeries []int, duration int) int {

	res := 0
	pt := 0
	for _, t := range timeSeries {
		if t >= pt {
			res += duration
		} else {
			res += t + duration - pt
		}
		pt = t + duration

	}

	return res
}

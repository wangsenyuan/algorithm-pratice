package main

import (
	"fmt"
	"math"
)

func constructRectangle(area int) []int {

	w := int(math.Sqrt(float64(area)))
	if w*w == area {
		return []int{w, w}
	}

	for area%w != 0 {
		w--
	}

	return []int{area / w, w}
}

func main() {
	fmt.Println(constructRectangle(4))
	fmt.Println(constructRectangle(5))
	fmt.Println(constructRectangle(6))
	fmt.Println(constructRectangle(7))
	fmt.Println(constructRectangle(8))
	fmt.Println(constructRectangle(9))
	fmt.Println(constructRectangle(10))
	fmt.Println(constructRectangle(10000000))

}

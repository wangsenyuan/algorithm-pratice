package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%d\n", reverse(-321))
}

func reverse(x int) int {
	var sign int64 = 1
	var _x = int64(x)
	if x < 0 {
		sign = -1
		_x = sign * _x
	}
	var y int64
	for _x > 0 {
		z := _x % 10
		y = y*10 + z
		_x = _x / 10
	}

	y = sign * y
	if y > math.MaxInt32 {
		y = 0
	} else if y < math.MinInt32 {
		y = 0
	}

	return int(y)
}

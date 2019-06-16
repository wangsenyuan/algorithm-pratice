package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {

	var x = 0

	for i := 0; i < len(chars); {
		j := i + 1
		for j < len(chars) && chars[j] == chars[i] {
			j++
		}
		cnt := j - i
		if cnt == 1 {
			chars[x] = byte(chars[i])
			x++
			i++
		} else {
			chars[x] = byte(chars[i])
			x++
			cntStr := strconv.Itoa(cnt)
			y := len(cntStr)
			copy(chars[x:x+y], cntStr)
			x += y
			i = j
		}
	}

	return x
}

func main() {
	//a := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	//a := []byte{'a'}
	//a := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	a := []byte{'G', 't', 'W', 'Y', 'v', '&', ':', 'd', '#', 'k'}
	l := compress(a)
	a = a[:l]
	fmt.Println(a)
}

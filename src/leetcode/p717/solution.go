package main

import "fmt"

func isOneBitCharacter(bits []int) bool {
	i := 0

	for i < len(bits)-1 {
		if bits[i] == 1 {
			i += 2
		} else {
			i += 1
		}
	}

	return i == len(bits)-1
}

func main() {
	fmt.Println(isOneBitCharacter([]int{1, 0, 0}))
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 0}))

}

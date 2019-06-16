package main

import "fmt"

func main() {
	data := []int{197, 130, 1}
	//data := []int{235, 140, 4}
	fmt.Println(validUtf8(data))
}

func validUtf8(data []int) bool {
	if len(data) == 0 {
		return true
	}

	first := data[0] & 0x000000ff

	if (first>>7)&1 == 0 {
		return validUtf8(data[1:])
	}

	if isTwoLen(first, data) {
		return validUtf8(data[2:])
	}

	if isThreeLen(first, data) {
		return validUtf8(data[3:])
	}

	if isFourLen(first, data) {
		return validUtf8(data[4:])
	}
	return false
}

func start10(x int) bool {
	return ((x >> 6) & 3) == 2
}
func isTwoLen(first int, data []int) bool {
	if (first>>5)&7 != 6 {
		return false
	}

	return len(data) > 1 && start10(data[1])
}

func isThreeLen(first int, data []int) bool {
	if (first>>4)&15 != 14 {
		return false
	}
	return len(data) > 2 && start10(data[1]) && start10(data[2])
}

func isFourLen(first int, data []int) bool {
	if (first>>3)&31 != 30 {
		return false
	}
	return len(data) > 3 && start10(data[1]) && start10(data[2]) && start10(data[3])
}

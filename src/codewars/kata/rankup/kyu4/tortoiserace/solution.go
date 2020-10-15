package main

import "fmt"

func Race(v1, v2, g int) [3]int {
	if g == 0 {
		return [...]int{0, 0, 0}
	}

	if v2 <= v1 {
		return [...]int{-1, -1, -1}
	}

	seconds := g * 60 * 60 / (v2 - v1)

	minutes := seconds / 60
	seconds = seconds % 60
	hours := minutes / 60
	minutes = minutes % 60

	return [...]int{hours, minutes, seconds}
}

func main() {
	fmt.Println(Race(720, 850, 70))
	fmt.Println(Race(820, 81, 550))
	fmt.Println(Race(80, 91, 37))
}

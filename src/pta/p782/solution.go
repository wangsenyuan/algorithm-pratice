package main

import "fmt"

func main() {
	start, diff := 0, 0

	fmt.Scanf("%d %d", &start, &diff)

	if diff < 0 {
		diff += 1440
	}

	diff = diff % 1440

	hour := start / 100
	mins := start % 100

	dh := diff / 60
	dm := diff % 60

	hour += dh
	mins += dm
	if mins >= 60 {
		mins -= 60
		hour++
	}
	hour %= 24

	if hour > 0 {
		fmt.Printf("%d", hour)
	}
	fmt.Printf("%02d\n", mins)
}

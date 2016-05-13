package main

import "fmt"

func main() {
	var years, workHours int
	fmt.Scanf("%d %d", &years, &workHours)

	overHours := max(workHours-40, 0)
	workHours = workHours - overHours
	salary := 30
	if years >= 5 {
		salary = 50
	}

	var total float64 = 1.5*float64(salary)*float64(overHours) + float64(salary)*float64(workHours)
	fmt.Printf("%.2f\n", total)
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

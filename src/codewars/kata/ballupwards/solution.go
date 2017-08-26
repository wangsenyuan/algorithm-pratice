package main

import "fmt"

func MaxBall(v0 int) int {
	// your code
	v := float64(v0) * 1000 / 3600

	g := 9.81

	t0 := v/g + 0.05

	return int(t0 * 10)
}

func main() {
	fmt.Println(MaxBall(15))
	fmt.Println(MaxBall(25))
	fmt.Println(MaxBall(37))
	fmt.Println(MaxBall(45))

	fmt.Println(MaxBall(99))
	fmt.Println(MaxBall(85))

	fmt.Println(MaxBall(136))
}

/*
testequal(37, 10)
        testequal(45, 13)
        testequal(99, 28)
        testequal(85, 24)
        testequal(136, 39)
 */

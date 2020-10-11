package main

import "fmt"

func Finance(n int) int {

	var ans int
	for i := 1; i <= n; i++ {
		ans += 3 * i * (i + 1) / 2
	}

	return ans
}

func main() {
	fmt.Println(Finance(5))
	fmt.Println(Finance(6))
	fmt.Println(Finance(7))
	fmt.Println(Finance(5000))

}

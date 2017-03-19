package main

import "fmt"

func main() {
	fmt.Println(findContestMatch(2))
	fmt.Println(findContestMatch(4))
	fmt.Println(findContestMatch(8))
	fmt.Println(findContestMatch(16))
	fmt.Println(findContestMatch(32))

}

func findContestMatch(n int) string {

	var match func(teams []int, help []int) string

	match = func(teams []int, help []int) string {
		if len(teams) == 2 {
			return fmt.Sprintf("(%d,%d)", teams[0], teams[1])
		}

		n := len(teams)

		i, a, b := 0, 0, n/2

		for i < n {
			help[a] = teams[i]
			a++
			i++
			help[b] = teams[i]
			b++
			i++
			help[b] = teams[i]
			b++
			i++
			help[a] = teams[i]
			a++
			i++
		}
		return "(" + match(help[:n/2], teams[:n/2]) + "," + match(help[n/2:], teams[n/2:]) + ")"
	}
	teams := make([]int, n)
	for i := 0; i < n; i++ {
		teams[i] = i + 1
	}

	return match(teams, make([]int, n))
}

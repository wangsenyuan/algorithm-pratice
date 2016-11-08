package main

import "fmt"

func main() {
	fmt.Println(poorPigs(1000, 15, 60))
}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	if buckets <= 1 {
		return 0
	}

	if minutesToDie >= minutesToTest {
		return -1
	}

	base := minutesToTest/minutesToDie + 1
	cnt := 0
	for buckets > 0 {
		buckets /= base
		cnt++
	}
	return cnt
}

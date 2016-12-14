package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d\n", &t)

	for t > 0 {
		var k int
		fmt.Scanf("%d\n", &k)

		leaves := make([]int, k)

		for i := 0; i < k; i++ {
			fmt.Scanf("%d", &leaves[i])
		}

		if count(leaves) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}

		t--
	}
}

func count(leaves []int) bool {
	stems := 1
	for i := 0; i < len(leaves); i++ {
		if leaves[i] > stems {
			return false
		}

		stems = 2 * (stems - leaves[i])
		if stems < 0 || (stems == 0 && i < len(leaves)-1) {
			return false
		}
	}

	return stems == 0
}

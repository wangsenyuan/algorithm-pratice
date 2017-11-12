package main

import "fmt"

func main() {
	fmt.Println(canIWin(10, 11))
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	var play func(used int, player int, sum int) int
	cache := make(map[int]int)

	play = func(used int, player int, sum int) int {
		if win, ok := cache[used]; ok {
			return win
		}
		win := 1 - player
		for i := 1; i <= maxChoosableInteger && win != player; i++ {
			if used&(1<<uint(i)) > 0 {
				continue
			}
			if sum+i >= desiredTotal {
				win = player
			} else {
				win = play(used|(1<<uint(i)), 1-player, sum+i)
			}
		}

		cache[used] = win
		return win
	}

	if desiredTotal == 0 {
		return true
	}

	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}

	return play(0, 0, 0) == 0
}

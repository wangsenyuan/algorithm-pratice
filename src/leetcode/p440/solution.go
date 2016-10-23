package main

import "fmt"
import "strconv"

func main() {
	fmt.Println(findKthNumber(1, 1))
	//fmt.Println(findKthNumber(13, 13))
	//fmt.Println(findKthNumber(957747794, 424238336))
}

func findKthNumber(n int, k int) int {

	var find func(prefix string, k int) int

	find = func(prefix string, k int) int {
		if len(prefix) > 0 && k == 1 {
			y, _ := strconv.Atoi(prefix)
			return y
		}
		i := 1
		cnt := 0
		if len(prefix) > 0 {
			i = 0
			cnt = 1
		}
		for ; i < 10; i++ {
			c := count(prefix, i, n)
			if cnt + c < k {
				cnt += c
				continue
			}
			return find(prefix + strconv.Itoa(i), k - cnt)
		}
		return -1
	}

	return find("", k)
}

func count(prefix string, i int, n int) int {
	cur, _ := strconv.Atoi(prefix + strconv.Itoa(i))
	cnt := 0
	num := 1
	for cur <= n {
		cnt += num
		cur *= 10
		num *= 10
	}
	if n < (cur / 10 + num / 10 - 1) {
		cnt -= (cur / 10 + num / 10 - 1) - n
	}
	return cnt
}

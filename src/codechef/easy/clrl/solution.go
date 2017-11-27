package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		bs := scanner.Bytes()
		var n, reziba int
		x := readInt(bs, 0, &n)
		readInt(bs, x+1, &reziba)
		scores := make([]int, n)
		scanner.Scan()
		bs = scanner.Bytes()
		x = 0
		for i := 0; i < n; i++ {
			x = readInt(bs, x, &scores[i]) + 1
		}
		res := solve(n, reziba, scores)
		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func solve(n int, reziba int, scores []int) bool {
	left, right := 0, 1000000001

	for i := 0; i < n-1; i++ {
		x := scores[i]
		if x <= left || x >= right {
			return false
		}
		if x == reziba {
			return false
		}
		if x < reziba {
			left = x
		} else {
			right = x
		}
	}

	return left < reziba && reziba < right
}

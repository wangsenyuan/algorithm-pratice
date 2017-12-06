package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bs []byte, from int, val *int) int {
	tmp := 0
	i := from
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int(bs[i]-'0')
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
		var x, y int
		scanner.Scan()
		s := scanner.Bytes()
		scanner.Scan()
		bs := scanner.Bytes()
		readInt(bs, readInt(bs, 0, &x)+1, &y)
		res := solve(x, y, s)
		fmt.Println(res)
	}
}

func solve(x, y int, s []byte) string {
	// fmt.Printf("[debug] %d %d %s\n", x, y, string(s))
	a, b, n := 0, 0, len(s)
	for i := 0; i < n; i++ {
		if s[i] == 'a' {
			a++
		} else {
			b++
		}
	}

	var A, B int
	res := make([]byte, 0, n)
	for a+b > 0 {
		giveApple := false
		if B == y {
			giveApple = true
		} else if A == x {
			giveApple = false
		} else if a == 0 {
			giveApple = false
		} else if b == 0 {
			giveApple = true
		} else {
			giveApple = f(a-1, b, x, y, A+1, 0) < f(a, b-1, x, y, 0, B+1)
		}
		if giveApple {
			if a > 0 {
				a--
				res = append(res, 'a')
			} else {
				res = append(res, '*')
			}
			A++
			B = 0
		} else {
			if b > 0 {
				b--
				res = append(res, 'b')
			} else {
				res = append(res, '*')
			}
			B++
			A = 0
		}
	}
	return string(res)
}

func f(a, b, x, y, A, B int) int {
	return max((b+B+y-1)/y-1, a) + max((a+A+x-1)/x-1, b)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

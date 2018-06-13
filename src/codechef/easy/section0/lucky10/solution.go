package main

import (
	"bufio"
	"fmt"
	"os"
)

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := toInt(scanner.Bytes())

	for i := 0; i < t; i++ {
		scanner.Scan()
		a := scanner.Bytes()
		scanner.Scan()
		b := scanner.Bytes()
		res := solve(a, b)
		fmt.Println(res)
	}
}

func solve(a, b []byte) string {
	n := len(a)
	a0, a4, a5, a7 := 0, 0, 0, 0
	b0, b4, b5, b7 := 0, 0, 0, 0

	for i := 0; i < n; i++ {
		if a[i] < '4' {
			a0++
		} else if a[i] == '4' {
			a4++
		} else if a[i] < '7' {
			a5++
		} else if a[i] == '7' {
			a7++
		}
		if b[i] < '4' {
			b0++
		} else if b[i] == '4' {
			b4++
		} else if b[i] < '7' {
			b5++
		} else if b[i] == '7' {
			b7++
		}
	}

	c7 := 0
	for a7 > 0 {
		if b5 > 0 {
			b5--
			c7++
		} else if b0 > 0 {
			b0--
			c7++
		} else if b4 > 0 {
			b4--
			c7++
		} else if b7 > 0 {
			b7--
			c7++
		}
		a7--
	}

	for b7 > 0 {
		if a5 > 0 {
			a5--
			c7++
		} else if a0 > 0 {
			a0--
			c7++
		} else if a4 > 0 {
			a4--
			c7++
		}
		b7--
	}

	c4 := 0
	for a4 > 0 {
		if b0 > 0 {
			b0--
			a4--
			c4++
		} else if b4 > 0 {
			b4--
			a4--
			c4++
		} else {
			break
		}
	}
	for b4 > 0 {
		if a0 > 0 {
			a0--
			b4--
			c4++
		} else if a4 > 0 {
			a4--
			b4--
			c4++
		} else {
			break
		}
	}

	res := make([]byte, c7+c4)
	i := 0
	for i < c7 {
		res[i] = '7'
		i++
	}
	for i < c7+c4 {
		res[i] = '4'
		i++
	}
	return string(res)
}

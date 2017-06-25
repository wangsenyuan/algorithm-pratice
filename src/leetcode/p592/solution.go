package main

import "fmt"

func main() {
	fmt.Println(fractionAddition("-1/2+1/2"))
	fmt.Println(fractionAddition("-1/2+1/2+1/3"))
	fmt.Println(fractionAddition("1/3-1/2"))
	fmt.Println(fractionAddition("5/3+1/3"))

}

func fractionAddition(expression string) string {
	a, b, left := read(expression)

	for len(left) > 0 {
		c, d, left2 := read(left)

		x := a*d + b*c
		y := b * d
		sign := 1
		if x < 0 {
			sign = -1
			x = -x
		}

		g := gcd(x, y)
		x /= g
		y /= g

		x *= sign
		a, b, left = x, y, left2
	}

	return fmt.Sprintf("%d/%d", a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func read(str string) (int, int, string) {
	sign := 1
	if str[0] == '-' {
		sign = -1
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}

	num := 0
	i := 0
	for ; i < len(str) && str[i] != '/'; i++ {
		num = num*10 + int(str[i]-'0')
	}

	num = sign * num

	i++
	den := 0
	for ; i < len(str) && (str[i] != '+' && str[i] != '-'); i++ {
		den = den*10 + int(str[i]-'0')
	}

	return num, den, str[i:]
}

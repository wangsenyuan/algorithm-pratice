package main

import "fmt"

func main() {
	fmt.Println("137=2(2(2)+2+2(0))+2(2+2(0))+2(0)")
	fmt.Println("1315=2(2(2+2(0))+2)+2(2(2+2(0)))+2(2(2)+2(0))+2+2(0)")
	fmt.Println("73=2(2(2)+2)+2(2+2(0))+2(0)")
	fmt.Println("136=2(2(2)+2+2(0))+2(2+2(0))")
	fmt.Println("255=2(2(2)+2+2(0))+2(2(2)+2)+2(2(2)+2(0))+2(2(2))+2(2+2(0))+2(2)+2+2(0)")
	fmt.Println("1384=2(2(2+2(0))+2)+2(2(2+2(0)))+2(2(2)+2)+2(2(2)+2(0))+2(2+2(0))")
	fmt.Println("16385=2(2(2+2(0))+2(2)+2)+2(0)")
}

func doExpr() {
	fmt.Printf("%d=%s\n", 137, expr(137))
	fmt.Printf("%d=%s\n", 1315, expr(1315))
	fmt.Printf("%d=%s\n", 73, expr(73))
	fmt.Printf("%d=%s\n", 136, expr(136))
	fmt.Printf("%d=%s\n", 255, expr(255))
	fmt.Printf("%d=%s\n", 1384, expr(1384))
	fmt.Printf("%d=%s\n", 16385, expr(16385))
}
func expr(num int) string {
	x, y := 0, 1
	for 2*y <= num {
		x++
		y *= 2
	}

	if y == num {
		return expr1(x)
	}
	return expr1(x) + "+" + expr(num-y)
}

func expr1(x int) string {
	if x == 0 {
		return "2(0)"
	}
	if x == 1 {
		return "2"
	}
	return "2(" + expr(x) + ")"
}

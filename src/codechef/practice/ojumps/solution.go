package main

import "fmt"

func main() {
	a := int64(0)
	fmt.Scanf("%d\n", &a)
	b := a % int64(6)
	if b == 0 || b == 1 || b == 3 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

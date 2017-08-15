package main

import "fmt"

func Fortune(f0 int, p float64, c0 int, n int, i float64) bool {
	//fmt.Printf("%d %.2f %d %d %.2f\n", f0, p, c0, n, i)
	if f0 < 0 {
		return false
	}

	if n == 1 {
		return true
	}

	return Fortune(f0+int(0.01*p*float64(f0))-c0, p, c0+int(float64(c0)*i*0.01), n-1, i)
}

func main() {
	fmt.Println(Fortune(100000, 1.0, 2000, 15, 1.0))
	fmt.Println(Fortune(100000, 1.0, 9185, 12, 1.0)) // false
	fmt.Println(Fortune(100000000, 1.0, 100000, 50, 1.0))
	fmt.Println(Fortune(100000000, 1.5, 10000000, 50, 1.0)) // false
	fmt.Println(Fortune(100000000, 5.0, 1000000, 50, 1.0))

}

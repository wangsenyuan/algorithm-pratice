package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for t > 0 {
		t--
		time1 := readTime()
		time2 := readTime()
		var dist int
		fmt.Scanf("%d", &dist)

		//plan A
		fmt.Printf("%d.0 ", time1+dist-time2)

		//plan B
		if time2+2*dist <= time1 {
			fmt.Printf("%d.0\n", time1-time2)
			continue
		}

		time3 := 0.5*float64(time1 - time2) + float64(dist)
		fmt.Printf("%0.1f\n", time3)
	}
}

func readTime() int {
	var a, b int
	fmt.Scanf("%d:%d", &a, &b)
	return a*60 + b
}

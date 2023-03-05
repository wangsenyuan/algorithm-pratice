package p2579

import "sort"

func splitNum(num int) int {
	var arr []int

	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}

	sort.Ints(arr)
	var num1, num2 int
	for i := 0; i < len(arr); i += 2 {
		num1 = num1*10 + arr[i]
		if i+1 < len(arr) {
			num2 = num2*10 + arr[i+1]
		}
	}

	return num1 + num2
}

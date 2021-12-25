package main

import "fmt"

func readInt() int {
	var x int
	fmt.Scanf("%d", &x)
	return x
}

func main() {
	t := readInt()

	for i := 0; i < t; i++ {
		solve()
	}
}

func solve() {
	n := readInt()

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = -1
	}

	k := getK(n, arr)
	//fmt.Printf("[debug] k is %d\n", k)
	if k == 1 {
		fmt.Printf("2 %d\n", arr[0])
		return
	}

	if 2*k > n {
		if arr[n-1] < 0 {
			fmt.Printf("1 %d\n", n)
			arr[n-1] = readInt()
		}
		fmt.Printf("2 %d\n", arr[n-1])
		return
	}

	if arr[k] < 0 {
		fmt.Printf("1 %d\n", 1+k)
		arr[k] = readInt()
	}
	if arr[2*k] < 0 {
		fmt.Printf("1 %d\n", 2*k+1)
		arr[2*k] = readInt()
	}
	if arr[k] == arr[2*k] {
		// k + 1 is the K
		fmt.Printf("2 %d\n", arr[0])
		return
	}

	res := find(n, k, arr)

	fmt.Printf("2 %d\n", res)
}

func find(n int, k int, arr []int) int {
	left, right := 0, n/k+1

	for left < right {
		mid := (left + right) / 2
		if mid == left {
			break
		}
		kth := mid * k
		if arr[kth] < 0 {
			fmt.Printf("1 %d\n", kth+1)
			arr[kth] = readInt()
		}
		if arr[kth-1] < 0 {
			fmt.Printf("1 %d\n", kth)
			arr[kth-1] = readInt()
		}

		if arr[kth-1] < arr[kth] {
			left = mid
		} else {
			right = mid
		}
	}

	return arr[left*k]
}

func getK(n int, arr []int) int {
	fmt.Printf("1 %d\n", 1)
	arr[0] = readInt()
	var ask func(left int, right int) int

	ask = func(left int, right int) int {
		if left == right {
			return left
		}
		mid := left + (right-left)/2
		if arr[mid] < 0 {
			fmt.Printf("1 %d\n", mid+1)
			arr[mid] = readInt()
		}
		if arr[0] < arr[mid] {
			return ask(left, mid)
		}
		return ask(mid+1, right)
	}
	return ask(0, n)
}

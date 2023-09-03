package p2843

func countSymmetricIntegers(low int, high int) int {
	var ans int

	for i := low; i <= high; i++ {
		if check(i) {
			ans++
		}
	}
	return ans
}

func check(num int) bool {
	var arr []int
	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}
	if len(arr)%2 != 0 {
		return false
	}
	var sum int
	for i := 0; i < len(arr); i++ {
		if i < len(arr)/2 {
			sum += arr[i]
		} else {
			sum -= arr[i]
		}
	}
	return sum == 0
}

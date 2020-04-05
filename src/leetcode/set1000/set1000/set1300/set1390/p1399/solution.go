package p1399

func countLargestGroup(n int) int {
	cnt := make(map[int]int)

	var max int
	for i := 1; i <= n; i++ {
		d := sumOfDigits(i)
		cnt[d]++
		if cnt[d] > max {
			max = cnt[d]
		}
	}

	var ans int

	for _, v := range cnt {
		if max == v {
			ans++
		}
	}

	return ans
}

func sumOfDigits(num int) int {
	var res int
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}

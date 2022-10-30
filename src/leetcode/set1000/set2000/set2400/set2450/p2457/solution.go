package p2457

func makeIntegerBeautiful(n int64, target int) int64 {
	// target <= 150
	// 把 n 用 x[0...m] 表示
	var arr []int
	num := n
	for num > 0 {
		arr = append(arr, int(num%10))
		num /= 10
	}

	// 在某几位数上增加，使得 sum(arr) <= target
	m := len(arr)
	add := make([]int, m)
	var sum int

	for i := 0; i < m; i++ {
		sum += arr[i]
	}

	for i := 0; i < m && sum > target; i++ {
		x := arr[i]
		if x == 0 {
			continue
		}
		sum -= x
		sum++
		add[i] = 10 - x
		if i+1 < m {
			arr[i+1]++
		}
	}

	var res int64

	for i := m - 1; i >= 0; i-- {
		res = res*10 + int64(add[i])
	}
	return res
}

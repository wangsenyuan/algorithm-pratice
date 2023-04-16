package p2644

func maxDivScore(nums []int, divisors []int) int {
	res := divisors[0]

	score := divScore(nums, divisors[0])

	for i := 1; i < len(divisors); i++ {
		tmp := divScore(nums, divisors[i])
		if tmp > score || tmp == score && res > divisors[i] {
			score = tmp
			res = divisors[i]
		}
	}
	return res
}

func divScore(nums []int, div int) int {
	var res int
	for _, num := range nums {
		if num%div == 0 {
			res++
		}
	}
	return res
}

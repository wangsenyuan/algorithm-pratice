package p2827

func numberOfBeautifulIntegers(low int, high int, k int) int {
	res := countBeautifulNums(high, k)
	res -= countBeautifulNums(low-1, k)
	return res
}

func countBeautifulNums(high int, k int) int {
	if high < 10 {
		return 0
	}

	digits := getDigits(high)
	n := len(digits)

	dp := make([][][]int, 3)
	fp := make([][][]int, 3)

	for i := 0; i < 3; i++ {
		dp[i] = make([][]int, 2*n+1)
		fp[i] = make([][]int, 2*n+1)
		for j := 0; j <= 2*n; j++ {
			dp[i][j] = make([]int, k+1)
			fp[i][j] = make([]int, k+1)
		}
	}

	set := func() {
		for i := 0; i < 3; i++ {
			for j := 0; j <= 2*n; j++ {
				for x := 0; x <= k; x++ {
					dp[i][j][x] = fp[i][j][x]
					fp[i][j][x] = 0
				}
			}
		}
	}

	// 0 for 0 for now
	// 1 for equals for now
	// 2 for less than
	dp[0][n][0] = 1

	for p, d := range digits {
		for i := 0; i < 3; i++ {
			for j := -n; j < n; j++ {
				for x := 0; x <= k; x++ {
					if dp[i][j+n][x] == 0 {
						continue
					}
					if i == 0 {
						ed := d
						if p > 0 {
							ed = 9
						}
						for nd := 1; nd <= ed; nd++ {
							nx := (x*10 + nd) % k
							ni := 2
							if nd == d && p == 0 {
								ni = 1
							}
							nj := j
							if nd%2 == 0 {
								nj++
							} else {
								nj--
							}
							fp[ni][nj+n][nx] += dp[i][j+n][x]
						}
					} else {
						// i == 1 or i == 2
						ed := d
						if i == 2 {
							ed = 9
						}

						for nd := 0; nd <= ed; nd++ {
							nx := (x*10 + nd) % k
							ni := 2
							if i == 1 && nd == ed {
								ni = 1
							}
							nj := j
							if nd%2 == 0 {
								nj++
							} else {
								nj--
							}
							fp[ni][nj+n][nx] += dp[i][j+n][x]
						}
					}
				}
			}
		}
		set()
		dp[0][n][0]++
	}

	var ans int

	for i := 1; i <= 2; i++ {
		ans += dp[i][n][0]
	}

	return ans
}

func getDigits(num int) []int {
	var arr []int
	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}
	reverse(arr)

	return arr
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

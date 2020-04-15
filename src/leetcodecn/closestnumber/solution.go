package closestnumber

func findClosedNumbers(num int) []int {
	x := findGt(num)
	y := findLt(num)

	return []int{x, y}
}

func findGt(num int) int {
	if num == 2147483647 {
		return -1
	}
	var i int

	for num&(1<<i) == 0 {
		i++
	}
	var cnt int = 1
	for i < 31 && num&(1<<(i+1)) > 0 {
		cnt++
		i++
	}
	if i == 31 {
		return -1
	}
	cnt--
	// 1 at pos i, left shift 1
	num ^= 1 << i
	num ^= 1 << (i + 1)

	for j := i - 1; j >= 0; j-- {
		if num&(1<<j) > 0 {
			num ^= 1 << j
		}
	}

	for j := 0; cnt > 0; j++ {
		num ^= 1 << j
		cnt--
	}

	return num
}

func findLt(num int) int {
	if num == 1 {
		return -1
	}

	// find right-most 1, whose right-side is 0
	var i int = 1
	var cnt int
	for i < 32 {
		if num&(1<<(i-1)) == 0 && num&(1<<i) > 0 {
			break
		}

		if num&(1<<(i-1)) > 0 {
			cnt++
		}

		i++
	}

	if i == 32 {
		return -1
	}

	num ^= 1 << i
	num ^= 1 << (i - 1)

	for j := i - 2; j >= 0; j-- {
		if num&(1<<j) > 0 {
			if cnt == 0 {
				num ^= (1 << j)
			} else {
				cnt--
			}
		} else {
			if cnt > 0 {
				num ^= 1 << j
				cnt--
			}
		}
	}

	return num
}

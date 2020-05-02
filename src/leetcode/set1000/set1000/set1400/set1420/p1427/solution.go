package p1427

func maxDiff(num int) int {
	digits := make([]int, 9)
	var h int
	for n := num; n > 0; n /= 10 {
		digits[h] = n % 10
		h++
	}
	// try change it to like 999

	for i, j := 0, h-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	backup := make([]int, h)
	copy(backup, digits)

	var x int
	for x < h && digits[x] == 9 {
		x++
	}

	if x < h {
		//change digits[x] to 9
		tmp := digits[x]
		for x < h {
			if digits[x] == tmp {
				digits[x] = 9
			}
			x++
		}
	}

	a := toNum(digits, h)

	copy(digits, backup)

	if digits[0] > 1 {
		// change digits[0] to 1
		tmp := digits[0]
		for i := 0; i < h; i++ {
			if digits[i] == tmp {
				digits[i] = 1
			}
		}
	} else {
		// change next to 0
		var i int
		for i < h && (digits[i] == 1 || digits[i] == 0) {
			i++
		}
		if i < h {
			tmp := digits[i]
			for i < h {
				if digits[i] == tmp {
					digits[i] = 0
				}
				i++
			}
		}
	}

	b := toNum(digits, h)

	return a - b
}

func toNum(digits []int, h int) int {
	var res int

	for i := 0; i < h; i++ {
		res = res*10 + digits[i]
	}

	return res
}

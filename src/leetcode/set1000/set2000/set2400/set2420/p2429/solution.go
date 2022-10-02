package p2429

const H = 30

func minimizeXor(num1 int, num2 int) int {
	if num2 == 0 {
		return num1
	}

	var cnt int

	for num2 > 0 {
		cnt++
		num2 -= num2 & -num2
	}

	var res int

	for i := H - 1; i >= 0; i-- {
		if (num1>>i)&1 == 1 {
			res |= 1 << i
			cnt--
		}
		if cnt == 0 {
			break
		}
	}

	for i := 0; i < H && cnt > 0; i++ {
		if (num1>>i)&1 == 0 {
			res |= 1 << i
			cnt--
		}
	}

	return res
}

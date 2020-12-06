package p5620

const MOD = 1000000007

func concatenatedBinary(n int) int {
	var cur int64
	digits := make([]int, 20)
	for i := 1; i <= n; i++ {
		p := binary(i, digits)
		for j := p - 1; j >= 0; j-- {
			cur = cur*2%MOD + int64(digits[j])
			cur %= MOD
		}
	}
	return int(cur)
}

func binary(num int, digits []int) int {
	var n int
	for num > 0 {
		digits[n] = num & 1
		n++
		num >>= 1
	}
	return n
}

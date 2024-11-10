package p3345

func smallestNumber(n int, t int) int {

	for {
		if digitProd(n)%t == 0 {
			return n
		}
		n++
	}

}

func digitProd(num int) int {
	res := 1
	for num > 0 {
		res *= num % 10
		num /= 10
	}
	return res
}

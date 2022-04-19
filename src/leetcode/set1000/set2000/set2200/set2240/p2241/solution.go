package p2241

type ATM struct {
	qs []int
}

var factors = []int{20, 50, 100, 200, 500}

func Constructor() ATM {
	qs := make([]int, 5)
	return ATM{qs}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i := 0; i < 5; i++ {
		this.qs[i] += banknotesCount[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
	res := make([]int, 5)
	tmp := amount
	for i := 4; i >= 0 && tmp > 0; i-- {
		if tmp < factors[i] {
			continue
		}
		x := this.qs[i]
		y := tmp / factors[i]
		z := min(x, y)
		res[i] = z
		tmp -= factors[i] * z
		this.qs[i] -= z
	}

	if tmp == 0 {
		return res
	}

	this.Deposit(res)
	return []int{-1}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */

package p0242

type Bank struct {
	n       int
	balance []int64
}

func Constructor(balance []int64) Bank {
	n := len(balance)
	return Bank{n, balance}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	account1--
	account2--
	if account1 < 0 || account1 >= this.n {
		return false
	}
	if account2 < 0 || account2 >= this.n {
		return false
	}
	if money > this.balance[account1] {
		return false
	}
	this.balance[account1] -= money
	this.balance[account2] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	account--
	if account < 0 || account >= this.n {
		return false
	}

	this.balance[account] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	account--
	if account < 0 || account >= this.n {
		return false
	}
	if this.balance[account] < money {
		return false
	}
	this.balance[account] -= money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */

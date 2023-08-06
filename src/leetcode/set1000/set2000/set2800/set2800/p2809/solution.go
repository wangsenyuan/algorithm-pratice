package p2809

func accountBalanceAfterPurchase(purchaseAmount int) int {
	// 100
	x := purchaseAmount / 10
	if x == 10 {
		return 0
	}
	a := purchaseAmount - x*10
	b := (x+1)*10 - purchaseAmount

	if a < b {
		return 100 - x*10
	}

	return 100 - (x+1)*10
}

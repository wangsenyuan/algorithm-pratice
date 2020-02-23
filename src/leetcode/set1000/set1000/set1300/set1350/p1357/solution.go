package p1357

type Cashier struct {
	n        int
	discount float64
	products map[int]int
	cur      int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	mem := make(map[int]int)

	for i := 0; i < len(products); i++ {
		mem[products[i]] = prices[i]
	}
	return Cashier{n, float64(discount), mem, 0}
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	this.cur++

	x := totalPrice(*this, product, amount)

	if this.cur%this.n == 0 {
		x -= x * this.discount / 100
	}

	return x
}

func totalPrice(cashier Cashier, product []int, amount []int) float64 {
	var res float64

	for i := 0; i < len(product); i++ {
		id := product[i]
		cnt := amount[i]
		price := cashier.products[id]

		res += float64(cnt) * float64(price)
	}

	return res
}

/**
 * Your Cashier object will be instantiated and called as such:
 * obj := Constructor(n, discount, products, prices);
 * param_1 := obj.GetBill(product,amount);
 */

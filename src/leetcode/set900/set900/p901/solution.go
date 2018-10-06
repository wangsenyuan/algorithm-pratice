package p901

const MAX_PRICE = 1e5 + 10

type StockSpanner struct {
	arr []int
	cnt []int
	sz  int
}

func Constructor() StockSpanner {
	arr := make([]int, MAX_PRICE)
	cnt := make([]int, MAX_PRICE)
	return StockSpanner{arr, cnt, 0}
}

func (this *StockSpanner) Next(price int) int {
	arr := this.arr
	cnt := this.cnt
	sz := this.sz

	res := 1

	for sz > 0 && arr[sz-1] <= price {
		res += cnt[sz-1]
		sz--
	}

	arr[sz] = price
	cnt[sz] = res
	sz++
	this.sz = sz

	return res
}

package p1801

import "container/heap"

const MOD = 1000000007

func getNumberOfBacklogOrders(orders [][]int) int {
	n := len(orders)
	buy := make(Buy, 0, n)
	sell := make(Sell, 0, n)

	for _, order := range orders {
		price := order[0]
		amount := order[1]

		if order[2] == 0 {
			// a buy order
			for sell.Len() > 0 && sell[0].price <= price && amount > 0 {
				tmp := heap.Pop(&sell).(*Order)
				if amount >= tmp.amount {
					amount -= tmp.amount
					continue
				}
				tmp.amount -= amount
				amount = 0
				heap.Push(&sell, tmp)
			}

			if amount > 0 {
				heap.Push(&buy, &Order{price: price, amount: amount})
			}
		} else {
			for buy.Len() > 0 && buy[0].price >= price && amount > 0 {
				tmp := heap.Pop(&buy).(*Order)
				if amount >= tmp.amount {
					amount -= tmp.amount
					continue
				}
				tmp.amount -= amount
				amount = 0
				heap.Push(&buy, tmp)
			}
			if amount > 0 {
				heap.Push(&sell, &Order{price: price, amount: amount})
			}
		}
	}
	var res int
	for buy.Len() > 0 {
		cur := heap.Pop(&buy).(*Order)
		res += cur.amount
		if res >= MOD {
			res -= MOD
		}
	}

	for sell.Len() > 0 {
		cur := heap.Pop(&sell).(*Order)
		res += cur.amount
		if res >= MOD {
			res -= MOD
		}
	}
	return res
}

type Order struct {
	price  int
	amount int
	index  int
}

type Buy []*Order

func (pq Buy) Len() int { return len(pq) }

func (pq Buy) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].price > pq[j].price
}

func (pq Buy) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *Buy) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Order)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *Buy) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type Sell []*Order

func (pq Sell) Len() int { return len(pq) }

func (pq Sell) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].price < pq[j].price
}

func (pq Sell) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *Sell) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Order)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *Sell) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

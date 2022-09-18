package p2411

import "sort"

const N_INF = -(1 << 60)

func minimumMoney(transactions [][]int) int64 {
	// money - cost + cash money >= cost
	// money 有可能减少
	n := len(transactions)
	// 关键是找到 <= money 的 transactions中，收益最大的那个
	txs := make([]Tx, 0, n)

	var most_positive int

	for i := 0; i < n; i++ {
		if transactions[i][0] <= transactions[i][1] {
			most_positive = max(most_positive, transactions[i][0])
		} else {
			txs = append(txs, Tx{transactions[i][0], transactions[i][1]})
		}
	}

	sort.Slice(txs, func(i, j int) bool {
		return txs[i].Less(txs[j])
	})

	var money int64
	var cur int64
	for i := 0; i < len(txs); i++ {
		if cur < int64(txs[i].cost) {
			money += int64(txs[i].cost) - cur
			cur = int64(txs[i].cost)
		}
		cur -= int64(txs[i].cost - txs[i].back)
	}

	if cur < int64(most_positive) {
		money += int64(most_positive) - cur
	}
	return money
}

func search(n int, f func(int) bool) int {
	l, r := 0, n

	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Tx struct {
	cost int
	back int
}

func (this Tx) Less(that Tx) bool {
	x := this.cost
	if this.back < that.cost {
		x += that.cost - this.back
	}
	y := that.cost
	if that.back < this.cost {
		y += this.cost - that.back
	}
	return x >= y
}

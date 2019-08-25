package p1169

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func invalidTransactions(transactions []string) []string {
	n := len(transactions)
	all_txs := make([]*Transaction, n)
	byPersons := make(map[string][]*Transaction)

	for i := 0; i < n; i++ {
		tx := parseAsTransaction(transactions[i])
		all_txs[i] = &tx

		if _, found := byPersons[tx.person]; !found {
			byPersons[tx.person] = make([]*Transaction, 0, 10)
		}
		byPersons[tx.person] = append(byPersons[tx.person], all_txs[i])
	}

	for person := range byPersons {
		txs := byPersons[person]
		if len(txs) > 1 {
			sort.Sort(Txs(txs))

			for i := 0; i < len(txs); i++ {
				j := sort.Search(i, func(j int) bool {
					return txs[j].time >= txs[i].time-60
				})

				for j < i {
					if txs[j].city != txs[i].city {
						txs[j].invalid = true
						txs[i].invalid = true
					}
					j++
				}
			}
		}
	}
	var res []string

	for i := 0; i < n; i++ {
		if all_txs[i].invalid || all_txs[i].amount > 1000 {
			res = append(res, formatTransaction(*all_txs[i]))
		}
	}

	return res
}

type Transaction struct {
	person  string
	time    int
	amount  int
	city    string
	invalid bool
}

func parseAsTransaction(s string) Transaction {
	ss := strings.Split(s, ",")
	person := ss[0]
	time, _ := strconv.Atoi(ss[1])
	amount, _ := strconv.Atoi(ss[2])
	city := ss[3]
	return Transaction{person, time, amount, city, false}
}

func formatTransaction(tx Transaction) string {
	return fmt.Sprintf("%s,%d,%d,%s", tx.person, tx.time, tx.amount, tx.city)
}

type Txs []*Transaction

func (this Txs) Len() int {
	return len(this)
}

func (this Txs) Less(i, j int) bool {
	a := this[i]
	b := this[j]
	if a.person < b.person {
		return true
	} else if a.person == b.person && a.time < b.time {
		return true
	}
	return false
}

func (this Txs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

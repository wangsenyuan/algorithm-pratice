package p2411

import "testing"

func runSample(t *testing.T, transactions [][]int, expect int64) {
	res := minimumMoney(transactions)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	transactions := [][]int{{2, 1}, {5, 0}, {4, 2}}
	var expect int64 = 10
	runSample(t, transactions, expect)
}

func TestSample2(t *testing.T) {
	transactions := [][]int{{3, 0}, {0, 3}}
	var expect int64 = 3
	runSample(t, transactions, expect)
}

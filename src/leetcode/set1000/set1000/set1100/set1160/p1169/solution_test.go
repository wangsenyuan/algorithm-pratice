package p1169

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, transactions []string, expect []string) {
	res := invalidTransactions(transactions)
	sort.Strings(expect)
	sort.Strings(res)
	if !reflect.DeepEqual(expect, res) {
		t.Errorf("Sample %v, expect %v, but got %v", transactions, expect, res)
	}
}

func TestSample(t *testing.T) {
	txs := []string{"alice,20,800,mtv", "alice,50,100,beijing"}
	expect := []string{"alice,20,800,mtv", "alice,50,100,beijing"}
	runSample(t, txs, expect)
}

func TestSample2(t *testing.T) {
	txs := []string{"alice,20,800,mtv", "alice,50,1200,mtv"}
	expect := []string{"alice,50,1200,mtv"}
	runSample(t, txs, expect)
}

func TestSample3(t *testing.T) {
	txs := []string{"alice,20,800,mtv", "bob,50,1200,mtv"}
	expect := []string{"bob,50,1200,mtv"}
	runSample(t, txs, expect)
}

package p1418

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, orders [][]string, expect [][]string) {
	res := displayTable(orders)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", orders, expect, res)
	}
}

func TestSample1(t *testing.T) {
	orders := [][]string{
		{"David", "3", "Ceviche"}, {"Corina", "10", "Beef Burrito"}, {"David", "3", "Fried Chicken"}, {"Carla", "5", "Water"}, {"Carla", "5", "Ceviche"}, {"Rous", "3", "Ceviche"},
	}
	expect := [][]string{
		{"Table", "Beef Burrito", "Ceviche", "Fried Chicken", "Water"}, {"3", "0", "2", "1", "0"}, {"5", "0", "1", "0", "1"}, {"10", "1", "0", "0", "0"},
	}
	runSample(t, orders, expect)
}

func TestSample2(t *testing.T) {
	orders := [][]string{
		{"James", "12", "Fried Chicken"}, {"Ratesh", "12", "Fried Chicken"}, {"Amadeus", "12", "Fried Chicken"}, {"Adam", "1", "Canadian Waffles"}, {"Brianna", "1", "Canadian Waffles"},
	}
	expect := [][]string{
		{"Table", "Canadian Waffles", "Fried Chicken"}, {"1", "2", "0"}, {"12", "0", "3"},
	}
	runSample(t, orders, expect)
}

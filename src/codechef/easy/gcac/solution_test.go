package main

import "testing"

func TestSample(t *testing.T) {
	n, m := 5, 6

	minSalary := []int{5000, 10000, 3000, 20, 100}
	offeredSalary := []int{10000, 800, 600, 10, 1000, 2000}
	maxJobOffers := []int{2, 2, 1, 8, 9, 10}

	qual := [][]byte{
		[]byte(string("111111")),
		[]byte(string("100000")),
		[]byte(string("000000")),
		[]byte(string("000001")),
		[]byte(string("100100")),
	}

	employee, salary, employer := solve(n, minSalary, m, offeredSalary, maxJobOffers, qual)

	if employee != 3 {
		t.Errorf("the sample should have 3 employeed, but got %d", employee)
	}

	if salary != 22000 {
		t.Errorf("the sample should have total 22000 salary, but got %d", salary)
	}

	if employer != 4 {
		t.Errorf("the sample should have 2 employers, but got %d", employer)
	}
}

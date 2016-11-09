package main

import (
	"reflect"
	"testing"
)

func TestFindItinerary(t *testing.T) {
	cases := []struct {
		tickets [][]string
		expect  []string
	}{
		{
			[][]string{[]string{"MUC", "LHR"}, []string{"JFK", "MUC"}, []string{"SFO", "SJC"}, []string{"LHR", "SFO"}},
			[]string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			[][]string{[]string{"JFK", "SFO"}, []string{"JFK", "ATL"}, []string{"SFO", "ATL"}, []string{"ATL", "JFK"}, []string{"ATL", "SFO"}},
			[]string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
	}

	for _, ex := range cases {
		tickets := ex.tickets
		expect := ex.expect
		got := findItinerary(tickets)
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("%v should get %v, but get %v\n", tickets, expect, got)
		}
	}
}

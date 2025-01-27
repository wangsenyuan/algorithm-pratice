package p3433

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, events [][]string, expect []int) {
	res := countMentions(n, events)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	events := [][]string{
		{"MESSAGE", "10", "id1 id0"}, {"OFFLINE", "11", "0"}, {"MESSAGE", "71", "HERE"},
	}
	expect := []int{2, 2}
	runSample(t, n, events, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	events := [][]string{
		{"MESSAGE", "10", "id1 id0"}, {"OFFLINE", "11", "0"}, {"MESSAGE", "12", "ALL"},
	}
	expect := []int{2, 2}
	runSample(t, n, events, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	events := [][]string{
		{"OFFLINE", "10", "0"}, {"MESSAGE", "12", "HERE"},
	}
	expect := []int{0, 1}
	runSample(t, n, events, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	events := [][]string{
		{"MESSAGE", "2", "HERE"}, {"OFFLINE", "2", "1"}, {"OFFLINE", "1", "0"}, {"MESSAGE", "61", "HERE"},
	}
	expect := []int{1, 0, 2}
	runSample(t, n, events, expect)
}

func TestSample5(t *testing.T) {
	n := 10
	events := [][]string{
		{"OFFLINE", "51", "2"}, {"MESSAGE", "100", "ALL"}, {"MESSAGE", "81", "id1 id7 id5 id0 id6 id7 id5 id2"}, {"MESSAGE", "74", "id5 id9 id7 id6 id9 id1"}, {"OFFLINE", "99", "7"}, {"MESSAGE", "83", "ALL"}, {"MESSAGE", "75", "id4 id3 id1 id0"}, {"OFFLINE", "9", "3"}, {"OFFLINE", "86", "8"}, {"OFFLINE", "62", "0"}, {"MESSAGE", "25", "HERE"}, {"OFFLINE", "13", "6"}, {"MESSAGE", "69", "HERE"}, {"MESSAGE", "25", "id9"}, {"MESSAGE", "16", "ALL"}, {"OFFLINE", "20", "4"}, {"MESSAGE", "85", "ALL"}, {"MESSAGE", "65", "id9 id8 id3 id1 id0 id6 id1 id9 id5 id1"}, {"OFFLINE", "69", "3"},
	}
	expect := []int{8, 12, 6, 6, 5, 10, 7, 9, 7, 11}
	runSample(t, n, events, expect)
}

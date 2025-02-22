package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, records []string, expect []string) {
	res := solve(records)

	sort.Strings(expect)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	records := []string{
		"vasya 100",
		"vasya 200",
		"artem 100",
		"kolya 200",
		"igor 250",
	}
	expect := []string{
		"artem noob",
		"igor pro",
		"kolya random",
		"vasya random",
	}
	runSample(t, records, expect)
}

func TestSample2(t *testing.T) {
	records := []string{
		"vasya 200",
		"kolya 1000",
		"vasya 1000",
	}
	expect := []string{
		"kolya pro",
		"vasya pro",
	}
	runSample(t, records, expect)
}

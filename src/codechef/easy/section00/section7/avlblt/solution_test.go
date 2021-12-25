package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, T int, blocks [][]string, expect [][]string) {
	res := solve(T, blocks)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", T, blocks, expect, res)
	}
}

func TestSample1(t *testing.T) {
	T := 12
	blocks := [][]string{
		{"00:00 07:30", "08:30 15:30", "16:30 00:00"},
		{},
		{},
		{},
		{},
		{},
		{},
	}
	expect := [][]string{
		{"00:00 03:30", "04:30 12:00"},
		{},
		{},
		{},
		{},
		{},
		{"12:00 19:30", "20:30 00:00"},
	}

	runSample(t, T, blocks, expect)
}

func TestSample2(t *testing.T) {
	T := 12
	blocks := [][]string{
		{"00:00 07:30", "08:30 15:30", "16:30 00:00"},
		{},
		{},
		{},
		{},
		{},
		{"16:30 00:00"},
	}
	expect := [][]string{
		{"00:00 03:30", "04:30 12:00"},
		{},
		{},
		{},
		{},
		{},
		{"04:30 19:30", "20:30 00:00"},
	}

	runSample(t, T, blocks, expect)
}

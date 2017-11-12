package main

import (
	"testing"
	"reflect"
)

func TestCase1(t *testing.T) {
	findNums := []int{4, 1, 2}
	nums := []int{1, 3, 4, 2}
	res := nextGreaterElement(findNums, nums)
	exp := []int{-1, 3, -1}
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("answer should be %v, but is %v\n", exp, res)
	}
}

func TestCase2(t *testing.T) {
	findNums := []int{2, 4}
	nums := []int{1, 2, 3, 4}
	res := nextGreaterElement(findNums, nums)
	exp := []int{3, -1}
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("answer should be %v, but is %v\n", exp, res)
	}
}

func TestCase3(t *testing.T) {
	findNums := []int{}
	nums := []int{}
	res := nextGreaterElement(findNums, nums)
	if len(res) != 0 {
		t.Errorf("answer should be empty, but is %v\n", res)
	}
}

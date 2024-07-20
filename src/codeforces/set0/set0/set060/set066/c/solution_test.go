package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, lines []string, expect []int) {
	res := solve(lines)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	lines := []string{
		"C:\\folder1\\file1.txt",
	}
	expect := []int{0, 1}
	runSample(t, lines, expect)
}

func TestSample2(t *testing.T) {
	lines := []string{
		"C:\\folder1\\folder2\\folder3\\file1.txt",
		"C:\\folder1\\folder2\\folder4\\file1.txt",
		"D:\\folder1\\file1.txt",
	}
	expect := []int{3, 2}
	runSample(t, lines, expect)
}

func TestSample3(t *testing.T) {
	lines := []string{
		"C:\\file\\file\\file\\file\\file.txt",
		"C:\\file\\file\\file\\file2\\file.txt",
	}
	expect := []int{4, 2}
	runSample(t, lines, expect)
}

func TestSample4(t *testing.T) {
	lines := []string{
		"C:\\a\\b\\c\\d\\d.txt",
		"C:\\a\\b\\c\\e\\f.txt",
	}
	expect := []int{4, 2}
	runSample(t, lines, expect)
}

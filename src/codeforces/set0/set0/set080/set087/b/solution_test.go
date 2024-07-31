package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, statements []string, expect []string) {
	res := solve(statements)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	statements := []string{
		"typedef void* ptv",
		"typeof ptv",
		"typedef &&ptv node",
		"typeof node",
		"typeof &ptv",
	}
	expect := []string{
		"void*",
		"errtype",
		"void",
	}
	runSample(t, statements, expect)
}

func TestSample2(t *testing.T) {
	statements := []string{
		"typedef void* b",
		"typedef b* c",
		"typeof b",
		"typeof c",
		"typedef &b b",
		"typeof b",
		"typeof c",
		"typedef &&b* c",
		"typeof c",
		"typedef &b* c",
		"typeof c",
		"typedef &void b",
		"typeof b",
		"typedef b******* c",
		"typeof c",
		"typedef &&b* c",
		"typeof c",
	}
	expect := []string{
		"void*",
		"void**",
		"void",
		"void**",
		"errtype",
		"void",
		"errtype",
		"errtype",
		"errtype",
	}
	runSample(t, statements, expect)
}

package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
3
Hamilton
Vettel
Webber
2
Webber
Vettel
2
Hamilton
Vettel
`
	expect := []string{"Vettel", "Hamilton"}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
7
Prost
Surtees
Nakajima
Schumacher
Button
DeLaRosa
Buemi
8
Alonso
Prost
NinoFarina
JimClark
DeLaRosa
Nakajima
Patrese
Surtees
`
	expect := []string{"Prost", "Prost"}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
10
Hamilton
Vettel
Webber
Button
Kubica
Petrov
Roseberg
Schumacher
Alonson
Massa
11
Senna
Kobayashi
DeLaRosa
Hakkinen
Raikkonen
Prost
JimClark
Nakajima
Berger
AlbertoAscari
Hamilton
`
	expect := []string{"Hamilton", "Hamilton"}
	runSample(t, s, expect)
}

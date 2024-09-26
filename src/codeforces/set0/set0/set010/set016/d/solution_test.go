package main

import "testing"

func runSample(t *testing.T, logs []string, expect int) {
	res := solve(logs)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	logs := []string{
		"[05:00 a.m.]: Server is started",
		"[05:00 a.m.]: Rescan initialized",
		"[01:13 p.m.]: Request processed",
		"[01:10 p.m.]: Request processed",
		"[11:40 p.m.]: Rescan completed",
	}
	expect := 2
	runSample(t, logs, expect)
}

func TestSample2(t *testing.T) {
	logs := []string{
		"[09:00 a.m.]: User logged in",
		"[08:00 a.m.]: User logged in",
		"[07:00 a.m.]: User logged in",
	}
	expect := 3
	runSample(t, logs, expect)
}

func TestSample3(t *testing.T) {
	logs := []string{
		"[07:23 p.m.]: y vTNVMa VWxb  rpE",
		"[12:00 a.m.]: wkr EcZc",
		"[05:16 a.m.]: nWf lypg NT",
		"[04:22 a.m.]: UQIXmL",
	}
	expect := 3
	runSample(t, logs, expect)
}

func TestSample4(t *testing.T) {
	logs := []string{
		"[09:37 p.m.]: pamammapampmapaa",
		"[09:37 p.m.]: ppmm",
		"[09:37 p.m.]: aapapppaampmappppppm",
		"[09:37 p.m.]: pmppmpmmpm",
		"[09:37 p.m.]: mmppppamamaa",
		"[09:37 p.m.]: mm",
		"[09:37 p.m.]: apamppmaaapaa",
		"[09:37 p.m.]: pmaammpaa",
		"[09:37 p.m.]: m",
		"[09:37 p.m.]: pppmppa",
		"[09:37 p.m.]: ppmpmm",
		"[09:37 p.m.]: mpamappmpmpamaampmpm",
		"[05:10 a.m.]: a",
		"[05:10 a.m.]: aaapamppaaamppapa",
	}
	expect := 3
	runSample(t, logs, expect)
}

func TestSample5(t *testing.T) {
	logs := []string{
		"[06:53 a.m.]: wmaflY",
		"[04:14 a.m.]: OGcky HfWLWqj",
		"[08:03 a.m.]: A RX kpH   D   Ov",
		"[03:31 a.m.]: jaoQKHYO TLawGsD",
		"[04:20 a.m.]: e s",
		"[04:20 a.m.]: mjNaqjnQf",
		"[04:20 a.m.]: bpS  A  HQDyVjGM",
		"[04:20 a.m.]: G  oUz",
		"[04:20 a.m.]: S r HSEFIphbGM",
		"[04:20 a.m.]: xfEf KEHn PuFGjSa",
		"[04:20 a.m.]: X Dh y",
		"[04:20 a.m.]: rWVOKl",
		"[04:20 a.m.]: c pOTbKJW P  Sno",
		"[04:20 a.m.]: yMhXqbALYINf",
		"[04:20 a.m.]: pzwJ K",
		"[04:20 a.m.]: i DA l L",
		"[04:20 a.m.]: GmA zX   D",
		"[04:20 a.m.]: yM OQ u GSoYVRkcRS",
		"[04:20 a.m.]: z Thbf...",
	}
	expect := 4
	runSample(t, logs, expect)
}

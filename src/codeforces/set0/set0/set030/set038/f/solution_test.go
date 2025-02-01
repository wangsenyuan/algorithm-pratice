package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, w string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	winner, res := process(reader)

	if w != winner || !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %s %v, but got %s %v", w, expect, winner, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
aba
abac`
	winner := "Second"
	score := []int{29, 35}
	runSample(t, s, winner, score)
}

func TestSample2(t *testing.T) {
	s := `3
artem
nik
max`
	winner := "First"
	score := []int{2403, 1882}
	runSample(t, s, winner, score)
}

func TestSample3(t *testing.T) {
	s := `10
ndqlxtrxiftvtji
aoblenbunumdge
lgkmt
x
x
bg
ds
nlhdlxeh
ugxufipnaxvkxl
k`
	winner := "First"
	score := []int{26161, 23191}
	runSample(t, s, winner, score)
}

func TestSample4(t *testing.T) {
	s := `23
dcpnjubpjzsl
ngjriwa
usoctpzjlnm
a
cdwgavsnnxfyxra
xpbbbmvsveen
aoimjmsbjedryewog
duivkwfrxxkrfqcyb
advttzqrztsoysddhg
dowgfmzitsxzfwcrl
oedzctovdfcmnyxo
dugtkselxsad
nwopexs
crsbpqtwynunf
zoue
impwuhy
pvlsjchpic
kprgfngigrvpnhjh
vmxpckkagbvwarig
hxemegdijylqcfevnvv
mwln
qwtos
tpcgunugerj`
	winner := "Second"
	score := []int{40314, 43585}
	runSample(t, s, winner, score)
}

func TestSample5(t *testing.T) {
	s := `14
ym
bi
yu
yb
vwtb
semn
byr
c
ir
qyx
gnk
ao
dzeo
cd`
	winner := "First"
	score := []int{2228, 1226}
	runSample(t, s, winner, score)
}

func TestSample6(t *testing.T) {
	s := `27
aa
aa
aaaaa
aaa
a
aaaaa
aa
aaaa
aaaaa
aaaaa
a
aa
aa
aaa
aa
aaaa
aa
aaaaaa
aaaaaa
aa
aaaaaa
aaaaa
aaaaaa
a
aaaaaa
aaaaa
aaaaaa`
	winner := "Second"
	score := []int{64, 56}
	runSample(t, s, winner, score)
}

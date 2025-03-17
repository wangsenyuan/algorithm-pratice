package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `ababa
aba
2
c->ba
c->cc
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `ababa
aba
7
c->ba
c->cc
e->ab
z->ea
b->ba
d->dd
d->ab
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `ababa
aba
1
c->ba
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `oaujn
oaemvsbb
50
y->zr
w->qg
f->ep
w->pd
q->gf
g->sn
w->ug
e->yj
n->bb
d->ko
i->ua
f->wi
c->mf
l->mh
m->vy
j->xt
p->hp
f->ok
c->ht
g->kn
w->zj
s->tc
e->bd
l->ym
k->ra
j->lw
v->kq
v->xd
u->em
t->wa
j->vs
h->vs
e->gg
o->bo
m->ka
z->sq
t->hx
x->rz
e->tu
r->ff
b->vf
y->vq
e->um
q->pz
e->xe
w->ci
c->bg
a->kh
x->zf
m->rk
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
50
a->aa
a->bb
a->cc
a->dd
a->ee
a->ff
a->gg
a->hh
a->ii
a->jj
a->kk
a->ll
a->mm
a->nn
a->oo
a->pp
a->qq
a->rr
a->ss
a->tt
a->uu
a->vv
a->ww
a->xx
a->yy
a->aa
b->aa
c->aa
d->aa
e->aa
f->aa
g->aa
h->aa
i->aa
j->aa
k->aa
l->aa
m->aa
n->aa
o->aa
p->aa
q->aa
r->aa
s->aa
t->aa
u->aa
v->aa
w->aa
x->aa
y->aa
`
	expect := 1
	runSample(t, s, expect)
}

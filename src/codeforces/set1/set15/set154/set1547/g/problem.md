You are given a directed graph 𝐺
which can contain loops (edges from a vertex to itself). Multi-edges are absent in 𝐺
which means that for all ordered pairs (𝑢,𝑣)
exists at most one edge from 𝑢
to 𝑣
. Vertices are numbered from 1
to 𝑛
.

A path from 𝑢
to 𝑣
is a sequence of edges such that:

vertex 𝑢
is the start of the first edge in the path;
vertex 𝑣
is the end of the last edge in the path;
for all pairs of adjacent edges next edge starts at the vertex that the previous edge ends on.
We will assume that the empty sequence of edges is a path from 𝑢
to 𝑢
.

For each vertex 𝑣
output one of four values:

0
, if there are no paths from 1
to 𝑣
;
1
, if there is only one path from 1
to 𝑣
;
2
, if there is more than one path from 1
to 𝑣
and the number of paths is finite;
−1
, if the number of paths from 1
to 𝑣
is infinite.

### ideas

1. 一旦进入一个loop，从这个loop能够到达的，就是inf
You have a simple undirected graph consisting of 𝑛
 vertices and 𝑚
 edges. The graph doesn't contain self-loops, there is at most one edge between a pair of vertices. The given graph can be disconnected.

Let's make a definition.

Let 𝑣1
 and 𝑣2
 be two some nonempty subsets of vertices that do not intersect. Let 𝑓(𝑣1,𝑣2)
 be true if and only if all the conditions are satisfied:

There are no edges with both endpoints in vertex set 𝑣1
.
There are no edges with both endpoints in vertex set 𝑣2
.
For every two vertices 𝑥
 and 𝑦
 such that 𝑥
 is in 𝑣1
 and 𝑦
 is in 𝑣2
, there is an edge between 𝑥
 and 𝑦
.
Create three vertex sets (𝑣1
, 𝑣2
, 𝑣3
) which satisfy the conditions below;

All vertex sets should not be empty.
Each vertex should be assigned to only one vertex set.
𝑓(𝑣1,𝑣2)
, 𝑓(𝑣2,𝑣3)
, 𝑓(𝑣3,𝑣1)
 are all true.
Is it possible to create such three vertex sets? If it's possible, print matching vertex set for each vertex.


You are given an undirected graph consisting of 𝑛
vertices. A number is written on each vertex; the number on vertex 𝑖
is 𝑎𝑖
. Initially there are no edges in the graph.

You may add some edges to this graph, but you have to pay for them. The cost of adding an edge between vertices 𝑥
and 𝑦
is 𝑎𝑥+𝑎𝑦
coins. There are also 𝑚
special offers, each of them is denoted by three numbers 𝑥
, 𝑦
and 𝑤
, and means that you can add an edge connecting vertices 𝑥
and 𝑦
and pay 𝑤
coins for it. You don't have to use special offers: if there is a pair of vertices 𝑥
and 𝑦
that has a special offer associated with it, you still may connect these two vertices paying 𝑎𝑥+𝑎𝑦
coins for it.

What is the minimum number of coins you have to spend to make the graph connected? Recall that a graph is connected if
it's possible to get from any vertex to any other vertex using only the edges belonging to this graph.

### ideas

1. 先尽量使用special offer去连接， 这样子会得到一些独立的component
2. 然后用每个component中最小的去连接其他的（component中最小的）
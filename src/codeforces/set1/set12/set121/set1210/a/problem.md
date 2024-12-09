Anadi has a set of dominoes. Every domino has two parts, and each part contains some dots. For every 𝑎
and 𝑏
such that 1≤𝑎≤𝑏≤6
, there is exactly one domino with 𝑎
dots on one half and 𝑏
dots on the other half. The set contains exactly 21
dominoes. Here is an exact illustration of his set:

Also, Anadi has an undirected graph without self-loops and multiple edges. He wants to choose some dominoes and place
them on the edges of this graph. He can use at most one domino of each type. Each edge can fit at most one domino. It's
not necessary to place a domino on each edge of the graph.

When placing a domino on an edge, he also chooses its direction. In other words, one half of any placed domino must be
directed toward one of the endpoints of the edge and the other half must be directed toward the other endpoint. There's
a catch: if there are multiple halves of dominoes directed toward the same vertex, each of these halves must contain the
same number of dots.

How many dominoes at most can Anadi place on the edges of his graph?

### ideas

1. 没个点数，最多出现在6个dominoes中（上下一起算）
2. 所以，也就是是说，如果要在一个节点的边上都放置dominoes，它的deg不能超过6
3. 这个条件似乎始终是存在的，因为只有7个节点，所以deg最多位6
4. 其实就是把6个节点编号分配给7个节点，分配后，再按照6！进行排列，看是否ok
5. 
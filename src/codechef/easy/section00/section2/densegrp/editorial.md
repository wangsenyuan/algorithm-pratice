QUICK EXPLANATION:
First, we compress all the vertices given in the input, so that the number of vertices becomes O(M)O(M).
Build a graph as shown in the figure below, of size equal to the number of compressed vertices. The tree edges are weighted 00. Then for each valid ii, decompose [a_i, b_i][a
i
​
,b
i
​
] into subsegments in the lower half and draw an edge from each of them to each subsegment of [c_i, d_i][c
i
​
,d
i
​
] in the upper half. All these edges should be of weight 11. In this way, the graph consists of O(M\log^2 M)O(Mlog
2
M) edges. Now, we can find the answer by running 0-1 BFS from source [x, x][x,x], and finding the distance to [y, y][y,y]. This solution takes O(M\log^2 M)O(Mlog
2
M) time.
Instead of drawing an edge from each subsegment of [a_i, b_i][a
i
​
,b
i
​
] to each of [c_i, d_i][c
i
​
,d
i
​
], we can define an extra vertex v_iv
i
​
and draw an edge from each subsegment of [a_i, b_i][a
i
​
,b
i
​
] to v_iv
i
​
, followed by an edge from v_iv
i
​
to each subsegment of [c_i, d_i][c
i
​
,d
i
​
]. The graph now consists of a total of O(M\log M)O(MlogM) edges. We can then run 0-1 BFS as before, but divide by 22 to get the final answer. Thus, the problem is solved in O(M\log M)O(MlogM) time.
EXPLANATION:
Firstly, we can compress the input so that instead of dealing with NN vertices, which can be upto 10^910
9
, we only need to deal with O(M)O(M) vertices in our graph, as there are only 4M + 24M+2 vertices to compress. Let us represent the number of vertices after compression to be KK.

![./sample.png](./sample.png?raw=true)

Motivation

We can start by naively building the graph as given in the problem, followed by a BFS. The problem with this approach is that for each valid ii, we require O(M^2)O(M
2
) edges to be added to the graph in the worst case. Clearly, we need a structure which can represent that same information with far fewer edges. We can recognize that a segment tree is well-suited for such a role. For instance, let us build a segment tree of size KK. Then for each valid ii, we can decompose [a_i, b_i][a
i
​
,b
i
​
] into O(\log K)O(logK) subsegments of the segment tree, and draw an edge from all those subsegments to all subsegments of [c_i, d_i][c
i
​
,d
i
​
]. After some experimentation with similar kind of structures, we can realize that the following graph works for us.


Let us construct this graph, i.e. the two segment trees of size KK joined together as shown in the figure. We initially assign all the tree edges weight 00. Now, we start iterating over each valid ii. As we do in any segment tree, we can partition every segment [a, b][a,b] into O(\log K)O(logK) subsegments of the segment tree. Thus, for every [a_i, b_i][a
i
​
,b
i
​
] we draw an edge from each of its subsegments in the lower half to each subsegment of [c_i, d_i][c
i
​
,d
i
​
] in the upper half. We draw these edges with weight 11. So why does this graph work?

We show that if there exists an edge from vertex pp to qq in the original graph, then and only then there exists a path of distance 11 from segment [p, p][p,p] to [q, q][q,q] in our constructed graph. Note that if there is an edge from pp to qq, there must be an ii such that p \in [a_i, b_i]p∈[a
i
​
,b
i
​
] and q \in [c_i, d_i]q∈[c
i
​
,d
i
​
]. Thus, there must be a subsegment of [a_i, b_i][a
i
​
,b
i
​
] in the segment tree which includes pp, let us call it [a'_i, b'_i][a
i
′
​
,b
i
′
​
], and a subsegment of [c_i, d_i][c
i
​
,d
i
​
] including qq, which we can call [c'_i, d'_i][c
i
′
​
,d
i
′
​
].

It is easy to see that all segments in a segment tree that include pp, lie on the path from [p, p][p,p] to the root. Thus, we can move downwards from [p, p][p,p] in our graph and reach [a'_i, b'_i][a
i
′
​
,b
i
′
​
] in the lower half of the graph. Further, as we had drawn an edge from each subsegment of [a_i, b_i][a
i
​
,b
i
​
] to each subsegment of [c_i, d_i][c
i
​
,d
i
​
], there exists an edge from [a'_i, b'_i][a
i
′
​
,b
i
′
​
] in the lower half to [c'_i, d'_i][c
i
′
​
,d
i
′
​
] to the upper half. So we can move to [c'_i, d'_i][c
i
′
​
,d
i
′
​
] and from there on we can move downwards in the tree until we reach [q, q][q,q]. Thus, we have shown that a path exists from [p, p][p,p] to [q, q][q,q]. Also, notice that we used exactly 11 unit of distance in doing so. The only-if part of the proof is very similar.

Hence, any path from xx to yy of length ll in our original graph corresponds to a path from [x, x][x,x] to [y, y][y,y] of the same distance ll in our constructed graph, and vice versa. So how can we find the shortest path from [x, x][x,x] to [y, y][y,y]? The simplest way is to use 0-1 BFS with source vertex [x, x][x,x], which is possible since all our edges are of weight 00 or 11. Alternatively, we can use Dijkstra’s but that will result in an extra log factor in the time complexity.

Time complexity

Remember that the compression phase took O(M\log M)O(MlogM) time. Now, as building a segment tree of size KK requires O(K)O(K) edges our initial graph consisted of O(K)O(K) edges as well, or equivalently O(M)O(M) edges. Then, for each ii we divided the segments [a_i, b_i][a
i
​
,b
i
​
] and [c_i, d_i][c
i
​
,d
i
​
] into their O(\log M)O(logM) subsegments, and we defined an edge from each subsegment of [a_i, b_i][a
i
​
,b
i
​
] to each of [c_i, d_i][c
i
​
,d
i
​
].

This means that for every ii, we inserted O(\log^2 M)O(log
2
M) edges in our graph. The total edges in our graph becomes O(M + M\log^2M)O(M+Mlog
2
M), and it is not difficult to see that the time taken was O(M\log^2M)O(Mlog
2
M) as well. Further, as 0-1 BFS takes O(V + E)O(V+E) time, this part also takes O(M\log^2M)O(Mlog
2
M) time. Hence, our final time complexity is O(M\log^2M)O(Mlog
2
M).

Further reducing the time complexity

We can further reduce the time complexity to O(M\log M)O(MlogM) very simply. For each valid ii, instead of drawing an edge from each subsegment of [a_i, b_i][a
i
​
,b
i
​
] to each of [c_i, d_i][c
i
​
,d
i
​
], we can define an extra vertex v_iv
i
​
and draw an edge from each subsegment of [a_i, b_i][a
i
​
,b
i
​
] to v_iv
i
​
, followed by an edge from v_iv
i
​
to each subsegment of [c_i, d_i][c
i
​
,d
i
​
]. As before, we draw these edges of weight 1. Note that we only used O(\log M)O(logM) edges here instead of O(\log M^2)O(logM
2
) previously.

With this modification, we can easily show that any edge from pp to qq in the original graph now corresponds to a path from segment [p, p][p,p] to [q, q][q,q] of distance 22 in our new graph. Indeed, the only change in the previous proof is that the distance of [a'_i, b'_i][a
i
′
​
,b
i
′
​
] to [c'_i, d'_i][c
i
′
​
,d
i
′
​
] is now 22, but otherwise the same path exists. Thus, we can run 0-1 BFS as before from [x, x][x,x] to [y, y][y,y] in this graph, but divide by 22 to get the final answer. Since the graph now consists of only O(M\log M)O(MlogM) edges, our final time complexity is O(M\log M)O(MlogM) as well.


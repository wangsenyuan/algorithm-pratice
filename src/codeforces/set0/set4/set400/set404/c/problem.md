Valera had an undirected connected graph without self-loops and multiple edges consisting of n vertices. The graph had
an interesting property: there were at most k edges adjacent to each of its vertices. For convenience, we will assume
that the graph vertices were indexed by integers from 1 to n.

One day Valera counted the shortest distances from one of the graph vertices to all other ones and wrote them out in
array d. Thus, element d[i] of the array shows the shortest distance from the vertex Valera chose to vertex number i.

Then something irreparable terrible happened. Valera lost the initial graph. However, he still has the array d. Help him
restore the lost graph.

Input
The first line contains two space-separated integers n and k (1 ≤ k<n ≤ 105). Number n shows the number of vertices in
the original graph. Number k shows that at most k edges were adjacent to each vertex in the original graph.

The second line contains space-separated integers d[1], d[2], ..., d[n] (0 ≤ d[i]<n). Number d[i] shows the shortest
distance from the vertex Valera chose to the vertex number i.

Output
If Valera made a mistake in his notes and the required graph doesn't exist, print in the first line number -1.
Otherwise, in the first line print integer m (0 ≤ m ≤ 106) — the number of edges in the found graph.

In each of the next m lines print two space-separated integers ai and bi (1 ≤ ai, bi ≤ n; ai ≠ bi), denoting the edge
that connects vertices with numbers ai and bi. The graph shouldn't contain self-loops and multiple edges. If there are
multiple possible answers, print any of them.

### ideas

1. 几个观察
2. d[i] = 0 的只能有一个节点，它就是root，
3. 然后d[i] = 1的挂在root上
4. 下一步处理d[i] = 2...
5. 但是必须保证挂在一个节点上的度数 <= k
6. got
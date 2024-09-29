The clique problem is one of the most well-known NP-complete problems. Under some simplification it can be formulated as follows. Consider an undirected graph G. It is required to find a subset of vertices C of the maximum size such that any two of them are connected by an edge in graph G. Sounds simple, doesn't it? Nobody yet knows an algorithm that finds a solution to this problem in polynomial time of the size of the graph. However, as with many other NP-complete problems, the clique problem is easier if you consider a specific type of a graph.

Consider n distinct points on a line. Let the i-th point have the coordinate xi and weight wi. Let's form graph G, whose vertices are these points and edges connect exactly the pairs of points (i, j), such that the distance between them is not less than the sum of their weights, or more formally: |xi - xj| ≥ wi + wj.

Find the size of the maximum clique in such graph.

### ideas
1. 题解是把 (x, w)看作是一个以w为半径，(x, 0)为中心的圆
2. 题目的条件(x[i] - x[j]) >= w[i] + w[j]就是找出最多的这样的圆，且没有相交的地方（除了切点）
3. 那么这样子看的话，就按照右端点升序排，每次选择最靠左的圆
4. 这样子的好处是，假设右边还有更多的选择的时候，现在越靠左，对后续答案的限制越小
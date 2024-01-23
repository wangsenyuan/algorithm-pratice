Given a tree consisting of 𝑛
vertices. A tree is a connected undirected graph without cycles. Each edge of the tree has its weight, 𝑤𝑖
.

Your task is to count the number of different graphs that satisfy all four conditions:

The graph does not have self-loops and multiple edges.
The weights on the edges of the graph are integers and do not exceed 𝑆
.
The graph has exactly one minimum spanning tree.
The minimum spanning tree of the graph is the given tree.
Two graphs are considered different if their sets of edges are different, taking into account the weights of the edges.

The answer can be large, output it modulo 998244353
.

### thoughts

1. 这个tree是唯一的spanning tree
2. 任意两个点（u, v) 它们之间的weight 必须是最小的吗？
3. 它们的weight > max(weight in the path from u to v)
4. 需要知道任意两点间的max path weight, 假设为x
5. 那么就是 sum of (s - x)
6. 计算贡献
7. 从边最大到最小开始算起，然后分成两边，继续
8. 然后就是怎么构造这个以边为节点的树？
9. 边要按照最大到最小进行排序，然后计算出u的sz，v的sz
10. 
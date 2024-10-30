You've got a undirected graph G, consisting of n nodes. We will consider the nodes of the graph indexed by integers from 1 to n. We know that each node of graph G is connected by edges with at least k other nodes of this graph. Your task is to find in the given graph a simple cycle of length of at least k + 1.

A simple cycle of length d (d > 1) in graph G is a sequence of distinct graph nodes v1, v2, ..., vd such, that nodes v1 and vd are connected by an edge of the graph, also for any integer i (1 ≤ i < d) nodes vi and vi + 1 are connected by an edge of the graph.


### ideas
1. 先放置节点1，然后把和他相连的k个节点都放进去
2. 然后从1开始，开始，在这k个节点中，寻找一个，和前一个节点没有连线的点，比如u
3. 后一个节点是v，但是（u, v)没有链接，那么肯定存在一个点w（因为u的deg也至少是k）
4. 然后把u的剩余的节点加入进来
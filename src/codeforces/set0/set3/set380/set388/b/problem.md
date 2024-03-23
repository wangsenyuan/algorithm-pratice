Fox Ciel wants to write a task for a programming contest. The task is: "You are given a simple undirected graph with n
vertexes. Each its edge has unit length. You should calculate the number of shortest paths between vertex 1 and vertex
2."

Same with some writers, she wants to make an example with some certain output: for example, her birthday or the number
of her boyfriend. Can you help her to make a test case with answer equal exactly to k?

Input
The first line contains a single integer k (1 ≤ k ≤ 109).

Output
You should output a graph G with n vertexes (2 ≤ n ≤ 1000). There must be exactly k shortest paths between vertex 1 and
vertex 2 of the graph.

The first line must contain an integer n. Then adjacency matrix G with n rows and n columns must follow. Each element of
the matrix must be 'N' or 'Y'. If Gij is 'Y', then graph G has a edge connecting vertex i and vertex j. Consider the
graph vertexes are numbered from 1 to n.

The graph must be undirected and simple: Gii = 'N' and Gij = Gji must hold. And there must be at least one path between
vertex 1 and vertex 2. It's guaranteed that the answer exists. If there multiple correct answers, you can output any of
them.

### ideas

1. 如果k=1, 非常直接
2. 如果k=2, 也比较清楚
3. 如果k=3呢？ 那么这时候距离至少是多少呢？至少要是2， 1-3-2, 1-4-2, 1-5-2
4. 理论上，只要在1-2中间放入k个点，即可，但是这样子n至少要k+2个
5. 让k=a * b * c * d, 1...3....4....5...2
6. 在1.。。3中间放入a个点，3.。。4中间放入b个点。。。。
7. 如果k是个质数，上面这个策略就不对了
8. 所以，这个时候，还要加一条路径
9. 假设k=1 + 2 + 4.... + 2 ** i
10. 2 ** i 可以用2 * 2 * 2 * 2... i个2相乘，不考虑节点，需要2 * i个点，中间的节点i-1个
11. 找到最大的i，那么这个路径上需要i+2个汇点（1 + 2 + i-1) 再加2 * i个点
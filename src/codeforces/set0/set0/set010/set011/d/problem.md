Given a simple graph, output the number of simple cycles in it. A simple cycle is a cycle with no repeated vertices or edges.

Input
The first line of input contains two integers n and m (1 ≤ n ≤ 19, 0 ≤ m) – respectively the number of vertices and edges of the graph. Each of the subsequent m lines contains two integers a and b, (1 ≤ a, b ≤ n, a ≠ b) indicating that vertices a and b are connected by an undirected edge. There is no more than one edge connecting any pair of vertices.

### ideas
1. 估计是没有理解这个题目
2. 在由4个点组成的完全图中，为啥会有3个长度为4的cycle呢？
3. (1, 2, 3, 4) 
4. (1, 3, 2, 4) (1 2 4 3) 和这个似乎是等价的
5. (1, 4, 2, 3)
6. 如果用state来表示一条path，信息似乎是不够的
7. 比如 1 4 2 和 1 2 4 其实是两条路径
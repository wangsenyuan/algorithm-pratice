The Arms Factory needs a poster design pattern and finds Kevin for help.

A poster design pattern is a bipartite graph with 2𝑛
 vertices in the left part and 𝑚
 vertices in the right part, where there is an edge between each vertex in the left part and each vertex in the right part, resulting in a total of 2𝑛𝑚
 edges.

Kevin must color each edge with a positive integer in the range [1,𝑛]
. A poster design pattern is good if there are no monochromatic cycles∗
 in the bipartite graph.

Kevin needs your assistance in constructing a good bipartite graph or informing him if it is impossible.

### ideas
1. 什么情况下会产生cycle呢？
2. L1 <-> R1 <-> L2 <-> R2 <-> L1 (4条边组成的)
3. L1 - R1 - L2 - R2 - L3 - R3 - L1 (6条边组成的)
4. 左边x个点, 右边y个点, 组成 2 * max(x, y) 条边组成 (也能组成一个环)
5. 且x,y都必须是偶数，这样才能回的去
6. so?
7. 如果有2个R的节点，R1/R2, 每个节点连接了L1,L2...L2n
8. 如果m <= n 那比较简单的，只要给每个m出来的点使用不同的编号就可
9. 如果m > 2 * n, 是不是不可能有答案？
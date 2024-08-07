She does her utmost to flawlessly carry out a person's last rites and preserve the world's balance of yin and yang.

Hu Tao, being the little prankster she is, has tried to scare you with this graph problem! You are given a connected undirected graph of 𝑛
 nodes with 𝑚
 edges. You also have 𝑞
 queries. Each query consists of two nodes 𝑎
 and 𝑏
.

Initially, all edges in the graph have a weight of 0
. For each query, you must choose a simple path starting from 𝑎
 and ending at 𝑏
. Then you add 1
 to every edge along this path. Determine if it's possible, after processing all 𝑞
 queries, for all edges in this graph to have an even weight. If so, output the choice of paths for each query.

If it is not possible, determine the smallest number of extra queries you could add to make it possible. It can be shown that this number will not exceed 1018
 under the given constraints.

A simple path is defined as any path that does not visit a node more than once.

An edge is said to have an even weight if its value is divisible by 2
.

### ideas. 
1. 如果这个graph是一个tree，就非常简单了。
2. 但是这是一个graph，咋搞呢？
3. 我们期望的结果是每条边，都被cover到偶数次
4. 考虑两个query (a, b) (c, d)
5. 如果(c, d) 能够包含 (a, b)是不是就是对(a, b)来说是good的？
6. 这样还不行呐，因为(c, a) (b, d) 还是一次
7. 所以应该是 (a, b), (b, d) (d, e) ... (x, y)
8. 如果存在两个额外的 (a, y) 那么就是good的
9. 除了这种case，有其他的解吗？
10. 但是这个query本身也不大好处理， 比如对于路径 (a, b)
11. 可以拆解成(a, c) (c, b)
12. 也可以拆解成（a, d) (d, b)
13. 如果此时存在两个query (a, d) (d, b)
14. 那么就应该按照d这种case进行拆解
15. 算了，看看solution。能不能理解

### solution

Let 𝑓𝑣
 be the number of times 𝑣
 appears in the 𝑞
 queries. If 𝑓𝑣
 is odd for any 1≤𝑣≤𝑛
, then there does not exist an assignment of paths that will force all even edge weights. To see why, notice that one query will correspond to exactly one edge adjacent to 𝑣
. If an odd number of paths are adjacent to 𝑣
, this implies that at least one edge adjacent to 𝑣
 will have an odd degree.

It turns out that this is the only condition that we need to check. In other words, if 𝑓𝑣
 is even for all 𝑣
, then there will exist an assignment of paths that will force all edge weights to be even.

Let's assume all 𝑓𝑣
 is even. We can find a solution by doing the following: take any spanning tree of the graph and assign each query to be the path from 𝑎
 to 𝑏
 in this tree.

An intuitive way of thinking about this is the following. Consider the case if the spanning tree is a line. Then each query becomes a range and we're checking if all points in this range are covered an even number of times. For all points to be covered an even number of times, every point should occur an even number of times in the queries. To generalize this to a tree, when the first path 𝑎1
 to 𝑏1
 is incremented, in order to make these values even again, we will need later paths to also overlap the segment from 𝑎1
 to 𝑏1
. One way this can be done is if we use two paths 𝑎1
 to 𝑐
 and 𝑐
 to 𝑏1
. Notice that even if a new segment that makes the 𝑎1
 to 𝑏1
 path even makes some other edges odd, the later queries will always fix these edges.
You are given a tree with 𝑛
 nodes. In the beginning, 0
 is written on all edges. In one operation, you can choose any 2
 distinct leaves 𝑢
, 𝑣
 and any real number 𝑥
 and add 𝑥
 to values written on all edges on the simple path between 𝑢
 and 𝑣
.

For example, on the picture below you can see the result of applying two operations to the graph: adding 2
 on the path from 7
 to 6
, and then adding −0.5
 on the path from 4
 to 5
.


Is it true that for any configuration of real numbers written on edges, we can achieve it with a finite number of operations?

Leaf is a node of a tree of degree 1
. Simple path is a path that doesn't contain any node twice.

### ideas
1. 肯定有一个简单的性质，是什么呢～
2. 假设最后的结果是 e1, e2, .... en-1
3. 那么就是选择任何两个节点x, y => f(x, y)
4. f(x1, x2) + f(x2, x3) + ... f(xi, x1) = e1 (如果这个叶子节点的边没有经过e1, 那么就为0)
5. 。。。。
6. 所以这里是n-1个等式，要解n-1个等式，需要至少n个方程
7. 也就是说 m * (m - 1) / 2 >= n 即可
8. 3个节点是 2 * (2 - 1) / 2 >= 2 不成立
Vanya has a graph with 𝑛
 vertices (numbered from 1
 to 𝑛
) and an array 𝑎
 of 𝑛
 integers; initially, there are no edges in the graph. Vanya got bored, and to have fun, he decided to perform 𝑛−1
 operations.

Operation number 𝑥
 (operations are numbered in order starting from 1
) is as follows:

Choose 2
 different numbers 1≤𝑢,𝑣≤𝑛
, such that |𝑎𝑢−𝑎𝑣|
 is divisible by 𝑥
.
Add an undirected edge between vertices 𝑢
 and 𝑣
 to the graph.
Help Vanya get a connected∗
 graph using the 𝑛−1
 operations, or determine that it is impossible.

 ### ideas
 1. x = 操作的序号 （1， 2，.... n - 1)
 2. 那就有点麻烦了 a[u] - a[v] 可以整除 n - 1, 那么可以把它们记录在[n-1]的队列中
 3. 假设 u/v可以在i时连接，自然也可以在i的因数j时连接，这时应该尽可能的让u/v在i时被连接
 4. 这是因为，在j时，让给必须在j时进行连接的边
 5. 先排个序，可以去掉绝对值
 6. (a[u] - a[v]) % x = 0, 
 7. a[u] % x = a[v] % x 对于x同余
 8. 也就是说在第x步时，可以连接那些需要同余的节点
 9. 假设在第x步时选择了u, v, 那么为啥必须是u，v？或者采取什么策略来选择？
 10. 假设存在3个点u, v, w, a[u] % x = a[v] %x, a[u] % y = a[w] % y
 11. 但是这里的关键是假设 a[u] % x = a[v] % x, 且 a[u] % y = a[v] % y
 12. 那么是不是尽量的延迟A[u], a[v]的连接是不是更好
 13. 怎么连接也是个问题
 14. 这里是不是没有关系，只要能添加一条新的边，就添加上去就可以了？
 15. 根据鸽子洞原理，肯定是有解的

### solutions

Note that we have only 𝑛−1
 edges, so after each operation, the number of connectivity components must decrease by 1
.

Since the order of operations is not important, we will perform the operations in reverse order. Then after the operation with number 𝑥
, there will be 𝑥+1
 connectivity components in the graph. For each component, let's take any of its vertices 𝑣
 and look at the number that corresponds to it. Note that we have chosen 𝑥+1
 numbers, so by the pigeonhole principle, some two numbers will be equal modulo 𝑥
. This means that we find two vertices 𝑢
 and 𝑣
 from different components such that |𝑎𝑢−𝑎𝑣|
 is a multiple of 𝑥
. By drawing an edge between 𝑢
 and 𝑣
, we will achieve what we want — the number of connectivity components will become 1
 less.

Now it remains to go through the operations in reverse order and print the answer.


You are given a tree, which consists of 𝑛
 vertices. Recall that a tree is a connected undirected graph without cycles.

Example of a tree.
Vertices are numbered from 1
 to 𝑛
. All vertices have weights, the weight of the vertex 𝑣
 is 𝑎𝑣
.

Recall that the distance between two vertices in the tree is the number of edges on a simple path between them.

Your task is to find the subset of vertices with the maximum total weight (the weight of the subset is the sum of weights of all vertices in it) such that there is no pair of vertices with the distance 𝑘
 or less between them in this subset.


 ### ideas
 1. dp + tree, so hard
 2. n <= 200, k <= 200
 3. dp[u][i] 表示在子树u中，且距离u的距离>= i的最优值
 4. dp[u][i] = sum(dp[v][i-1] when i > 0)
 5. 应该还不够，比如在距离i+1处的，也是可以算到i处的
 6. dp[u][i] = max(dp[u][i], dp[u][i+1])
 7. 考虑是一条直线， dp[u] = max(dp[u-1], a[i] + seg max of dp[...u - k])
 8. dp[u]是最优的结果， dp[u] = max(a[u] + 距离为k的dp[c], sum of dp[v])
 9. 但是tree不对， 因为没法把两颗子树的对应层级直接加起来
 10. 比如k = 2， 那么如果选择了第一颗子树v的第一层，就不能选择第二课子树的第一层
 11. dp[u][k] 表示在处理子树u时，它的parent[k]被选中时的最优解
 12. dp[u][0] 表示u自己被选中了， dp[u][1]表示它的直接上级被选中了， dp[u][k]表示第k个上级被选中了
 13. dp[u][i] = sum of dp[v][i+1] + a[u] 如果 i = 0
 14. 但是问题出在 i = k 的时候，一方面来自上层的压力没有了，但是子树之间开始出现关联了
 15. 所以，问题的关键还是两颗子树之间的关联要怎么处理
 16. 还是选中某个点，它肯定被选中了
 17. 然后删除它距离范围内的所有的点
 18. 那么剩下的那些，又变成了一个子问题
 19. dp[u][k] 可以描述这个子问题
 20. 可不可以这样，新增一个图，就是把所有距离k以内的点，连一条线，然后做个二部图？似乎不大对
 21. dp[u][k] = 。。。。
 22. 完全没有想法～～～～
 23. dp[u][k] 表示在u子树中，距离u最近的距离为k的点被选中时的最优解
 24. dp[u][k] = dp[v][k+1] (只有一个子节点)
 25.   dp[u][min(i, j) + 1]   =     dp[x][i] + dp[y][j] (i + j > k)
 26. dp[u][0] = a[u] + dp[v][k+1]
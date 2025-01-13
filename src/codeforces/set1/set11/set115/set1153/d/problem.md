Now Serval is a junior high school student in Japari Middle School, and he is still thrilled on math as before.

As a talented boy in mathematics, he likes to play with numbers. This time, he wants to play with numbers on a rooted tree.

A tree is a connected graph without cycles. A rooted tree has a special vertex called the root. A parent of a node 𝑣
 is the last different from 𝑣
 vertex on the path from the root to the vertex 𝑣
. Children of vertex 𝑣
 are all nodes for which 𝑣
 is the parent. A vertex is a leaf if it has no children.

The rooted tree Serval owns has 𝑛
 nodes, node 1
 is the root. Serval will write some numbers into all nodes of the tree. However, there are some restrictions. Each of the nodes except leaves has an operation max
 or min
 written in it, indicating that the number in this node should be equal to the maximum or minimum of all the numbers in its sons, respectively.

Assume that there are 𝑘
 leaves in the tree. Serval wants to put integers 1,2,…,𝑘
 to the 𝑘
 leaves (each number should be used exactly once). He loves large numbers, so he wants to maximize the number in the root. As his best friend, can you help him?


 ### ideas
  1. 假设当前节点的操作是max, 那么我们只需要保证从一个子节点上能够得到max, 其他节点没有关系
  2. dp[u] = max(dp[v]) + 其他的放弃即可
  3. 如果选择了v, 那么无论它是max/min，都应该把后面的数分配给它
  4. 如果当前节点是min，那么就应该保证所有节点的结果应该足够的大
  5. 如果有一个子节点是max(那么就把最小的那些分配给它)
  6. 如果有一个子节点是min，那么就把最大的那些分配给它
  7. dp[u] = 在给这个子树[1...sz[u]]节点后的最大值
  8. dp[u] max  = max(dp[v] + sz[u] - sz[v]) 
  9. dp[u] min =  这个比较难搞了，这个是不是要二分呢？一个点在于，分配给v子树的数字越大，结果肯定越大
  10.    假设期望得到x，
  11. 按照dp[v]升序排，dp[v] < x, 那么给它分配的一定是连续的一段吗？
  12. got 了，如果sz[v] = 10, dp[v] = 5, 那么只要分配5个数给它，就可以增加dp[v]
  13. 因为 dp[v]表示的是在1...sz[v]区间的。能够得到的最大的第几个数，那么要增加它，必须同时增加dp[v]...sz[v]的数
  14. 
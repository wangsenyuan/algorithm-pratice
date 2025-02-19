You are given a tree∗
 with 𝑛
 nodes and values 𝑙𝑖,𝑟𝑖
 for each node. You can choose an initial value 𝑎𝑖
 satisfying 𝑙𝑖≤𝑎𝑖≤𝑟𝑖
 for the 𝑖
-th node. A tree is balanced if all node values are equal, and the value of a balanced tree is defined as the value of any node.

In one operation, you can choose two nodes 𝑢
 and 𝑣
, and increase the values of all nodes in the subtree†
 of node 𝑣
 by 1
 while considering 𝑢
 as the root of the entire tree. Note that 𝑢
 may be equal to 𝑣
.

Your goal is to perform a series of operations so that the tree becomes balanced. Find the minimum possible value of the tree after performing these operations. Note that you don't need to minimize the number of operations.

### ideas
1. 最后的value会是除去 l/r外的其他值吗？
2. 假设没有任何限制，可以给节点设置任何小于x的值，那么有没有办法通过操作，得到x呢？
3. 有一个情况是无法得到x的，就是所有的叶子节点都是x，但是内部节点小于x，那么无论选择哪个分支，都会增加某个叶子节点的值，从而超过x
4. 考虑某个内部节点u, 考虑它和所有叶子节点的关系
5. 考虑一种中心节点，加一圈的叶子节点的结构
6. 假设最后的结果是x，那么所有节点的操作次数 = x - a[u]
7. 有两种方案，每个叶子节点单独变化，或者除了某个叶子节点，所有其他的节点单独变化（相当于某个叶子节点单独变化）
8. 比如我想增加中心节点，那么可以全部增加dx, 然后其他的所有节点，再而外操作
9. 首先，选择u/v是同一个似乎没啥好处，这是因为这个操作，相当于所有的节点+1，但是这个操作可以表明，最后的结果是符合2分的，就是如果x成立，那么x+1也是成立的
10. u/v，对v进行+1，那么相当于 v/u对u子树进行 -1
11. 也就是说，选择任何一个节点，可以对它的任何一个子树，单独进行 +1/-1操作（-1操作，就是对除v外的，其他子树进行+1）
12. 对于任何一个子树，固定节点u的情况下，是否可以判断出它最小能达到的数字？
13. 那么结果貌似是，确定root u后，max(dp[u][v], a[u])
14. dp[v]就是以v为root时，这整颗子树能达到的最小相同值
15. dp[u] = max(dp[v], a[u])?
16. 不大对啊。 如果 a[u]是这颗子树的最大值，没有问题。 可以先选择u, 再选择某个子树v，让那个子树增加到a[u]
17. 但是如果a[u]不是这棵树的最大值，比如它其实是最小值，那么就有问题（任何选择u的尝试，都会增加整颗子树u， 从而a[u]还是最小值）
18. 那么此时意味着，必须选择a[u]某个子树v，反过来选，直到把a[u]变成最大值，才可以
19. 那么此时就选择u子树中的最大值所在的v，
20. 如果只有一个这样的v，那么u可以在有限次操作后，变成a[v], 如果有两个子树（以上）有最大值v，就麻烦了
21. 这是因为增加u的同时，也在增加它的某个最大值的子节点
22. 有点麻烦
23. 还有一个点，这里给定一个范围的作用是什么？它怎么影响这个结果？
24. 假设某个内部节点u，它能达到的最小值 = 至少是叶子节点的最大值
You are given two integers 𝑛
 and 𝑑
. You need to construct a rooted binary tree consisting of 𝑛
 vertices with a root at the vertex 1
 and the sum of depths of all vertices equals to 𝑑
.

A tree is a connected graph without cycles. A rooted tree has a special vertex called the root. A parent of a vertex 𝑣
 is the last different from 𝑣
 vertex on the path from the root to the vertex 𝑣
. The depth of the vertex 𝑣
 is the length of the path from the root to the vertex 𝑣
. Children of vertex 𝑣
 are all vertices for which 𝑣
 is the parent. The binary tree is such a tree that no vertex has more than 2
 children.

You have to answer 𝑡
 independent test cases.

 ### ideas
 1. n个节点，最大的depth的sum = n - 1 + n - 2 + ... + 1 + 0 = (n - 1) * n / 2
 2. 最小的depth的sum = 1 * 2 + 2 * 4 + 3 * 8 
 3. 在这中间的都是可以的
 4. 假设m层节点点，可以组成一个完整的二叉树 f(m) = 他的depth 之和
 5. f(m) = f(m-1) + (m-1) * (1 << (m - 1))
 6. f(1) = 0
 7. f(2) = f(1) + 1 * (1 << 1)
 8. f(3) = f(2) + 2 * (1 << 2)
 9. f(4) = f(3) + 3 * (1 << 3)
 10. 找到最大的f(m) <= d
 11. d - f(m) 
 12. 假设还需要x个的节点，d - x * m
 13. 这里d比较小，所以可以用 dp[d][x]来表示，是否能用x个节点来得到sum 为d的的tree
 14. 不大对
 15. 上面方法构造的树，depth的变化太剧烈了，不能保证所有的结构都能取到
 16. 假设这里有两部分组成，上/下
 17. 上面的树，有n1, d1组成， 下面的树有 n2, d2
 18. 那么有n = n1 + n2
 19. 但是 d = d1 + d2 + n2 * x (x是下面的树连接的节点的层级 + 1)
 20. 迭代n1/d1, 那么x是可以算出来的
 21. 但是此时如何保证一定存在一个x-1的层级的点来对接呢？
 22. 可以倒也是可以弄的。但感觉似乎有点麻烦
 23. 这里，整体还是应该符合这样一个结构，就是存在一个头顶，是完全二叉树
 24. 然后剩余的，分别挂在叶子节点上，
 25. (s1, n1), (s2 + n2 * m, n2), (s3 + n3 * m, n3)...
 26. 是这样一个结构
 27. dp[s][n]是可行的，当且仅当存在这样一个划分的时候
 28. s1, n1是定的(F[m], 1 << m - 1)
 29. 假设使用n个节点组成了sum的树，且目前最深的深度是x
 30. 如果把x-1深度的一个节点移动到x+1深度，那么变化是sum + 2 (-(x - 1) + (x + 1) = 2)
 31. 感觉要换一个思路，假设一共有深度是x （0, 1, 2, ... x)
 32. 且每个深度有w[i]个节点
 33. 那么有 sum w[i] = n
 34. sum(w[i] * i) = d
 35. 且 w[i] 有些限制（w[i] <= 2 * w[i-1])
 36. dp[i][s][j]= 表示前i层使用了j个节点，且depth和为s时，是否ok
 37. 那么在第i层使用y个节点时的变化 = dp[i-1][s - y * i][j - y]
 38. 但这里有个问题，这个y的范围如何限定（它不能超过上一层节点数的2倍）
 39. 而且也不能太小。如果增加上上一层节点使用的个数，那么状态似乎又太多了。（现在也有问题，跑不出来的）
 40. j是有上限的，如果是连续的多层的话，j的上限就是 2048 (这个时候是11层)
 41. 先使n个节点构造一颗完全平衡树（最下面一层有可能少了几个）
 42. 如果此时sum > d => false 
 43. 如果 sum == d => done
 44. sum < d, 那么此时需要调整这棵树
 45. 继续增加节点,知道 sum > d
 46. 然后，开始删除节点。
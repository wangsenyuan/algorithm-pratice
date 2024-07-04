Julia's 𝑛
 friends want to organize a startup in a new country they moved to. They assigned each other numbers from 1 to 𝑛
 according to the jobs they have, from the most front-end tasks to the most back-end ones. They also estimated a matrix 𝑐
, where 𝑐𝑖𝑗=𝑐𝑗𝑖
 is the average number of messages per month between people doing jobs 𝑖
 and 𝑗
.

Now they want to make a hierarchy tree. It will be a binary tree with each node containing one member of the team. Some member will be selected as a leader of the team and will be contained in the root node. In order for the leader to be able to easily reach any subordinate, for each node 𝑣
 of the tree, the following should apply: all members in its left subtree must have smaller numbers than 𝑣
, and all members in its right subtree must have larger numbers than 𝑣
.

After the hierarchy tree is settled, people doing jobs 𝑖
 and 𝑗
 will be communicating via the shortest path in the tree between their nodes. Let's denote the length of this path as 𝑑𝑖𝑗
. Thus, the cost of their communication is 𝑐𝑖𝑗⋅𝑑𝑖𝑗
.

Your task is to find a hierarchy tree that minimizes the total cost of communication over all pairs: ∑1≤𝑖<𝑗≤𝑛𝑐𝑖𝑗⋅𝑑𝑖𝑗
.

### ideas
1. 这棵树在中序访问的时候，得到的就是1,2,3....n 
2. 假设i是root，那么[1...i-1]是它的左子树，[i+1...n]是它的右子树
3. dp[i...j]是这段序列组成树的最小值，
4. 那么dp[i][j] = 以k为root， dp[i...k-1] + cost from k + dp[k+1...j] + cost from k 的最小值
5. cost from k = c[i][k] * d[i][k] + c[i+1][k] * d[i+1..k] + ..
6.   c[i][k]这个是已知的，问题是 d[i][k], 
7. 这里因为不知道[i...k-1]的root是哪个，所以也无从知道这个高度是什么
8. 还得变形
9. c[i][j] * d[i][j] = c[i][j] * (d[i] - d[j]) (j是root)
10. = c[i][j] * d[i] - c[i][j] * d[j]
11. 上面的分析不大对。 所有的（i， j）pair 都需要沟通，并不是只有上下级需要沟通
12. 假设由[i..j]段组成的子树，它对整体的贡献是多少呢？
13. 它内部的，假设已经知道了。不去管它
14. 它和外部的呢？假设有一个k
15. c[i][k] * d[i][k] + c[i+1][k] * d[i+1][k] + .... + c[j][k] * d[j][k]
16. c[i][k] * (d[i] - d[r] + d[r][k]) 的距离
17. c[i][k] * (d[i] - d[r]) + c[i][k] * d[r][k]的距离
18. 前半部分完全由这颗子树的结构决定(和k是无关的)
19. 感觉这里搅在一起了，有点乱。缓缓
20. 考虑一个子树u，它的左右子树中的节点的贡献要怎么计算呢？
21. c[i][j] * (d[i] + d[j] - 2 * d[u]) (u是子树根节点)
22. 也就是说在确定了子树，且这个子树的根节点是u的情况下， c[i][j] * (d[i] + d[j]) 
23. 我们希望这个越小越好，
24. c[i][j]是确定的，其实就是d[i] + d[j]越小越好
25. 这里的问题就在于，没有办法同时知道(i, j)
26. 且只有当出现了j的时候
27. 在确定u的时候，对于左子树的root，其实是要求，所有 c[i][j] 的 和最大的那个i最为左子树的root
28. 还是不大通 假设 a, b 是root左边的两个节点
29. s[a][..] > s[b][..], 但是在最优答案中b是root， 而a不是, 通过交换ab是不是能得到更好的答案？
30. 因为b是root，所以在外部的节点到b的距离要小于到a的距离，假设a就是b的一个child
31. 交换后的收益(对外部节点来说) = -s[a][..] + s[b][..] < 0 （所以cost更小了）
32.  但是对于内部来说，b要变成a的新的右子树，原来a的右子树变成了b的左子树
33.  发生变化的，只有a的原来的右子树部分，它们离a的距离增加了，但是到b的距离减少了
34.  +c[x][a] - c[x][b] (x其实就是 a和b中间的节点)
35.   如果选a作b的parent， 2 * c[x][a] + c[x][b], 如果选b作a的parent, c[x][a] + 2 * c[x][b]
36.   所以如果 c[x][a] < c[x][b] 的话， a作为parent更有利
37.   但是同时在子树u外部的节点也有影响～
38. 如果a是b的parent，在子树u外部的到a的距离比到b的距离小1，所以c[y][a] + 2 * c[y][b], 反过来的话就是 2 * c[y][a] + c[y][b]
39. 同样的如果 c[y][a] > c[y][b] 的话，让a作为parent更有利
40. 那这里就有点麻烦了，首先y这个不好确定（它依赖当前的所说的子树的root是什么）然后对于两个a,b （不相邻)
41. 如果 s[y][a] > s[y][b], 且s[x][a] < s[x][b] （y是root外面的节点，x是a，b中间的节点）时a作为root更好
42. s[y][a] - s[x][a] > s[y][b] - s[x][b]时，a更应该高一点
43. 这里还是不大对的。。。再想想
44. 考虑root，应该是什么样的，假设有两个a，b它们做root有什么不同
45. 首先a前面和b后面的，似乎不受影响。只有中间的部分受影响。这个就是上面的如果s[a][x] < s[b][x] 那么a作为root更合理
46. 那么这样是不是就可以递归下去了？
47. 对于[l...r]中间的a和b，对所有的a，b进行比对，找到它们的关系，然后根据这个关系，找到入度为0的节点
48. 且a,b的关系，仅需要根据它们中间的部分进行计算。所以这个是可以提前算好的。
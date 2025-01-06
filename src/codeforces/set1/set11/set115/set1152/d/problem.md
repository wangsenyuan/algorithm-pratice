Neko is playing with his toys on the backyard of Aki's house. Aki decided to play a prank on him, by secretly putting catnip into Neko's toys. Unfortunately, he went overboard and put an entire bag of catnip into the toys...

It took Neko an entire day to turn back to normal. Neko reported to Aki that he saw a lot of weird things, including a trie of all correct bracket sequences of length 2𝑛
.

The definition of correct bracket sequence is as follows:

The empty sequence is a correct bracket sequence,
If 𝑠
 is a correct bracket sequence, then (𝑠)
 is a correct bracket sequence,
If 𝑠
 and 𝑡
 are a correct bracket sequence, then 𝑠𝑡
 is also a correct bracket sequence.
For example, the strings "(())", "()()" form a correct bracket sequence, while ")(" and "((" not.

Aki then came up with an interesting problem: What is the size of the maximum matching (the largest set of edges such that there are no two edges with a common vertex) in this trie? Since the answer can be quite large, print it modulo 109+7
.

### ideas
1. trie是一棵树，一棵树的最大匹配 = 
2. dp[u][0] 在子树u中，u的出边都未被选中的最大值
3. dp[u][1] 表示u的某条出被选中的最大值
4. dp[u][0] = sum(max(dp[v][0], dp[v][1])) (u未被选中，所以对v是没有限制的)
5. dp[u][1] = dp[v][0] + 1 + max(dp[w][0], dp[w][1]) (选择一个子节点v, 其他的求和)
6. 但问题是这个trie 有 pow(2, 1000)条边。
7. 或者没有这么多的边，因为有些边是无效的, 比如 >>> 就是无效的
8. 而且它似乎是个递归结构
9. 比如 f(n) =》 < f(n) > 或者 <> f(n)
10. 
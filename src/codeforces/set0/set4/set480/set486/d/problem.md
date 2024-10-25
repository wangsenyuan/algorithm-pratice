As you probably know, an undirected connected graph with n vertices and n  - 1 edges is called a tree . You are given an integer d and a tree consisting of n vertices. Each vertex i has an associated value a i .

We call a set S of tree vertices admissible if the following conditions are met:

S is a non-empty set.
S is a connected set. In other words, if vertices u and v are contained in S , then all vertices lying on a simple path between u and v must also lie in S .
.
Your task is to count the number of admissible sets. Since the result can be very large, print its remainder modulo 1000000007 ( 10 9  + 7 ).

### ideas
1. 考虑一个简单的情况，一条线
2. 固定r的情况下，l是最左边，满足max(a[l...r]) - min(a[l...r]) <= d 的位置
3. 这种情况下似乎，只需要考虑 abs(a[r] - val(a[l...r-1])) 的情况就可以了。
4. 如果没有d的限制，可以计算出所有的集合数量吗？
5. dp[u]表示子树u， 且包含u的的集合数 = sum of dp[v] v是u的子节点 + 1(单独的u节点)
6.  + dp[v] * (sum of fp[c] + 1) fp[..] dp[v]的前缀和
7. res = sum of dp[u]
8. 然后从里面除掉 bad的？
9. 如果按照a[u]从小到大处理呢？
10. 加入a[u]时，就是包含a[u]的联通块的，并把那些a[v] < a[u] - d的都给关闭掉（就产生了一个只包含a[u]）的联通块
11. 好像这样子，不会重复计数，因为每次至少有一个新的节点被考虑进去
12. 然后就是如何计算这个联通块的计数，就是上面那个子树方案
13. got
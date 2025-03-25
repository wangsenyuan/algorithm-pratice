You are given a rooted tree∗
 containing 𝑛
 vertices rooted at vertex 1
. A pair of vertices (𝑢,𝑣)
 is called a good pair if 𝑢
 is not an ancestor†
 of 𝑣
 and 𝑣
 is not an ancestor of 𝑢
. For any two vertices, dist(𝑢,𝑣)
 is defined as the number of edges on the unique simple path from 𝑢
 to 𝑣
, and lca(𝑢,𝑣)
 is defined as their lowest common ancestor.

A function 𝑓(𝑢,𝑣)
 is defined as follows.

If (𝑢,𝑣)
 is a good pair, 𝑓(𝑢,𝑣)
 is the number of distinct integer values 𝑥
 such that there exists a non-degenerate triangle‡
 formed by side lengths dist(𝑢,lca(𝑢,𝑣))
, dist(𝑣,lca(𝑢,𝑣))
, and 𝑥
.
Otherwise, 𝑓(𝑢,𝑣)
 is 0
.
You need to find the following value:
∑𝑖=1𝑛−1∑𝑗=𝑖+1𝑛𝑓(𝑖,𝑗).



### ideas
1. p = lca(u, v) and p != u, p != v
2. a = dist(p, u), b = dist(p, v) and a <= b
3. 那么 b - a < x < a + b
4. cnt = b + a - (b - a) - 1 = 2 * a - 1 （也就是小的边的长度 * 2 - 1）
5. 固定p和u的情况下，要知道有多少条个点v, dist(p, v) >= dist(p, u)
6. 考虑处理p节点
7. 迭代它的子节点u，在每个子节点上维护这样一个结构freq[d], 表示和u的距离为d的子节点的数量
8. 并且已经知道目前为止在节点p上的这样一个结构
9. 那么 (2 * (d + 1) - 1) * freq[d] * suf[d] (suf 表示离p距离至少为d的子节点的数量)
10. 逻辑是成立的。但问题在与怎么把v的信息合并到p里面去呢？
11. 所有v的freq[d] 要变成 freq[d+1], 还要更新suf
12. 如果不使用d, 而是使用h(高度)， 那么就是把小的高度合并到高的高度上去？
13. 先选择最高的那颗子树，然后把其他的都合并进去？ 
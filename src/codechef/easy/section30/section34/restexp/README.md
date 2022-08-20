### [codechef practice problem expend restarant](https://www.codechef.com/problems-old/RESTEXP)

### 思路
> 1. constraints

   ```
    1 <= N <= 30
    1 <= C <= 30
    1 <= D <= 30
    -1000 <= Si <= 1000
   ```

> 2. 基本的想法是在S越大的节点上派驻一个chef即可

> 3. 先将chef通过transfer的方式运送到节点上，然后build
 
> 4. 需要在D-1天内，运送chef到对应的节点去

> 5. 使用dp[u][c][d]表示在d天内运送c个chef到节点u上所能得到的最大profit, p是u的parent node

  ```
    dp[u][c][d] = max(dp[p][c+1][d-1] + S[p], dp[p][c][d-1])
  ```

  > * 有个问题就是d不一定是连续的，就是可以先处理这个节点，但是下个节点可以从另外一个地方开始
  > * 所以上面的表达式不能正确的执行, 换一个思路，在d天内，包括transfer，build的情况下，子树u所能获得的最大值是多少

  ```
    dp[p][c][d] = sum(dp[u][cu][du]) + (S[p] when S[p] > 0)

    sum(du) + cnt(u) = d
    sum(cu) = c or sum(cu) + 1 = c 
  ``` 

### explaination

We present two important observations of the problem:

The corresponding graph is a tree, because it is connected and has N-1 edges.
After we move a chef from city a to city b, it is useless to move it to city a again.
Those observations make the problem possible to be solved using DP in tree. To simplify the DP, we first transform the tree into child-sibling representation rooted at node 1. Let DP(n, c, d) be the maximum profit achieved in the subtree rooted at the parent[n], with c chefs available in parent[n], and we are given d days remaining. The recurrence is given in this pseudocode:

DP(n, c, d) = max (
DP(sibling[n], c, d), // doesn’t transfer any chefs from parent[n] to n
max ({DP(child[n], rc, rd) + DP(sibling[n], c - rc, d - 1 - rd) : 1 <= rc <= c and 0 <= rd <= d - 1}) // transfers rc chefs to n and gives rd days
)
The base case is when n is null (sentinel node that is the sibling of the last sibling of a node). We can build a restaurant in parent[n] if c and d are positive, and also if S[parent[n]] is not negative. So, DP(n, c, d) = max(0, S[parent[n]]) if n is sentinel node and c and d are positive.
The solution to the problem is then available in DP(child1 1, C, D). We can then use all information in the DP table to reconstruct an optimal expansion plan.
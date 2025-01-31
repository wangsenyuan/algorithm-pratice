### ideas
1. 考虑确定的k开始
2. f(k) = 以k为root时，访问整棵树的方式数
3. dp(u) 表示访问以u（root = 1）的子树时的方式数
4. dp[u] = nCr(sz[u], sz[v]) * dp[v] * nCr(sz[u] - sz[v], sz[w]) * dp[w] ...
5. 所以是个换根dp
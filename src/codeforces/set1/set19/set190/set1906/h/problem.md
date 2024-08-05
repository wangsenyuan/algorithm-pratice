You meet two new friends who are twins. The name of the elder twin is 𝐴
, which consists of 𝑁
 characters. While the name of the younger twin is 𝐵
, which consists of 𝑀
 characters. It is known that 𝑁≤𝑀
.

You want to call each of them with a nickname. For the elder twin, you want to pick any permutation of 𝐴
 as the nickname. For the younger twin, you want to remove exactly 𝑀−𝑁
 characters from any permutation of 𝐵
. Denote the nicknames of the elder twin and the younger twin as 𝐴′
 and 𝐵′
, respectively.

You want the nicknames to satisfy the following requirement. For each 𝑖
 that satisfies 1≤𝑖≤𝑁
, 𝐵′𝑖
 must be equal to either 𝐴′𝑖
 or the next letter that follows alphabetically after 𝐴′𝑖
 (if such a next letter exists).

Determine the number of different pairs of nicknames (𝐴′,𝐵′)
 that satisfy the requirement. Two pairs of nicknames are considered different if at least one of the nicknames are different. As the result might be large, find the answer modulo 998244353
.

### ideas
1. 一头雾水～
2. 先考虑，a有多少种可能性？
3. a是A的排列组合, C(n, cnt[a]) * C(n - cnt[a], cnt[b]) ....
4. 先选择cnt[a]个位置放置a, 然后cnt[b]个位置放置b....
5. 在确定a的情况下，如何得到b？
6. 假设现在要处理x,(和x匹配的b)
7. 如果这些位置的下面放置x（从b中获取）是符合条件的
8. 或者b这里要放置x+1
9. dp[x][i]表示在放置完[a,b....x],且使用了b中i个x+1的计数
10. dp[x+1][0] = sum dp[x][i] * nCr(cnt_b[x+1] - i, cnt_a[x+1]) 也就是在剩余的cnt_b[x+1] - i 中，选择cnt_a[x+1]个x+1
11. dp[x+1][j] = sum dp[x][i] * nCr(cnt_b[x+1] - i, cnt_a[x+1] - j) * nCr(cnt_b[x+2], j)
12. 
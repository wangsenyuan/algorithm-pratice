The Fair Nut found a string 𝑠
. The string consists of lowercase Latin letters. The Nut is a curious guy, so he wants to find the number of strictly
increasing sequences 𝑝1,𝑝2,…,𝑝𝑘
, such that:

For each 𝑖
(1≤𝑖≤𝑘
), 𝑠𝑝𝑖=
'a'.
For each 𝑖
(1≤𝑖<𝑘
), there is such 𝑗
that 𝑝𝑖<𝑗<𝑝𝑖+1
and 𝑠𝑗=
'b'.
The Nut is upset because he doesn't know how to find the number. Help him.

This number should be calculated modulo 109+7
.

### ideas

1. 只有ab是有用的
2. dp[a] 表示到目前以a结尾的计数
3. dp[b] 表示以b结尾
4. dp[a] += dp[b] 如果遇到a
5. dp[b] += dp[a] 如果遇到b
6. 上面的不大对， 这是因为多个b会计算多次，但其实只应该计算一次（或者0次)
7. dp[i] = 到i为止的以a结尾的计数
8. dp[i+1] = dp[i] 如果 s[i] != a
9. else dp[i+1] = dp[i] + dp[j] j是最近的b的位置？
Ecrade has an integer 𝑥
. He will show you this number in the form of a binary number of length 𝑛
.

There are two kinds of operations.

Replace 𝑥
 with ⌊𝑥/2⌋, where ⌊𝑥/2⌋ is the greatest integer ≤𝑥2
.
Replace 𝑥 with ⌈𝑥/2⌉, where ⌈𝑥/2⌉ is the smallest integer ≥𝑥2
.
Ecrade will perform several operations until 𝑥
 becomes 1
. Each time, he will independently choose to perform either the first operation or the second operation with probability 12
.

Ecrade wants to know the expected number of operations he will perform to make 𝑥
 equal to 1
, modulo 109+7
. However, it seems a little difficult, so please help him!


### ideas
1. dp[i] = 1 + dp[i-1] (如果 s[i]是0)
2. dp[i] = 1 + dp[i-1] / 2 + fp[j] / 2 如果 s[i] = 1, j是最靠近它的左边的s[j] = 0的地方
3. fp[j] = dp[j-1]
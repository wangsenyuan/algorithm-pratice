Let 𝑓
(𝑥
) be the floor of the binary logarithm of 𝑥
. In other words, 𝑓
(𝑥
) is largest non-negative integer 𝑦
, such that 2𝑦
does not exceed 𝑥
.

Let 𝑔
(𝑥
) be the floor of the logarithm of 𝑥
with base 𝑓
(𝑥
). In other words, 𝑔
(𝑥
) is the largest non-negative integer 𝑧
, such that 𝑓(𝑥)𝑧
does not exceed 𝑥
.

You are given 𝑞
queries. The 𝑖
-th query consists of two integers 𝑙𝑖
and 𝑟𝑖
. The answer to the query is the sum of 𝑔
(𝑘
) across all integers 𝑘
, such that 𝑙𝑖≤𝑘≤𝑟𝑖
. Since the answers might be large, print them modulo 109+7
.

### thoughts

1. f(x)是非递减的，且对于x，不超过60
2. 但是g(x)在相同f(x)时是，非递减；
3. 对于相同的给定的f(x) = y, 找到它的下界lx, 和上界rx
4. 算出对应的g
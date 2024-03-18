You have a vertical strip with 𝑛
cells, numbered consecutively from 1
to 𝑛
from top to bottom.

You also have a token that is initially placed in cell 𝑛
. You will move the token up until it arrives at cell 1
.

Let the token be in cell 𝑥>1
at some moment. One shift of the token can have either of the following kinds:

Subtraction: you choose an integer 𝑦
between 1
and 𝑥−1
, inclusive, and move the token from cell 𝑥
to cell 𝑥−𝑦
.
Floored division: you choose an integer 𝑧
between 2
and 𝑥
, inclusive, and move the token from cell 𝑥
to cell ⌊𝑥𝑧⌋
(𝑥
divided by 𝑧
rounded down).
Find the number of ways to move the token from cell 𝑛
to cell 1
using one or more shifts, and print it modulo 𝑚
. Note that if there are several ways to move the token from one cell to another in one shift, all these ways are
considered distinct (check example explanation for a better understanding).

### ideas

1. 考虑一种brute-force的处理dp[n] = dp[n-1] + dp[n-2] + ... dp[1]  (使用sub)
2. dp[n] = dp[n/2] + dp[n/3] + ... dp[1] 使用方式2
3. subtraction这个很好处理，只要维护一个前面的dp的sum即可
4. 但是方式2有点麻烦，比如 5 / 2 = 2, 4 / 2 = 2
5. 2的两倍包括[4..5], 3倍包括[6...7], 其实这个还是连续的
6. fp[i] 表示前面的操作2的计数之和, fp[i] += fp[i-1], dp[i] += fp[i]
7. fp[i]不大对，比如 3 / 2 = 1， 3 / 3 = 1， 意味着dp[1]对fp[3]的贡献其实是2 * dp[1]
8. 7 / 3 = 2, 7 / 2 = 3, 7 / 4 = 1
9. 也就是说dp[1]对于某个i的贡献 = i - i / 2, 其他的好像还是1
10. 2 对于 10 的贡献也是2， 10 / 5 = 2， 10 / 4 = 2
11. 3对于15的贡献是2， 15 / 5 = 3， 15 / 4 = 12
12. 对于数字n来说，n / n = 1.... n / (h + 1) = 1
13. n / h = 2, n / (i+1)= 2
14. n / i = 3, n / (j + 1) = 3
15. ... 
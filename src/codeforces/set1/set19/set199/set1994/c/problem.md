Yaroslav is playing a computer game, and at one of the levels, he encountered 𝑛
 mushrooms arranged in a row. Each mushroom has its own level of toxicity; the 𝑖
-th mushroom from the beginning has a toxicity level of 𝑎𝑖
. Yaroslav can choose two integers 1≤𝑙≤𝑟≤𝑛
, and then his character will take turns from left to right to eat mushrooms from this subsegment one by one, i.e., the mushrooms with numbers 𝑙,𝑙+1,𝑙+2,…,𝑟
.

The character has a toxicity level 𝑔
, initially equal to 0
. The computer game is defined by the number 𝑥
 — the maximum toxicity level at any given time. When eating a mushroom with toxicity level 𝑘
, the following happens:

The toxicity level of the character is increased by 𝑘
.
If 𝑔≤𝑥
, the process continues; otherwise, 𝑔
 becomes zero and the process continues.
Yaroslav became interested in how many ways there are to choose the values of 𝑙
 and 𝑟
 such that the final value of 𝑔
 is not zero. Help Yaroslav find this number!

Input
Each test consists of multiple test cases. The first line contains an integer 𝑡
 (1≤𝑡≤104
) — the number of test cases. Then follows the description of the test cases.

The first line of each test case contains two integers 𝑛
, 𝑥
 (1≤𝑛≤2⋅105,1≤𝑥≤109
) — the number of mushrooms and the maximum toxicity level.

The second line of each test case contains 𝑛
 numbers 𝑎1,𝑎2,…,𝑎𝑛
 (1≤𝑎𝑖≤109
).

It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 2⋅105
.

### ideas
1. 从后往前， dp[i]表示，如果从i开始，(g = 0), 能够有多少good的区间
2. dp[i] = dp[j] + (j - i) 如果 sum(j...i) > x
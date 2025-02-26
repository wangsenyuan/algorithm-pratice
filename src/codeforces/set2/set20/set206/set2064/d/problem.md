There are 𝑛
 slimes on a line, the 𝑖
-th of which has weight 𝑤𝑖
. Slime 𝑖
 is able to eat another slime 𝑗
 if 𝑤𝑖≥𝑤𝑗
; afterwards, slime 𝑗
 disappears and the weight of slime 𝑖
 becomes 𝑤𝑖⊕𝑤𝑗
∗
.

The King of Slimes wants to run an experiment with parameter 𝑥
 as follows:

Add a new slime with weight 𝑥
 to the right end of the line (after the 𝑛
-th slime).
This new slime eats the slime to its left if it is able to, and then takes its place (moves one place to the left). It will continue to do this until there is either no slime to its left or the weight of the slime to its left is greater than its own weight. (No other slimes are eaten during this process.)
The score of this experiment is the total number of slimes eaten.
The King of Slimes is going to ask you 𝑞
 queries. In each query, you will be given an integer 𝑥
, and you need to determine the score of the experiment with parameter 𝑥
.

Note that the King does not want you to actually perform each experiment; his slimes would die, which is not ideal. He is only asking what the hypothetical score is; in other words, the queries are not persistent.


### ideas
1. 从最高位开始考虑，如果d的最右边的是i，且x[d] = 0, 那么x不会超过i（到达位置i时，它就会被吃掉）
2. 假设x[d]是x的最高位，且在更高位确定它的范围不会超过l
3. 如果x[d] = 0, 且l后面所有的数 y[d] = 0, 那么x可以到到l
4. 否则的话，在到达最右边的y[d] = 1 的位置，且需要更新 x ^ sum(y...)
5. 然后再从x[d] = 0, 然后找max(prev[d], l) ... 直到遇到新的x的最高位
6. 
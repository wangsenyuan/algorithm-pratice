You are given a binary string 𝑠
 of length 𝑛
, consisting of zeros and ones. You can perform the following operation exactly once:

Choose an integer 𝑝
 (1≤𝑝≤𝑛
).
Reverse the substring 𝑠1𝑠2…𝑠𝑝
. After this step, the string 𝑠1𝑠2…𝑠𝑛
 will become 𝑠𝑝𝑠𝑝−1…𝑠1𝑠𝑝+1𝑠𝑝+2…𝑠𝑛
.
Then, perform a cyclic shift of the string 𝑠
 to the left 𝑝
 times. After this step, the initial string 𝑠1𝑠2…𝑠𝑛
 will become 𝑠𝑝+1𝑠𝑝+2…𝑠𝑛𝑠𝑝𝑠𝑝−1…𝑠1
.
For example, if you apply the operation to the string 110001100110 with 𝑝=3
, after the second step, the string will become 011001100110, and after the third step, it will become 001100110011.

A string 𝑠
 is called 𝑘
-proper if two conditions are met:

𝑠1=𝑠2=…=𝑠𝑘
;
𝑠𝑖+𝑘≠𝑠𝑖
 for any 𝑖
 (1≤𝑖≤𝑛−𝑘
).
For example, with 𝑘=3
, the strings 000, 111000111, and 111000 are 𝑘
-proper, while the strings 000000, 001100, and 1110000 are not.

You are given an integer 𝑘
, which is a divisor of 𝑛
. Find an integer 𝑝
 (1≤𝑝≤𝑛
) such that after performing the operation, the string 𝑠
 becomes 𝑘
-proper, or determine that it is impossible. Note that if the string is initially 𝑘
-proper, you still need to apply exactly one operation to it.


### ideas
1. k-proper 就是k个1/0交替出现
2. 0和1的个数就是确定的
3. a = n / k, 有这么多段
4. 其中 a/2个 k0, a - a / 2个k1, 或者反过来
5. 选择一个p，并进行操作2后，
6. 后续的必须满足00001111这样子
7. 也就是需要检查的位置，是确定的，
8. reverse前面，再翻转，前缀其实也要满足条件
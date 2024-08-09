Consider a sequence of integers 𝑎1,𝑎2,…,𝑎𝑛
. In one move, you can select any element of the sequence and delete it. After an element is deleted, all elements to the right are shifted to the left by 1
 position, so there are no empty spaces in the sequence. So after you make a move, the sequence's length decreases by 1
. The indices of the elements after the move are recalculated.

E. g. let the sequence be 𝑎=[3,2,2,1,5]
. Let's select the element 𝑎3=2
 in a move. Then after the move the sequence will be equal to 𝑎=[3,2,1,5]
, so the 3
-rd element of the new sequence will be 𝑎3=1
 and the 4
-th element will be 𝑎4=5
.

You are given a sequence 𝑎1,𝑎2,…,𝑎𝑛
 and a number 𝑘
. You need to find the minimum number of moves you have to make so that in the resulting sequence there will be at least 𝑘
 elements that are equal to their indices, i. e. the resulting sequence 𝑏1,𝑏2,…,𝑏𝑚
 will contain at least 𝑘
 indices 𝑖
 such that 𝑏𝑖=𝑖
.

Input
The first line contains one integer 𝑡
 (1≤𝑡≤100
) — the number of test cases. Then 𝑡
 test cases follow.

Each test case consists of two consecutive lines. The first line contains two integers 𝑛
 and 𝑘
 (1≤𝑘≤𝑛≤2000
). The second line contains a sequence of integers 𝑎1,𝑎2,…,𝑎𝑛
 (1≤𝑎𝑖≤𝑛
). The numbers in the sequence are not necessarily different.

It is guaranteed that the sum of 𝑛
 over all test cases doesn't exceed 2000
.

### ideas
1. dp[i][j] 表示到i为止删除了j个元素后得到的最大值
2. dp[i][j] = dp[i-1][j] + (a[i] = i - j + 1 ?) 不删除时，a[i]是否等于它的下标
3.          = dp[i-1][j-1] 删除掉a[i] 
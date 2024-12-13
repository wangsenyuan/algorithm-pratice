### description

Let's denote the following function 𝑓
. This function takes an array 𝑎
of length 𝑛
and returns an array. Initially the result is an empty array. For each integer 𝑖
from 1
to 𝑛
we add element 𝑎𝑖
to the end of the resulting array if it is greater than all previous elements (more formally, if 𝑎𝑖>max1≤𝑗<𝑖𝑎𝑗
). Some examples of the function 𝑓
:

if 𝑎=[3,1,2,7,7,3,6,7,8]
then 𝑓(𝑎)=[3,7,8]
;
if 𝑎=[1]
then 𝑓(𝑎)=[1]
;
if 𝑎=[4,1,1,2,3]
then 𝑓(𝑎)=[4]
;
if 𝑎=[1,3,1,2,6,8,7,7,4,11,10]
then 𝑓(𝑎)=[1,3,6,8,11]
.
You are given two arrays: array 𝑎1,𝑎2,…,𝑎𝑛
and array 𝑏1,𝑏2,…,𝑏𝑚
. You can delete some elements of array 𝑎
(possibly zero). To delete the element 𝑎𝑖
, you have to pay 𝑝𝑖
coins (the value of 𝑝𝑖
can be negative, then you get |𝑝𝑖|
coins, if you delete this element). Calculate the minimum number of coins (possibly negative) you have to spend for
fulfilling equality 𝑓(𝑎)=𝑏
.

### thoughts

1. 先考虑一个dp[i][j] 表示b[..i] 由a[...j]得到时的最小cost
2. 状态转移，dp[i+1][j] = dp[i][j-1] 如果 b[i] = a[j]
   or dp[i+1][j] = dp[i+1][j-1] + cost[j]
3. 这个复杂度时 m * n
4. dp[i] = {j, cost}
    - 状态转移为 dp[i+1] = {k, newcost},
    - 满足条件是 b[i+1] = a[k]
    - 且在 j...k中间不存在 > a[j]的值 （删除它们就是新的cost）
5. 这里k比较容易找到，就是 b[i+1]对应的下个位置，
6. 但是如何计算j...k中间的cost是个问题
    - 并不是删除越少越好，因为cost有可能是负值
    - 在中间的负值全部删掉（这个可以搞个前缀和）
    - 中间的正值，应该仅删除那部分 > b[i]的部分（越少越好）
7. 用persistent二叉树，是可以处理的。但是似乎也太复杂了吧～
8. b是递增的，处理b[i]的时候，只需要关心b[i-1]的情况，所以此时可以将b[i-1]前面的记录给清理掉
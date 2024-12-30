You are given an array of integers 𝑎
 of length 𝑛
. You can perform the following operation zero or more times:

In one operation choose an index 𝑖
 (1≤𝑖≤𝑛
), assign 𝑎𝑖:=𝑎𝑖+1
, and then move 𝑎𝑖
 to the back of the array (to the rightmost position). For example, if 𝑎=[3,5,1,9]
, and you choose 𝑖=2
, the array becomes [3,1,9,6]
.
Find the lexicographically smallest∗
 array you can get by performing these operations.

### ideas
1. 假设结果是b，那么b[1]肯定没有被操作过
2. 所以b[1]必然是最小的a[?]
3. 如果a[i]没有被操作过的话，a[i+1]可以被操作吗？是有可能的，比如a[i+2]比较小
4. 所以这里可以那些没有被操作的数，应该是应该非降序的，其他的都是被操作过的
5. got
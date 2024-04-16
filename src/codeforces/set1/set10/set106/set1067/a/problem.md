Ivan unexpectedly saw a present from one of his previous birthdays. It is array of 𝑛
numbers from 1
to 200
. Array is old and some numbers are hard to read. Ivan remembers that for all elements at least one of its neighbours ls
not less than it, more formally:

𝑎1≤𝑎2
,

𝑎𝑛≤𝑎𝑛−1
and

𝑎𝑖≤𝑚𝑎𝑥(𝑎𝑖−1,𝑎𝑖+1)
for all 𝑖
from 2
to 𝑛−1
.

Ivan does not remember the array and asks to find the number of ways to restore it. Restored elements also should be
integers from 1
to 200
. Since the number of ways can be big, print it modulo 998244353
.

Input
First line of input contains one integer 𝑛
(2≤𝑛≤105
) — size of the array.

Second line of input contains 𝑛
integers 𝑎𝑖
— elements of array. Either 𝑎𝑖=−1
or 1≤𝑎𝑖≤200
. 𝑎𝑖=−1
means that 𝑖
-th element can't be read.

Output
Print number of ways to restore the array modulo 998244353
.

### ideas

1. [1, -1, 2] 中间的数只能是2, 这是因为 a[1] <= max(a[0], a[2]) = 2, a[1] >= a[2] = 2
2. dp, dp[i][x] 表示前i个数，且a[i] = x时的计数
3. 但是这里有个问题是i+1, 也会受到 i+ 2的影响， 如果i+2是不确定的（-1）时，就有问题了
4. a[1] <= a[2]是肯定成立的
5. 如果a[1] < a[2] 那么必须有 a[2] <= a[3]
6. 如果a[1] = a[2], 那么 a[2]对a[3]是没有限制的
7. dp[i][x][0/1] 表示第i个数是x时，a[i]是否 <= a[i-1] (0 不满足，1 满足)
8. dp[i+1][y][0] = dp[i][x][1] if y > x
9. '+' dp[i][x][0] if y > x
10. dp[i+1][y][1] = dp[i][x][1] if y <= x
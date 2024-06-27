You are given an array 𝑎
 of length 𝑛
. Start with 𝑐=0
. Then, for each 𝑖
 from 1
 to 𝑛
 (in increasing order) do exactly one of the following:

Option 1
: set 𝑐
 to 𝑐+𝑎𝑖
.
Option 2
: set 𝑐
 to |𝑐+𝑎𝑖|
, where |𝑥|
 is the absolute value of 𝑥
.
Let the maximum final value of 𝑐
 after the procedure described above be equal to 𝑘
. Find the number of unique procedures that result in 𝑐=𝑘
. Two procedures are different if at any index 𝑖
, one procedure chose option 1
 and another chose option 2
, even if the value of 𝑐
 is equal for both procedures after that turn.

Since the answer may be large, output it modulo 998244353
.


#### ideas
1. 如果a[i] > 0, c > 0 => c + a[i]， 操作1和操作2是等价的
2.    如果 c < 0, c + a[i] < a[i], abs(c + a[i]) 翻转了下符号
3.    a[i] < 0, c < 0 => c + a[i] 和 abs(c + a[i]) 值相同，但是一个正值，一个负值 
4.  c + a[i] 如果 c > 0, a[i] < 0
5.  -1, 4, -3 
6.  可以得到 0, 0, 2
7.  到目前为止c必须是最大值吗？
8.  -1, -2, -3 => 这个的结果是6，在3前面的c不是最大值，其实是最小值
9.  那是不是c必须是到目前为止的最大追，或者是最小值才行？
10. [a, b] 表示到i为止的最大值或者最小值
11. 然后根据操作得到新的[a, b]最后看哪个能得到最大值？
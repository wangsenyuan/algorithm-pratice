You are given an array of integers 𝑎1,𝑎2,…,𝑎𝑛
 and an integer 𝑘
. You need to make it beautiful with the least amount of operations.

Before applying operations, you can shuffle the array elements as you like. For one operation, you can do the following:

Choose an index 1≤𝑖≤𝑛
,
Make 𝑎𝑖=𝑎𝑖+𝑘
.
The array 𝑏1,𝑏2,…,𝑏𝑛
 is beautiful if 𝑏𝑖=𝑏𝑛−𝑖+1
 for all 1≤𝑖≤𝑛
.

Find the minimum number of operations needed to make the array beautiful, or report that it is impossible.


### ideas
1. 要左右对称
2. a[i] = a[j], when i = n - 1 - j
3. 操作只能是 a[i] + k, 那么 a[i] % k = a[j] % k
4. 按照余数分组
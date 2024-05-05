Ralph has a magic field which is divided into n × m blocks. That is to say, there are n rows and m columns on the field.
Ralph can put an integer in each block. However, the magic field doesn't always work properly. It works only if the
product of integers in each row and each column equals to k, where k is either 1 or -1.

Now Ralph wants you to figure out the number of ways to put numbers in each block in such a way that the magic field
works properly. Two ways are considered different if and only if there exists at least one block where the numbers in
the first way and in the second way are different. You are asked to output the answer modulo 1000000007 = 109 + 7.

Note that there is no range of the numbers to put in the blocks, but we can prove that the answer is not infinity.

### ideas

1. pow(2, (n - 1) * (m - 1))
2. 有可能为0
3. 如果n和m都是奇数，且k = -1
4. 比如(n, m) = (3, 3) 这个是没问题（因为前面n-1行，都是-1，前面m-1列也是-1，最后一个始终可以根据需要保证最后一行是-1还是1）
5. 但是如果n,m的parity不一样，会怎样呢？
6. 任何一个-1，都会给该行/列贡献一次-1，如果那么区域(n-1)(m-1)，

### solution

First, it's obvious that the numbers put can be only 1 or -1. If k equals to -1 and the parity of n and m differ, the
answer is obviously 0.

Otherwise, for the first (n - 1) lines and the first (m - 1) columns, we can put either 1 or -1 in it, and there're pow(
2,[(n - 1) * (m - 1)]) ways in total. Then it's obvious that the remaining numbers are uniquely determined because the
product of each row and each column is known already. So in this case the answer is pow(2,[(n - 1) * (m - 1)]) .

### discussion

suppose that n is even and m is odd. if an answer exists then we have even rows each with odd -1s (even -1 in total) and
odd columns each with odd -1s (odd -1 in total) => contradiction


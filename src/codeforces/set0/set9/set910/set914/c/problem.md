The Travelling Salesman spends a lot of time travelling so he tends to get bored. To pass time, he likes to perform
operations on numbers. One such operation is to take a positive integer x and reduce it to the number of bits set to 1
in the binary representation of x. For example for number 13 it's true that 1310 = 11012, so it has 3 bits set and 13
will be reduced to 3 in one operation.

He calls a number special if the minimum number of operations to reduce it to 1 is k.

He wants to find out how many special numbers exist which are not greater than n. Please help the Travelling Salesman,
as he is about to reach his destination!

Since the answer can be large, output it modulo 109 + 7.

### ideas

1. 这个题目好费解
2. One such operation is to take a positive integer x and reduce it to the number of bits set to 1 in the binary
   representation of x
3. 翻译: 一次操作，是将x，变成它二进制表示中1的个数
4. 比如 13 (1101) => 3 (11) => 2 (10) => 1
5. 所以上面13用了3步变成了1
6. 不论这个数多大，它在一步内都会变成1000以内的数， 所以1000以内的数，可以brute force
7. n很大的时候，就需要知道有多少个数可以在一步内变成，比如x (而x需要k-1步)
8. 
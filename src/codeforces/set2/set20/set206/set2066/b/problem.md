We call a sequence 𝑎1,𝑎2,…,𝑎𝑛
 magical if for all 1≤𝑖≤𝑛−1
 it holds that: min(𝑎1,…,𝑎𝑖)≥mex(𝑎𝑖+1,…,𝑎𝑛)
. In particular, any sequence of length 1
 is considered magical.

The minimum excluded (MEX) of a collection of integers 𝑎1,𝑎2,…,𝑎𝑘
 is defined as the smallest non-negative integer 𝑡
 which does not occur in the collection 𝑎
.

You are given a sequence 𝑎
 of 𝑛
 non-negative integers. Find the maximum possible length of a magical subsequence∗
 of the sequence 𝑎
.


### ideas
1. 是找a的子序列，不是a本身
2. 而且在一个序列中，必须对于所有的i 都满足 f(1...i) >= m(i+1...)
3. 如果有一个序列 11111111 那么貌似这个序列肯定是好的
4. 4, 4, 3, 3, 2, 2, 1, 1 0 呢？不行的
5. 不能连续出现两个1，这时候它的mex = 2， 肯定> 1了
6. 所以，只能是 4, 3, 2, 1, 0, 后面呢？必须是 〉 5 的
7. 所以，肯定是一个连续递减的序列
8. 如果我们限定mex = 5， 那么显然在4的前面只能是6（否则它就到6了）
9. 然后，4的后面可以6以上的数（但不能是5）
10. 似乎也不对， 假设 32104 那么这个序列貌似是good的
11. 32140呢？也是ok的
12. 21340呢？也是ok的
13. 20134  也是ok的
14. 有点混乱了～
15. 如果包含0， 只能包括1个0
16. 两个00， 都肯定是不对的
17. 进而可以得出，如果包含一个序列 ,那这个序列中，只能有一个数字 0, 1, 2, 3, 4
18. 
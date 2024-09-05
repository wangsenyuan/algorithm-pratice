You wrote down all integers from 0
 to 10𝑛−1
, padding them with leading zeroes so their lengths are exactly 𝑛
. For example, if 𝑛=3
 then you wrote out 000, 001, ..., 998, 999.

A block in an integer 𝑥
 is a consecutive segment of equal digits that cannot be extended to the left or to the right.

For example, in the integer 00027734000
 there are three blocks of length 1
, one block of length 2
 and two blocks of length 3
.

For all integers 𝑖
 from 1
 to 𝑛
 count the number of blocks of length 𝑖
 among the written down integers.

Since these integers may be too large, print them modulo 998244353
.

### ideas
1. 计算贡献吗。
2. 对于位置l...r （长度为 r - l + 1) 
3. 它们的值是同样的，所有有9种，那么它对 i = r - l + 1 的贡献 = 10 * pow(10, n - i)
4. pow(10, n - i + 1)?
5. 当i固定是 r = i - 1 ... n
6. (n - i) * pow(10, n - i + 1)?
7. 但是这样子感觉有重复计算的情况，且没有控制住长度是i（而不是 i+1等）
8. 不考虑限制，那么数字的个数 = pow(10, n)
9.  如果 r = i - 1, 那么 10 * 9 * pow(10, n - i - 1) (位置i，必须是一个不一样的数)
10. 好像不用考虑重复的问题
11. 因为这里设定了位置（r），所以只需要考虑这段在多少个数里面出现了就可以
12. 10 * 9 * 9 * pow(10, n - i - 2)
13. 
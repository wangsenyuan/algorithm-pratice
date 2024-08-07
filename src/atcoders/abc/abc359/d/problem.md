You are given a string 
S of length 
N consisting of characters A, B, and ?.

You are also given a positive integer 
K. A string 
T consisting of A and B is considered a good string if it satisfies the following condition:

No contiguous substring of length 
K in 
T is a palindrome.
Let 
q be the number of ? characters in 
S. There are 
2 
q
  strings that can be obtained by replacing each ? in 
S with either A or B. Find how many of these strings are good strings.

The count can be very large, so find it modulo 
998244353.

### ideas
1. 如果已经有连续k长的substring 是回文了 => 0
2. dp[l...r] 表示使得l...r为回文时的计数
3. 但是这里主要是会出现重叠的计数，比如l1...r1是超过k的 dp[l1....r1] 和 dp[l2....r2]它们会产生重复的计数
4. dp[1...n] 这个可以很容易计算出来
5. dp[1...n-1], dp[2.....n] 分别表示1.。。n-1是回文时的总计数, 2...n是回文时的总计数
6. 那是不是可以减去 1...n是回文时的总计数，就没有重叠了？
7. 换个思路，如何保证，在s中不出现k个字符长的回文串呢？
8. 假设l...r是一个回文，那么如果r+1 != l - 1， 那么就不会出现
9. 假设到i为止，的前k个字符的情况被记录了下来
10. 如果i是A/B，判断是否ok，如果不ok =》 0
11. 如果ok，更新状态
12. 如果i是?,同样可以更新状态
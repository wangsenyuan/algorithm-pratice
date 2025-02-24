You are given an array 𝑏
 of length 𝑚
. You can perform the following operation any number of times (possibly zero):

Choose two distinct indices 𝑖
 and 𝑗
 where 1≤𝐢<𝐣≤𝐦
 and 𝑏𝑖
 is even, divide 𝑏𝑖
 by 2
 and multiply 𝑏𝑗
 by 2
.
Your task is to maximize the sum of the array after performing any number of such operations. Since it could be large, output this sum modulo 109+7
.
Since this problem is too easy, you are given an array 𝑎
 of length 𝑛
 and need to solve the problem for each prefix of 𝑎
.

In other words, denoting the maximum sum of 𝑏
 after performing any number of such operations as 𝑓(𝑏)
, you need to output 𝑓([𝑎1])
, 𝑓([𝑎1,𝑎2])
, …
, 𝑓([𝑎1,𝑎2,…,𝑎𝑛])
 modulo 109+7
 respectively.

 ### ideas
 1. 考虑a, b
 2. a + b 和 a/2 + 2 * b 
 3. 假设a是偶数
 4. 2 * a + b 与 a + 2 * b 的差 = a - b, 也就是说当 a < b 的时候， 实施这个操作是更合理的
 5. 比如 2 + 3 = 5， 1 + 6 = 7
 6. 对于一个给定的数组，就是把最后一个数a[n]给一直往上加，前面所有的偶数，都用在这个数上
 7. 似乎也不对， 比如 4 + 3 = 7, 4 / 2 + 3 * 2 = 8 
 8. 所以是把所有的偶数给找出来？然后在最大的奇数上进行操作吗？
 9. 4 + 6 = 10, 4 / 2 + 6 * 2 = 14
 10. 6 / 2 + 2 * 4 = 11
 11. a / 2 + 2 * b - (a + b) = a / 2 - a + b
 12. 只要 b > a - a / 2 就能获得正收益
 13. 而且这个收益，b越大，越高
 14. 这个数组中，所有的数字，我们都可以计算2的pow数，这个是可以操作的数量
 15. 然后看把这个数加到某个数上面 got了
 16. 搞错了一点，就是数字只能往后面加，而不能往前加
 17. dp[i] = dp[i-1] + a[i] 在前面的基础上得到的结果

### solution
Consider how to solve the problem for an entire array. We iterate backwards on the array, and for each 𝑎𝑖
 we consider, we give all of its 2
 divisors to the largest 𝑎𝑗>𝑎𝑖
 with 𝑗>𝑖
. If 𝑎𝑖
 is the largest, we do nothing.

To do this for every prefix 𝑥
, notice that if we after we have iterated backwards from 𝑥
 over 31
 2
 divisors, the largest 𝑎𝑖
 will always be larger than 109
 and thus all remaining 2
's will be given to that largest 𝑎𝑖
. Thus, we can iterate over these 31
 2
's and assign them manually, then multiply the largest 𝑎𝑖
 by the number of 2
's we have remaining. Thus, if we maintain the previous 31
 even numbers for every prefix, we can solve the problem in 𝑂(𝑛log𝑎𝑖)
.




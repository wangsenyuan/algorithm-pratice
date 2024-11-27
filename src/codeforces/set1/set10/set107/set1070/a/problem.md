You are given two positive integers d
 and s
. Find minimal positive integer n
 which is divisible by d
 and has sum of digits equal to s
.

Input
The first line contains two positive integers d
 and s
 (1≤d≤500,1≤s≤5000
) separated by space.

### ideas
1. digits sum = s, and can divide d
2. minimal 
3. dp[i][x][r] 是否可以达到
4. 先考虑一下，什么时候没有答案， -1
5. d的倍数，肯定是没有限制的，可以随便加上去
6. 但是s是有限制的。 10 * d = d0, 也就是说，如果d的digits sum 不等于s, 那么 10 * d 也不能等于s
7. 如果 digits_sum(d) = x, 如果 s % x = 0, 那么 s / x 个 d的连接，肯定满足条件（不一定最小）
8. 比如 11, 20
9. 但是这个似乎不是唯一的方案
10. 比如13， 50
11. digits_sum(num) = x
12. num % d = 0
13. num = n * d
14. 如果一个数是2进制的，那么它的digit_sum 不管乘多少，都是不变的
15. 那么可以先变成2进制吗？
16. 13 = 8 + 4 + 1 = 1101
17. 699998 = 10101010111001011110
18. 也看不出啥规律
19. 我们考虑这样一个情况， 尽量的添加9这样的数，（保证最小，用的数越大，长度越短，从而结果越小）
20. 然后得到 w = digits(?) <= s
21. 然后剩余的几个数（它们的digits sum = s - w, 但是它们要求整除的结果是什么呢)
22. 假如能够计算出d1, 似乎就是进入了一个更小的子问题（新的s肯定更小了）
23. d1是不是可以迭代的？比如从0到d-1?
24. d1, s1 得到了一个结果，那么只要在这个结果的基础上，把其他的拼上去就可以了？
25. 但是怎么拼上去，是个问题
26. 9添加在尾巴上，不一定能满足条件，但是添加到头部，却不一定能得到最小值
27. 如果添加在两头，这个d要怎么应用又是个问题
28. dp[d][s] = x （在要求是d，且ds = s时的最小值）
29. 放置足够的数字，从s得到s0是容易的（这个和位置无关，只和具体的数字有关）
30. 
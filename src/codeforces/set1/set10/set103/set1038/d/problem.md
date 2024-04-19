There are 𝑛
slimes in a row. Each slime has an integer value (possibly negative or zero) associated with it.

Any slime can eat its adjacent slime (the closest slime to its left or to its right, assuming that this slime exists).

When a slime with a value 𝑥
eats a slime with a value 𝑦
, the eaten slime disappears, and the value of the remaining slime changes to 𝑥−𝑦
.

The slimes will eat each other until there is only one slime left.

Find the maximum possible value of the last slime.

### ideas

1. 要么被前面的吃掉，要么吃掉前面的
2. dp[i][0] 表示到i为止剩下最大的slime
3. dp[i][1] 表示到i为止剩下的最小的slime
4. dp[i+1][0] = max(dp[i][0] - a[i+1], a[i+1] - dp[i][1])
5. dp[i+1][1] = min(dp[i-1][1] - a[i+1], a[i+1] - dp[i-1][0])
6. 上面的做法不对
7. 是不是一定是从最大值开始操作的？
8. 如果是两个数，那么结果 = max - min
9. 假设在n个数的时候成立，现在是n+1个数，
10. 假设最大的数在位置i，那么它将数组分成两部分pref, suf
11. 肯定希望pref/suf能够有最大值/最小值
12. 但是可以分成前后两端，不过前面怎么分配的，希望是前面的最大值 - 后面的最小值/或者反过来
13. 直觉上应该是有个O(n)的算法的
14. 没个数在最后的结果里，要么 +/-
15. 假设连续的负数，遇到一个正数，那么就可以把前面的分配-，从而增大
16. 如果是连续的正数呢？
17. 如果都是负数，貌似答案显而易见，选择一个最大的负数,然后去减去其他数即可
18. 或者只有一个正数也可以用上面的办法
19. 如果有两个正数，只有两个正数，必须牺牲一个，
20. 但是如果有三个数呢？-1, 2, 3 => (-1 - 2), 3 => 3 - (-1 - 2)即可
21. 更多的负数数，似乎也没问题
22. 如果有三个正数 1， 2， 3 也还是要牺牲一个，
23. 但是如果是-1, 1, 2, 3呢？还是要牺牲一个1
24. 1， -1， 2， 3呢？
25. 多余1个（0或者）正数的情况下，必须牺牲一个数
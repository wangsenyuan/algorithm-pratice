There are n psychos standing in a line. Each psycho is assigned a unique integer from 1 to n. At each step every psycho who has an id greater than the psycho to his right (if exists) kills his right neighbor in the line. Note that a psycho might kill and get killed at the same step.

You're given the initial arrangement of the psychos in the line. Calculate how many steps are needed to the moment of time such, that nobody kills his neighbor after that moment. Look notes to understand the statement more precise.

### ideas
1. 如果 a[i] > a[i+1]
2. dp[i] = max(dp[i+1], dp[j] + 1) j是右边第一个 a[j-1] < a[j]的数
3. 如果a[i] < a[i+1] 这个是没有定义的
4. 不大对，比如 5, 1, 4, 2, 3
5. 一轮后 => 5, 4, 3
6. 再一轮后 => 5
7. 如果dp[i] >= dp[j], dp[i] 不变
8. 如果 dp[i] < dp[j] 会怎样？
9. 6, 1, 5, 2, 3, 4 => 6, 5, 3, 4 => 6, 4 => 6
10. 不过倒是有点眉目，先要处理n，看它需要几次，可以把它右边的全部kill掉
11. 然后看左边下一个更小的数， 依次类推
12. 
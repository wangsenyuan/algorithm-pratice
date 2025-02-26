You will play a game against a dealer. The game uses a 
D-sided die (dice) that shows an integer from 
1 to 
D with equal probability, and two variables 
x and 
y initialized to 
0. The game proceeds as follows:

You may perform the following operation any number of times: roll the die and add the result to 
x. You can choose whether to continue rolling or not after each roll.

Then, the dealer will repeat the following operation as long as 
y<L: roll the die and add the result to 
y.

If 
x>N, you lose. Otherwise, you win if 
y>N or 
x>y and lose if neither is satisfied.

Determine the probability of your winning when you act in a way that maximizes the probability of winning.

### ideas
1. x > y 成立 => x越大越好
2. 但是如果 x 超过了n => lose
3. 也就是说，x最好在 （y, n]
4. 如果 L >= n => win (1.0)
5. y 要到 [L, n) 这个范围的概率 * x 在 （y, n] 这个范围的概率
6. y到一个数字的概率？dp[i] = dp[j] * 1/D if i - j < D
7. dp[i] = sum of (dp[i-D+1]...dp[i-1]) / D 
8. fp[i][j] 表示y在j时，x为i的概率 = dp[j] * dp[i]
9. fp[i] = x = i时的失败概率 = dp[i] * (1 - dp[i....n)) = x达到i的概率 * y到达[i...n)的概率
10. 
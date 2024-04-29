Alice and Bob begin their day with a quick game. They first choose a starting number X0 ≥ 3 and try to reach one million
by the process described below.

Alice goes first and then they take alternating turns. In the i-th turn, the player whose turn it is selects a prime
number smaller than the current number, and announces the smallest multiple of this prime number that is not smaller
than the current number.

Formally, he or she selects a prime p<Xi - 1 and then finds the minimum Xi ≥ Xi - 1 such that p divides Xi. Note that if
the selected prime p already divides Xi - 1, then the number does not change.

Eve has witnessed the state of the game after two turns. Given X2, help her determine what is the smallest possible
starting number X0. Note that the players don't necessarily play optimally. You should consider all possible game
evolutions.

Input

### ideas

1. x0, x1, x2
2. 那么先考虑一种特殊情况x2 = x1 = x0
3. 这种情况肯定是成立的，是因为x2肯定能找到一个质数整除它
4. 然后考虑 x2 = x1 > x0
5. let x0 = lps[x2] + 1
6. 第三种情况 x2 > x1 > x0
7. 可以确定的一点是 x0 = prime + 1 (2, 3, 5, 7)
8. 假设 p0 = x0 - 1, x1 = p0 的倍数，然后 x0 和 x1中间需要存在一个prime数，能够整除x2
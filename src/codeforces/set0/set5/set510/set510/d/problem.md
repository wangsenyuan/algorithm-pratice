Fox Ciel is playing a game. She stands on an endless long tape with cells numbered with whole numbers (positive, negative and zero). At the beginning, she stands in cell 0.

There are also n cards, each of which is characterized by a number l i and a cost c i . If the fox pays c i dollars, then it will be able to use the i -th card. After using the i -th card, the fox will be able to make jumps of length l i , that is, from cell x to cell ( x  -  l i ) or cell ( x  +  l i ) .

The fox wants to learn how to jump to any cell on the tape (possibly visiting some intermediate cells). To achieve this goal, she wants to buy a certain number of cards, paying as little money as possible.

If this can be achieved, calculate the minimum cost.

### ideas
1. 显然如果存在l? = 1, 就可以到达所有的位置
2. 或者至少有 gcd(li) = 1
3. dp[i] = gcd(set) = i 的最小的cost
4. 答案是 dp[1]
5. 不过这里的问题是c[i]太大了。但是，可以先算出它们的质因数
6. c[i] < 1e9 这样的质因数不会超过10个
7. dp[i][x]表示前i个数中，gcd = x的最小的cost
8. dp[i+1][gcd(x, num)] = dp[i][x] + cost[i+1]
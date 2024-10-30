There are n balls. They are arranged in a row. Each ball has a color (for convenience an integer) and an integer value. The color of the i-th ball is ci and the value of the i-th ball is vi.

Squirrel Liss chooses some balls and makes a new sequence without changing the relative order of the balls. She wants to maximize the value of this sequence.

The value of the sequence is defined as the sum of following values for each ball (where a and b are given constants):

If the ball is not in the beginning of the sequence and the color of the ball is same as previous ball's color, add (the value of the ball)  ×  a.
Otherwise, add (the value of the ball)  ×  b.
You are given q queries. Each query contains two integers ai and bi. For each query find the maximal value of the sequence she can make when a = ai and b = bi.

Note that the new sequence can be empty, and the value of an empty sequence is defined as zero.

Input
The first line contains two integers n and q (1 ≤ n ≤ 105; 1 ≤ q ≤ 500). The second line contains n integers: v1, v2, ..., vn (|vi| ≤ 105). The third line contains n integers: c1, c2, ..., cn (1 ≤ ci ≤ n).

The following q lines contain the values of the constants a and b for queries. The i-th of these lines contains two integers ai and bi (|ai|, |bi| ≤ 105).

In each line integers are separated by single spaces.

Output
For each query, output a line containing an integer — the answer to the query. The i-th line contains the answer to the i-th query in the input order.

Please, do not write the %lld specifier to read or write 64-bit integers in С++. It is preferred to use the cin, cout streams or the %I64d specifier.

### ideas
1. dp思考一下
2. dp[i][c] 是到i为止，颜色为c时的最大收益
3. dp[i+1][c] = (c 是v[i+1]的颜色)  
4.    dp[i+1][c] = dp[i][c] + a * v[i+1] 相同颜色时
5.    dp[i+1][c] = max(dp[i][d]) + b * v[i+1] where d != c
6. 然后就是怎么快速的知道 max(dp[i][d])
7. 可以用segment tree
8. 剩下担心的就是时间复杂度 n * q * lg(n) 
9. 果然时间超了
10. 但是有个隐约的想法，就是貌似具体的数值没有关系，但是正负，才有关系
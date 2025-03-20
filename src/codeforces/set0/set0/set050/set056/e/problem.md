Vasya is interested in arranging dominoes. He is fed up with common dominoes and he uses the dominoes of different heights. He put n dominoes on the table along one axis, going from left to right. Every domino stands perpendicular to that axis so that the axis passes through the center of its base. The i-th domino has the coordinate xi and the height hi. Now Vasya wants to learn for every domino, how many dominoes will fall if he pushes it to the right. Help him do that.

Consider that a domino falls if it is touched strictly above the base. In other words, the fall of the domino with the initial coordinate x and height h leads to the fall of all dominoes on the segment [x + 1, x + h - 1].


### ideas
1. 如果i跌倒后，能够压到j, 也就是说 x[j] <= x[i] + h[i] - 1
2. 那么dp[i] = dp[j] （dp[i]表示能够压倒的最远的牌的编号）
3. range query / point update
4. 
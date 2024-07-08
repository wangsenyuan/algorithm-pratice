Alice has 𝑛
 coins and wants to shop at Bob's jewelry store. Today, although Bob has not set up the store yet, Bob wants to make sure Alice will buy exactly 𝑘
 jewels. To set up the store, Bob can erect at most 60
 stalls (each containing an unlimited amount of jewels) and set the price per jewel for each stall to be an integer number of coins between 1
 and 1018
.

Fortunately, Bob knows that Alice buys greedily: and she will go to stall 1
, buy as many jewels as possible, then go to stall 2
, buy as many jewels as possible, and so on until the last stall. Knowing this, Bob can choose the number of stalls to set up, as well as set the price for each stall so that Alice buys exactly 𝑘
 jewels. Help Bob fulfill the task, or determine if it is impossible to do so.

Note that Alice does not need to spend all her coins.

### ideas
1. 这个题目有点奇怪。
2. 假设每个stall的价格是 p1, p2, ...pi
3. 且每个stall购买的数量是 x1, x2, ...xi
4. 那么有 x1 + x2 + .. + xi = k
5. x1 * p1 + x2 * p2 + ... + xi * pi <= n
6. 且 n - x1 * p1 < p1 (否则alice会再购买更多的1号)
7.   n1 - x2 * p2 < p2
8.   ...
9.   n(i-1) - xi * pi < pi
10.  n1 < p1 
11.  n1 >= x2 * p2 => x2 * p2 < p1 如果 x2 >= 1, p2 < p1
12.  n - x2 * p1 < p1
13.  n - x1 * p1 < x2 * p2 + p2 = (x2 + 1) * p2
14.  n < x1 * p1 + (x2 + 1) * p2
15.  假设最后一个就是1（如果没法的话，后面再考虑）
16.  那么就有 n(i-1) = x[i] （剩下的可以全部买完）
17.  n(i-1) < p[i-1] => x[i] < p[i-1] 的
18.  每次剩下的，都不够买一个当前的
19.  假设n个，买k个，刚好在第一次里面用完, n - k * p1 <= p1
20.  n <= (k + 1) * p1
21.  p1 >= (n + k) / (k + 1)
22.  这个是个数学题目，没啥意思
23.  
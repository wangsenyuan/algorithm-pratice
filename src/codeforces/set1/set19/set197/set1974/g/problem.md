You can never buy enough happiness, so here we go again! In this version, you can only buy ℎ𝑖=1
 unit of happiness each month, but the number of months is hugely increased. We are in the realm of quantum happiness and time dilation.

Being a physicist, Charlie likes to plan his life in simple and precise terms.

For the next 𝑚
 months, starting with no money, Charlie will work hard and earn 𝑥
 pounds per month. For the 𝑖
-th month (1≤𝑖≤𝑚)
, there'll be a single opportunity of paying cost 𝑐𝑖
 pounds to obtain one unit of happiness. You cannot buy more than one unit each month.

Borrowing is not allowed. Money earned in the 𝑖
-th month can only be spent in a later 𝑗
-th month (𝑗>𝑖
).

Since physicists don't code, help Charlie find the maximum reachable units of happiness.

### ideas
1. 现在在用f的算法就不行了。因为h的总和 = m
2. 有个观察，就是同样是1个单位的happiness，
3. 比如连续的两个月，c[i], c[i+1]
4. 如果 c[i] <= c[i+1], 那么在都能够购买的情况下，貌似越早买越好
5. 假设在i的时候的钱是w，(w >= c[i]), 那么到达 i+1,(如果不购买) w + x >= c[i+1]
6. 如果 c[i] + x >= c[i+1] 那么如果能够满足i+1的时候，节省i，而选择i+1至少是一个不差的选择
7. 如果 c[i] + x < c[i+1], 那么选择i是一个更好的选择
A conspiracy of ancient sages, who decided to redirect rivers for their own convenience, has put the world on the brink. But before implementing their grand plan, they decided to carefully think through their strategy — that's what sages do.

There are 𝑛
 countries, each with exactly 𝑘
 regions. For the 𝑗
-th region of the 𝑖
-th country, they calculated the value 𝑎𝑖,𝑗
, which reflects the amount of water in it.

The sages intend to create channels between the 𝑗
-th region of the 𝑖
-th country and the 𝑗
-th region of the (𝑖+1)
-th country for all 1≤𝑖≤(𝑛−1)
 and for all 1≤𝑗≤𝑘
.

Since all 𝑛
 countries are on a large slope, water flows towards the country with the highest number. According to the sages' predictions, after the channel system is created, the new value of the 𝑗
-th region of the 𝑖
-th country will be 𝑏𝑖,𝑗=𝑎1,𝑗|𝑎2,𝑗|...|𝑎𝑖,𝑗
, where |
 denotes the bitwise "OR" operation.

After the redistribution of water, the sages aim to choose the most suitable country for living, so they will send you 𝑞
 queries for consideration.

Each query will contain 𝑚
 requirements.

Each requirement contains three parameters: the region number 𝑟
, the sign 𝑜
 (either "<
" or ">
"), and the value 𝑐
. If 𝑜
 = "<
", then in the 𝑟
-th region of the country you choose, the new value must be strictly less than the limit 𝑐
, and if 𝑜
 = ">
", it must be strictly greater.

In other words, the chosen country 𝑖
 must satisfy all 𝑚
 requirements. If in the current requirement 𝑜
 = "<
", then it must hold that 𝑏𝑖,𝑟<𝑐
, and if 𝑜
 = ">
", then 𝑏𝑖,𝑟>𝑐
.

In response to each query, you should output a single integer — the number of the suitable country. If there are multiple such countries, output the smallest one. If no such country exists, output −1
.
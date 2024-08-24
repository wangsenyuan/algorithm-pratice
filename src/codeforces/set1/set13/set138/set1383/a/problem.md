Koa the Koala has two strings 𝐴
 and 𝐵
 of the same length 𝑛
 (|𝐴|=|𝐵|=𝑛
) consisting of the first 20
 lowercase English alphabet letters (ie. from a to t).

In one move Koa:

selects some subset of positions 𝑝1,𝑝2,…,𝑝𝑘
 (𝑘≥1;1≤𝑝𝑖≤𝑛;𝑝𝑖≠𝑝𝑗
 if 𝑖≠𝑗
) of 𝐴
 such that 𝐴𝑝1=𝐴𝑝2=…=𝐴𝑝𝑘=𝑥
 (ie. all letters on this positions are equal to some letter 𝑥
).
selects a letter 𝑦
 (from the first 20
 lowercase letters in English alphabet) such that 𝑦>𝑥
 (ie. letter 𝑦
 is strictly greater alphabetically than 𝑥
).
sets each letter in positions 𝑝1,𝑝2,…,𝑝𝑘
 to letter 𝑦
. More formally: for each 𝑖
 (1≤𝑖≤𝑘
) Koa sets 𝐴𝑝𝑖=𝑦
.
Note that you can only modify letters in string 𝐴
.

Koa wants to know the smallest number of moves she has to do to make strings equal to each other (𝐴=𝐵
) or to determine that there is no way to make them equal. Help her!


### ideas
1. 如果存在i, a[i] > b[i] => -1 
2. 然后剩余的假设对于 a[i] -> b[i] 
3. 根据 pair{a, b} 当然可以用一次完成；但这样子不一定是最优的
4. 如何最优呢？
5. 把上面组成的图，变成树？每棵树，操作 n-1,次，有几棵树树？
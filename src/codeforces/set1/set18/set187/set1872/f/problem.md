You are the owner of a menagerie consisting of 𝑛
animals numbered from 1
to 𝑛
. However, maintaining the menagerie is quite expensive, so you have decided to sell it!

It is known that each animal is afraid of exactly one other animal. More precisely, animal 𝑖
is afraid of animal 𝑎𝑖
(𝑎𝑖≠𝑖
). Also, the cost of each animal is known, for animal 𝑖
it is equal to 𝑐𝑖
.

You will sell all your animals in some fixed order. Formally, you will need to choose some permutation†
𝑝1,𝑝2,…,𝑝𝑛
, and sell animal 𝑝1
first, then animal 𝑝2
, and so on, selling animal 𝑝𝑛
last.

When you sell animal 𝑖
, there are two possible outcomes:

If animal 𝑎𝑖
was sold before animal 𝑖
, you receive 𝑐𝑖
money for selling animal 𝑖
.
If animal 𝑎𝑖
was not sold before animal 𝑖
, you receive 2⋅𝑐𝑖
money for selling animal 𝑖
. (Surprisingly, animals that are currently afraid are more valuable).
Your task is to choose the order of selling the animals in order to maximize the total profit.

For example, if 𝑎=[3,4,4,1,3]
, 𝑐=[3,4,5,6,7]
, and the permutation you choose is [4,2,5,1,3]
, then:

The first animal to be sold is animal 4
. Animal 𝑎4=1
was not sold before, so you receive 2⋅𝑐4=12
money for selling it.
The second animal to be sold is animal 2
. Animal 𝑎2=4
was sold before, so you receive 𝑐2=4
money for selling it.
The third animal to be sold is animal 5
. Animal 𝑎5=3
was not sold before, so you receive 2⋅𝑐5=14
money for selling it.
The fourth animal to be sold is animal 1
. Animal 𝑎1=3
was not sold before, so you receive 2⋅𝑐1=6
money for selling it.
The fifth animal to be sold is animal 3
. Animal 𝑎3=4
was sold before, so you receive 𝑐3=5
money for selling it.
Your total profit, with this choice of permutation, is 12+4+14+6+5=41
. Note that 41
is not the maximum possible profit in this example.

### thoughts

1. 这些动物或者形成环，或者一条线连接到某个环上
2. 不在环上的，double it，
3. 在环上的，选择最小的那个，其他的double
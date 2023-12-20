### [description](https://codeforces.com/problemset/problem/1826/E)

A fashion tour consists of 𝑚
identical runway shows in different cities. There are 𝑛
models willing to participate in the tour, numbered from 1
to 𝑛
. People in different cities have different views on the fashion industry, so they rate each model differently. In
particular, people in city 𝑖
rate model 𝑗
with rating 𝑟𝑖,𝑗
.

You are to choose some number of 𝑘
models, and their order, let the chosen models have indices 𝑗1,𝑗2,…,𝑗𝑘
in the chosen order. In each city, these 𝑘
models will walk the runway one after another in this order. To make the show exciting, in each city, the ratings of
models should be strictly increasing in the order of their performance. More formally, for any city 𝑖
and index 𝑡
(2≤𝑡≤𝑘
), the ratings must satisfy 𝑟𝑖,𝑗𝑡−1<𝑟𝑖,𝑗𝑡
.

After all, the fashion industry is all about money, so choosing model 𝑗
to participate in the tour profits you 𝑝𝑗
money. Compute the maximum total profit you can make by choosing the models and their order while satisfying all the
requirements.

### thoughts

1. 假设要选择模特(i, j)pair, 那么必然要满足在所有的城市都满足 ri < rj
2. 可以先建立这样一个图，这个图肯定是一个dag
3. 然后bfs处理即可

### solution
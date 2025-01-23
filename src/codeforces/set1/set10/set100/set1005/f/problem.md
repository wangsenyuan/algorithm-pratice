There are 𝑛
 cities in Berland. Some pairs of cities are connected by roads. All roads are bidirectional. Each road connects two different cities. There is at most one road between a pair of cities. The cities are numbered from 1
 to 𝑛
.

It is known that, from the capital (the city with the number 1
), you can reach any other city by moving along the roads.

The President of Berland plans to improve the country's road network. The budget is enough to repair exactly 𝑛−1
 roads. The President plans to choose a set of 𝑛−1
 roads such that:

it is possible to travel from the capital to any other city along the 𝑛−1
 chosen roads,
if 𝑑𝑖
 is the number of roads needed to travel from the capital to city 𝑖
, moving only along the 𝑛−1
 chosen roads, then 𝑑1+𝑑2+⋯+𝑑𝑛
 is minimized (i.e. as minimal as possible).
In other words, the set of 𝑛−1
 roads should preserve the connectivity of the country, and the sum of distances from city 1
 to all cities should be minimized (where you can only use the 𝑛−1
 chosen roads).

The president instructed the ministry to prepare 𝑘
 possible options to choose 𝑛−1
 roads so that both conditions above are met.

Write a program that will find 𝑘
 possible ways to choose roads for repair. If there are fewer than 𝑘
 ways, then the program should output all possible valid ways to choose roads.

 ### ideas
 1. 使用bfs生成最小生成树就是最少的cost 
 2. cost = ...
 3. 然后主要的问题是怎么找出k个符合树 = cost
 4. 每个d[v] = 是固定的（从1到v的最短距离）
 5. 如果v有w个前继节点，那么这里就可以贡献w (乘法)
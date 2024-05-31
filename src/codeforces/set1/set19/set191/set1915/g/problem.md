All of Slavic's friends are planning to travel from the place where they live to a party using their bikes. And they all have a bike except Slavic. There are 𝑛
 cities through which they can travel. They all live in the city 1
 and want to go to the party located in the city 𝑛
. The map of cities can be seen as an undirected graph with 𝑛
 nodes and 𝑚
 edges. Edge 𝑖
 connects cities 𝑢𝑖
 and 𝑣𝑖
 and has a length of 𝑤𝑖
.

Slavic doesn't have a bike, but what he has is money. Every city has exactly one bike for sale. The bike in the 𝑖
-th city has a slowness factor of 𝑠𝑖
. Once Slavic buys a bike, he can use it whenever to travel from the city he is currently in to any neighboring city, by taking 𝑤𝑖⋅𝑠𝑗
 time, considering he is traversing edge 𝑖
 using a bike 𝑗
 he owns.

Slavic can buy as many bikes as he wants as money isn't a problem for him. Since Slavic hates traveling by bike, he wants to get from his place to the party in the shortest amount of time possible. And, since his informatics skills are quite rusty, he asks you for help.

What's the shortest amount of time required for Slavic to travel from city 1
 to city 𝑛
? Slavic can't travel without a bike. It is guaranteed that it is possible for Slavic to travel from city 1
 to any other city.

 ### ideas
 1. 如果到达了某个城市i，如果他之前到过某个j，且s[j] < s[i]， 那么使用bike j更优
 2. dp[u][su] = 在城市u使用bike su所能达到的最优解
 3. dp[j][min(su, s[i])] = dp[i][su] + min(su, s[i]) * w[i]
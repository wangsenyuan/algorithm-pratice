In Ancient Berland there were n cities and m two-way roads of equal length. The cities are numbered with integers from 1 to n inclusively. According to an ancient superstition, if a traveller visits three cities ai, bi, ci in row, without visiting other cities between them, a great disaster awaits him. Overall there are k such city triplets. Each triplet is ordered, which means that, for example, you are allowed to visit the cities in the following order: ai, ci, bi. Vasya wants to get from the city 1 to the city n and not fulfil the superstition. Find out which minimal number of roads he should take. Also you are required to find one of his possible path routes.

Input
The first line contains three integers n, m, k (2 ≤ n ≤ 3000, 1 ≤ m ≤ 20000, 0 ≤ k ≤ 105) which are the number of cities, the number of roads and the number of the forbidden triplets correspondingly.

Then follow m lines each containing two integers xi, yi (1 ≤ xi, yi ≤ n) which are the road descriptions. The road is described by the numbers of the cities it joins. No road joins a city with itself, there cannot be more than one road between a pair of cities.

Then follow k lines each containing three integers ai, bi, ci (1 ≤ ai, bi, ci ≤ n) which are the forbidden triplets. Each ordered triplet is listed mo more than one time. All three cities in each triplet are distinct.

City n can be unreachable from city 1 by roads.

Output
If there are no path from 1 to n print -1. Otherwise on the first line print the number of roads d along the shortest path from the city 1 to the city n. On the second line print d + 1 numbers — any of the possible shortest paths for Vasya. The path should start in the city 1 and end in the city n.

### ideas
1. 如果出现了连续的 a -> b, 那么下一个位置就被限制了
2. dp[a][b] 表示最后两个位置是a, b时，处在b时的最短距离
3. 但似乎还不够，比如a, b, c（c != a）都被禁止了，
4. 但是 a, b, a（回头路，是否可行呢？）
5. 是可行的，因为此时，a就会出现 dp[b][a]这样的有效状态
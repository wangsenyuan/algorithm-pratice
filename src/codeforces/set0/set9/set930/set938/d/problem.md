Musicians of a popular band "Flayer" have announced that they are going to "make their exit" with a world tour. Of
course, they will visit Berland as well.

There are n cities in Berland. People can travel between cities using two-directional train routes; there are exactly m
routes, i-th route can be used to go from city vi to city ui (and from ui to vi), and it costs wi coins to use this
route.

Each city will be visited by "Flayer", and the cost of the concert ticket in i-th city is ai coins.

You have friends in every city of Berland, and they, knowing about your programming skills, asked you to calculate the
minimum possible number of coins they have to pay to visit the concert. For every city i you have to compute the minimum
number of coins a person from city i has to spend to travel to some city j (or possibly stay in city i), attend a
concert there, and return to city i (if j ≠ i).

Formally, for every you have to calculate , where d(i, j) is the minimum number of coins you have to spend to travel
from city i to city j. If there is no way to reach city j from city i, then we consider d(i, j) to be infinitely large.

### ideas

1. 对于city i, 需要知道某个最便宜的city j, d[i][j] * 2 + a[j]
2. 首先，对于city i来讲，只有那些a[j] < a[i] 的城市，才需要关注，
3. 但是n也很大，即使是n*n的算法也是不行的， 假设i到j是最优的方案， 有可能出现j到其他的地方，比如k是最优的方案吗？
4. 假设ok, 那么有 a[j] > 2 * d(j, k) + a[k]
5. 我们知道有 a[i] > 2 * d(i, j) + a[j] 显然从i的角度看，它到达j后，继续到k，是更优的方案
6. 这里就存在这样一种情况，如果从i出发，到达了j，那么j就是那些能够到达i的最优的方案
7. 存在一种可能的方式是，假设ans[i] = j表示i应该到j, ans[x] = j if x 能够到达i
8. 假设从j出发，j显然是局部最小的，然后，如果a[j] + 2 * d(i, j) < a[i]， 那么它们都应该到j去
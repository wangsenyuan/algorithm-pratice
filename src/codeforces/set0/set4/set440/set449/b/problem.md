Jzzhu is the president of country A. There are n cities numbered from 1 to n in his country. City 1 is the capital of A.
Also there are m roads connecting the cities. One can go from city ui to vi (and vise versa) using the i-th road, the
length of this road is xi. Finally, there are k train routes in the country. One can use the i-th train route to go from
capital of the country to city si (and vise versa), the length of this route is yi.

Jzzhu doesn't want to waste the money of the country, so he is going to close some of the train routes. Please tell
Jzzhu the maximum number of the train routes which can be closed under the following condition: the length of the
shortest path from every city to the capital mustn't change.

### ideas

1. 先算出每个城市的最短距离，dijistra
2. 然后如果某个城市的最短距离 < y[i], 那么这条铁路就可以关闭掉
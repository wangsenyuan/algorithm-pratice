There are 𝑛
 cities in Berland. The city numbered 1
 is the capital. Some pairs of cities are connected by a one-way road of length 1.

Before the trip, Polycarp for each city found out the value of 𝑑𝑖
 — the shortest distance from the capital (the 1
-st city) to the 𝑖
-th city.

Polycarp begins his journey in the city with number 𝑠
 and, being in the 𝑖
-th city, chooses one of the following actions:

Travel from the 𝑖
-th city to the 𝑗
-th city if there is a road from the 𝑖
-th city to the 𝑗
-th and 𝑑𝑖<𝑑𝑗
;
Travel from the 𝑖
-th city to the 𝑗
-th city if there is a road from the 𝑖
-th city to the 𝑗
-th and 𝑑𝑖≥𝑑𝑗
;
Stop traveling.
Since the government of Berland does not want all people to come to the capital, so Polycarp no more than once can take the second action from the list. in other words, he can perform the second action 0
 or 1
 time during his journey. Polycarp, on the other hand, wants to be as close to the capital as possible.


For example, if 𝑛=6
 and the cities are connected, as in the picture above, then Polycarp could have made the following travels (not all possible options):

2→5→1→2→5
;
3→6→2
;
1→3→6→2→5
.
Polycarp wants for each starting city 𝑖
 to find out how close he can get to the capital. More formally: he wants to find the minimal value of 𝑑𝑗
 that Polycarp can get from the city 𝑖
 to the city 𝑗
 according to the rules described above.

 ### ideas
 1. 先计算出d[i]
 2. 可以按照规则1，从较小的d[i]往较大的d[j]运动，但是这样除非在某个k处能到达更小的d[x]， 否则是没有好处的
 3. 所以，按照d[i]从大到小运动f[i]表示i能达到的最近的d[x], 那么 f[j] = min(d[i] (使用规则2), f[i] 使用规则1)
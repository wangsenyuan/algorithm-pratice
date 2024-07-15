There are 𝑛
 cities located on the number line, the 𝑖
-th city is in the point 𝑎𝑖
. The coordinates of the cities are given in ascending order, so 𝑎1<𝑎2<⋯<𝑎𝑛
.

The distance between two cities 𝑥
 and 𝑦
 is equal to |𝑎𝑥−𝑎𝑦|
.

For each city 𝑖
, let's define the closest city 𝑗
 as the city such that the distance between 𝑖
 and 𝑗
 is not greater than the distance between 𝑖
 and each other city 𝑘
. For example, if the cities are located in points [0,8,12,15,20]
, then:

the closest city to the city 1
 is the city 2
;
the closest city to the city 2
 is the city 3
;
the closest city to the city 3
 is the city 4
;
the closest city to the city 4
 is the city 3
;
the closest city to the city 5
 is the city 4
.
The cities are located in such a way that for every city, the closest city is unique. For example, it is impossible for the cities to be situated in points [1,2,3]
, since this would mean that the city 2
 has two closest cities (1
 and 3
, both having distance 1
).

You can travel between cities. Suppose you are currently in the city 𝑥
. Then you can perform one of the following actions:

travel to any other city 𝑦
, paying |𝑎𝑥−𝑎𝑦|
 coins;
travel to the city which is the closest to 𝑥
, paying 1
 coin.
You are given 𝑚
 queries. In each query, you will be given two cities, and you have to calculate the minimum number of coins you have to spend to travel from one city to the other city.


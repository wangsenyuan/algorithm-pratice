Asya loves animals very much. Recently, she purchased 𝑛
 kittens, enumerated them from 1
 and 𝑛
 and then put them into the cage. The cage consists of one row of 𝑛
 cells, enumerated with integers from 1
 to 𝑛
 from left to right. Adjacent cells had a partially transparent partition wall between them, hence there were 𝑛−1
 partitions originally. Initially, each cell contained exactly one kitten with some number.

Observing the kittens, Asya noticed, that they are very friendly and often a pair of kittens in neighboring cells wants to play together. So Asya started to remove partitions between neighboring cells. In particular, on the day 𝑖
, Asya:

Noticed, that the kittens 𝑥𝑖
 and 𝑦𝑖
, located in neighboring cells want to play together.
Removed the partition between these two cells, efficiently creating a single cell, having all kittens from two original cells.
Since Asya has never putted partitions back, after 𝑛−1
 days the cage contained a single cell, having all kittens.

For every day, Asya remembers numbers of kittens 𝑥𝑖
 and 𝑦𝑖
, who wanted to play together, however she doesn't remember how she placed kittens in the cage in the beginning. Please help her and find any possible initial arrangement of the kittens into 𝑛
 cells.



### ideas
1. 似乎清楚的，但是又有点混乱
2. 假设，目前处理(x, y)， 且和x在一起的小猫有5只，和y在一起的有3只，
3. 
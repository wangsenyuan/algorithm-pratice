There are 𝑁
 cities in the country of Numbata, numbered from 1
 to 𝑁
. Currently, there is no road connecting them. Therefore, each of these 𝑁
 cities proposes a road candidate to be constructed.

City 𝑖
 likes to connect with city 𝐴𝑖
, so city 𝑖
 proposes to add a direct bidirectional road connecting city 𝑖
 and city 𝐴𝑖
. It is guaranteed that no two cities like to connect with each other. In other words, there is no pair of integers 𝑖
 and 𝑗
 where 𝐴𝑖=𝑗
 and 𝐴𝑗=𝑖
. It is also guaranteed that any pair of cities are connected by a sequence of road proposals. In other words, if all proposed roads are constructed, then any pair of cities are connected by a sequence of constructed road.

City 𝑖
 also prefers the road to be constructed using a specific material. Each material can be represented by an integer (for example, 0
 for asphalt, 1
 for wood, etc.). The material that can be used for the road connecting city 𝑖
 and city 𝐴𝑖
 is represented by an array 𝐵𝑖
 containing 𝑀𝑖
 integers: [(𝐵𝑖)1,(𝐵𝑖)2,…,(𝐵𝑖)𝑀𝑖]
. This means that the road connecting city 𝑖
 and city 𝐴𝑖
 can be constructed with either of the material in 𝐵𝑖
.

There are 𝐾
 workers to construct the roads. Each worker is only familiar with one material, thus can only construct a road with a specific material. In particular, the 𝑖𝑡ℎ
 worker can only construct a road with material 𝐶𝑖
. Each worker can only construct at most one road. You want to assign each worker to construct a road such that any pair of cities are connected by a sequence of constructed road.

Input
Input begins with a line containing two integers: 𝑁
 𝐾
 (3≤𝑁≤2000
; 1≤𝐾≤2000
) representing the number of cities and the number of workers, respectively. The next 𝑁
 lines each contains several integers: 𝐴𝑖
 𝑀𝑖
 (𝐵𝑖)1
, (𝐵𝑖)2
, ⋯
, (𝐵𝑖)𝑀𝑖
 (1≤𝐴𝑖≤𝑁
; 𝐴𝑖≠𝑖
; 1≤𝑀𝑖≤10000
; 0≤(𝐵𝑖)1<(𝐵𝑖)2<⋯<(𝐵𝑖)𝑀𝑖≤109
) representing the bidirectional road that city 𝑖
 likes to construct. It is guaranteed that the sum of 𝑀𝑖
 does not exceed 10000
. It is also guaranteed that no two cities like to connect with each other and any pair of cities are connected by a sequence of road proposals. The next line contains 𝐾
 integers: 𝐶𝑖
 (0≤𝐶𝑖≤109
) representing the material that is familiarized by the workers.

Output
If it is not possible to assign each worker to construct a road such that any pair of cities are connected by a sequence of constructed road, simply output -1 in a line. Otherwise, for each worker in the same order as input, output in a line two integers (separated by a single space): 𝑢
 and 𝑣
 in any order. This means that the worker constructs a direct bidirectional road connecting city 𝑢
 and 𝑣
. If the worker does not construct any road, output "0 0" (without quotes) instead. Each pair of cities can only be assigned to at most one worker. You may output any assignment as long as any pair of cities are connected by a sequence of constructed road.

### ideas

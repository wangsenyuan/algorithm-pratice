King Kendrick is a sovereign ruler of Kotlin Kingdom. He is getting ready for the next session of the
government. Kotlin Kingdom consists of n cities. These cities need to be connected by several bidirectional
roads. Since ministries are responsible for aspects of safety and comfort of the kingdom’s residents, some
of them have made the following requirements:
• “All the cities should be connected by new roads, i.e. there should be a path from any city to any
other city via the roads” — Ministry of Transport and Digital Infrastructure.
• “There may not be a loop road — a road that connects a city with itself” — Ministry of Environment.
• “There should be at most one road between a pair of cities” — Treasury Department.
• “If ai
is the number of roads connected to i-th city, then the set {a1, . . . , an} should consist of exactly
k distinct numbers” — Ministry of ICPC.
King Kendrick has issues with the requirements from the Ministry of ICPC. He asks you to help him.
Find any set of roads that suits all the requirements above or say that it is impossible.


### ideas
1. 规则1、2、3要求这是一个没有回环、重边的联通图
2. 重要的是规则4，要求所有节点的deg组成的set的size = k
3. 先组成一棵树（如果n = 2, 那么只有k = 1 时成立，否则答案为false)
4. 如果 k= 1, 组成一个环 (deg = 2)
5. n > 2, k > 2, 先组成一条线（deg_size = 2)中间节点的deg = 2， 两头的deg = 1
6. 如果 1 + 2 + 3 ... + k <= n， 那么是有答案的。否则是没有答案的
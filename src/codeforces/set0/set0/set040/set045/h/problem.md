The Berland capital (as you very well know) contains n junctions, some pairs of which are connected by two-way roads. Unfortunately, the number of traffic jams in the capital has increased dramatically, that's why it was decided to build several new roads. Every road should connect two junctions.

The city administration noticed that in the cities of all the developed countries between any two roads one can drive along at least two paths so that the paths don't share any roads (but they may share the same junction). The administration decided to add the minimal number of roads so that this rules was fulfilled in the Berland capital as well. In the city road network should exist no more than one road between every pair of junctions before or after the reform.

Input
The first input line contains a pair of integers n, m (2 ≤ n ≤ 900, 1 ≤ m ≤ 100000), where n is the number of junctions and m is the number of roads. Each of the following m lines contains a description of a road that is given by the numbers of the connected junctions ai, bi (1 ≤ ai, bi ≤ n, ai ≠ bi). The junctions are numbered from 1 to n. It is possible to reach any junction of the city from any other one moving along roads.

Output
On the first line print t — the number of added roads. Then on t lines print the descriptions of the added roads in the format of the input data. You can use any order of printing the roads themselves as well as the junctions linked by every road. If there are several solutions to that problem, print any of them.

If the capital doesn't need the reform, print the single number 0.

If there's no solution, print the single number -1.

### ideas
1. 这个题目咋读不通呢。
2. 一条road, 用它的两个端点(u, v)来表示，
3. 那么任何两条路(u1, v1), (u2, v2), 条件 (u1, u2) and (v1, v2) 能够联通，且不经过另外两个点
4. 那是不是，所有的点，都需要至少在一个环上？
5. 任何一个compoent，size 大于1的那些，deg = 1的节点的数目 / 2 的ceil？
6. 这是一个component的情况
7. 如果有多个component， 可以把它们用其中一个deg = 1 的点，把它们都连起来（这样子可以减少 deg = 1的节点）
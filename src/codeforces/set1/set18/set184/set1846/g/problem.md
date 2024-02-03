A new virus called "CodeVid-23" has spread among programmers. Rudolf, being a programmer, was not able to avoid it.

There are 𝑛
symptoms numbered from 1
to 𝑛
that can appear when infected. Initially, Rudolf has some of them. He went to the pharmacy and bought 𝑚
medicines.

For each medicine, the number of days it needs to be taken is known, and the set of symptoms it removes. Unfortunately,
medicines often have side effects. Therefore, for each medicine, the set of symptoms that appear when taking it is also
known.

After reading the instructions, Rudolf realized that taking more than one medicine at a time is very unhealthy.

Rudolph wants to be healed as soon as possible. Therefore, he asks you to calculate the minimum number of days to remove
all symptoms, or to say that it is impossible.

Input
The first line contains a single integer 𝑡
(1≤𝑡≤100)
— the number of test cases.

Then follow the descriptions of the test cases.

The first line of each test case contains two integers 𝑛,𝑚
(1≤𝑛≤10,1≤𝑚≤103)
— the number of symptoms and medicines, respectively.

The second line of each test case contains a string of length 𝑛
consisting of the characters 0
and 1
— the description of Rudolf's symptoms. If the 𝑖
-th character of the string is 1
, Rudolf has the 𝑖
-th symptom, otherwise he does not.

Then follow 3⋅𝑚
lines — the description of the medicines.

The first line of each medicine description contains an integer 𝑑
(1≤𝑑≤103)
— the number of days the medicine needs to be taken.

The next two lines of the medicine description contain two strings of length 𝑛
, consisting of the characters 0
and 1
— the description of the symptoms it removes and the description of the side effects.

In the first of the two lines, 1
at position 𝑖
means that the medicine removes the 𝑖
-th symptom, and 0
otherwise.

In the second of the two lines, 1
at position 𝑖
means that the 𝑖
-th symptom appears after taking the medicine, and 0
otherwise.

Different medicines can have the same sets of symptoms and side effects. If a medicine relieves a certain symptom, it
will not be among the side effects.

The sum of 𝑚
over all test cases does not exceed 103
.

### thoughts

1. 第一个观察n <= 10, 所以可以用一个state来表示 (1 << 10) 表示症状
2. dp[x][y] 表示使用某些药物后，消除x代表的症状，但产生y代表的症状的最小天数
3. 考虑第i中药物 dp[x ^ (r[i] & x)][y | g[i]] = min(dp[x][y] + day[i])
4. 但这个复杂度是 m * 1000 * 1000 = 1e9
5. 考虑一个比较自然的想法，假设当前的症状是x， 要使用某种药物，一个自然的想法，是使用能够解决症状x，且产生副作用少的药物
6. 如果某个药物不能处理症状x，完全不用考虑
7. 把症状x代表节点，那么就是从一开始的症状到 0 的转换过程，药物，就是两个状态间的连线
8. 但是每个药物，应该是把它包含的子集和y连起来
9. 然后计算最短路径

### solution

Let's denote Rudolf's state as a binary mask of length 𝑛
consisting of 0
and 1
, similar to how it is given in the input data. Then each medicine transforms Rudolf from one state to another.

Let's construct a weighted directed graph, where the vertices will represent all possible states of Rudolf. There will
be 2𝑛
such vertices. Two vertices will be connected by an edge if there exists a medicine that transforms Rudolf from the
state corresponding to the first vertex to the state corresponding to the second vertex. The weight of the edge will be
equal to the number of days that this medicine needs to be taken. Note that in this case, we simply need to find the
shortest path in this graph from the vertex 𝑠
, corresponding to the initial state of Rudolf, to the vertex 𝑓
, corresponding to the state without symptoms.

To find the shortest path in a weighted graph, we will use Dijkstra's algorithm. We will run it from the vertex 𝑠
and if, as a result, we visit the vertex 𝑓
, output the distance to it, otherwise −1
.

The time complexity is 𝑂(𝑛⋅𝑚⋅2𝑛)
.
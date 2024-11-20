Boy Dima gave Julian a birthday present — set 𝐵
 consisting of positive integers. However, he didn't know, that Julian hates sets, but enjoys bipartite graphs more than anything else!

Julian was almost upset, but her friend Alex said, that he can build an undirected graph using this set in such a way: let all integer numbers be vertices, then connect any two 𝑖
 and 𝑗
 with an edge if |𝑖−𝑗|
 belongs to 𝐵
.

Unfortunately, Julian doesn't like the graph, that was built using 𝐵
. Alex decided to rectify the situation, so he wants to erase some numbers from 𝐵
, so that graph built using the new set is bipartite. The difficulty of this task is that the graph, Alex has to work with, has an infinite number of vertices and edges! It is impossible to solve this task alone, so Alex asks you for help. Write a program that erases a subset of minimum size from 𝐵
 so that graph constructed on the new set is bipartite.

Recall, that graph is bipartite if all its vertices can be divided into two disjoint sets such that every edge connects a vertex from different sets.

### ideas
1. [1, 2, 3]删掉1也是可以的
2. (a, b, c) 如果(a, b)，(b, c), (a, c) 有条边
3. 如果处理c的时候，如果b存在，且a = c -b也存在，那么就出现了长度为3的环
4. 所以处理c的时候，b存在，但是a = c - b 不存在，那就可以把c加进去
5. 就是这个集合中，要大家都是孤立的点，如果存在边，也是类似2,4,8,16 这样的，序列
6. 3，6，12，24.。。也是ok的

### solution

If you take some pair of numbers (𝑎,𝑏)
 from B you can alwas find a cycle. Let's say you start at 0
 then you'll have the path 0,𝑎,2𝑎,3𝑎,...,𝑙𝑐𝑚(𝑎,𝑏)
. From there you can decrease 𝑏
 until you hit 0
 again.

E. g., if 𝑎=3,𝑏=5
, you have: 0→3→6→9→12→15→10→5→0
.

The length of the cycle length 𝑙𝑒𝑛
 is equal to the number of ascending 𝑎
 steps plus the number of descending 𝑏
 steps. So 𝑙𝑒𝑛=𝑙𝑐𝑚(𝑎,𝑏)𝑎+𝑙𝑐𝑚(𝑎,𝑏)𝑏
.

From 𝑎⋅𝑏=𝑙𝑐𝑚(𝑎,𝑏)⋅𝑔𝑐𝑑(𝑎,𝑏)
, it follows: 𝑙𝑐𝑚(𝑎,𝑏)𝑎=𝑏𝑔𝑐𝑑(𝑎,𝑏)
.

So 𝑙𝑒𝑛=𝑙𝑐𝑚(𝑎,𝑏)𝑎+𝑙𝑐𝑚(𝑎,𝑏)𝑏=𝑏𝑔𝑐𝑑(𝑎,𝑏)+𝑎𝑔𝑐𝑑(𝑎,𝑏)
.

If 𝑎
 and 𝑏
 are both divisible by the same power of 2
, then 𝑏𝑔𝑐𝑑(𝑎,𝑏)
 and 𝑏𝑔𝑐𝑑(𝑎,𝑏)
 are gonna be odd, so 𝑙𝑒𝑛
 will be even.

Otherwise, one of the terms will be even so 𝑙𝑒𝑛
 will be odd.

Finally, you can show that you can't have odd length cycles in a bipartite graph. So solutions won't have numbers with different number of powers of 2
 in the set.
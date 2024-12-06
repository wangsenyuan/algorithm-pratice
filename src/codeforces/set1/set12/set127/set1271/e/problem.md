At first, let's define function 𝑓(𝑥)
 as follows:
𝑓(𝑥)={𝑥2𝑥−1if 𝑥 is evenotherwise 

We can see that if we choose some value 𝑣
 and will apply function 𝑓
 to it, then apply 𝑓
 to 𝑓(𝑣)
, and so on, we'll eventually get 1
. Let's write down all values we get in this process in a list and denote this list as 𝑝𝑎𝑡ℎ(𝑣)
. For example, 𝑝𝑎𝑡ℎ(1)=[1]
, 𝑝𝑎𝑡ℎ(15)=[15,14,7,6,3,2,1]
, 𝑝𝑎𝑡ℎ(32)=[32,16,8,4,2,1]
.

Let's write all lists 𝑝𝑎𝑡ℎ(𝑥)
 for every 𝑥
 from 1
 to 𝑛
. The question is next: what is the maximum value 𝑦
 such that 𝑦
 is contained in at least 𝑘
 different lists 𝑝𝑎𝑡ℎ(𝑥)
?

Formally speaking, you need to find maximum 𝑦
 such that |{𝑥 | 1≤𝑥≤𝑛,𝑦∈𝑝𝑎𝑡ℎ(𝑥)}|≥𝑘
.

### ideas
1.   考虑一个有n个节点的树，其中每个节点的父节点 = f(i) = i/2 if i is even
2.   else f(i) = i-1
3.   那么i出现的次数，就是这个节点的子节点的数目
4.   好像有戏
5.   对于节点i，如何知道这个节点的子节点的数目呢？
6.   f(i) = ...
7.   别人的代码完全看不懂呐～～～～
8.   
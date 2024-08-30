You are given an undirected graph without self-loops or multiple edges which consists of 𝑛
 vertices and 𝑚
 edges. Also you are given three integers 𝑛1
, 𝑛2
 and 𝑛3
.

Can you label each vertex with one of three numbers 1, 2 or 3 in such way, that:

Each vertex should be labeled by exactly one number 1, 2 or 3;
The total number of vertices with label 1 should be equal to 𝑛1
;
The total number of vertices with label 2 should be equal to 𝑛2
;
The total number of vertices with label 3 should be equal to 𝑛3
;
|𝑐𝑜𝑙𝑢−𝑐𝑜𝑙𝑣|=1
 for each edge (𝑢,𝑣)
, where 𝑐𝑜𝑙𝑥
 is the label of vertex 𝑥
.
If there are multiple valid labelings, print any of them.

### ideas
1. 看起来2比较重要，同一条边不能用同样的label，且其中一个必须是2，且1\3不能直接链接
2. 所以，先看图是否能变成bipartitle
3. 然后其中一个的数量必须等于n2
4. 这样子肯定是充分条件，也就是如果2分+数量等于n2，可以得到一个有效的结果；但是是必要条件吗？
5. 这个图有可能有很多个component组成
6. 那还设计到dp？
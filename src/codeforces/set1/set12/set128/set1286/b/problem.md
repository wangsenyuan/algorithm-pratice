Evlampiy was gifted a rooted tree. The vertices of the tree are numbered from 1
 to 𝑛
. Each of its vertices also has an integer 𝑎𝑖
 written on it. For each vertex 𝑖
, Evlampiy calculated 𝑐𝑖
 — the number of vertices 𝑗
 in the subtree of vertex 𝑖
, such that 𝑎𝑗<𝑎𝑖
.

Illustration for the second example, the first integer is 𝑎𝑖
 and the integer in parentheses is 𝑐𝑖
After the new year, Evlampiy could not remember what his gift was! He remembers the tree and the values of 𝑐𝑖
, but he completely forgot which integers 𝑎𝑖
 were written on the vertices.

Help him to restore initial integers!

Input
The first line contains an integer 𝑛
 (1≤𝑛≤2000)
 — the number of vertices in the tree.

The next 𝑛
 lines contain descriptions of vertices: the 𝑖
-th line contains two integers 𝑝𝑖
 and 𝑐𝑖
 (0≤𝑝𝑖≤𝑛
; 0≤𝑐𝑖≤𝑛−1
), where 𝑝𝑖
 is the parent of vertex 𝑖
 or 0
 if vertex 𝑖
 is root, and 𝑐𝑖
 is the number of vertices 𝑗
 in the subtree of vertex 𝑖
, such that 𝑎𝑗<𝑎𝑖
.

It is guaranteed that the values of 𝑝𝑖
 describe a rooted tree with 𝑛
 vertices.

Output
If a solution exists, in the first line print "YES", and in the second line output 𝑛
 integers 𝑎𝑖
 (1≤𝑎𝑖≤109)
. If there are several solutions, output any of them. One can prove that if there is a solution, then there is also a solution in which all 𝑎𝑖
 are between 1
 and 109
.

If there are no solutions, print "NO".



### ideas
1. 考虑u->v, 且v已经处理好了
2. 如果c[u] > c[v]
3. 好像有办法了。假设v的节点已经全部处理好了，且每个节点有不一样的值
4. 现在处理u，希望在子树v中能够出现x个节点 < a[u]
5. 那么很简单，就是v子树中的第x个节点的值+1
6. 如果c[u] >= sz[v], 那么就是这个子树的最大值+1
7. 这里可以尽量的使用那些sz小的树（这样子，这些子树对u的上限没有限制)
8. 然后找到某个子树v, c[u] < sz[v]
9. 找到第c[u]个值，然后设置a[u] = x + 1, 然后把剩余的节点的所有 >x的都+1即可
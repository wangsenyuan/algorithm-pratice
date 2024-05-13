Mitya has a rooted tree with 𝑛
 vertices indexed from 1
 to 𝑛
, where the root has index 1
. Each vertex 𝑣
 initially had an integer number 𝑎𝑣≥0
 written on it. For every vertex 𝑣
 Mitya has computed 𝑠𝑣
: the sum of all values written on the vertices on the path from vertex 𝑣
 to the root, as well as ℎ𝑣
 — the depth of vertex 𝑣
, which denotes the number of vertices on the path from vertex 𝑣
 to the root. Clearly, 𝑠1=𝑎1
 and ℎ1=1
.

Then Mitya erased all numbers 𝑎𝑣
, and by accident he also erased all values 𝑠𝑣
 for vertices with even depth (vertices with even ℎ𝑣
). Your task is to restore the values 𝑎𝑣
 for every vertex, or determine that Mitya made a mistake. In case there are multiple ways to restore the values, you're required to find one which minimizes the total sum of values 𝑎𝑣
 for all vertices in the tree.

Input
The first line contains one integer 𝑛
 — the number of vertices in the tree (2≤𝑛≤105
). The following line contains integers 𝑝2
, 𝑝3
, ... 𝑝𝑛
, where 𝑝𝑖
 stands for the parent of vertex with index 𝑖
 in the tree (1≤𝑝𝑖<𝑖
). The last line contains integer values 𝑠1
, 𝑠2
, ..., 𝑠𝑛
 (−1≤𝑠𝑣≤109
), where erased values are replaced by −1
.


### ideas
1. 如果存在偶数h的叶子结点，可以直接忽略(设置a[v] = 0)
2. 剩下的叶子节点都是奇数， s[v] = s[p[v]] + a[v]
3. s[v]是知道的，但是 s[p[v]]不知道(它在偶数层)，
4. s[p[v]] = s[p[p[v]]] + a[p[v]]
5. s[v] = s[p[p[v]]] + a[p[v]] + a[v]
6. a[p[v]] + a[v] = s[v] - s[p[p[v]]]
7. 也就是说，其实不用知道a[v]的值，大概能够确定最小的sum
8. let u = p[v]
9. a[u] + a[v] = s[v] - s[p[u]]
10. a[u] + a[w] = s[w] - s[p[u]]
11. 上面 + 下面的 , 2 * a[u] + a[v] + a[w] = ... 后面的是确定的 = X
12. a[u] + a[v] + a[w] = X - a[u]
13. 为了上面的a[u] + a[v] + a[w] 越小，需要 a[u]越大越大
14. a[u] <= s[v] - s[p[u]], s[w] - ...
15. 所以有a[u]的最大值， 如果 a[u] < 0  => 无解， 否则就是这个解
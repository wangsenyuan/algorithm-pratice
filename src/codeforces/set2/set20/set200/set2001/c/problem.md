Misuki has chosen a secret tree with 𝑛
 nodes, indexed from 1
 to 𝑛
, and asked you to guess it by using queries of the following type:

"? a b" — Misuki will tell you which node 𝑥
 minimizes |𝑑(𝑎,𝑥)−𝑑(𝑏,𝑥)|
, where 𝑑(𝑥,𝑦)
 is the distance between nodes 𝑥
 and 𝑦
. If more than one such node exists, Misuki will tell you the one which minimizes 𝑑(𝑎,𝑥)
.
Find out the structure of Misuki's secret tree using at most 15𝑛
 queries!

### ideas
1. ask(a, b) 返回了ab的中点（当a，b中间有偶数个点时，返回靠近a的那个点）
2. ask(a, b) 返回了 x, 如果 a = x, 那么（a, b)是一条边
3. 否则的话，用x替代a，继续,知道 a = x
4. 一共是n-1条边，问出一条边，不会超过10次，
5. 上面的做法是错误的。
6. 先找到一个中点，比如1，用它去和其他的节点进行一次ask
7. 比如 ask(1, 10) 得到3， 那么可以得到一个信息是 3是(1, 10)的中点
8. 也就是说，这样子可以得到一些点之间的关系，同时也能得到1的直接连接的边
9. 这些返回的x点，肯定是靠近1的那些点
10. 比如和1直接连接的边有3, 如果3不是叶子节点，它肯定也是某些问题的答案
11. 
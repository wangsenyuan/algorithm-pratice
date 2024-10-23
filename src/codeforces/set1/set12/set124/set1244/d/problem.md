You are given a tree consisting of n
 vertices. A tree is an undirected connected acyclic graph.

Example of a tree.
You have to paint each vertex into one of three colors. For each vertex, you know the cost of painting it in every color.

You have to paint the vertices so that any path consisting of exactly three distinct vertices does not contain any vertices with equal colors. In other words, let's consider all triples (x,y,z)
 such that x≠y,y≠z,x≠z
, x
 is connected by an edge with y
, and y
 is connected by an edge with z
. The colours of x
, y
 and z
 should be pairwise distinct. Let's call a painting which meets this condition good.

You have to calculate the minimum cost of a good painting and find one of the optimal paintings. If there is no good painting, report about it.

### ideas
1. 考虑tree的特殊形式，一条线
2. 编号从1....n，依次连接在一起，形成一条线， 每个节点只需要知道前面（两个）节点的着色，然后第三个节点的着色就确定了
3. 如果一个节点的deg > 2, 似乎是没有答案的
4. 假设这个中心点w, 又x, y, z 和它连接，那么就没法给x,y,z进行着色
5. 所以，只能是一条线
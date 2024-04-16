Someone give a strange birthday present to Ivan. It is hedgehog — connected undirected graph in which one vertex has
degree at least 3
(we will call it center) and all other vertices has degree 1. Ivan thought that hedgehog is too boring and decided to
make himself 𝑘
-multihedgehog.

Let us define 𝑘
-multihedgehog as follows:

1
-multihedgehog is hedgehog: it has one vertex of degree at least 3
and some vertices of degree 1.
For all 𝑘≥2
, 𝑘
-multihedgehog is (𝑘−1)
-multihedgehog in which the following changes has been made for each vertex 𝑣
with degree 1: let 𝑢
be its only neighbor; remove vertex 𝑣
, create a new hedgehog with center at vertex 𝑤
and connect vertices 𝑢
and 𝑤
with an edge. New hedgehogs can differ from each other and the initial gift.
Thereby 𝑘
-multihedgehog is a tree. Ivan made 𝑘
-multihedgehog but he is not sure that he did not make any mistakes. That is why he asked you to check if his tree is
indeed 𝑘
-multihedgehog.

### ideas

1. 是递归的，显示最外层的 deg = 1 的叶子节点，它们的parent 的deg >= 3
2. 然后删除最外层的，剩下的是重新成为叶子节点，直到剩下一个点
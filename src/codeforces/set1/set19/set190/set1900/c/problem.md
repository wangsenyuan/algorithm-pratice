Keksic keeps getting left on seen by Anji. Through a mutual friend, he's figured out that Anji really likes binary trees
and decided to solve her problem in order to get her attention.

Anji has given Keksic a binary tree with 𝑛
vertices. Vertex 1
is the root and does not have a parent. All other vertices have exactly one parent. Each vertex can have up to 2
children, a left child, and a right child. For each vertex, Anji tells Keksic index of both its left and its right child
or tells him that they do not exist.

Additionally, each of the vertices has a letter 𝑠𝑖
on it, which is either 'U', 'L' or 'R'.

Keksic begins his journey on the root, and in each move he does the following:

If the letter on his current vertex is 'U', he moves to its parent. If it doesn't exist, he does nothing.
If the letter on his current vertex is 'L', he moves to its left child. If it doesn't exist, he does nothing.
If the letter on his current vertex is 'R', he moves to its right child. If it doesn't exist, he does nothing.
Before his journey, he can perform the following operations: choose any node, and replace the letter written on it with
another one.
You are interested in the minimal number of operations he needs to do before his journey, such that when he starts his
journey, he will reach a leaf at some point. A leaf is a vertex that has no children. It does not matter which leaf he
reaches. Note that it does not matter whether he will stay in the leaf, he just needs to move to it. Additionally, note
that it does not matter how many times he needs to move before reaching a leaf.


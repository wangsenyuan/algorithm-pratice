You are given a complete directed graph 𝐾𝑛
 with 𝑛
 vertices: each pair of vertices 𝑢≠𝑣
 in 𝐾𝑛
 have both directed edges (𝑢,𝑣)
 and (𝑣,𝑢)
; there are no self-loops.

You should find such a cycle in 𝐾𝑛
 that visits every directed edge exactly once (allowing for revisiting vertices).

We can write such cycle as a list of 𝑛(𝑛−1)+1
 vertices 𝑣1,𝑣2,𝑣3,…,𝑣𝑛(𝑛−1)−1,𝑣𝑛(𝑛−1),𝑣𝑛(𝑛−1)+1=𝑣1
 — a visiting order, where each (𝑣𝑖,𝑣𝑖+1)
 occurs exactly once.

Find the lexicographically smallest such cycle. It's not hard to prove that the cycle always exists.

Since the answer can be too large print its [𝑙,𝑟]
 segment, in other words, 𝑣𝑙,𝑣𝑙+1,…,𝑣𝑟
.

### ideas
1. 考虑一个3点的图
2. (1 -> 2 -> 1 -> 3 -> 2 -> 3 -> 1)
3. 考虑4点的图
4. (1 -> 2 -> 1 -> 3 -> 1 -> 4 -> 2 -> 3 -> 2 -> 4 -> 3 -> 4 -> 1)
5. 

### solution
The solution of the problem can be found clearly in constructive way. An example for 𝑛=5
: (1 2 1 3 1 4 1 5 (2 3 2 4 2 5 (3 4 3 5 (4 5 ()))) 1) where brackets mean that we call here some recursive function 𝑐𝑎𝑙𝑐
.

Since on each level of recursion we have only 𝑂(𝑛)
 elements and there 𝑂(𝑛)
 levels then the generation of the certificate is quite easy: if on the currect level of recursion we can skip the whole part — let's just skip it. Otherwise let's build this part. Anyway, the built part of the cycle will have only 𝑂(𝑛+(𝑟−𝑙))
 length so the whole algorithm has 𝑂(𝑛+(𝑟−𝑙))
 complexity.

The answer is lexicographically minimum by the construction, since on each level of recursion there is no way to build lexicographically smaller sequence.
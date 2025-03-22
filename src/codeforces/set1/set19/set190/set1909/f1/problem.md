# Good Permutation Count

You are given an integer $n$ and an array $a_1, a_2, \ldots, a_n$ of integers in the range $[0, n]$.

A permutation $p_1, p_2, \ldots, p_n$ of $[1, 2, \ldots, n]$ is **good** if, for each $i$, the following condition is true:

* The number of values $\leq i$ in $[p_1, p_2, \ldots, p_i]$ is exactly $a_i$.

Count the good permutations of $[1, 2, \ldots, n]$, modulo $998244353$.

Andrey is just starting to come up with problems, and it's difficult for him. That's why he came up with a strange
problem about permutations†
and asks you to solve it. Can you do it?

Let's call the cost of a permutation 𝑝
of length 𝑛
the value of the expression:

(∑𝑛𝑖=1𝑝𝑖⋅𝑖)−(max𝑛𝑗=1𝑝𝑗⋅𝑗)
.
Find the maximum cost among all permutations of length 𝑛
.

†
A permutation of length 𝑛
is an array consisting of 𝑛
distinct integers from 1
to 𝑛
in arbitrary order. For example, [2,3,1,5,4]
is a permutation, but [1,2,2]
is not a permutation (2
appears twice in the array), and [1,3,4]
is also not a permutation (𝑛=3
but there is 4
in the array).

### thoughts

1. 为了让前面的数大，一种可能的做法是 let p[i] = i
2. 但是这样，同样会让 后面一项更大
3. result = 1 * p[1] + 2 * p[2] + .. + n * p[n] - i * p[i]
4. 如果 p[i] != i, 那把i 和 p[i] 挑出来，剩余的部分仍然是 j * p[j]更优
5. 有问题，因为这样无法保证 i * j 是最大值
6. i * j > n * k
7. 先固定最大值，然后再分配？
8. 对于i，如果对应的j是最大值，然后保证其他的不超过它时
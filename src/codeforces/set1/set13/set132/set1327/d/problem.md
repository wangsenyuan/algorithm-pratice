You are given a colored permutation 𝑝1,𝑝2,…,𝑝𝑛
. The 𝑖
-th element of the permutation has color 𝑐𝑖
.

Let's define an infinite path as infinite sequence 𝑖,𝑝[𝑖],𝑝[𝑝[𝑖]],𝑝[𝑝[𝑝[𝑖]]]…
 where all elements have same color (𝑐[𝑖]=𝑐[𝑝[𝑖]]=𝑐[𝑝[𝑝[𝑖]]]=…
).

We can also define a multiplication of permutations 𝑎
 and 𝑏
 as permutation 𝑐=𝑎×𝑏
 where 𝑐[𝑖]=𝑏[𝑎[𝑖]]
. Moreover, we can define a power 𝑘
 of permutation 𝑝
 as 𝑝𝑘=𝑝×𝑝×⋯×𝑝𝑘 times
.

Find the minimum 𝑘>0
 such that 𝑝𝑘
 has at least one infinite path (i.e. there is a position 𝑖
 in 𝑝𝑘
 such that the sequence starting from 𝑖
 is an infinite path).

It can be proved that the answer always exists.

### ideas
1. a * b 代表什么？
2. 当p = [1,2,3], p**k = p
3. 当p = [1, 3, 2], p * p = [1, 2, 3]
4. 似乎p * p 就是将那个loopshift一次
5. p = [3, 1, 2], p * p = [2, 3, 1]
6. 所以经过lengh - 1 次，它就变成了1, 2, 3
7. 也就是说这个答案不会超过最大的loop的size - 1
8. infinite path 表示的就是，同一个圈的颜色要一致，如果一开始这个圈就是同种颜色的，可以直接忽略，否则的话，它至少需要 size - 1
9. p = [7 4 5 6 1 8 3 2], p2 = [3,6,1,8,7,2,5,4]
10. [7, ?, 5, ?, 1, ?, 3, ?] 组成了一个环，[7, 3, 5, 1]
11. [....] 另外一个也组成了一个环
12. 
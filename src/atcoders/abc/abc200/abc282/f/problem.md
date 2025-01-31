**Prompt:**

This is an interactive task where your and the judge's programs interact via Standard Input and Output.

You and the judge will follow the procedure below, which consists of two phases: Phase 1 followed immediately by Phase 2.

**Phase 1:**

1. The judge gives you an integer N.
2. You print an integer M between 1 and 50000, inclusive.
3. You also print M pairs of integers (l₁, r₁), (l₂, r₂), ..., (lₘ, rₘ) such that 1 ≤ lᵢ ≤ rᵢ ≤ N for every i = 1, 2, ..., M (the M pairs do not have to be distinct).

**Phase 2:**

1. The judge gives you an integer Q.
2. Repeat the following Q times:
   - The judge gives you two integers L and R as a query.
   - You respond with two integers a and b between 1 and M, inclusive (possibly with a = b).
   - a and b must satisfy the condition below:
     - The union of the set {lₐ, lₐ + 1, ..., rₐ} and the set {l♭, l♭ + 1, ..., r♭} equals the set {L, L + 1, ..., R}.

After the procedure above, terminate the program immediately to be judged correct.


### ideas
1. 有点懵
2. 这里n <= 4000
3. 然后可以准备m(50000)个区间，保证对于任何区间查询(l <= r <= n)能够使用两个区间(a, b)使得他们的union = (l...r)
4. 如果这里m = 50000 >= n * (n + 1) / 2, 那么有个简单的做法，就是把所有的区间都列一遍，即可
5. 但这个只能处理 n <= 300的情况
6. 首先要能回答 l == r 的情况，所以必须有n个单节点区间
7. 任何 l + 1 = r 的情况也能被cover住
8. 现在处理l + 1 < r 的情况
9. (1, 1), (1, 2), (1, 4), ... (1, x)
10. n * lg(n)即可 4000 * 12 = 480000 good
11. 
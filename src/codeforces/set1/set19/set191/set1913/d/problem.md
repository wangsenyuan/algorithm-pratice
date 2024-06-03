You are given an array [𝑝1,𝑝2,…,𝑝𝑛]
, where all elements are distinct.

You can perform several (possibly zero) operations with it. In one operation, you can choose a contiguous subsegment of 𝑝
 and remove all elements from that subsegment, except for the minimum element on that subsegment. For example, if 𝑝=[3,1,4,7,5,2,6]
 and you choose the subsegment from the 3
-rd element to the 6
-th element, the resulting array is [3,1,2,6]
.

An array 𝑎
 is called reachable if it can be obtained from 𝑝
 using several (maybe zero) aforementioned operations. Calculate the number of reachable arrays, and print it modulo 998244353
.

Input
The first line of the input contains one integer 𝑡
 (1≤𝑡≤104
) — the number of test cases.

Each test case consists of two lines. The first line contains one integer 𝑛
 (1≤𝑛≤3⋅105
). The second line contains 𝑛
 distinct integers 𝑝1,𝑝2,…,𝑝𝑛
 (1≤𝑝𝑖≤109
).

Additional constraint on the input: the sum of 𝑛
 over all test cases does not exceed 3⋅105
.

Output
For each test case, print one integer — the number of reachable arrays, taken modulo 998244353
.

### ideas
1. 考虑例子 [2, 4, 1, 3]
2. 这里一共有 n * (n + 1) / 2 = 4 * 5 / 2 = 10 个区间
3. > 1 的有6个区间
4. 可以得到 [2, 4, 1, 3], [2, 1, 3], [2, 4, 1], [1, 3], [2, 1], [1] 六个不同的集合

### solution
Let's rephrase the problem a bit. Instead of counting the number of arrays, let's count the number of subsequences of elements that can remain at the end.

One of the classic methods for counting subsequences is dynamic programming of the following kind: dp𝑖
 is the number of subsequences such that the last taken element is at position 𝑖
. Counting exactly this is that simple. It happens that element 𝑖
 cannot be the last overall because it is impossible to remove everything after it. Thus, let's say it a little differently: the number of good subsequences on a prefix of length 𝑖
. Now we are not concerned with what comes after 𝑖
 — after the prefix.

If we learned to calculate such dp, we could get the answer from it. We just need to understand when we can remove all elements after a fixed one. It is easy to see that it is enough for all these elements to be smaller than the fixed one. Then they can be removed one by one from left to right. If there is at least one larger element, then the maximum of such elements cannot be removed.

Therefore, the answer is equal to the sum of dp for all 𝑖
 such that 𝑎𝑖
 is the largest element on a suffix of length 𝑛−𝑖+1
.

How to calculate such dp?

Let's look at the nearest element to the left that is larger than 𝑎𝑖
. Let its position be 𝑗
. Then any subsequence that ended with an element between 𝑗
 and 𝑖
 can have the element 𝑖
 appended to it. It is only necessary to remove the elements standing before 𝑖
. This can also be done one by one.

Can 𝑗
 be removed as well? It can, but it's not that simple. Only the element that is the nearest larger one to the left for 𝑎𝑗
 or someone else even more to the left can do this. Let 𝑓[𝑖]
 be the position of the nearest larger element to the left. Then the element 𝑖
 can also extend subsequences ending in 𝑎[𝑓[𝑖]],𝑎[𝑓[𝑓[𝑖]]],𝑎[𝑓[𝑓[𝑓[𝑖]]]],…
.

If there are no larger elements to the left of the element — the element is the maximum on the prefix — then 1
 is added to its dp
 value. This is a subsequence consisting of only this element.

The position of the nearest larger element to the left can be found using a monotonic stack: keep a stack of elements; while the element at the top is smaller, remove it; then add the current one to the stack.

Counting the dp currently works in 𝑂(𝑛2)
 in the worst case. How to optimize this? The first type of transitions is optimized by prefix sums, as it is simply the sum of dp on a segment. For the second type of transitions, you can maintain the sum of dp on a chain of jumps to the nearest larger element: dpsum𝑖=dpsum𝑓[𝑖]+dp𝑖
. Now, both transitions can be done in 𝑂(1)
.

Overall complexity: 𝑂(𝑛)
 for each testcase.
https://codeforces.com/problemset/problem/1906/E

### problem

You are currently researching the Merge Sort algorithm. Merge Sort is a sorting algorithm that is based on the principle
of Divide and Conquer. It works by dividing an array into two subarrays of equal length, sorting each subarrays, then
merging the sorted subarrays back together to form the final sorted array.

You are particularly interested in the merging routine. Common merge implementation will combine two subarrays by
iteratively comparing their first elements, and move the smaller one to a new merged array. More precisely, the merge
algorithm can be presented by the following pseudocode.

```
    Merge(A[1..N], B[1..N]):
        C = []
        i = 1
        j = 1
        while i <= N AND j <= N:
            if A[i] < B[j]:
                append A[i] to C
                i = i + 1
            else:
                append B[j] to C
                j = j + 1 
        while i <= N:
            append A[i] to C
            i = i + 1 
        while j <= N:
            append B[j] to C
            j = j + 1 
        return C

```

During your research, you are keen to understand the behaviour of the merge algorithm when arrays 𝐴
and 𝐵
are not necessarily sorted. For example, if 𝐴=[3,1,6]
and 𝐵=[4,5,2]
, then Merge(𝐴,𝐵)=[3,1,4,5,2,6]
.

To further increase the understanding of the merge algorithm, you decided to work on the following problem. You are
given an array 𝐶
of length 2⋅𝑁
such that it is a permutation of 1
to 2⋅𝑁
. Construct any two arrays 𝐴
and 𝐵
of the same length 𝑁
, such that Merge(𝐴,𝐵)=𝐶
, or determine if it is impossible to do so.

### thoughts

1. 要将C分成两个相等的数组A, B, 使得它们可以merge出C
2. 如果 c[i] > c[i+1] 它们两个必须在同一个数组中
3. c[i] <= c[i+1] 可以在两个数组中
4. dp[i][j] 表示前i个，其中j个和最后一个（包括最后一个）在同一个数组中
5. dp[i+1][j] = dp[i][j-1] 如果 c[i] > c[i+1] (i+1必须和i在同一个数组中)
6. dp[i+1][j] = dp[i][j-1] c[i] <= c[i+1] (i + 1必须和i在同一个数组中)
7.            = dp[i][i+1 - j]  或者 i和i+1不在同一个分组中
8. dp[2 * n][n]

### solution

Observe what happens if we merge any two arrays A and B of length N consisting of a permutation
of integers from 1 to 2N.
• Let A[P] be the first element in A[1..N] which is greater than A[1].
• Let B[Q] be the first element in B[1..N] which is greater than B[1].
• If A[1] < B[1], then the “block” A[1..P − 1] must be appended to C, as all of them must be
less than B[1].
• Else, the “block” B[1..Q − 1] will be appended to C instead.
• Repeat the procedure with the new A and B.
In the end, the C will consist of a series of appended blocks. We label these blocks to be C =
C1 · C2 · C3 · · · · . Observe that, no matter which the block of Ci comes from (either A or B), the first
element of Ci+1 must be the first element which is greater than the first element of Ci
. Therefore,
given the input C, our first step is to split it into blocks, which can be done in O(N). To construct the
A and B, we can take note of the length of each block and then assign each block to be either from
A or B such that the total length is exactly N for both. This can be done using a simple Knapsack
DP in O(N2
).
As an additional info, the original constraint for this problem is actually N ≤ 100, 000. Indeed, there
exists an O(N
√
N) algorithm to solve the problem, and this is left to reader as an exercise.
The total time complexity for this problem is O(N2
).
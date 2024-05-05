You are given an array a with n distinct integers. Construct an array b by permuting a such that for every non-empty
subset of indices S = {x1, x2, ..., xk} (1 ≤ xi ≤ n, 0<k<n) the sums of elements on that positions in a and b are
different, i. e.

### ideas

1. dp[state]表示是否能将state表示的下标进行处理
2. 不大对
3. 但是也不能进行 22!次检查吧

### solution

Sort the array and shift it by one. This array will be an answer.

Proof:

When we shift the sorted array all of the elements become greater except the first one, consider f = {1, 2, ..., n} and
t = {x1, x2, ..., xk} if 1 wasn't in t we would have

otherwise consider q = {y1, y2, ..., yn - k} = f - t then 1 can't be in q and we have

so

and we are done!
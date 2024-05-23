Devu wants to decorate his garden with flowers. He has purchased n boxes, where the i-th box contains fi flowers. All flowers in a single box are of the same color (hence they are indistinguishable). Also, no two boxes have flowers of the same color.

Now Devu wants to select exactly s flowers from the boxes to decorate his garden. Devu would like to know, in how many different ways can he select the flowers from each box? Since this number may be very large, he asks you to find the number modulo (109 + 7).

Devu considers two ways different if there is at least one box from which different number of flowers are selected in these two ways.


### ideas
1. dp[i][x] = dp[i-1][x - y] for y >= 0 y <= a[i]
2. 所以这里会出现 s * s 的情况
3. dp[i][x] = dp[i-1][x - a[i]] + dp[i-1][x - a[i] + 1] + ... dp[i-1][x]
4. let fp[i][x] = dp[i][0] + ... + dp[i][x]
5. dp[i][x] = fp[i-1][x] - fp[i-1][x-a[i]-1]
6. 😢， s[i] <= 1e14 
7. 所以上面的dp是不行的，要换个思路
8. x1 + x2 + x3 ... + x4 = s
9. 且 x1 <= a1, x2 <= a2 ....
10. C(S, x1) * C(s - x1, x2) * C(s - x1 - x2, x3)... * C(., xk) 

### solutions

The number of ways to choose N items out of R groups where each item in a group is identical is equal to the number of integral solutions to x1 + x2 + x3...xR = N, where 0 ≤ xi ≤ Li, where Li is the number of items in ith group. Number of integral solutions are coefficient of xN in [Product of (1 + x + x * x + ...xLi) over all $i$].

You need to find coefficient of xs in (1 + x + x2 + x3 +  + ..xf1) *  *  * (1 + x + x2 + x3 +  + ..xfn).

Using sum of Geometric progression we can say that (1 + x + x2 + x3 +  + ..xf1) = (1 - x(f1 + 1)) / (1 - x).

Substituting in the expression, we get (1 - x(f1 + 1)) / (1 - x) *  *  * (1 - x(fn + 1)) / (1 - x).

= (1 - x(f1 + 1)) * .. * (1 - x(fn + 1)) * (1 - x)( - n).

Now we can find xs in (1 - x) - n easily. It is .

You can have a look at following link. to understand it better.

So now as s is large, we can not afford to iterate over s.

But n is small, we notice that (1 - x(f1 + 1)) * .. * (1 - x(fn + 1)) can have at most 2n terms.

So we will simply find all those terms, they can be very easily computed by maintaining a vector<pair<int, int> > containing pairs of coefficients and their corresponding powers. You can write a recursive function for doing this.

How to find  % p. As n + s - 1 is large and s is very small. You can use lucas's theorem. If you understand lucas's theorem, you can note that we simply have to compute .

Complexity: O(n * 2n).


A ticket is a non-empty string of digits from 1
to 9
.

A lucky ticket is such a ticket that:

it has an even length;
the sum of digits in the first half is equal to the sum of digits in the second half.
You are given 𝑛
ticket pieces 𝑠1,𝑠2,…,𝑠𝑛
. How many pairs (𝑖,𝑗)
(for 1≤𝑖,𝑗≤𝑛
) are there such that 𝑠𝑖+𝑠𝑗
is a lucky ticket? Note that it's possible that 𝑖=𝑗
.

Here, the + operator denotes the concatenation of the two strings. For example, if 𝑠𝑖
is 13, and 𝑠𝑗
is 37, then 𝑠𝑖+𝑠𝑗
is 1337.

### thoughts

1. si + sj is good
2. 如果 len(si) >= len(sj)
3. x + y = si, x - y = sj
4. 2 * x = si + sj
5. sj = 2 * x - si
6. 如果 len(si) < len(sj)
7. x - y = si, x + y = sj
8. 2 * y = sj - si
9. sj = 2 * y + si
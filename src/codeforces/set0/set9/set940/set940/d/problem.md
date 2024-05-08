The array b of length n is obtained from the array a of length n and two integers l and r (l ≤ r) using the following
procedure:

b1 = b2 = b3 = b4 = 0.

For all 5 ≤ i ≤ n:

bi = 0 if ai, ai - 1, ai - 2, ai - 3, ai - 4>r and bi - 1 = bi - 2 = bi - 3 = bi - 4 = 1
bi = 1 if ai, ai - 1, ai - 2, ai - 3, ai - 4<l and bi - 1 = bi - 2 = bi - 3 = bi - 4 = 0
bi = bi - 1 otherwise
You are given arrays a and b' of the same length. Find two integers l and r (l ≤ r), such that applying the algorithm
described above will yield an array b equal to b'.

It's guaranteed that the answer exists.

### ideas

1. 有点混乱
2. 如果有连续4个1+一个0 => r的上界
3. 如果有连续4个0+一个1 => l的下界
4. 不对， 4个连续的0 + r <= min(a[i-4...i]) 这个是b[i] = 1充分条件，不是必要条件
5. 如果b[5] = 0, 那么 l >= max(1...5)
6. ....知道遇到一个b[i] = 1, 此时可以确定的是l的上界 <= max ...
7. 然后b[i] = 1, l >  max(a[i-4....i])
8. 到此时，没有r的信息
9. 如果 b[i+1] = 
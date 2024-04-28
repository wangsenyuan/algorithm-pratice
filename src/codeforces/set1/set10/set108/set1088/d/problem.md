This is an interactive problem!

Ehab plays a game with Laggy. Ehab has 2 hidden integers (𝑎,𝑏)
. Laggy can ask a pair of integers (𝑐,𝑑)
and Ehab will reply with:

1 if 𝑎⊕𝑐>𝑏⊕𝑑
.
0 if 𝑎⊕𝑐=𝑏⊕𝑑
.
-1 if 𝑎⊕𝑐<𝑏⊕𝑑
.
Operation 𝑎⊕𝑏
is the bitwise-xor operation of two numbers 𝑎
and 𝑏
.

Laggy should guess (𝑎,𝑏)
with at most 62 questions. You'll play this game. You're Laggy and the interactor is Ehab.

It's guaranteed that 0≤𝑎,𝑏<230
.

### ideas

1. 从只能问62次，可以看出来应该是按位来问
2. 假设只有一位的情况下(a, b) = (0, 0) (0, 1) (1, 0), (1, 1)
3. 如果let (c, d) = (0, 1)
4. 根据a ^ 0 ><= b ^ 1 的关系，确实可以判断处a, b
5. 如果 a ^ 0 > b ^ 1, 那么 (a, b) = (1, 0)
6. 如果 a ^ 0 < b ^ 1 那么 (a, b) = (0, 0)
7. 如果 a ^ 0 = b ^ 1, 那么 (a, b) 有可能 = （0，1）， 或者 (1, 0) 所以还要再问一次
8. 但前提是知道高位，并且使用能够将高位相等的c，d
9. ok， 从高位到低位
10. 假设到(i)的时候，已经知道了a/b >= i的情况， 只要让c的高位和a的高位一致，b和d的一致，就可以询问当前位
Little X has n distinct integers: p1, p2, ..., pn. He wants to divide all of them into two sets A and B. The following two conditions must be satisfied:

If number x belongs to set A, then number a - x must also belong to set A.
If number x belongs to set B, then number b - x must also belong to set B.
Help Little X divide the numbers into two sets or determine that it's impossible.

### ideas
1. 在a中，为true，那么第一个相当于 x => a-x, 第二个条件相当于 !x => b - x
2. 就是说，如果x在A中，那么a-x必须在A中，也就是说如果 a - x 不在A中（在B中，那么x必须在B中）
3. 
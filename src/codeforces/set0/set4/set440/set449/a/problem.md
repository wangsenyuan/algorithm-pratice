Jzzhu has a big rectangular chocolate bar that consists of n × m unit squares. He wants to cut this bar exactly k times.
Each cut must meet the following requirements:

each cut should be straight (horizontal or vertical);
each cut should go along edges of unit squares (it is prohibited to divide any unit chocolate square with cut);
each cut should go inside the whole chocolate bar, and all cuts must be distinct.

Imagine Jzzhu have made k cuts and the big chocolate is splitted into several pieces. Consider the smallest (by area)
piece of the chocolate, Jzzhu wants this piece to be as large as possible. What is the maximum possible area of smallest
piece he can get with exactly k cuts? The area of a chocolate piece is the number of unit squares in it.

### ideas

1. 为了让最小的面积最大，应该要尽可能的平均
2. 进行k次cut后，一共有多少块面积呢？
3. 假设最后一共有l块，那么平均每块有 n * m / l
4. 为了最后的面积大， 需要l尽可能的小
5. 假设横切了a,竖切了b
6. l = (a + 1) * (b + 1)
7. 然后需要考虑 a > n, b > m的情况
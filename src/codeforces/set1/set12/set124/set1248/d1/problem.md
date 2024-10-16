Vasya is an experienced developer of programming competitions' problems. As all great minds at some time, Vasya faced a creative crisis. To improve the situation, Petya gifted him a string consisting of opening and closing brackets only. Petya believes, that the beauty of the bracket string is a number of its cyclical shifts, which form a correct bracket sequence.

To digress from his problems, Vasya decided to select two positions of the string (not necessarily distinct) and swap characters located at this positions with each other. Vasya will apply this operation exactly once. He is curious what is the maximum possible beauty he can achieve this way. Please help him.

We remind that bracket sequence 𝑠
 is called correct if:

𝑠
 is empty;
𝑠
 is equal to "(𝑡
)", where 𝑡
 is correct bracket sequence;
𝑠
 is equal to 𝑡1𝑡2
, i.e. concatenation of 𝑡1
 and 𝑡2
, where 𝑡1
 and 𝑡2
 are correct bracket sequences.
For example, "(()())", "()" are correct, while ")(" and "())" are not.

The cyclical shift of the string 𝑠
 of length 𝑛
 by 𝑘
 (0≤𝑘<𝑛
) is a string formed by a concatenation of the last 𝑘
 symbols of the string 𝑠
 with the first 𝑛−𝑘
 symbols of string 𝑠
. For example, the cyclical shift of string "(())()" by 2
 equals "()(())".

Cyclical shifts 𝑖
 and 𝑗
 are considered different, if 𝑖≠𝑗
.

### ideas
1. 考虑一个string的beauty = 可以把字符串分成个有效的（level = 0）的序列 
2. 那如何在修改的情况下知道最大的beauty呢？
3. 假设确定了开始端点，那么结束端点也是定的(i + n), 然后，这些区间内，level = 0 的位置
4. 且没有 < 0 的位置出现
5. 在n 〈= 500 的时候，可以进行brute force
6. 交换(i, j) 相当于在位置 (i, j)区间内+1/-1 （ 这个是范围更新）
7. 主要是0的数量不好搞～


### solution(easy)
Note first that the number of opening brackets must be equal to the number of closing brackets, otherwise the answer is always 0
. Note that the answer to the question about the number of cyclic shifts, which are correct bracket sequences, equals the number of minimal prefix balances. For example, for string )(()))()((, the array of prefix balances is [-1, 0, 1, 0, -1, -2, -1, -2, -1, 0], and the number of cyclic shifts, 2
 – the number of minimums in it (−2
). Now we have a solution of complexuty (𝑛3)
: let's iterate over all pairs of symbols that can be swapped. Let's do this and find the number of cyclic shifts that are correct bracket sequences according to the algorithm described above.


### solution(hard)
Note first that the number of opening brackets must be equal to the number of closing brackets, otherwise the answer is always 0
. The answer for a bracket sequence is the same as the answer for any of its cyclic shifts. Then find 𝑖
— index of the minimum balance and perform a cyclic shift of the string by i to the left. The resulting line never has a balance below zero, which means it is a correct bracket sequence. Further we will work only with it

Let us draw on the plane the prefix balances (the difference between the number of opening and closing brackets) of our correct bracket sequence in the form of a polyline. It will have a beginning at (0,0)
, an end — at (2𝑛,0)
 and its points will be in the upper half-plane. It can easily be shown that the answer to the question of the number of cyclic shifts being correct bracket sequence is the number of times how much minimum balance is achieved in the array of prefix balances. For example, for string )(()))()((, the array of prefix balances is [-1, 0, 1, 0, -1, -2, -1, -2, -1, 0], and the number of cyclic shifts, 2
 – the number of minimums in it (−2
).

After swapping two different brackets (')' and '('), either on some sub-segment in an array of balances 2
 is added, or on some sub-segment −2
 is added.


In the first case (as you can see from the figure above), the minimum remains the same as it was — 0
, and its number does not increase due to the addition of 2
 to some sub-section of the array. So, there is no profit in swapping brackets in this case.

In the second case, there are three options: the minimum becomes equal to −2
, −1
 or 0
. In the first case, the minimum reaches a value equal to −2
 only at those points at which there was previously value equal to 0
, so this answer will less or equal than the one that would have turned out, if no operations were applied to the array at all.

If the minimum becomes equal to −1
, then on the segment in the balance array on which 2
 was deducted as a result of this operation there could not be balances equal to 0
, otherwise the new minimum would become equal to −2
. So, if 0=𝑋0<𝑋1<⋯<𝑋𝑘=2𝑛
 — positions of minimums in the array of prefix balances, then the operation −2
 was performed on the segment [𝐿,𝑅]
 (𝑋𝑖<𝐿≤𝑅<𝑋𝑖+1
 for some 𝑖
). After this operation, the number of local minimums will be equal to the number of elements of the array of balance on the segment [𝐿;𝑅]
, equal to 1
. Since this number shall be maximized, the right border swall be maximised and the left border shall be minimised. So, for some fixed 𝑖
 it is optimal to swap 𝑋𝑖
-th and 𝑋𝑖+1−1
-th brackets (see following figure)


If the minimum remains equal to 0
, then using the reasoning similar to the reasoning above, it can be proven that for some 𝑖
 if we consider points 𝑋𝑖+1=𝑌0<𝑌1<⋯<𝑌𝑚=𝑋𝑖+1−1
 — positions of 1
 in the array of balances on the segment [𝑋𝑖;𝑋𝑖+1]
, it is optimal for some 𝑗
 to swap 𝑌𝑗
 and 𝑌𝑗+1−1
 brackets (see following figure)


Thus, we have the solution of (𝑛)
 complexity — we shall find all 𝑋𝑖
, 𝑌𝑖
 and count on each of the segments of the form [𝑋𝑖;𝑋𝑖+1]
 or [𝑌𝑖;𝑌𝑖+1]
 the number of elements equal to 0
, 1
, 2
, then according to reasonings above, maximal answer can be found.
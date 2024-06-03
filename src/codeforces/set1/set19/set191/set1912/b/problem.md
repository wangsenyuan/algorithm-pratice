An aircraft manufacturing company wants to optimize their products for passenger airlines. The company's latest research shows that most of the delays happen because of slow boarding.

Most of the medium-sized aircraft are designed with 3-3 seat layout, meaning each row has 6 seats: 3 seats on the left side, a single aisle, and 3 seats on the right side. At each of the left and right sides there is a window seat, a middle seat, and an aisle seat. A passenger that boards an aircraft assigned to an aisle seat takes significantly less time than a passenger assigned to a window seat even when there is no one else in the aircraft.

The company decided to compute an inconvenience of a layout as the total sum of distances from each of the seats of a single row to the closest aisle. The distance from a seat to an aisle is the number of seats between them. For a 3-3 layout, a window seat has a distance of 2, a middle seat — 1, and an aisle seat — 0. The inconvenience of a 3-3 layout is (2+1+0)+(0+1+2)=6
. The inconvenience of a 3-5-3 layout is (2+1+0)+(0+1+2+1+0)+(0+1+2)=10
.

Formally, a layout is a sequence of positive integers 𝑎1,𝑎2,…,𝑎𝑘+1
 — group 𝑖
 having 𝑎𝑖
 seats, with 𝑘
 aisles between groups, the 𝑖
-th aisle being between groups 𝑖
 and 𝑖+1
. This means that in a layout each aisle must always be between two seats, so no aisle can be next to a window, and no two aisles can be next to each other.

The company decided to design a layout with a row of 𝑛
 seats, 𝑘
 aisles and having the minimum inconvenience possible. Help them find the minimum inconvenience among all layouts of 𝑛
 seats and 𝑘
 aisles, and count the number of such layouts modulo 998244353
.

### ideas
1. 有点糊涂，不是除去两边后，其他的尽量相等吗？
2. n = 10, k = 3
3. 考虑几种layout 1,4,4,1, 2,3,3,2
4. 0 + (1 + 1) + (1 + 1) + 0 = 4
5. 1 + 1 + 1 + 1 = 4
6. (1, 3, 4, 2) = (0 + 1 + (1 + 1) + 1) = 4 怎么也是4
7. 没招了～
8. 这个最小值，貌似是可以被计算出的。尽量平均分配，两头最多到平均值
9. 但是怎么在最小值的情况下，计算能够得到最小值的数量呢？


### solution
Consider any group not located next to a window. If the group consists of t seats, then the inconvenience
part of the group is s(t1) + s(t2), where t1 and t2 are chosen optimally such that t1 + t2 = t, and
s(x) = x(x−1)
2
.
For a group of t seats next to a window, the inconvenience part is s(t).
So the inconvenience is the sum of some s(r1) + s(r2) + . . . + s(r2k). Each ri
is either one of the two sizes
of groups next to a window, or one of the 2 · (k − 1) subgroups created by division of middle groups.
Consider we’ve chosen r, a = max ri and b = min ri such that a−b ≥ 2, then s(a−1)+s(b+1) < s(a)+s(b),
because
2 · (s(a − 1) + s(b + 1)) = (a − 1) · (a − 2) + (b + 1) · b
= (a − 1) · a − 2 · (a − 1) + (b − 1) · b + 2b
= 2 · (s(a) + s(b)) + 2 · (b − a + 1)
< 2 · (s(a) + s(b))
That’s because (b − a + 1) is negative.
So the only way to divide is almost equally. The multiset of r: n mod (2k) times ⌈
n
2k
⌉
and 2k−(n mod (2k))
times ⌊
n
2k
⌋
.
We can iterate over four ways to choose the window groups, not considering, when the window group is
of size of 0. Then we need to calculate how to arrange middle groups.
We have (k − 1) middle groups, and x subgroups of size ⌊
n
2k
⌋
= t and y subgroups of size t + 1.
Each group can be of size either 2t, or 2t+1, or 2t+2. If we iterate with number of groups of size 2t+1 say
g, then x−g
2 = f groups will be of size 2t and y−g
2 = h groups of size 2t + 2. The number of arrangements
with the fixed g is the product of some binomial coefficients, for example (
f+g+h
g
)
·
(
f+h
f
)
. We should only
consider cases, when f and h are non-negative integers. And we should handle the case when 2t equals
to zero, it means f needs to be zero as well.
The binomial coefficient can either be recalculated when iterating g, by several multiplications and
divisions. Another way to do it is to pre-calculate all factorials and their inverse modulo the given prime
number, and use them to calculate binomial coefficient in O(1). For a linear algorithm to pre-calculate
all inverse values modulo prime p you can use the following recurrent formula:
x
−1 ≡ −(p mod x)
−1
·
⌊
p
x
⌋
(mod p)
note that (p mod x)
−1
can be calculated before x
−1
. The formula comes from relation p mod x = p−
⌊
p
x
⌋
·x.
The total time complexity for a single test case could be O(k) or O(k log k), depending on how you handle
the binomial coefficients.
Page 1 of 9
IC
You are given two integers 𝑛
 and 𝑚
. Calculate the number of pairs of arrays (𝑎,𝑏)
 such that:

the length of both arrays is equal to 𝑚
;
each element of each array is an integer between 1
 and 𝑛
 (inclusive);
𝑎𝑖≤𝑏𝑖
 for any index 𝑖
 from 1
 to 𝑚
;
array 𝑎
 is sorted in non-descending order;
array 𝑏
 is sorted in non-ascending order.
As the result can be very large, you should print it modulo 109+7
.


### ideas
1. a1 <= b1, a2 <= b2
2. a1 <= a2 and b1 >= b2
3. 那么考虑最后两个数 a[m] <= b[m]
4. 那么 a[m]前面的数 C(a[m], m - 1)
5. b[m]后面的数  = C(n - b[m] + 1, m - 1)
6. 不对，a[m] < m - 1 就没有结果了，显然不对
7. 数组a可以使a[m]个数
8. 1，2，3.。。 a[m]，C(a[m] + m - 1, m - 1)
9. 
For each positive integer n consider the integer ψ(n) which is obtained from n by replacing every digit a in the decimal notation of n with the digit (9  -  a). We say that ψ(n) is the reflection of n. For example, reflection of 192 equals 807. Note that leading zeros (if any) should be omitted. So reflection of 9 equals 0, reflection of 91 equals 8.

Let us call the weight of the number the product of the number and its reflection. Thus, the weight of the number 10 is equal to 10·89 = 890.

Your task is to find the maximum weight of the numbers in the given range [l, r] (boundaries are included).

Input
Input contains two space-separated integers l and r (1 ≤ l ≤ r ≤ 109) — bounds of the range.

Output
Output should contain single integer number: maximum value of the product n·ψ(n), where l ≤ n ≤ r.

### ideas
1. 44 * 55 = 2420
2. 88 * 11 = 968
3. 基本上还是越靠近的越好
4. 如果l.r中存在一个数n = 5555 或者 4444..那么这个是可以提供出最大的结果
5. 5...16   15 * 84 = 1260
6. 14 * 85 = 1190
7. 如果r最高位是9
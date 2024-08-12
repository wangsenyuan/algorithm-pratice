You are given two integers 𝑙
 and 𝑟
, where 𝑙<𝑟
. We will add 1
 to 𝑙
 until the result is equal to 𝑟
. Thus, there will be exactly 𝑟−𝑙
 additions performed. For each such addition, let's look at the number of digits that will be changed after it.

For example:

if 𝑙=909
, then adding one will result in 910
 and 2
 digits will be changed;
if you add one to 𝑙=9
, the result will be 10
 and 2
 digits will also be changed;
if you add one to 𝑙=489999
, the result will be 490000
 and 5
 digits will be changed.
Changed digits always form a suffix of the result written in the decimal system.

Output the total number of changed digits, if you want to get 𝑟
 from 𝑙
, adding 1
 each time.

### ideas
1. ?999 + 1 => 4
2. 假设能够计算出从0到x的结果
3. 那么solve(l, r) = f(r) - f(l-1)
4. 假设要计算 f(1000) = f(99) + 3
5. f(2000) = f(1000) + f(1000) = 2 * f(1000)
6. f(abcd) = a * f(1000) + b * f(100) + c * f(10) + f(d)
7. f(1000) = f(99) + 3 = 9 * f(10) + 9 + 3 = 9 * (9 + 2) + 9 + 3
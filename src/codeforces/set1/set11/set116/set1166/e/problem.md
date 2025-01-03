Dora the explorer has decided to use her money after several years of juicy royalties to go shopping. What better place to shop than Nlogonia?

There are 𝑛
 stores numbered from 1
 to 𝑛
 in Nlogonia. The 𝑖
-th of these stores offers a positive integer 𝑎𝑖
.

Each day among the last 𝑚
 days Dora bought a single integer from some of the stores. The same day, Swiper the fox bought a single integer from all the stores that Dora did not buy an integer from on that day.

Dora considers Swiper to be her rival, and she considers that she beat Swiper on day 𝑖
 if and only if the least common multiple of the numbers she bought on day 𝑖
 is strictly greater than the least common multiple of the numbers that Swiper bought on day 𝑖
.

The least common multiple (LCM) of a collection of integers is the smallest positive integer that is divisible by all the integers in the collection.

However, Dora forgot the values of 𝑎𝑖
. Help Dora find out if there are positive integer values of 𝑎𝑖
 such that she beat Swiper on every day. You don't need to find what are the possible values of 𝑎𝑖
 though.

Note that it is possible for some values of 𝑎𝑖
 to coincide in a solution.

 ### ideas
 1. 如果有一个i，在所有的天数内，都被购买了，那么就让它等于无穷，其他的都等于1
 2. 这里不需要具体的a[i],那么我们可以假设a[i] 相互之间可以整除（但是这样子，不一定能work）
 3. 比如 2, 3, 4那么选择2, 3，就比4更大
 4. 如果存在两个不相交的集合 => false
 5. 所以答案存在的前提是，所有集合，两两间都有重复的元素
 6. ab, bc, ac 如果是这样的一个组合，貌似是可以的
 7. 234   23/4   24/3, 34/2
 8. 假设第一个集合和第二个集合的公共元素是a
 9. a, b, c, d, e 
 10. 假设选中了a, b => lcm(a, b) > lcm(c, d, e) = lcm(c, lcm(d, e)) = c * lcm(d, e) / gcd(c, lcm(d, e))
 11.   = c * (d * e / gcd(d, e)) / gcd(c, lcm(d, e))
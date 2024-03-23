Roman is a young mathematician, very famous in Uzhland. Unfortunately, Sereja doesn't think so. To make Sereja change
his mind, Roman is ready to solve any mathematical problem. After some thought, Sereja asked Roma to find, how many
numbers are close to number n, modulo m.

Number x is considered close to number n modulo m, if:

it can be obtained by rearranging the digits of number n,
it doesn't have any leading zeroes,
the remainder after dividing number x by m equals 0.
Roman is a good mathematician, but the number of such numbers is too huge for him. So he asks you to help him.

### ideas

1. x的digits和n的digits相同（排除0的情况)
2. x % m = 0
3. m <= 100
4. 假设从高位到低位dp[i][r] 是前i个，对m余数为r的计数
5. 这里缺个信息，就是digits要保证和n的一致
6. 假设n的digits可以用0...9的个数表示
7. 好像用一个10维的数组是可以的，如果所有的数都相同，2 * 2 * 2 。。。 * 2 = 2 ** 10 = 1e6

### solution

This problem we can be attributed to the dynamic programming. We must using mask and dynamic.

We have dynamic dp[i][x], when i — mask of reshuffle and x — remainder on dividing by m.

if we want to add number a[j], we must using it:

dp[i or (1 shl (j-1)),(x*10+a[j]) mod m] := dp[i or (1 shl (j-1)),(x*10+a[j]) mod m] + dp[i,x];
In the end we must answer to divide by the factorial number of occurrences of each digit.

for i:=0 to 9 do

for j:=2 to b[i] do

     ans:=ans div int64(j);

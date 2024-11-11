Employees of JebTrains are on their way to celebrate the 256-th day of the year! There are 𝑛
 employees and 𝑘
 teams in JebTrains. Each employee is a member of some (exactly one) team. All teams are numbered from 1
 to 𝑘
. You are given an array of numbers 𝑡1,𝑡2,…,𝑡𝑛
 where 𝑡𝑖
 is the 𝑖
-th employee's team number.

JebTrains is going to rent a single bus to get employees to the feast. The bus will take one or more rides. A bus can pick up an entire team or two entire teams. If three or more teams take a ride together they may start a new project which is considered unacceptable. It's prohibited to split a team, so all members of a team should take the same ride.

It is possible to rent a bus of any capacity 𝑠
. Such a bus can take up to 𝑠
 people on a single ride. The total cost of the rent is equal to 𝑠⋅𝑟
 burles where 𝑟
 is the number of rides. Note that it's impossible to rent two or more buses.

Help JebTrains to calculate the minimum cost of the rent, required to get all employees to the feast, fulfilling all the conditions above.

### ideas
1. 首先bus的容量不能小于最大team的人数（否则，就必须split这个team，但是这是不允许的）
2. 但是也不能超过2 * 最大team的人数，这样就浪费了，（不能由3个team在一趟bus上）
3. 感觉可以brute force？
4. 在s固定的时候，希望跑的次数越少越好
5. 
There are students studying at your university𝑛
students. Programming skills𝑖
-th student is equal to𝑎𝑖
. As a coach, you want to divide them into teams to prepare them for the upcoming ICPC finals. Just imagine how good this university is if it has2 ⋅105
students ready for the finals!

Each team must consist of at least three students . Each student must belong to exactly one team . The diversity of a team is the difference between the maximum programming skill of a student belonging to that team and the minimum programming skill of a student belonging to that team (in other words, if a team consists of𝑘
students with programming skills𝑎 [𝑖1] , 𝑎 [𝑖2] , … , 𝑎 [𝑖𝑘]
, then the diversity of this team is equal tomax𝑗 = 1𝑘𝑎 [𝑖𝑗] −min𝑗 = 1𝑘𝑎 [𝑖𝑗]
).

Total diversity is the sum of the diversity across all assembled teams.

Your task is to minimize the total diversity of student divisions and find the optimal way to divide them.

### ideas
1. 这个貌似没办法2分。
2. 先排个序，如果只有一个team = a[n] - a[1]
3. 假设在i处分成两个组，那么就有 a[i] - a[1] + a[n] - a[i+1] = a[n] - a[] - (a[i+1] - a[i])
4. 貌似，只要a[i+1] - a[i] > 0, 那么在i处分离就是好的
5. 这里另外一个限制就是每个组必须有3个人。
6. 
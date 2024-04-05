There are n students at Berland State University. Every student has two skills, each measured as a number: ai — the
programming skill and bi — the sports skill.

It is announced that an Olympiad in programming and sports will be held soon. That's why Berland State University should
choose two teams: one to take part in the programming track and one to take part in the sports track.

There should be exactly p students in the programming team and exactly s students in the sports team. A student can't be
a member of both teams.

The university management considers that the strength of the university on the Olympiad is equal to the sum of two
values: the programming team strength and the sports team strength. The strength of a team is the sum of skills of its
members in the corresponding area, so the strength of the programming team is the sum of all ai and the strength of the
sports team is the sum of all bi over corresponding team members.

Help Berland State University to compose two teams to maximize the total strength of the university on the Olympiad.

### ideas

1. 这里有个限制条件是参加program的人是p个人，参加sport的有s个人， 还有一部分人，不参加 (p + s < n)
2. dp[i][j][k][x] 表示在前i个人中，其中j个人参加项目1，k个人参加项目2，且项目1的能力总值为x时，项目2的最大能力值
3. dp[i][j][k][x] = dp[i-1][j][k][x], dp[i-1][j-1][k][x-a[i]], dp[i-1][j][k-1][x] + b[i]
4. 分别表示，当前用户不参与，当前用户参加项目1，当前用户参加项目2
5. TLE。
6. 这里哪些信息没有被考虑到？
7. 假设一个人的编程能力很强，当时sport很弱，那么让他参加项目1更好，反之项目2更好
8. 一个人只要参加了，那么他就有贡献（a[i]/b[i])
9. 先按照a[i]降序排列，那么假设选择了p个人参加项目1，s个人参加项目2，那么剩余的(n - p - s)
   个人的编程能力，一定是小于（等于）这(p)
   个人
10. 可以反证，如果不是，那么可以通过将他和其中的最后一名交换，从而得到更优的答案
11. 同样的，也可以得到，它们的sport的能力值，也必然小于选中的s个人
12. 那先选出最多的p个人，然后再剩余的人中，选出最大的s个人
13. 假设通过某种方式已经选中了p + s 个人，考虑其中某个选中项目1的i和项目2的j
14. 那么必然有 a[i] - b[i] > a[j] - b[j]
15. a[i] - b[j] > b[i] - b[j]
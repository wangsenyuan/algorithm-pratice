Polycarp is a head of a circus troupe. There are 𝑛
 — an even number — artists in the troupe. It is known whether the 𝑖
-th artist can perform as a clown (if yes, then 𝑐𝑖=1
, otherwise 𝑐𝑖=0
), and whether they can perform as an acrobat (if yes, then 𝑎𝑖=1
, otherwise 𝑎𝑖=0
).

Split the artists into two performances in such a way that:

each artist plays in exactly one performance,
the number of artists in the two performances is equal (i.e. equal to 𝑛2
),
the number of artists that can perform as clowns in the first performance is the same as the number of artists that can perform as acrobats in the second performance.


### ideas
1. 假设在team1中分配x个小丑，m - x 中不能能有小丑
2. 那么在team2中分配x个acrobat， m-x中不能有acrobat
3. 可以肯定的一点是， 两种技能都会的，不能出现在后半段。 否则就不能满足条件
4. 也就是有 2 * x >= 两种技能都会的人
5. 假设有4个人，两种技能都会，但是其中三个人分配给了team2， 只有一个人分配给了team1， x >= 3
6. team1还有两个小丑（只会小丑技能）才能保证小丑的人数和team2 acrobat的人数相等
7. 那么交换一个小丑和下方都会的人， 会造成结果不对吗？
8. 似乎是不能交换的，交换后，上方没有少小丑（新交换进入的，仍然可以扮演小丑），但是交换下去的，不能扮演acrobat，所以少了1个
9. x个小丑，m-x中不能有小丑
10. 考虑后半段，其中u个人可以扮演acrobat，v个人啥都不会
11. m-x考虑，其中k个人可以扮演小丑，但是w个人啥都不大会
12. 如果使用1个交换一个可以扮演小丑（对结果是没有影响的）=》因为上下都+1
13. 那么肯定会进入一个状态，u = 0,或者 k = 0
14. 如果小丑的人数少余acrobat的人数，那么就把小丑都放在team1中
15. 然后放置 acrobat - clown的人数的，两个都会的人
16. 
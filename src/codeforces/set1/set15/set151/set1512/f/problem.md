Polycarp is wondering about buying a new computer, which costs 𝑐
 tugriks. To do this, he wants to get a job as a programmer in a big company.

There are 𝑛
 positions in Polycarp's company, numbered starting from one. An employee in position 𝑖
 earns 𝑎[𝑖]
 tugriks every day. The higher the position number, the more tugriks the employee receives. Initially, Polycarp gets a position with the number 1
 and has 0
 tugriks.

Each day Polycarp can do one of two things:

If Polycarp is in the position of 𝑥
, then he can earn 𝑎[𝑥]
 tugriks.
If Polycarp is in the position of 𝑥
 (𝑥<𝑛
) and has at least 𝑏[𝑥]
 tugriks, then he can spend 𝑏[𝑥]
 tugriks on an online course and move to the position 𝑥+1
.
For example, if 𝑛=4
, 𝑐=15
, 𝑎=[1,3,10,11]
, 𝑏=[1,2,7]
, then Polycarp can act like this:

On the first day, Polycarp is in the 1
-st position and earns 1
 tugrik. Now he has 1
 tugrik;
On the second day, Polycarp is in the 1
-st position and move to the 2
-nd position. Now he has 0
 tugriks;
On the third day, Polycarp is in the 2
-nd position and earns 3
 tugriks. Now he has 3
 tugriks;
On the fourth day, Polycarp is in the 2
-nd position and is transferred to the 3
-rd position. Now he has 1
 tugriks;
On the fifth day, Polycarp is in the 3
-rd position and earns 10
 tugriks. Now he has 11
 tugriks;
On the sixth day, Polycarp is in the 3
-rd position and earns 10
 tugriks. Now he has 21
 tugriks;
Six days later, Polycarp can buy himself a new computer.
Find the minimum number of days after which Polycarp will be able to buy himself a new computer.

### ideas
1. 一个观察，如果要在第i个职位上获取到足够的钱，那么似乎是越快得到职位越好
2. 这个主要在于，到达职位i，所需要花费的总的成本是固定的sum(b[1], ... b[i-1])
3. 但是收益，是在职位i最高
4. 也就是说，应该计算到到职位i，所需要的最小天数是多少。这个可以贪心计算出来
5. 假设到到职位i-1的最快的天数是 f(i-1) = {剩余的钱数,天数}
6. f(i) = f(i-1).first + d * a[i-1] >= a[i], 从而计算出新的结果
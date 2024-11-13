Imagine a game where you play as a character that has two attributes: "Strength" and "Intelligence", that are at zero level initially.

During the game, you'll acquire 𝑚
 attribute points that allow you to increase your attribute levels — one point will increase one of the attributes by one level. But sometimes, you'll encounter a so-called "Attribute Checks": if your corresponding attribute is high enough, you'll pass it; otherwise, you'll fail it.

Spending some time, you finally prepared a list which contains records of all points you got and all checks you've met. And now you're wondering: what is the maximum number of attribute checks you can pass in a single run if you'd spend points wisely?

Note that you can't change the order of records.

Input
The first line contains two integers 𝑛
 and 𝑚
 (1≤𝑚≤5000
; 𝑚<𝑛≤2⋅106
) — the number of records in the list and the total number of points you'll get during the game.

The second line contains 𝑛
 integers 𝑟1,𝑟2,…,𝑟𝑛
 (−𝑚≤𝑟𝑖≤𝑚
), where 𝑟𝑖
 encodes the 𝑖
-th record:

If 𝑟𝑖=0
, then the 𝑖
-th record is an acquiring one attribute point. You can spend to level up either Strength or Intelligence;
If 𝑟𝑖>0
, then it's an Intelligence check: if your Intelligence level is greater than or equal to |𝑟𝑖|
, you pass.
If 𝑟𝑖<0
, then it's a Strength check: if your Strength level is greater than or equal to |𝑟𝑖|
, you pass.
Additional constraint on the input: the sequence 𝑟1,𝑟2,…,𝑟𝑛
 contains exactly 𝑚
 elements equal to 0
.


### ideas
1. dp[i][j] = 到i为止有x=j时的最大等分
2. 那么如果 r[i] > 0, 那么 j, j+1, j + 2. ... +1 （range update ?)
3. 如果 r[i] < 0, 那么（假设当前共有v个0， 那么就是 y = (v - x) + 1， 
4. 那么 0....v + r[i] (也是range update)
5. 当r[i] =0, v++ 
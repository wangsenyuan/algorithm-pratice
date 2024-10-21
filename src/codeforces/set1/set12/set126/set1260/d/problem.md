https://codeforces.com/problemset/problem/1260/D

You are playing a computer game in which you control𝑚
soldiers. Each soldier has his own dexterity.𝑎𝑖
.

The level you are going through can be represented as a straight line segment between points0
(where you and your squad are originally) and𝑛 + 1
(where the level boss is located).

The level is full𝑘
traps. Each trap is characterized by three numbers𝑙𝑖
,𝑟𝑖
And𝑑𝑖
.𝑙𝑖
— the point at which the trap is located, and𝑑𝑖
- danger of a trap: when a soldier with less dexterity𝑑𝑖
steps on it (that is, moves to the point𝑙𝑖
), the trap kills him immediately. Luckily, you can disarm traps: if you reach the point𝑟𝑖
, you can disarm the trap and it will no longer pose a threat. Traps do not affect you, only your soldiers.

Do you have𝑡
seconds to complete the level — that is, to be able to deliver some soldiers to the boss. Before the level starts, you choose which soldiers you will lead and which will remain outside the level. After that, you need to bring all the selected soldiers to the boss. To do this, you can do the following:

if you are at the point𝑥
, you can move to𝑥 + 1
or in𝑥 − 1
This action takes one second;
if you are at the point𝑥
and your squad is at the point𝑥
, you can move to𝑥 + 1
or𝑥 − 1
with the squad for one second. You cannot perform this action if it is dangerous for soldiers (i.e. the point you are moving to contains an unarmed trap with a magnitude of𝑑𝑖
more dexterity of any soldier in the squad). This action takes one second;
if you are at the point𝑥
and there is a trap𝑖
With𝑟𝑖= 𝑥
, you can disarm this trap. This action does not waste time.
Note that after each action, both your coordinate and your squad's coordinate are integers.

You must select the maximum number of soldiers you can bring from the point.0
right on target𝑛 + 1
no more than for𝑡
seconds.

### ideas
1. 可以二分；如果能够把a[i]及以后的士兵带过去，那么就可以把a[i+1]带过去
2. 假设目前希望带过去的士兵的抗毒性为w，所有毒性超过w的陷阱，都需要被处理掉
3. 且为了减少时间，应该把士兵尽可能带到近的位置
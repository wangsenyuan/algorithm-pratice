Polycarp plays a computer game (yet again). In this game, he fights monsters using magic spells.

There are two types of spells: fire spell of power 𝑥
 deals 𝑥
 damage to the monster, and lightning spell of power 𝑦
 deals 𝑦
 damage to the monster and doubles the damage of the next spell Polycarp casts. Each spell can be cast only once per battle, but Polycarp can cast them in any order.

For example, suppose that Polycarp knows three spells: a fire spell of power 5
, a lightning spell of power 1
, and a lightning spell of power 8
. There are 6
 ways to choose the order in which he casts the spells:

first, second, third. This order deals 5+1+2⋅8=22
 damage;
first, third, second. This order deals 5+8+2⋅1=15
 damage;
second, first, third. This order deals 1+2⋅5+8=19
 damage;
second, third, first. This order deals 1+2⋅8+2⋅5=27
 damage;
third, first, second. This order deals 8+2⋅5+1=19
 damage;
third, second, first. This order deals 8+2⋅1+2⋅5=20
 damage.
Initially, Polycarp knows 0
 spells. His spell set changes 𝑛
 times, each time he either learns a new spell or forgets an already known one. After each change, calculate the maximum possible damage Polycarp may deal using the spells he knows.

 ### ideas
 1. 在给定的一组spells中，最优的策略是什么？
 2. 如果同时有fire, lighting, 那么先使用lighting是更优的策略，这是因为使用lighting，可以double 后面（不管是lighting，还是fire）spell的伤害
 3. 假设有x个lighting，那么最多可以double x个spell的强度。所以，应该找出这个set中，最大的x个spell，然后double它们
 4. 当然如果全部是lighting，那么应该找x-1个
 5. 如果维护了全部spell的集合，那么要能快速的知道有多少个lighting的spell（这个可以单独维护一个计数）然后从这个spell集合中找出，最大的x个spell
 6. 使用avl tree。
 7. 也可以用segment tree（avl tree似乎更容易实现）
Polycarpus has a complex electronic device. The core of this device is a circuit board. The board has 109
 contact points which are numbered from 1
 to 109
. Also there are 𝑛
 wires numbered from 1
 to 𝑛
, each connecting two distinct contact points on the board. An electric signal can pass between wires 𝐴
 and 𝐵
 if:

either both wires share the same contact point;
or there is a sequence of wires starting with 𝐴
 and ending with 𝐵
, and each pair of adjacent wires in the sequence share a contact point.
The picture shows a circuit board with 5
 wires. Contact points with numbers 2,5,7,8,10,13
 are used. Here an electrical signal can pass from wire 2
 to wire 3
, but not to wire 1
.
Currently the circuit board is broken. Polycarpus thinks that the board could be fixed if the wires were re-soldered so that a signal could pass between any pair of wires.

It takes 1
 minute for Polycarpus to re-solder an end of a wire. I.e. it takes one minute to change one of the two contact points for a wire. Any contact point from range [1,109]
 can be used as a new contact point. A wire's ends must always be soldered to distinct contact points. Both wire's ends can be re-solded, but that will require two actions and will take 2
 minutes in total.

Find the minimum amount of time Polycarpus needs to re-solder wires so that a signal can pass between any pair of wires. Also output an optimal sequence of wire re-soldering.

Input
The input contains one or several test cases. The first input line contains a single integer 𝑡
 — number of test cases. Then, 𝑡
 test cases follow.

The first line of each test case contains a single integer 𝑛
 (1≤𝑛≤105
) — the number of wires. The following 𝑛
 lines describe wires, each line containing two space-separated integers 𝑥𝑖,𝑦𝑖
 (1≤𝑥𝑖,𝑦𝑖≤109
, 𝑥𝑖≠𝑦𝑖
) — contact points connected by the 𝑖
-th wire. A couple of contact points can be connected with more than one wire.

Sum of values of 𝑛
 across all test cases does not exceed 105
.

Output
For each test case first print one line with a single integer 𝑘
 — the minimum number of minutes needed to re-solder wires so that a signal can pass between any pair of wires. In the following 𝑘
 lines print the description of re-solderings. Each re-soldering should be described by three integers 𝑤𝑗,𝑎𝑗,𝑏𝑗
 (1≤𝑤𝑗≤𝑛
, 1≤𝑎𝑗,𝑏𝑗≤109
). Such triple means that during the 𝑗
-th re-soldering an end of the 𝑤𝑗
-th wire, which was soldered to contact point 𝑎𝑗
, becomes soldered to contact point 𝑏𝑗
 instead. After each re-soldering of a wire it must connect two distinct contact points. If there are multiple optimal re-solderings, print any of them.

 ### ideas
 1. 不应该使用不存在的point，
 2. 假设共有x的component，必须把它们连起来变成一个component
 3. 有两类情况需要考虑，一种是当前的component有deg = 1的point
 4. 一类是没有，deg=1的始终是可以优先被拿来连接（消耗一次）这样子的好处是不会产生断开的情况
 5. 如果没有，那么必然是存在环的，这类的也可以被拿来使用（断开它后，也不会产生分裂的情况）
 6. 可以优先使用环的那种情况，这种情况下，会产生(有可能产生)deg = 1的节点
 7. 转来转去，头都快炸了～
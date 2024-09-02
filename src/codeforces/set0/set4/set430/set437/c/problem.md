On Children's Day, the child got a toy from Delayyy as a present. However, the child is so naughty that he can't wait to destroy the toy.

The toy consists of n parts and m ropes. Each rope links two parts, but every pair of parts is linked by at most one rope. To split the toy, the child must remove all its parts. The child can remove a single part at a time, and each remove consume an energy. Let's define an energy value of part i as vi. The child spend vf1 + vf2 + ... + vfk energy for removing part i where f1, f2, ..., fk are the parts that are directly connected to the i-th and haven't been removed.

Help the child to find out, what is the minimum total energy he should spend to remove all n parts.

### ideas
1. 感觉就应该按照v倒序排，然后依次删除
2. 假设不是这样子，存在一个最优的顺序，其中f(u) < f(v), 但是u在v的前面
3. 如果u和v相连，那么交换u、v是个更好的选择，因为u、v他们相邻，在交换他们前后、对其他的邻居节点是没有影响的
4. 但是先删除u，消耗时f(v), 先删除v消耗是f(u), 因为f(v) > f(u), 显然先删除v更合适
5. 如果他们不相连，交换他们的顺序显然不会更差
6. 前提是在(u, v)中间没有它们的公共相邻的点？
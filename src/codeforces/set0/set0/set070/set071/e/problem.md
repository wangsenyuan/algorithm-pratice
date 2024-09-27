There is the following puzzle popular among nuclear physicists.

A reactor contains a set of n atoms of some chemical elements. We shall understand the phrase "atomic number" as the number of this atom's element in the periodic table of the chemical elements.

You are allowed to take any two different atoms and fuse a new one from them. That results in a new atom, whose number is equal to the sum of the numbers of original atoms. The fusion operation can be performed several times.

The aim is getting a new pregiven set of k atoms.

The puzzle's difficulty is that it is only allowed to fuse two atoms into one, it is not allowed to split an atom into several atoms. You are suggested to try to solve the puzzle.

Input
The first line contains two integers n and k (1 ≤ k ≤ n ≤ 17). The second line contains space-separated symbols of elements of n atoms, which are available from the start. The third line contains space-separated symbols of elements of k atoms which need to be the result of the fusion. The symbols of the elements coincide with the symbols from the periodic table of the chemical elements. The atomic numbers do not exceed 100 (elements possessing larger numbers are highly unstable). Some atoms can have identical numbers (that is, there can be several atoms of the same element). The sum of numbers of initial atoms is equal to the sum of numbers of the atoms that need to be synthesized.

Output
If it is impossible to synthesize the required atoms, print "NO" without the quotes. Otherwise, print on the first line «YES», and on the next k lines print the way of synthesizing each of k atoms as equations. Each equation has the following form: "x1+x2+...+xt->yi", where xj is the symbol of the element of some atom from the original set, and yi is the symbol of the element of some atom from the resulting set. Each atom from the input data should occur in the output data exactly one time. The order of summands in the equations, as well as the output order does not matter. If there are several solutions, print any of them. For a better understanding of the output format, see the samples.


### ideas
1. dp[state] = 是否可行， state表示目标集合
2. 单个的比较容易判断， dp[state | j] = dp[state] = true and dp[j] = true and 它们需要的集合不重叠
3. 似乎有点麻烦～
4. 那是不是分成两半
5. 不大好分呐，主要是因为，假设按照目标atom分成两半，但是源头不一定恰好是一半（单至少是目标集合的大小）
6. 假设 a + b + c -> x
7. 同时 e + f -> x， 那么使用 a + b + c 和 e + f 对结果会有区别吗？
8. 最大流
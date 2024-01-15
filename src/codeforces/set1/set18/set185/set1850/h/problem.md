In order to win his toughest battle, Mircea came up with a great strategy for his army. He has 𝑛
soldiers and decided to arrange them in a certain way in camps. Each soldier has to belong to exactly one camp, and
there is one camp at each integer point on the 𝑥
-axis (at points ⋯,−2,−1,0,1,2,⋯
).

The strategy consists of 𝑚
conditions. Condition 𝑖
tells that soldier 𝑎𝑖
should belong to a camp that is situated 𝑑𝑖
meters in front of the camp that person 𝑏𝑖
belongs to. (If 𝑑𝑖<0
, then 𝑎𝑖
's camp should be −𝑑𝑖
meters behind 𝑏𝑖
's camp.)

Now, Mircea wonders if there exists a partition of soldiers that respects the condition and he asks for your help!
Answer "YES" if there is a partition of the 𝑛
soldiers that satisfies all of the 𝑚
conditions and "NO" otherwise.

Note that two different soldiers may be placed in the same camp.

### thoughts

1. 一条condition组成了一条有向边，u -> v 长度为w (w > 0)
2. 如果出现环肯定不行，否则的话，就是成立的
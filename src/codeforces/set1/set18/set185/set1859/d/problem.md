An incident occurred in Capygrad, the capital of Tyagoland, where all the capybaras in the city went crazy and started
throwing mandarins. Andrey was forced to escape from the city as far as possible, using portals.

Tyagoland is represented by a number line, and the city of Capygrad is located at point 0
. There are 𝑛
portals all over Tyagoland, each of which is characterised by four integers 𝑙𝑖
, 𝑟𝑖
, 𝑎𝑖
and 𝑏𝑖
(1≤𝑙𝑖≤𝑎𝑖≤𝑏𝑖≤𝑟𝑖≤109
). Note that the segment [𝑎𝑖,𝑏𝑖]
is contained in the segment [𝑙𝑖,𝑟𝑖]
.

If Andrey is on the segment [𝑙𝑖,𝑟𝑖]
, then the portal can teleport him to any point on the segment [𝑎𝑖,𝑏𝑖]
. Andrey has a pass that allows him to use the portals an unlimited number of times.

Andrey thinks that the point 𝑥
is on the segment [𝑙,𝑟]
if the inequality 𝑙≤𝑥≤𝑟
is satisfied.

Andrey has 𝑞
options for where to start his escape, each option is characterized by a single integer 𝑥𝑖
— the starting position of the escape. He wants to escape from Capygrad as far as possible (to the point with the
maximum possible coordinate). Help Andrey determine how far he could escape from Capygrad, starting at each of the 𝑞
positions.

### thoughts

1. 考虑seg[i] 和 seg[j]的关系
2. 假设i在j的前面, 如果 r[i] < l[j] 它们就是隔离的
3. 讨论 r[i] >= l[j]
4. 如果 b[i] < l[j], 那么在b[i]前面时，无法继续往前
5. 如果 b[i] >= l[j] 那么 就可以叨叨 b[j], 相当于i, j合并成了一个大区间
6. 先通过[a, b]将有重叠的区域进行union
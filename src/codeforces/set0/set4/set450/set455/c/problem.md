Andrew plays a game called "Civilization". Dima helps him.

The game has n cities and m bidirectional roads. The cities are numbered from 1 to n. Between any pair of cities there either is a single (unique) path, or there is no path at all. A path is such a sequence of distinct cities v1, v2, ..., vk, that there is a road between any contiguous cities vi and vi + 1 (1 ≤ i < k). The length of the described path equals to (k - 1). We assume that two cities lie in the same region if and only if, there is a path connecting these two cities.

During the game events of two types take place:

Andrew asks Dima about the length of the longest path in the region where city x lies.
Andrew asks Dima to merge the region where city x lies with the region where city y lies. If the cities lie in the same region, then no merging is needed. Otherwise, you need to merge the regions as follows: choose a city from the first region, a city from the second region and connect them by a road so as to minimize the length of the longest path in the resulting region. If there are multiple ways to do so, you are allowed to choose any of them.
Dima finds it hard to execute Andrew's queries, so he asks you to help him. Help Dima.
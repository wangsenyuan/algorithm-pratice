You are given a graph with 𝑛
nodes and 𝑚
directed edges. One lowercase letter is assigned to each node. We define a path's value as the number of the most
frequently occurring letter. For example, if letters on a path are "abaca", then the value of that path is 3
. Your task is find a path whose value is the largest.

### ideas

1. dp on dags
2. cycle => -1
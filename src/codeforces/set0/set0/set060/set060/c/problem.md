Pasha and Akim were making a forest map — the lawns were the graph's vertexes and the roads joining the lawns were its edges. They decided to encode the number of laughy mushrooms on every lawn in the following way: on every edge between two lawns they wrote two numbers, the greatest common divisor (GCD) and the least common multiple (LCM) of the number of mushrooms on these lawns. But one day Pasha and Akim had an argument about the laughy mushrooms and tore the map. Pasha was left with just some part of it, containing only m roads. Your task is to help Pasha — use the map he has to restore the number of mushrooms on every lawn. As the result is not necessarily unique, help Pasha to restore any one or report that such arrangement of mushrooms does not exist. It is guaranteed that the numbers on the roads on the initial map were no less that 1 and did not exceed 106.

### ideas
1. 节点u的值为a[u]
2. 那么对于每条边来说， a[u] * a[v] = g * l
3. 假设某个节点的值是1, 那么和它相邻的节点的值就可以算出来a[v] = g * l.
4. 如果出现了冲突，那么就没有解
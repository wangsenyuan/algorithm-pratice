Consider a solution which minimises the total debt.

Assume for contradiction that there is a triple of vertices 𝑢≠𝑣≠𝑤
, such that 𝑑(𝑢,𝑣)>0
 and 𝑑(𝑣,𝑤)>0
. We can use the first rule with 𝑎=𝑢
, 𝑏=𝑐=𝑣
, 𝑑=𝑤
 and 𝑧=𝑚𝑖𝑛(𝑑(𝑢,𝑣),𝑑(𝑣,𝑤))
 and then the second rule with 𝑎=𝑣
 and the same 𝑧
. We have just reduced the total debt by 𝑧
, which is a contradiction. So, there cannot be such a triple, in particular there cannot be a vertex 𝑣
 that has both incoming and outgoing debt. Hence, every vertex has only outgoing edges or only incoming ones.

Define 𝑏𝑎𝑙(𝑢)=∑𝑣𝑑(𝑢,𝑣)−∑𝑣𝑑(𝑣,𝑢)
 to be the balance of 𝑢
. Any application of either rule preserves balance of all vertices. It follows that any solution in which every vertex has either outgoing or incoming edges is constructible using finite number of applications of rules. This means that we can just find balances, and greedily match vertices with positive balance to vertices with negative balance.

The total debt is then Σ𝑑=∑𝑣|𝑏𝑎𝑙(𝑣)|2
 and it is clear that we cannot do better than that.
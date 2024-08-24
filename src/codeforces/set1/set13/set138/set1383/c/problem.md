The only difference between this problem and the previous problem is that the underlying graph might have cycles.

Each weakly connected component can be solved independently and the answer is 2⋅𝑛−|𝐿𝐷𝐴𝐺|−1
 where 𝑛
 is the number of nodes in the component and |𝐿𝐷𝐴𝐺|
 is the size of the largest Directed Acyclic Graph.

The largest directed acyclic graph can be computed using dynamic programming in 𝑂(2𝑛⋅𝑛)
.

For every node 𝑢
 store a mask 𝑟𝑒𝑎𝑐ℎ[𝑢]
 with all nodes it can reach directly. Then go through every possible mask from 1
 to 2𝑛−1
 and check whether this set of nodes is acyclic or not. It is acyclic if there exists at least one node 𝑢
 (the last node in one topological order) such that the set without this node is acyclic and this node doesn't reach any other node in this set:


is_dag[0] = 1
for mask in 1..2^n-1
   for u in mask:
      is_dag[mask] |= is_dag[mask ^ (1 « u)] && ((mask & reach[u]) == 0))
Proof by @eatmore:

Lemma 1: Suppose that there is a weakly connected component with 𝑛
 vertices, and there is a solution with 𝑘
 edges. Then, the size of the largest DAG is at least 2⋅𝑛−1−𝑘
.

Proof: Let's keep track of current weakly connected components. Also, in each of the components, let's keep track of some DAG. Initially, each vertex is in a separate component, and each DAG consists of a single vertex. So, there are 𝑛
 components, and the total size of all DAGs is 𝑛
.

Processing an edge (𝑢,𝑣)
:

If 𝑢
 and 𝑣
 are in the same component: if 𝑣
 is in the DAG, remove it. Number of components is unchanged, the total size of all DAGs is decreased by at most 1.
If 𝑢
 and 𝑣
 are in different components, join the components. Concatenate the DAGs (DAG of 𝑢
's component comes before DAG of 𝑣
's component). Number of components decreases by 1
, the total size of all DAGs is unchanged.
At the end, the number of components becomes 1
, so 𝑛−1
 edges are used to decrease the number of components. The remaining 𝑘−𝑛+1
 edges could decrease the size of DAGs, so the final size is at least 𝑛−(𝑘−𝑛+1)=2⋅𝑛−1−𝑘
.

From lemma 1 we know that 𝑘>=2⋅𝑛−1−|𝐷𝐴𝐺|
, then 𝑘
 is minimized when the size of the final DAG is maximized.

Time complexity: 𝑂(|𝐴|+|𝐵|+2𝑛⋅𝑛)
 per test case
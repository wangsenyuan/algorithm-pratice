You are given a node of the tree with index 1 and with weight 0. Let cnt be the number of nodes in the tree at any instant (initially, cnt is set to 1). Support Q queries of following two types:

 Add a new node (index cnt + 1) with weight W and add edge between node R and this node.
 Output the maximum length of sequence of nodes which
starts with R.
Every node in the sequence is an ancestor of its predecessor.
Sum of weight of nodes in sequence does not exceed X.
For some nodes i, j that are consecutive in the sequence if i is an ancestor of j then w[i] ≥ w[j] and there should not exist a node k on simple path from i to j such that w[k] ≥ w[j]
The tree is rooted at node 1 at any instant.

### ideas
1. 
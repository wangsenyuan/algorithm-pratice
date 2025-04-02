One night, having had a hard day at work, Petya saw a nightmare. There was a binary search tree in the dream. But it was not the actual tree that scared Petya. The horrifying thing was that Petya couldn't search for elements in this tree. Petya tried many times to choose key and look for it in the tree, and each time he arrived at a wrong place. Petya has been racking his brains for long, choosing keys many times, but the result was no better. But the moment before Petya would start to despair, he had an epiphany: every time he was looking for keys, the tree didn't have the key, and occured exactly one mistake. "That's not a problem!", thought Petya. "Why not count the expectation value of an element, which is found when I search for the key". The moment he was about to do just that, however, Petya suddenly woke up.

Thus, you are given a binary search tree, that is a tree containing some number written in the node. This number is called the node key. The number of children of every node of the tree is equal either to 0 or to 2. The nodes that have 0 children are called leaves and the nodes that have 2 children, are called inner. An inner node has the left child, that is the child whose key is less than the current node's key, and the right child, whose key is more than the current node's key. Also, a key of any node is strictly larger than all the keys of the left subtree of the node and strictly smaller than all the keys of the right subtree of the node.

Also you are given a set of search keys, all of which are distinct and differ from the node keys contained in the tree. For each key from the set its search in the tree is realised. The search is arranged like this: initially we are located in the tree root, if the key of the current node is larger that our search key, then we move to the left child of the node, otherwise we go to the right child of the node and the process is repeated. As it is guaranteed that the search key is not contained in the tree, the search will always finish in some leaf. The key lying in the leaf is declared the search result.

It is known for sure that during the search we make a mistake in comparing exactly once, that is we go the wrong way, but we won't make any mistakes later. All possible mistakes are equiprobable, that is we should consider all such searches where exactly one mistake occurs. Your task is to find the expectation (the average value) of the search result for every search key, considering that exactly one mistake occurs in the search. That is, for a set of paths containing exactly one mistake in the given key search, you should count the average value of keys containing in the leaves of those paths.


### ideas
1. 一旦移动错了，比如本来应该右移的，它左移了，那么这个result是确定的，就是左子树中的最大值
2. 这个概率 = 1/剩余的步骤数
3. 如果只有一个查询，那这样走下去，是没问题的
4. 但是这里有一批，就有点糟糕了
5. 但是假设到达某个节点，能到达这个节点的正确节点的批次，是可以计算出来的
6. 它们的在走错的情况下，结果是一样的。 所以可以一批进行更新。 而且这个范围是连续的
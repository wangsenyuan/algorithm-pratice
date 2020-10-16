You are given an undirected weighted connected graph with N vertices numbered from 1 to N. Chef knows all the edges of the graph but will only reveal information if you give him some money.

You have 104 coins. You are allowed to ask the chef a certain type of query multiple times. In a single query, you will provide 2 non-intersecting, non-empty set of vertices. Let A be the first set and B be the second set. Chef then will respond by providing the minimum (least) weight edge between any pair of nodes u and v, such that node u is in set A and node v is in set B. In particular, he will return the end points of the least weight edge along with the weight of this edge (in case of a tie, Chef will choose arbitrarily). If there is no such edge, Chef returns -1. A single query costs you |A| coins, where |X| denotes the size of the set X. Do note that the Chef does NOT like too many numbers, so he has allowed the sum of |A| + |B| over all the queries to be at most 2*106

Your aim is to determine the sum of all the weights in a minimum spanning tree of this graph using at most 104 coins.
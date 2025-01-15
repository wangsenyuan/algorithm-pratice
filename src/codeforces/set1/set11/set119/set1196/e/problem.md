You are given two integers 𝑏
 and 𝑤
. You have a chessboard of size 109×109
 with the top left cell at (1;1)
, the cell (1;1)
 is painted white.

Your task is to find a connected component on this chessboard that contains exactly 𝑏
 black cells and exactly 𝑤
 white cells. Two cells are called connected if they share a side (i.e. for the cell (𝑥,𝑦)
 there are at most four connected cells: (𝑥−1,𝑦),(𝑥+1,𝑦),(𝑥,𝑦−1),(𝑥,𝑦+1)
). A set of cells is called a connected component if for every pair of cells 𝐶1
 and 𝐶2
 from this set, there exists a sequence of cells 𝑐1
, 𝑐2
, ..., 𝑐𝑘
 such that 𝑐1=𝐶1
, 𝑐𝑘=𝐶2
, all 𝑐𝑖
 from 1
 to 𝑘
 are belong to this set of cells and for every 𝑖∈[1,𝑘−1]
, cells 𝑐𝑖
 and 𝑐𝑖+1
 are connected.

Obviously, it can be impossible to find such component. In this case print "NO". Otherwise, print "YES" and any suitable connected component.

You have to answer 𝑞
 independent queries.

### ideas
1. x = min(b, w), y = max(b, w)
2. y <= 3 * x + 1 的时候是有答案的
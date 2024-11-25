After getting AC after 13 Time Limit Exceeded verdicts on a geometry problem, Kuroni went to an Italian restaurant to celebrate this holy achievement. Unfortunately, the excess sauce disoriented him, and he's now lost!

The United States of America can be modeled as a tree (why though) with 𝑛
 vertices. The tree is rooted at vertex 𝑟
, wherein lies Kuroni's hotel.

Kuroni has a phone app designed to help him in such emergency cases. To use the app, he has to input two vertices 𝑢
 and 𝑣
, and it'll return a vertex 𝑤
, which is the lowest common ancestor of those two vertices.

However, since the phone's battery has been almost drained out from live-streaming Kuroni's celebration party, he could only use the app at most ⌊𝑛2⌋
 times. After that, the phone would die and there will be nothing left to help our dear friend! :(

As the night is cold and dark, Kuroni needs to get back, so that he can reunite with his comfy bed and pillow(s). Can you help him figure out his hotel's location?

### ideas
1. 考虑以1为root，形成一棵树
2. 考虑p = lca(u, v)， 如果r在p的不饱和u, 或者 v 的其他子树（包括从1...p)这条路径
3. 那么ask(u, v) 返回还是p
4. 如果在包含u的子树中，那么答案q,是这个子树的起点
5. 假设这时找到了一个新的点p,(从u，v找到的)
6. 如果u, v是一条线，那么 r = p
7. 如果不是，那么肯定存在额外的一个分支
8. 然后选择p = u, 然后在新的分支中找到一个叶子节点作为v
9. 如果 p != u, v， 可以相当于舍弃掉了两个分支
10. 考虑一个上面无法处理的一个结构，假设是一个中心点1,其他的点都连到这个中心点。但是r = 非中心点的一个
11. ask(u, v) = u / v =》 r = u (因为u，v始终是叶子节点)
12. 可以的
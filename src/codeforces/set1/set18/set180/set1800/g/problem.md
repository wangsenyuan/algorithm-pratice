Formally, a tree is symmetrical if there exists an order of children such that:

The subtree of the leftmost child of the root is a mirror image of the subtree of the rightmost child;
the subtree of the second-left child of the root is a mirror image of the subtree of the second-right child of the root;
...
if the number of children of the root is odd, then the subtree of the middle child should be symmetrical.

Note that if one subtree is a mirror image of another, then they are isomorphic (that is, equal without taking into
account the vertex numbers). To check the subtrees for isomorphism, we use hashing of root trees.

Now we just have to learn how to check trees for symmetry. To do this, let's calculate how many children of each type
our vertex has (let's denote the hash of its subtree by the vertex type). In order for the vertex subtree to be
symmetric, each child must have a pair of the same type, except perhaps one, which must also be symmetric. We can
calculate the symmetry of the subtrees while counting their hash to simplify this task.
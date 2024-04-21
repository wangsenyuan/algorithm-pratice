Maksim has 𝑛
objects and 𝑚
boxes, each box has size exactly 𝑘
. Objects are numbered from 1
to 𝑛
in order from left to right, the size of the 𝑖
-th object is 𝑎𝑖
.

Maksim wants to pack his objects into the boxes and he will pack objects by the following algorithm: he takes one of the
empty boxes he has, goes from left to right through the objects, and if the 𝑖
-th object fits in the current box (the remaining size of the box is greater than or equal to 𝑎𝑖
), he puts it in the box, and the remaining size of the box decreases by 𝑎𝑖
. Otherwise he takes the new empty box and continues the process above. If he has no empty boxes and there is at least
one object not in some box then Maksim cannot pack the chosen set of objects.

Maksim wants to know the maximum number of objects he can pack by the algorithm above. To reach this target, he will
throw out the leftmost object from the set until the remaining set of objects can be packed in boxes he has. Your task
is to say the maximum number of objects Maksim can pack in boxes he has.

Each time when Maksim tries to pack the objects into the boxes, he will make empty all the boxes he has before do it (
and the relative order of the remaining set of objects will not change).

Input
The first line of the input contains three integers 𝑛
, 𝑚
, 𝑘
(1≤𝑛,𝑚≤2⋅105
, 1≤𝑘≤109
) — the number of objects, the number of boxes and the size of each box.

The second line of the input contains 𝑛
integers 𝑎1,𝑎2,…,𝑎𝑛
(1≤𝑎𝑖≤𝑘
), where 𝑎𝑖
is the size of the 𝑖
-th object.

### ideas

1. binary search check the last x objects
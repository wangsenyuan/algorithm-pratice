Vlad found an array 𝑎
of 𝑛
integers and decided to sort it in non-decreasing order.

To do this, Vlad can apply the following operation any number of times:

Extract the first element of the array and insert it at the end;
Swap that element with the previous one until it becomes the first or until it becomes strictly greater than the
previous one.
Note that both actions are part of the operation, and for one operation, you must apply both actions.

For example, if you apply the operation to the array [4,3,1,2,6,4
], it will change as follows:

[4,3,1,2,6,4
];
[3,1,2,6,4,4
];
[3,1,2,6,4,4
];
[3,1,2,4,6,4
].
Vlad doesn't have time to perform all the operations, so he asks you to determine the minimum number of operations
required to sort the array or report that it is impossible.
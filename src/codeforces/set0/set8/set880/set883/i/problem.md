Evlampiy has found one more cool application to process photos. However the application has certain limitations.

Each photo i has a contrast vi. In order for the processing to be truly of high quality, the application must receive at
least k photos with contrasts which differ as little as possible.

Evlampiy already knows the contrast vi for each of his n photos. Now he wants to split the photos into groups, so that
each group contains at least k photos. As a result, each photo must belong to exactly one group.

He considers a processing time of the j-th group to be the difference between the maximum and minimum values of vi in
the group. Because of multithreading the processing time of a division into groups is the maximum processing time among
all groups.

Split n photos into groups in a such way that the processing time of the division is the minimum possible, i.e. that the
the maximum processing time over all groups as least as possible.

### ideas

1. sort first
2. then binary search on the result
### description

    As is known, Omsk is the capital of Berland. Like any capital, Omsk has a well-developed metro system. The Omsk metro consists of a certain number of stations connected by tunnels, and between any two stations there is exactly one path that passes through each of the tunnels no more than once. In other words, the metro is a tree.

    To develop the metro and attract residents, the following system is used in Omsk. Each station has its own weight
    𝑥∈{−1,1}
    . If the station has a weight of −1
    , then when the station is visited by an Omsk resident, a fee of 1
    burle is charged. If the weight of the station is 1
    , then the Omsk resident is rewarded with 1
    burle.
    
    Omsk Metro currently has only one station with number 1
    and weight 𝑥=1
    . Every day, one of the following events occurs:
    
    A new station with weight 𝑥
    is added to the station with number 𝑣𝑖
    , and it is assigned a number that is one greater than the number of existing stations.
    Alex, who lives in Omsk, wonders: is there a subsegment†
    (possibly empty) of the path between vertices 𝑢
    and 𝑣
    such that, by traveling along it, exactly 𝑘
    burles can be earned (if 𝑘<0
    , this means that 𝑘
    burles will have to be spent on travel). In other words, Alex is interested in whether there is such a subsegment of the
    path that the sum of the weights of the vertices in it is equal to 𝑘
    . Note that the subsegment can be empty, and then the sum is equal to 0
    .
    You are a friend of Alex, so your task is to answer Alex's questions.

### thoughts

1. 考虑一条直线的情况，因为x只能是+1/-1， 所以如果到v为止，所有节点的权重sum = w
2. 那么 0....w的情况都可以得到
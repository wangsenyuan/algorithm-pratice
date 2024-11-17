You are an environmental activist at heart but the reality is harsh and you are just a cashier in a cinema. But you can
still do something!

You have 𝑛
tickets to sell. The price of the 𝑖
-th ticket is 𝑝𝑖
. As a teller, you have a possibility to select the order in which the tickets will be sold (i.e. a permutation of the
tickets). You know that the cinema participates in two ecological restoration programs applying them to the order you
chose:

The 𝑥%
of the price of each the 𝑎
-th sold ticket (𝑎
-th, 2𝑎
-th, 3𝑎
-th and so on) in the order you chose is aimed for research and spreading of renewable energy sources.
The 𝑦%
of the price of each the 𝑏
-th sold ticket (𝑏
-th, 2𝑏
-th, 3𝑏
-th and so on) in the order you chose is aimed for pollution abatement.
If the ticket is in both programs then the (𝑥+𝑦)%
are used for environmental activities. Also, it's known that all prices are multiples of 100
, so there is no need in any rounding.

For example, if you'd like to sell tickets with prices [400,100,300,200]
and the cinema pays 10%
of each 2
-nd sold ticket and 20%
of each 3
-rd sold ticket, then arranging them in order [100,200,300,400]
will lead to contribution equal to 100⋅0+200⋅0.1+300⋅0.2+400⋅0.1=120
. But arranging them in order [100,300,400,200]
will lead to 100⋅0+300⋅0.1+400⋅0.2+200⋅0.1=130
.

Nature can't wait, so you decided to change the order of tickets in such a way, so that the total contribution to
programs will reach at least 𝑘
in minimum number of sold tickets. Or say that it's impossible to do so. In other words, find the minimum number of
tickets which are needed to be sold in order to earn at least 𝑘
.

### ideas

1. 
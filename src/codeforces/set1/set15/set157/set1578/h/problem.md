Helen studies functional programming and she is fascinated with a concept of higher order functions — functions that are
taking other functions as parameters. She decides to generalize the concept of the function order and to test it on some
examples.

For her study, she defines a simple grammar of types. In her grammar, a type non-terminal 𝑇
is defined as one of the following grammar productions, together with order(𝑇)
, defining an order of the corresponding type:

"()" is a unit type, order("()")=0
.
"(" 𝑇
")" is a parenthesized type, order("("𝑇")")=order(𝑇)
.
𝑇1
"->" 𝑇2
is a functional type, order(𝑇1"->"𝑇2)=𝑚𝑎𝑥(order(𝑇1)+1,order(𝑇2))
. The function constructor 𝑇1
"->" 𝑇2
is right-to-left associative, so the type "()->()->()" is the same as the type "()->(()->())" of a function returning a
function, and it has an order of 1
. While "(()->())->()" is a function that has an order-1 type "(()->())" as a parameter, and it has an order of 2
.
Helen asks for your help in writing a program that computes an order of the given type.
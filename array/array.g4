grammar array;

vec4: '[' ints[4] ']';
ints[int max]
locals [int i=1]
    : INT(',' {$i++;} {$i<=$max}?<fail={"exceeded max: "+ $max}> INT)*
    ;
INT : [0-9]+ ;

WS : [ \t\r\n]+ -> skip ;

/*
$  antlr4 array.g4
$  javac array*.java
$  grun array vec4 -tree
[1,2,3,4,5]
line 1:9 rule ints exceeded max: 4
(vec4 [ (ints 1 , 2 , 3 , 4 , 5) ])
*/
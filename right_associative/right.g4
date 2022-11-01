grammar right;

stat : expr;

expr : expr AddSub expr
     | <assoc=right> expr '=' expr
     | INT
     | ID
     ;

AddSub : '+' | '-' ;

INT  : [0-9]+ ;
ID : [a-zA-Z]+;
WS : [ \t\r\n]+ -> skip ;

grammar listen_err;

stat : expr;

expr : expr '-' expr
     | INT
     ;

SUB : '-' ;

INT  : [0-9]+ | '-' [0-9]+ ;
WS : [ \t\r\n]+ -> skip ;

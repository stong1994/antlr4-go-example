grammar listen_err;

stat : expr;

expr
     : expr ADD expr # add
     | INT # num
     ;

ADD : '+' ;

INT  : [0-9]+ ;
WS : [ \t\r\n]+ -> skip ;

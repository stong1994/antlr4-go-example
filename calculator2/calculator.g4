grammar calculator;

stat : expr;

expr returns [int value]
     : expr op=('*'|'/') expr # MulDiv
     | expr op=('+'|'-') expr # AddSub
     | INT # num
     | '(' expr ')' # parens
     ;

MUL : '*' ;
DIV : '/' ;
ADD : '+' ;
SUB : '-' ;

INT  : [0-9]+ | '-' [0-9]+ ;
WS : [ \t\r\n]+ -> skip ;

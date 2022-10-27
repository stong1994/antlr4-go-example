grammar calculator;

stat : expr;

expr : '-' expr # NegNum
     | expr op=('*'|'/') expr # MulDiv
     | expr op=('+'|'-') expr # AddSub
     | '(' expr ')' # parens
     | INT # num
     ;

MUL : '*' ;
DIV : '/' ;
ADD : '+' ;
SUB : '-' ;

INT  : [0-9]+ ;
WS : [ \t\r\n]+ -> skip ;

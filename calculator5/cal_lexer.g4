lexer grammar cal_lexer;

// 默认模式下的词法规则
OPEN : '<' -> mode(MARK) ; // 进入MARK模式
MUL : '*' ;
DIV : '/' ;
ADD : '+' ;
SUB : '-' ;

INT  : [0-9]+ ;
WS : [ \t\r\n]+ -> skip ;

// MARK模式下的词法规则
mode MARK;
CLOSE : '>' -> mode(DEFAULT_MODE) ; // 回到SEA模式
CONTENT : ~[>]+ ; // 匹配所有字符

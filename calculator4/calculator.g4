grammar calculator;

@parser::members {
func handleExpr(op, left, right int) int {
    switch op {
    case calculatorParserADD:
        return left+right
    case calculatorParserSUB:
        return left-right
    case calculatorParserMUL:
        return left*right
    case calculatorParserDIV:
        return left/right
    default:
        return 0
    }
}
}

stat : expr;

expr returns [int value]
     : a=expr op=('*'|'/') b=expr
     {
     $ctx.value = handleExpr($op.type, $a.value, $b.value)
     fmt.Printf("%d %s %d = %d\n",$a.value, $op.text, $b.value, $ctx.value)
     }
     # MulDiv
     | a=expr op=('+'|'-') b=expr
     {
     $ctx.value = handleExpr($op.type, $a.value, $b.value)
     fmt.Printf("got %s\n", $op.text)
     fmt.Printf("calculating:\t%d %s %d = %d\n",$a.value, $op.text, $b.value, $ctx.value)
     }
     # AddSub
     | '(' expr ')'
     {
     $ctx.value=$expr.value;
     }
     # parens
     | INT
     {
     $ctx.value = $INT.int;
     fmt.Println("got", $ctx.value)
     }
     # num
     ;

MUL : '*' ;
DIV : '/' ;
ADD : '+' ;
SUB : '-' ;

INT  : [0-9]+ ;
WS : [ \t\r\n]+ -> skip ;
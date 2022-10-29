parser grammar cal_parser;

options { tokenVocab=cal_lexer; }

@parser::members {
func handleExpr(op, left, right int) int {
    switch op {
    case cal_lexerADD:
        return left+right
    case cal_lexerSUB:
        return left-right
    case cal_lexerMUL:
        return left*right
    case cal_lexerDIV:
        return left/right
    default:
        return 0
    }
}
}

stat : (expr|mark)+;

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
     | INT
     {
     $ctx.value = $INT.int;
     fmt.Println("got", $ctx.value)
     }
     # num
     ;

mark : '<' CONTENT '>' {fmt.Println("comment: ", $CONTENT.text)};
lexer grammar SwiftLexer;

// --------------- Tokens
// types
INT: 'int';
FLOAT: 'float';
BOOL: 'bool';
CHARACTER: 'character';
PSTRING: 'String';
NIL: 'nil';

// reserved words
TRU: 'true';
FAL: 'false';
PRINT: 'print';
IF: 'if';
ELSE: 'else';
WHILE: 'while';
FOR: 'for';
IN: 'in';
SWITCH: 'switch';
CASE: 'case';
DEFAULT: 'default';
VAR: 'var';
LET: 'let';
BREAK: 'break';
RETURN: 'return';
FUNC: 'func';
COUNT: 'count';
ISEMPTY: 'IsEmpty';
APPEND: 'append';
REMOVE_LAST: 'removeLast';
REMOVE: 'remove';
AT: 'at';

// primitives
NUMBER : [0-9]+('.'[0-9]+)?;
STRING: '"'~["]*'"';
ID: ([a-zA-Z_])[a-zA-Z0-9_]*;

// symbols

DIF:      '!=';
IG_IG:          '==';
NOT:            '!';
OR:             '||';
AND:            '&&';
IG:          '=';
MAY_IG:     '>=';
MEN_IG:     '<=';
MAYOR:          '>';
MENOR:          '<';
MUL:            '*';
DIV:            '/';
ADD:            '+';
SUB:            '-';
MOD:            '%';
PARIZQ:         '(';
PARDER:         ')';
LLAVEIZQ:       '{';
LLAVEDER:       '}';
CORCHIZQ:       '[';
CORCHDER:       ']';
DOSPUNTOS:      ':';
COMA:           ',';
PTCOMA:      ';';
INTERROGACION:  '?';
PUNTO:          '.';

// skip
WHITESPACE: [ \\\r\n\t]+ -> skip;
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;
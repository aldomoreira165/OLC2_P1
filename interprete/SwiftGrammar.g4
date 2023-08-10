grammar SwiftGrammar; 
import SwiftLexer; 

s: block EOF;

block: (stmt)*
     ;

stmt: printstmt
    | typedDeclstmt
    | untypedDeclstmt
    | optionalTypedDeclstmt
    | asignstmt
    | ifstmt
    ;

printstmt
    : PRINT PARIZQ expr PARDER
    ;

ifstmt  
    : IF expr LLAVEIZQ block LLAVEDER
    ;

typedDeclstmt
    : VAR ID DOSPUNTOS tipo IG expr 
    ;

untypedDeclstmt
    : VAR ID IG expr                         
    ;

optionalTypedDeclstmt
    : VAR ID DOSPUNTOS tipo INTERROGACION                          
    ;

asignstmt
    : ID IG expr
    ;

expr
    : left=expr op=(MUL|DIV) right=expr         # OpExpr
    | left=expr op=(ADD|SUB) right=expr         # OpExpr
    | left=expr op=(MAY_IG|MAYOR) right=expr    # OpExpr
    | left=expr op=(MEN_IG|MENOR) right=expr    # OpExpr
    | left=expr op=(IG_IG|DIF) right=expr       # OpExpr
    | left=expr AND right=expr                  # OpExpr
    | left=expr OR right=expr                   # OpExpr
    | PARIZQ expr PARDER                        # ParExpr
    | primitivo                                 # PrimExpr
    ;

primitivo    
    : NUMBER                                    # NumExpr
    | ID                                        # IdExpr
    | STRING                                    # StrExpr    
    | (TRU | FAL)                               # BoolExpr
    ;

tipo
    : INT                                      
    | FLOAT                                     
    | BOOL                                      
    | CHARACTER                                 
    | PSTRING                                  
    ;
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
    | switchstmt
    | whilestmt
    | opasignstmt
    | breakstmt
    | funcdclstmt
    | accfuncstm
    | returnstmt
    ;

//de flujo

returnstmt
    : RETURN expr PTCOMA
    ;

//funciones

funcdclstmt
    : FUNC ID PARIZQ parametros? PARDER LLAVEIZQ block LLAVEDER #FuncionNormal
    | FUNC ID PARIZQ parametros? PARDER SUB MAYOR tipo LLAVEIZQ block LLAVEDER #FuncionRetorno
    ;

accfuncstm
    : ID PARIZQ parametroscall? PARDER
    ;

parametros
    : ID ID DOSPUNTOS tipo (COMA ID ID DOSPUNTOS tipo)*
    ;

parametroscall
    : ID DOSPUNTOS expr (COMA ID DOSPUNTOS expr)*
    ;

//=============================

printstmt
    : PRINT PARIZQ expr PARDER
    ;

//sententencias de transferencia

breakstmt
    : BREAK
    ;

//sentencias de control de flujo

ifstmt  
    : IF expr LLAVEIZQ block LLAVEDER (elseifstmt)* (ELSE LLAVEIZQ block LLAVEDER)?
    ;

elseifstmt: ELSE IF expr LLAVEIZQ block LLAVEDER;

switchstmt: SWITCH expr LLAVEIZQ  caseStmt+ defaultCase? LLAVEDER;

caseStmt: CASE expr DOSPUNTOS block;

defaultCase: DEFAULT DOSPUNTOS block;

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

whilestmt
    : WHILE expr LLAVEIZQ block LLAVEDER
    ;

opasignstmt
    : ID ADD IG expr                           #Incremento
    | ID SUB IG expr                           #Decremento  
    ;

expr
    : PARIZQ expr PARDER                        # ParExpr
    | SUB expr                                  # UnariaExpr
    | left=expr op=(MUL|DIV|MOD) right=expr     # OpExpr
    | left=expr op=(ADD|SUB) right=expr         # OpExpr
    | left=expr op=(MAY_IG|MAYOR) right=expr    # OpExpr
    | left=expr op=(MEN_IG|MENOR) right=expr    # OpExpr
    | left=expr op=(IG_IG|DIF) right=expr       # OpExpr
    | left=expr op=AND right=expr               # OpExpr
    | left=expr op=OR right=expr                # OpExpr
    | NUMBER                                    # NumExpr
    | ID                                        # IdExpr
    | STRING                                    # StrExpr    
    | (TRU | FAL)                               # BoolExpr
    | NIL                                       # NilExpr
    | accfuncstm                                # AccFuncExpr
    ;

tipo
    : INT                                      
    | FLOAT                                     
    | BOOL                                      
    | CHARACTER                                 
    | PSTRING                                  
    ;
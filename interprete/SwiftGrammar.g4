grammar SwiftGrammar; 
import SwiftLexer; 

s: block EOF;

block: (stmt PTCOMA?)*
     ;

stmt: printstmt
    | typedDeclstmt
    | untypedDeclstmt
    | optionalTypedDeclstmt
    | asignstmt
    | ifstmt
    | switchstmt
    | whilestmt
    | forstmt
    | guardstmt
    | opasignstmt
    | funcdclstmt
    | accfuncstm
    | breakstmt
    | continuestmt
    | returnstmt
    | declvectorstmt
    | accesovectorstmt
    | appendvectorstmt
    | removelastvectorstmt
    | removeatvectorstmt
    | asignvectorstmt
    | declmatrizstmt
    | asignmatrizstmt
    | defstructstmt
    | struct_expr
    ;

// struct
defstructstmt
    : STRUCT ID LLAVEIZQ lista_atributos* LLAVEDER #DefStruct
    ;

lista_atributos
    : (LET|VAR) ID (DOSPUNTOS (tipo|ID))? (IG expr)? #AtributoStruct
    ;

struct_expr
    : (STRUCT_VAR|STRUCT_LET) ID (DOSPUNTOS ID)? IG valor_struct_expr #StructExpr
    ;

valor_struct_expr
    : ID ( PARIZQ l_dupla PARDER)? #ValorStructExpr
    ;

l_dupla
    : ID DOSPUNTOS expr (COMA ID DOSPUNTOS expr)* #Duplastruct
    ;

accesostructstmt
    : ID (PUNTO ID)+ #AccesoStruct
    ;

//vectores

declvectorstmt
    : VAR ID DOSPUNTOS CORCHIZQ tipo CORCHDER defvectorstmt
    ;

defvectorstmt
    : IG CORCHIZQ listaexpresiones? CORCHDER #DefVector
    | IG ID                                 #DefVectorID
    ;

listaexpresiones
    : expr (COMA expr)*
    ;

accesovectorstmt
    : ID CORCHIZQ expr CORCHDER
    ;

asignvectorstmt
    : ID CORCHIZQ expr CORCHDER IG expr
    ;

appendvectorstmt
    : ID PUNTO APPEND PARIZQ expr PARDER
    ;

removeatvectorstmt
    : ID PUNTO REMOVE PARIZQ AT DOSPUNTOS expr PARDER
    ;

removelastvectorstmt
    : ID PUNTO REMOVE_LAST PARIZQ PARDER
    ;

countvectorstmt
    : ID PUNTO COUNT
    ;

isemptyvectorstmt
    : ID PUNTO ISEMPTY
    ;

//matrices

declmatrizstmt
    : VAR ID (DOSPUNTOS tipomatriz)? IG listavaloresmatriz #declmatrizstmt2
    | VAR ID (DOSPUNTOS tipomatriz)? IG listavaloresmatriz3 #declmatrizstmt3
    ;

tipomatriz
    : CORCHIZQ CORCHIZQ tipo CORCHDER CORCHDER #tipomatriz2
    | CORCHIZQ CORCHIZQ CORCHIZQ tipo CORCHDER CORCHDER CORCHDER #tipomatriz3
    ;

listavaloresmatriz
    : CORCHIZQ CORCHIZQ listaexpresiones CORCHDER (COMA CORCHIZQ listaexpresiones CORCHDER)+ CORCHDER #listavaloresmatriz2
    ;

listavaloresmatriz3
    : CORCHIZQ listavaloresmatriz (COMA listavaloresmatriz)* CORCHDER
    ;

accesomatriz
    : ID CORCHIZQ expr CORCHDER CORCHIZQ expr CORCHDER #accesomatriz2
    | ID CORCHIZQ expr CORCHDER CORCHIZQ expr CORCHDER CORCHIZQ expr CORCHDER #accesomatriz3
    ;

asignmatrizstmt
    : ID CORCHIZQ expr CORCHDER CORCHIZQ expr CORCHDER IG expr #asignmatrizstmt2
    | ID CORCHIZQ expr CORCHDER CORCHIZQ expr CORCHDER CORCHIZQ expr CORCHDER IG expr #asignmatrizstmt3
    ;

//de flujo

returnstmt
    : RETURN expr PTCOMA
    ;

breakstmt
    : BREAK
    ;

continuestmt
    : CONTINUE
    ;

//funciones embebidas

printstmt
    : PRINT PARIZQ expr PARDER
    ;

intstmt
    : INT PARIZQ expr PARDER
    ;

floatstmt
    : FLOAT PARIZQ expr PARDER
    ;

stringstmt
    : PSTRING PARIZQ expr PARDER
    ;

//==============

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

//sentencias de control de flujo

ifstmt  
    : IF expr LLAVEIZQ block LLAVEDER (elseifstmt)* (ELSE LLAVEIZQ block LLAVEDER)?
    ;

elseifstmt: ELSE IF expr LLAVEIZQ block LLAVEDER;

switchstmt: SWITCH expr LLAVEIZQ  caseStmt+ defaultCase? LLAVEDER;

caseStmt: CASE expr DOSPUNTOS block;

defaultCase: DEFAULT DOSPUNTOS block;

typedDeclstmt
    : (VAR|LET) ID DOSPUNTOS tipo IG expr 
    ;

untypedDeclstmt
    : (VAR|LET) ID IG expr                         
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

forstmt
    : FOR ID IN (expr|rangostmt) LLAVEIZQ block LLAVEDER
    ;

guardstmt
    : GUARD expr ELSE LLAVEIZQ block LLAVEDER
    ;

rangostmt
    : expr PUNTO PUNTO PUNTO expr
    ;

opasignstmt
    : ID ADD IG expr                           #Incremento
    | ID SUB IG expr                           #Decremento  
    ;

expr
    : PARIZQ expr PARDER                        # ParExpr
    | SUB expr                                  # UnariaExpr
    | NOT expr                                  # NotExpr
    | left=expr op=(DIV|MOD|MUL) right=expr     # OpExpr
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
    | intstmt                                   # IntExpr
    | floatstmt                                 # FloatExpr
    | stringstmt                                # StringExpr
    | accesovectorstmt                          # AccesoVectorExpr
    | countvectorstmt                           # CountVectorExpr
    | isemptyvectorstmt                         # IsEmptyVectorExpr
    | accesomatriz                              # AccesoMatrizExpr
    | accesostructstmt                          # AccesoStructExpr
    ;

tipo
    : INT                                      
    | FLOAT                                     
    | BOOL                                      
    | CHARACTER                                 
    | PSTRING                                  
    ;
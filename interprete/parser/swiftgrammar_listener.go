// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// SwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type SwiftGrammarListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterDeclvectorstmt is called when entering the declvectorstmt production.
	EnterDeclvectorstmt(c *DeclvectorstmtContext)

	// EnterDefVector is called when entering the DefVector production.
	EnterDefVector(c *DefVectorContext)

	// EnterDefVectorID is called when entering the DefVectorID production.
	EnterDefVectorID(c *DefVectorIDContext)

	// EnterListaexpresiones is called when entering the listaexpresiones production.
	EnterListaexpresiones(c *ListaexpresionesContext)

	// EnterAccesovectorstmt is called when entering the accesovectorstmt production.
	EnterAccesovectorstmt(c *AccesovectorstmtContext)

	// EnterAsignvectorstmt is called when entering the asignvectorstmt production.
	EnterAsignvectorstmt(c *AsignvectorstmtContext)

	// EnterAppendvectorstmt is called when entering the appendvectorstmt production.
	EnterAppendvectorstmt(c *AppendvectorstmtContext)

	// EnterRemoveatvectorstmt is called when entering the removeatvectorstmt production.
	EnterRemoveatvectorstmt(c *RemoveatvectorstmtContext)

	// EnterRemovelastvectorstmt is called when entering the removelastvectorstmt production.
	EnterRemovelastvectorstmt(c *RemovelastvectorstmtContext)

	// EnterCountvectorstmt is called when entering the countvectorstmt production.
	EnterCountvectorstmt(c *CountvectorstmtContext)

	// EnterIsemptyvectorstmt is called when entering the isemptyvectorstmt production.
	EnterIsemptyvectorstmt(c *IsemptyvectorstmtContext)

	// EnterReturnstmt is called when entering the returnstmt production.
	EnterReturnstmt(c *ReturnstmtContext)

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

	// EnterIntstmt is called when entering the intstmt production.
	EnterIntstmt(c *IntstmtContext)

	// EnterFloatstmt is called when entering the floatstmt production.
	EnterFloatstmt(c *FloatstmtContext)

	// EnterStringstmt is called when entering the stringstmt production.
	EnterStringstmt(c *StringstmtContext)

	// EnterFuncionNormal is called when entering the FuncionNormal production.
	EnterFuncionNormal(c *FuncionNormalContext)

	// EnterFuncionRetorno is called when entering the FuncionRetorno production.
	EnterFuncionRetorno(c *FuncionRetornoContext)

	// EnterAccfuncstm is called when entering the accfuncstm production.
	EnterAccfuncstm(c *AccfuncstmContext)

	// EnterParametros is called when entering the parametros production.
	EnterParametros(c *ParametrosContext)

	// EnterParametroscall is called when entering the parametroscall production.
	EnterParametroscall(c *ParametroscallContext)

	// EnterBreakstmt is called when entering the breakstmt production.
	EnterBreakstmt(c *BreakstmtContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterElseifstmt is called when entering the elseifstmt production.
	EnterElseifstmt(c *ElseifstmtContext)

	// EnterSwitchstmt is called when entering the switchstmt production.
	EnterSwitchstmt(c *SwitchstmtContext)

	// EnterCaseStmt is called when entering the caseStmt production.
	EnterCaseStmt(c *CaseStmtContext)

	// EnterDefaultCase is called when entering the defaultCase production.
	EnterDefaultCase(c *DefaultCaseContext)

	// EnterTypedDeclstmt is called when entering the typedDeclstmt production.
	EnterTypedDeclstmt(c *TypedDeclstmtContext)

	// EnterUntypedDeclstmt is called when entering the untypedDeclstmt production.
	EnterUntypedDeclstmt(c *UntypedDeclstmtContext)

	// EnterOptionalTypedDeclstmt is called when entering the optionalTypedDeclstmt production.
	EnterOptionalTypedDeclstmt(c *OptionalTypedDeclstmtContext)

	// EnterAsignstmt is called when entering the asignstmt production.
	EnterAsignstmt(c *AsignstmtContext)

	// EnterWhilestmt is called when entering the whilestmt production.
	EnterWhilestmt(c *WhilestmtContext)

	// EnterIncremento is called when entering the Incremento production.
	EnterIncremento(c *IncrementoContext)

	// EnterDecremento is called when entering the Decremento production.
	EnterDecremento(c *DecrementoContext)

	// EnterBoolExpr is called when entering the BoolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterStringExpr is called when entering the StringExpr production.
	EnterStringExpr(c *StringExprContext)

	// EnterNilExpr is called when entering the NilExpr production.
	EnterNilExpr(c *NilExprContext)

	// EnterFloatExpr is called when entering the FloatExpr production.
	EnterFloatExpr(c *FloatExprContext)

	// EnterIsEmptyVectorExpr is called when entering the IsEmptyVectorExpr production.
	EnterIsEmptyVectorExpr(c *IsEmptyVectorExprContext)

	// EnterIdExpr is called when entering the IdExpr production.
	EnterIdExpr(c *IdExprContext)

	// EnterOpExpr is called when entering the OpExpr production.
	EnterOpExpr(c *OpExprContext)

	// EnterAccesoVectorExpr is called when entering the AccesoVectorExpr production.
	EnterAccesoVectorExpr(c *AccesoVectorExprContext)

	// EnterUnariaExpr is called when entering the UnariaExpr production.
	EnterUnariaExpr(c *UnariaExprContext)

	// EnterCountVectorExpr is called when entering the CountVectorExpr production.
	EnterCountVectorExpr(c *CountVectorExprContext)

	// EnterNumExpr is called when entering the NumExpr production.
	EnterNumExpr(c *NumExprContext)

	// EnterParExpr is called when entering the ParExpr production.
	EnterParExpr(c *ParExprContext)

	// EnterStrExpr is called when entering the StrExpr production.
	EnterStrExpr(c *StrExprContext)

	// EnterAccFuncExpr is called when entering the AccFuncExpr production.
	EnterAccFuncExpr(c *AccFuncExprContext)

	// EnterIntExpr is called when entering the IntExpr production.
	EnterIntExpr(c *IntExprContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitDeclvectorstmt is called when exiting the declvectorstmt production.
	ExitDeclvectorstmt(c *DeclvectorstmtContext)

	// ExitDefVector is called when exiting the DefVector production.
	ExitDefVector(c *DefVectorContext)

	// ExitDefVectorID is called when exiting the DefVectorID production.
	ExitDefVectorID(c *DefVectorIDContext)

	// ExitListaexpresiones is called when exiting the listaexpresiones production.
	ExitListaexpresiones(c *ListaexpresionesContext)

	// ExitAccesovectorstmt is called when exiting the accesovectorstmt production.
	ExitAccesovectorstmt(c *AccesovectorstmtContext)

	// ExitAsignvectorstmt is called when exiting the asignvectorstmt production.
	ExitAsignvectorstmt(c *AsignvectorstmtContext)

	// ExitAppendvectorstmt is called when exiting the appendvectorstmt production.
	ExitAppendvectorstmt(c *AppendvectorstmtContext)

	// ExitRemoveatvectorstmt is called when exiting the removeatvectorstmt production.
	ExitRemoveatvectorstmt(c *RemoveatvectorstmtContext)

	// ExitRemovelastvectorstmt is called when exiting the removelastvectorstmt production.
	ExitRemovelastvectorstmt(c *RemovelastvectorstmtContext)

	// ExitCountvectorstmt is called when exiting the countvectorstmt production.
	ExitCountvectorstmt(c *CountvectorstmtContext)

	// ExitIsemptyvectorstmt is called when exiting the isemptyvectorstmt production.
	ExitIsemptyvectorstmt(c *IsemptyvectorstmtContext)

	// ExitReturnstmt is called when exiting the returnstmt production.
	ExitReturnstmt(c *ReturnstmtContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

	// ExitIntstmt is called when exiting the intstmt production.
	ExitIntstmt(c *IntstmtContext)

	// ExitFloatstmt is called when exiting the floatstmt production.
	ExitFloatstmt(c *FloatstmtContext)

	// ExitStringstmt is called when exiting the stringstmt production.
	ExitStringstmt(c *StringstmtContext)

	// ExitFuncionNormal is called when exiting the FuncionNormal production.
	ExitFuncionNormal(c *FuncionNormalContext)

	// ExitFuncionRetorno is called when exiting the FuncionRetorno production.
	ExitFuncionRetorno(c *FuncionRetornoContext)

	// ExitAccfuncstm is called when exiting the accfuncstm production.
	ExitAccfuncstm(c *AccfuncstmContext)

	// ExitParametros is called when exiting the parametros production.
	ExitParametros(c *ParametrosContext)

	// ExitParametroscall is called when exiting the parametroscall production.
	ExitParametroscall(c *ParametroscallContext)

	// ExitBreakstmt is called when exiting the breakstmt production.
	ExitBreakstmt(c *BreakstmtContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitElseifstmt is called when exiting the elseifstmt production.
	ExitElseifstmt(c *ElseifstmtContext)

	// ExitSwitchstmt is called when exiting the switchstmt production.
	ExitSwitchstmt(c *SwitchstmtContext)

	// ExitCaseStmt is called when exiting the caseStmt production.
	ExitCaseStmt(c *CaseStmtContext)

	// ExitDefaultCase is called when exiting the defaultCase production.
	ExitDefaultCase(c *DefaultCaseContext)

	// ExitTypedDeclstmt is called when exiting the typedDeclstmt production.
	ExitTypedDeclstmt(c *TypedDeclstmtContext)

	// ExitUntypedDeclstmt is called when exiting the untypedDeclstmt production.
	ExitUntypedDeclstmt(c *UntypedDeclstmtContext)

	// ExitOptionalTypedDeclstmt is called when exiting the optionalTypedDeclstmt production.
	ExitOptionalTypedDeclstmt(c *OptionalTypedDeclstmtContext)

	// ExitAsignstmt is called when exiting the asignstmt production.
	ExitAsignstmt(c *AsignstmtContext)

	// ExitWhilestmt is called when exiting the whilestmt production.
	ExitWhilestmt(c *WhilestmtContext)

	// ExitIncremento is called when exiting the Incremento production.
	ExitIncremento(c *IncrementoContext)

	// ExitDecremento is called when exiting the Decremento production.
	ExitDecremento(c *DecrementoContext)

	// ExitBoolExpr is called when exiting the BoolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitStringExpr is called when exiting the StringExpr production.
	ExitStringExpr(c *StringExprContext)

	// ExitNilExpr is called when exiting the NilExpr production.
	ExitNilExpr(c *NilExprContext)

	// ExitFloatExpr is called when exiting the FloatExpr production.
	ExitFloatExpr(c *FloatExprContext)

	// ExitIsEmptyVectorExpr is called when exiting the IsEmptyVectorExpr production.
	ExitIsEmptyVectorExpr(c *IsEmptyVectorExprContext)

	// ExitIdExpr is called when exiting the IdExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitOpExpr is called when exiting the OpExpr production.
	ExitOpExpr(c *OpExprContext)

	// ExitAccesoVectorExpr is called when exiting the AccesoVectorExpr production.
	ExitAccesoVectorExpr(c *AccesoVectorExprContext)

	// ExitUnariaExpr is called when exiting the UnariaExpr production.
	ExitUnariaExpr(c *UnariaExprContext)

	// ExitCountVectorExpr is called when exiting the CountVectorExpr production.
	ExitCountVectorExpr(c *CountVectorExprContext)

	// ExitNumExpr is called when exiting the NumExpr production.
	ExitNumExpr(c *NumExprContext)

	// ExitParExpr is called when exiting the ParExpr production.
	ExitParExpr(c *ParExprContext)

	// ExitStrExpr is called when exiting the StrExpr production.
	ExitStrExpr(c *StrExprContext)

	// ExitAccFuncExpr is called when exiting the AccFuncExpr production.
	ExitAccFuncExpr(c *AccFuncExprContext)

	// ExitIntExpr is called when exiting the IntExpr production.
	ExitIntExpr(c *IntExprContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)
}

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

	// EnterReturnstmt is called when entering the returnstmt production.
	EnterReturnstmt(c *ReturnstmtContext)

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

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

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

	// EnterUnariaExpr is called when entering the UnariaExpr production.
	EnterUnariaExpr(c *UnariaExprContext)

	// EnterBoolExpr is called when entering the BoolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterNilExpr is called when entering the NilExpr production.
	EnterNilExpr(c *NilExprContext)

	// EnterIdExpr is called when entering the IdExpr production.
	EnterIdExpr(c *IdExprContext)

	// EnterNumExpr is called when entering the NumExpr production.
	EnterNumExpr(c *NumExprContext)

	// EnterParExpr is called when entering the ParExpr production.
	EnterParExpr(c *ParExprContext)

	// EnterStrExpr is called when entering the StrExpr production.
	EnterStrExpr(c *StrExprContext)

	// EnterAccFuncExpr is called when entering the AccFuncExpr production.
	EnterAccFuncExpr(c *AccFuncExprContext)

	// EnterOpExpr is called when entering the OpExpr production.
	EnterOpExpr(c *OpExprContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitReturnstmt is called when exiting the returnstmt production.
	ExitReturnstmt(c *ReturnstmtContext)

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

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

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

	// ExitUnariaExpr is called when exiting the UnariaExpr production.
	ExitUnariaExpr(c *UnariaExprContext)

	// ExitBoolExpr is called when exiting the BoolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitNilExpr is called when exiting the NilExpr production.
	ExitNilExpr(c *NilExprContext)

	// ExitIdExpr is called when exiting the IdExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitNumExpr is called when exiting the NumExpr production.
	ExitNumExpr(c *NumExprContext)

	// ExitParExpr is called when exiting the ParExpr production.
	ExitParExpr(c *ParExprContext)

	// ExitStrExpr is called when exiting the StrExpr production.
	ExitStrExpr(c *StrExprContext)

	// ExitAccFuncExpr is called when exiting the AccFuncExpr production.
	ExitAccFuncExpr(c *AccFuncExprContext)

	// ExitOpExpr is called when exiting the OpExpr production.
	ExitOpExpr(c *OpExprContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)
}

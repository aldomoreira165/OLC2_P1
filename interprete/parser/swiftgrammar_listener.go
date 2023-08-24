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

	// EnterFuncdclstmt is called when entering the funcdclstmt production.
	EnterFuncdclstmt(c *FuncdclstmtContext)

	// EnterAccfuncstm is called when entering the accfuncstm production.
	EnterAccfuncstm(c *AccfuncstmContext)

	// EnterRetfuncdclstmt is called when entering the retfuncdclstmt production.
	EnterRetfuncdclstmt(c *RetfuncdclstmtContext)

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

	// ExitFuncdclstmt is called when exiting the funcdclstmt production.
	ExitFuncdclstmt(c *FuncdclstmtContext)

	// ExitAccfuncstm is called when exiting the accfuncstm production.
	ExitAccfuncstm(c *AccfuncstmContext)

	// ExitRetfuncdclstmt is called when exiting the retfuncdclstmt production.
	ExitRetfuncdclstmt(c *RetfuncdclstmtContext)

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

	// ExitOpExpr is called when exiting the OpExpr production.
	ExitOpExpr(c *OpExprContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)
}

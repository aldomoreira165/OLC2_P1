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

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

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

	// EnterParExpr is called when entering the ParExpr production.
	EnterParExpr(c *ParExprContext)

	// EnterPrimExpr is called when entering the PrimExpr production.
	EnterPrimExpr(c *PrimExprContext)

	// EnterOpExpr is called when entering the OpExpr production.
	EnterOpExpr(c *OpExprContext)

	// EnterNumExpr is called when entering the NumExpr production.
	EnterNumExpr(c *NumExprContext)

	// EnterIdExpr is called when entering the IdExpr production.
	EnterIdExpr(c *IdExprContext)

	// EnterStrExpr is called when entering the StrExpr production.
	EnterStrExpr(c *StrExprContext)

	// EnterBoolExpr is called when entering the BoolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

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

	// ExitParExpr is called when exiting the ParExpr production.
	ExitParExpr(c *ParExprContext)

	// ExitPrimExpr is called when exiting the PrimExpr production.
	ExitPrimExpr(c *PrimExprContext)

	// ExitOpExpr is called when exiting the OpExpr production.
	ExitOpExpr(c *OpExprContext)

	// ExitNumExpr is called when exiting the NumExpr production.
	ExitNumExpr(c *NumExprContext)

	// ExitIdExpr is called when exiting the IdExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitStrExpr is called when exiting the StrExpr production.
	ExitStrExpr(c *StrExprContext)

	// ExitBoolExpr is called when exiting the BoolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)
}

// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SwiftGrammarParser.
type SwiftGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SwiftGrammarParser#s.
	VisitS(ctx *SContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#funcdclstmt.
	VisitFuncdclstmt(ctx *FuncdclstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#accfuncstm.
	VisitAccfuncstm(ctx *AccfuncstmContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#retfuncdclstmt.
	VisitRetfuncdclstmt(ctx *RetfuncdclstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#printstmt.
	VisitPrintstmt(ctx *PrintstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#breakstmt.
	VisitBreakstmt(ctx *BreakstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#ifstmt.
	VisitIfstmt(ctx *IfstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#elseifstmt.
	VisitElseifstmt(ctx *ElseifstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#switchstmt.
	VisitSwitchstmt(ctx *SwitchstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#caseStmt.
	VisitCaseStmt(ctx *CaseStmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#defaultCase.
	VisitDefaultCase(ctx *DefaultCaseContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#typedDeclstmt.
	VisitTypedDeclstmt(ctx *TypedDeclstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#untypedDeclstmt.
	VisitUntypedDeclstmt(ctx *UntypedDeclstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#optionalTypedDeclstmt.
	VisitOptionalTypedDeclstmt(ctx *OptionalTypedDeclstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#asignstmt.
	VisitAsignstmt(ctx *AsignstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#whilestmt.
	VisitWhilestmt(ctx *WhilestmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#Incremento.
	VisitIncremento(ctx *IncrementoContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#Decremento.
	VisitDecremento(ctx *DecrementoContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#UnariaExpr.
	VisitUnariaExpr(ctx *UnariaExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#BoolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#NilExpr.
	VisitNilExpr(ctx *NilExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#NumExpr.
	VisitNumExpr(ctx *NumExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#ParExpr.
	VisitParExpr(ctx *ParExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#StrExpr.
	VisitStrExpr(ctx *StrExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#OpExpr.
	VisitOpExpr(ctx *OpExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#tipo.
	VisitTipo(ctx *TipoContext) interface{}
}

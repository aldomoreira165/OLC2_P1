// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseSwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type BaseSwiftGrammarListener struct{}

var _ SwiftGrammarListener = &BaseSwiftGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSwiftGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSwiftGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSwiftGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSwiftGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BaseSwiftGrammarListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BaseSwiftGrammarListener) ExitS(ctx *SContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSwiftGrammarListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSwiftGrammarListener) ExitBlock(ctx *BlockContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseSwiftGrammarListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseSwiftGrammarListener) ExitStmt(ctx *StmtContext) {}

// EnterReturnstmt is called when production returnstmt is entered.
func (s *BaseSwiftGrammarListener) EnterReturnstmt(ctx *ReturnstmtContext) {}

// ExitReturnstmt is called when production returnstmt is exited.
func (s *BaseSwiftGrammarListener) ExitReturnstmt(ctx *ReturnstmtContext) {}

// EnterFuncionNormal is called when production FuncionNormal is entered.
func (s *BaseSwiftGrammarListener) EnterFuncionNormal(ctx *FuncionNormalContext) {}

// ExitFuncionNormal is called when production FuncionNormal is exited.
func (s *BaseSwiftGrammarListener) ExitFuncionNormal(ctx *FuncionNormalContext) {}

// EnterFuncionRetorno is called when production FuncionRetorno is entered.
func (s *BaseSwiftGrammarListener) EnterFuncionRetorno(ctx *FuncionRetornoContext) {}

// ExitFuncionRetorno is called when production FuncionRetorno is exited.
func (s *BaseSwiftGrammarListener) ExitFuncionRetorno(ctx *FuncionRetornoContext) {}

// EnterAccfuncstm is called when production accfuncstm is entered.
func (s *BaseSwiftGrammarListener) EnterAccfuncstm(ctx *AccfuncstmContext) {}

// ExitAccfuncstm is called when production accfuncstm is exited.
func (s *BaseSwiftGrammarListener) ExitAccfuncstm(ctx *AccfuncstmContext) {}

// EnterPrintstmt is called when production printstmt is entered.
func (s *BaseSwiftGrammarListener) EnterPrintstmt(ctx *PrintstmtContext) {}

// ExitPrintstmt is called when production printstmt is exited.
func (s *BaseSwiftGrammarListener) ExitPrintstmt(ctx *PrintstmtContext) {}

// EnterBreakstmt is called when production breakstmt is entered.
func (s *BaseSwiftGrammarListener) EnterBreakstmt(ctx *BreakstmtContext) {}

// ExitBreakstmt is called when production breakstmt is exited.
func (s *BaseSwiftGrammarListener) ExitBreakstmt(ctx *BreakstmtContext) {}

// EnterIfstmt is called when production ifstmt is entered.
func (s *BaseSwiftGrammarListener) EnterIfstmt(ctx *IfstmtContext) {}

// ExitIfstmt is called when production ifstmt is exited.
func (s *BaseSwiftGrammarListener) ExitIfstmt(ctx *IfstmtContext) {}

// EnterElseifstmt is called when production elseifstmt is entered.
func (s *BaseSwiftGrammarListener) EnterElseifstmt(ctx *ElseifstmtContext) {}

// ExitElseifstmt is called when production elseifstmt is exited.
func (s *BaseSwiftGrammarListener) ExitElseifstmt(ctx *ElseifstmtContext) {}

// EnterSwitchstmt is called when production switchstmt is entered.
func (s *BaseSwiftGrammarListener) EnterSwitchstmt(ctx *SwitchstmtContext) {}

// ExitSwitchstmt is called when production switchstmt is exited.
func (s *BaseSwiftGrammarListener) ExitSwitchstmt(ctx *SwitchstmtContext) {}

// EnterCaseStmt is called when production caseStmt is entered.
func (s *BaseSwiftGrammarListener) EnterCaseStmt(ctx *CaseStmtContext) {}

// ExitCaseStmt is called when production caseStmt is exited.
func (s *BaseSwiftGrammarListener) ExitCaseStmt(ctx *CaseStmtContext) {}

// EnterDefaultCase is called when production defaultCase is entered.
func (s *BaseSwiftGrammarListener) EnterDefaultCase(ctx *DefaultCaseContext) {}

// ExitDefaultCase is called when production defaultCase is exited.
func (s *BaseSwiftGrammarListener) ExitDefaultCase(ctx *DefaultCaseContext) {}

// EnterTypedDeclstmt is called when production typedDeclstmt is entered.
func (s *BaseSwiftGrammarListener) EnterTypedDeclstmt(ctx *TypedDeclstmtContext) {}

// ExitTypedDeclstmt is called when production typedDeclstmt is exited.
func (s *BaseSwiftGrammarListener) ExitTypedDeclstmt(ctx *TypedDeclstmtContext) {}

// EnterUntypedDeclstmt is called when production untypedDeclstmt is entered.
func (s *BaseSwiftGrammarListener) EnterUntypedDeclstmt(ctx *UntypedDeclstmtContext) {}

// ExitUntypedDeclstmt is called when production untypedDeclstmt is exited.
func (s *BaseSwiftGrammarListener) ExitUntypedDeclstmt(ctx *UntypedDeclstmtContext) {}

// EnterOptionalTypedDeclstmt is called when production optionalTypedDeclstmt is entered.
func (s *BaseSwiftGrammarListener) EnterOptionalTypedDeclstmt(ctx *OptionalTypedDeclstmtContext) {}

// ExitOptionalTypedDeclstmt is called when production optionalTypedDeclstmt is exited.
func (s *BaseSwiftGrammarListener) ExitOptionalTypedDeclstmt(ctx *OptionalTypedDeclstmtContext) {}

// EnterAsignstmt is called when production asignstmt is entered.
func (s *BaseSwiftGrammarListener) EnterAsignstmt(ctx *AsignstmtContext) {}

// ExitAsignstmt is called when production asignstmt is exited.
func (s *BaseSwiftGrammarListener) ExitAsignstmt(ctx *AsignstmtContext) {}

// EnterWhilestmt is called when production whilestmt is entered.
func (s *BaseSwiftGrammarListener) EnterWhilestmt(ctx *WhilestmtContext) {}

// ExitWhilestmt is called when production whilestmt is exited.
func (s *BaseSwiftGrammarListener) ExitWhilestmt(ctx *WhilestmtContext) {}

// EnterIncremento is called when production Incremento is entered.
func (s *BaseSwiftGrammarListener) EnterIncremento(ctx *IncrementoContext) {}

// ExitIncremento is called when production Incremento is exited.
func (s *BaseSwiftGrammarListener) ExitIncremento(ctx *IncrementoContext) {}

// EnterDecremento is called when production Decremento is entered.
func (s *BaseSwiftGrammarListener) EnterDecremento(ctx *DecrementoContext) {}

// ExitDecremento is called when production Decremento is exited.
func (s *BaseSwiftGrammarListener) ExitDecremento(ctx *DecrementoContext) {}

// EnterUnariaExpr is called when production UnariaExpr is entered.
func (s *BaseSwiftGrammarListener) EnterUnariaExpr(ctx *UnariaExprContext) {}

// ExitUnariaExpr is called when production UnariaExpr is exited.
func (s *BaseSwiftGrammarListener) ExitUnariaExpr(ctx *UnariaExprContext) {}

// EnterBoolExpr is called when production BoolExpr is entered.
func (s *BaseSwiftGrammarListener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production BoolExpr is exited.
func (s *BaseSwiftGrammarListener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterNilExpr is called when production NilExpr is entered.
func (s *BaseSwiftGrammarListener) EnterNilExpr(ctx *NilExprContext) {}

// ExitNilExpr is called when production NilExpr is exited.
func (s *BaseSwiftGrammarListener) ExitNilExpr(ctx *NilExprContext) {}

// EnterIdExpr is called when production IdExpr is entered.
func (s *BaseSwiftGrammarListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production IdExpr is exited.
func (s *BaseSwiftGrammarListener) ExitIdExpr(ctx *IdExprContext) {}

// EnterNumExpr is called when production NumExpr is entered.
func (s *BaseSwiftGrammarListener) EnterNumExpr(ctx *NumExprContext) {}

// ExitNumExpr is called when production NumExpr is exited.
func (s *BaseSwiftGrammarListener) ExitNumExpr(ctx *NumExprContext) {}

// EnterParExpr is called when production ParExpr is entered.
func (s *BaseSwiftGrammarListener) EnterParExpr(ctx *ParExprContext) {}

// ExitParExpr is called when production ParExpr is exited.
func (s *BaseSwiftGrammarListener) ExitParExpr(ctx *ParExprContext) {}

// EnterStrExpr is called when production StrExpr is entered.
func (s *BaseSwiftGrammarListener) EnterStrExpr(ctx *StrExprContext) {}

// ExitStrExpr is called when production StrExpr is exited.
func (s *BaseSwiftGrammarListener) ExitStrExpr(ctx *StrExprContext) {}

// EnterAccFuncExpr is called when production AccFuncExpr is entered.
func (s *BaseSwiftGrammarListener) EnterAccFuncExpr(ctx *AccFuncExprContext) {}

// ExitAccFuncExpr is called when production AccFuncExpr is exited.
func (s *BaseSwiftGrammarListener) ExitAccFuncExpr(ctx *AccFuncExprContext) {}

// EnterOpExpr is called when production OpExpr is entered.
func (s *BaseSwiftGrammarListener) EnterOpExpr(ctx *OpExprContext) {}

// ExitOpExpr is called when production OpExpr is exited.
func (s *BaseSwiftGrammarListener) ExitOpExpr(ctx *OpExprContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BaseSwiftGrammarListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BaseSwiftGrammarListener) ExitTipo(ctx *TipoContext) {}

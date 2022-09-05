// Code generated from ./parser/JsonPath.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JsonPath

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJsonPathListener is a complete listener for a parse tree produced by JsonPathParser.
type BaseJsonPathListener struct{}

var _ JsonPathListener = &BaseJsonPathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJsonPathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJsonPathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJsonPathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJsonPathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterJsonpath is called when production jsonpath is entered.
func (s *BaseJsonPathListener) EnterJsonpath(ctx *JsonpathContext) {}

// ExitJsonpath is called when production jsonpath is exited.
func (s *BaseJsonPathListener) ExitJsonpath(ctx *JsonpathContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseJsonPathListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseJsonPathListener) ExitRoot(ctx *RootContext) {}

// EnterFieldAccess is called when production fieldAccess is entered.
func (s *BaseJsonPathListener) EnterFieldAccess(ctx *FieldAccessContext) {}

// ExitFieldAccess is called when production fieldAccess is exited.
func (s *BaseJsonPathListener) ExitFieldAccess(ctx *FieldAccessContext) {}

// EnterLocalFieldAccess is called when production localFieldAccess is entered.
func (s *BaseJsonPathListener) EnterLocalFieldAccess(ctx *LocalFieldAccessContext) {}

// ExitLocalFieldAccess is called when production localFieldAccess is exited.
func (s *BaseJsonPathListener) ExitLocalFieldAccess(ctx *LocalFieldAccessContext) {}

// EnterGetAll is called when production GetAll is entered.
func (s *BaseJsonPathListener) EnterGetAll(ctx *GetAllContext) {}

// ExitGetAll is called when production GetAll is exited.
func (s *BaseJsonPathListener) ExitGetAll(ctx *GetAllContext) {}

// EnterGetByIndices is called when production GetByIndices is entered.
func (s *BaseJsonPathListener) EnterGetByIndices(ctx *GetByIndicesContext) {}

// ExitGetByIndices is called when production GetByIndices is exited.
func (s *BaseJsonPathListener) ExitGetByIndices(ctx *GetByIndicesContext) {}

// EnterGetByRange is called when production GetByRange is entered.
func (s *BaseJsonPathListener) EnterGetByRange(ctx *GetByRangeContext) {}

// ExitGetByRange is called when production GetByRange is exited.
func (s *BaseJsonPathListener) ExitGetByRange(ctx *GetByRangeContext) {}

// EnterGetByQuery is called when production GetByQuery is entered.
func (s *BaseJsonPathListener) EnterGetByQuery(ctx *GetByQueryContext) {}

// ExitGetByQuery is called when production GetByQuery is exited.
func (s *BaseJsonPathListener) ExitGetByQuery(ctx *GetByQueryContext) {}

// EnterGetByIndexBackward is called when production GetByIndexBackward is entered.
func (s *BaseJsonPathListener) EnterGetByIndexBackward(ctx *GetByIndexBackwardContext) {}

// ExitGetByIndexBackward is called when production GetByIndexBackward is exited.
func (s *BaseJsonPathListener) ExitGetByIndexBackward(ctx *GetByIndexBackwardContext) {}

// EnterAllRange is called when production AllRange is entered.
func (s *BaseJsonPathListener) EnterAllRange(ctx *AllRangeContext) {}

// ExitAllRange is called when production AllRange is exited.
func (s *BaseJsonPathListener) ExitAllRange(ctx *AllRangeContext) {}

// EnterGivenStartEndIdxRange is called when production GivenStartEndIdxRange is entered.
func (s *BaseJsonPathListener) EnterGivenStartEndIdxRange(ctx *GivenStartEndIdxRangeContext) {}

// ExitGivenStartEndIdxRange is called when production GivenStartEndIdxRange is exited.
func (s *BaseJsonPathListener) ExitGivenStartEndIdxRange(ctx *GivenStartEndIdxRangeContext) {}

// EnterGivenOnlyEndIdxRange is called when production GivenOnlyEndIdxRange is entered.
func (s *BaseJsonPathListener) EnterGivenOnlyEndIdxRange(ctx *GivenOnlyEndIdxRangeContext) {}

// ExitGivenOnlyEndIdxRange is called when production GivenOnlyEndIdxRange is exited.
func (s *BaseJsonPathListener) ExitGivenOnlyEndIdxRange(ctx *GivenOnlyEndIdxRangeContext) {}

// EnterGivenOnlyStartIdxRange is called when production GivenOnlyStartIdxRange is entered.
func (s *BaseJsonPathListener) EnterGivenOnlyStartIdxRange(ctx *GivenOnlyStartIdxRangeContext) {}

// ExitGivenOnlyStartIdxRange is called when production GivenOnlyStartIdxRange is exited.
func (s *BaseJsonPathListener) ExitGivenOnlyStartIdxRange(ctx *GivenOnlyStartIdxRangeContext) {}

// EnterIntSequence is called when production intSequence is entered.
func (s *BaseJsonPathListener) EnterIntSequence(ctx *IntSequenceContext) {}

// ExitIntSequence is called when production intSequence is exited.
func (s *BaseJsonPathListener) ExitIntSequence(ctx *IntSequenceContext) {}

// EnterQueryFieldValueLocalAccess is called when production QueryFieldValueLocalAccess is entered.
func (s *BaseJsonPathListener) EnterQueryFieldValueLocalAccess(ctx *QueryFieldValueLocalAccessContext) {
}

// ExitQueryFieldValueLocalAccess is called when production QueryFieldValueLocalAccess is exited.
func (s *BaseJsonPathListener) ExitQueryFieldValueLocalAccess(ctx *QueryFieldValueLocalAccessContext) {
}

// EnterQueryFieldValueRootAccess is called when production QueryFieldValueRootAccess is entered.
func (s *BaseJsonPathListener) EnterQueryFieldValueRootAccess(ctx *QueryFieldValueRootAccessContext) {
}

// ExitQueryFieldValueRootAccess is called when production QueryFieldValueRootAccess is exited.
func (s *BaseJsonPathListener) ExitQueryFieldValueRootAccess(ctx *QueryFieldValueRootAccessContext) {}

// EnterQueryFieldDoubleValue is called when production QueryFieldDoubleValue is entered.
func (s *BaseJsonPathListener) EnterQueryFieldDoubleValue(ctx *QueryFieldDoubleValueContext) {}

// ExitQueryFieldDoubleValue is called when production QueryFieldDoubleValue is exited.
func (s *BaseJsonPathListener) ExitQueryFieldDoubleValue(ctx *QueryFieldDoubleValueContext) {}

// EnterQueryFieldStringValue is called when production QueryFieldStringValue is entered.
func (s *BaseJsonPathListener) EnterQueryFieldStringValue(ctx *QueryFieldStringValueContext) {}

// ExitQueryFieldStringValue is called when production QueryFieldStringValue is exited.
func (s *BaseJsonPathListener) ExitQueryFieldStringValue(ctx *QueryFieldStringValueContext) {}

// EnterQueryFieldBoolValue is called when production QueryFieldBoolValue is entered.
func (s *BaseJsonPathListener) EnterQueryFieldBoolValue(ctx *QueryFieldBoolValueContext) {}

// ExitQueryFieldBoolValue is called when production QueryFieldBoolValue is exited.
func (s *BaseJsonPathListener) ExitQueryFieldBoolValue(ctx *QueryFieldBoolValueContext) {}

// EnterFilterOperation is called when production filterOperation is entered.
func (s *BaseJsonPathListener) EnterFilterOperation(ctx *FilterOperationContext) {}

// ExitFilterOperation is called when production filterOperation is exited.
func (s *BaseJsonPathListener) ExitFilterOperation(ctx *FilterOperationContext) {}

// EnterQueryWithLogicalOp is called when production QueryWithLogicalOp is entered.
func (s *BaseJsonPathListener) EnterQueryWithLogicalOp(ctx *QueryWithLogicalOpContext) {}

// ExitQueryWithLogicalOp is called when production QueryWithLogicalOp is exited.
func (s *BaseJsonPathListener) ExitQueryWithLogicalOp(ctx *QueryWithLogicalOpContext) {}

// EnterQueryExistence is called when production QueryExistence is entered.
func (s *BaseJsonPathListener) EnterQueryExistence(ctx *QueryExistenceContext) {}

// ExitQueryExistence is called when production QueryExistence is exited.
func (s *BaseJsonPathListener) ExitQueryExistence(ctx *QueryExistenceContext) {}

// EnterQueryWithComparator is called when production QueryWithComparator is entered.
func (s *BaseJsonPathListener) EnterQueryWithComparator(ctx *QueryWithComparatorContext) {}

// ExitQueryWithComparator is called when production QueryWithComparator is exited.
func (s *BaseJsonPathListener) ExitQueryWithComparator(ctx *QueryWithComparatorContext) {}

// EnterDbl is called when production dbl is entered.
func (s *BaseJsonPathListener) EnterDbl(ctx *DblContext) {}

// ExitDbl is called when production dbl is exited.
func (s *BaseJsonPathListener) ExitDbl(ctx *DblContext) {}

// EnterBool_type is called when production bool_type is entered.
func (s *BaseJsonPathListener) EnterBool_type(ctx *Bool_typeContext) {}

// ExitBool_type is called when production bool_type is exited.
func (s *BaseJsonPathListener) ExitBool_type(ctx *Bool_typeContext) {}

// Code generated from ./parser/JsonPath.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JsonPath

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JsonPathListener is a complete listener for a parse tree produced by JsonPathParser.
type JsonPathListener interface {
	antlr.ParseTreeListener

	// EnterJsonpath is called when entering the jsonpath production.
	EnterJsonpath(c *JsonpathContext)

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterFieldAccess is called when entering the fieldAccess production.
	EnterFieldAccess(c *FieldAccessContext)

	// EnterLocalFieldAccess is called when entering the localFieldAccess production.
	EnterLocalFieldAccess(c *LocalFieldAccessContext)

	// EnterGetAll is called when entering the GetAll production.
	EnterGetAll(c *GetAllContext)

	// EnterGetByIndices is called when entering the GetByIndices production.
	EnterGetByIndices(c *GetByIndicesContext)

	// EnterGetByRange is called when entering the GetByRange production.
	EnterGetByRange(c *GetByRangeContext)

	// EnterGetByQuery is called when entering the GetByQuery production.
	EnterGetByQuery(c *GetByQueryContext)

	// EnterGetByIndexBackward is called when entering the GetByIndexBackward production.
	EnterGetByIndexBackward(c *GetByIndexBackwardContext)

	// EnterAllRange is called when entering the AllRange production.
	EnterAllRange(c *AllRangeContext)

	// EnterGivenStartEndIdxRange is called when entering the GivenStartEndIdxRange production.
	EnterGivenStartEndIdxRange(c *GivenStartEndIdxRangeContext)

	// EnterGivenOnlyEndIdxRange is called when entering the GivenOnlyEndIdxRange production.
	EnterGivenOnlyEndIdxRange(c *GivenOnlyEndIdxRangeContext)

	// EnterGivenOnlyStartIdxRange is called when entering the GivenOnlyStartIdxRange production.
	EnterGivenOnlyStartIdxRange(c *GivenOnlyStartIdxRangeContext)

	// EnterIntSequence is called when entering the intSequence production.
	EnterIntSequence(c *IntSequenceContext)

	// EnterQueryFieldValueLocalAccess is called when entering the QueryFieldValueLocalAccess production.
	EnterQueryFieldValueLocalAccess(c *QueryFieldValueLocalAccessContext)

	// EnterQueryFieldValueRootAccess is called when entering the QueryFieldValueRootAccess production.
	EnterQueryFieldValueRootAccess(c *QueryFieldValueRootAccessContext)

	// EnterQueryFieldDoubleValue is called when entering the QueryFieldDoubleValue production.
	EnterQueryFieldDoubleValue(c *QueryFieldDoubleValueContext)

	// EnterQueryFieldStringValue is called when entering the QueryFieldStringValue production.
	EnterQueryFieldStringValue(c *QueryFieldStringValueContext)

	// EnterFilterOperation is called when entering the filterOperation production.
	EnterFilterOperation(c *FilterOperationContext)

	// EnterQueryWithLogicalOp is called when entering the QueryWithLogicalOp production.
	EnterQueryWithLogicalOp(c *QueryWithLogicalOpContext)

	// EnterQueryExistence is called when entering the QueryExistence production.
	EnterQueryExistence(c *QueryExistenceContext)

	// EnterQueryWithComparator is called when entering the QueryWithComparator production.
	EnterQueryWithComparator(c *QueryWithComparatorContext)

	// EnterDbl is called when entering the dbl production.
	EnterDbl(c *DblContext)

	// ExitJsonpath is called when exiting the jsonpath production.
	ExitJsonpath(c *JsonpathContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitFieldAccess is called when exiting the fieldAccess production.
	ExitFieldAccess(c *FieldAccessContext)

	// ExitLocalFieldAccess is called when exiting the localFieldAccess production.
	ExitLocalFieldAccess(c *LocalFieldAccessContext)

	// ExitGetAll is called when exiting the GetAll production.
	ExitGetAll(c *GetAllContext)

	// ExitGetByIndices is called when exiting the GetByIndices production.
	ExitGetByIndices(c *GetByIndicesContext)

	// ExitGetByRange is called when exiting the GetByRange production.
	ExitGetByRange(c *GetByRangeContext)

	// ExitGetByQuery is called when exiting the GetByQuery production.
	ExitGetByQuery(c *GetByQueryContext)

	// ExitGetByIndexBackward is called when exiting the GetByIndexBackward production.
	ExitGetByIndexBackward(c *GetByIndexBackwardContext)

	// ExitAllRange is called when exiting the AllRange production.
	ExitAllRange(c *AllRangeContext)

	// ExitGivenStartEndIdxRange is called when exiting the GivenStartEndIdxRange production.
	ExitGivenStartEndIdxRange(c *GivenStartEndIdxRangeContext)

	// ExitGivenOnlyEndIdxRange is called when exiting the GivenOnlyEndIdxRange production.
	ExitGivenOnlyEndIdxRange(c *GivenOnlyEndIdxRangeContext)

	// ExitGivenOnlyStartIdxRange is called when exiting the GivenOnlyStartIdxRange production.
	ExitGivenOnlyStartIdxRange(c *GivenOnlyStartIdxRangeContext)

	// ExitIntSequence is called when exiting the intSequence production.
	ExitIntSequence(c *IntSequenceContext)

	// ExitQueryFieldValueLocalAccess is called when exiting the QueryFieldValueLocalAccess production.
	ExitQueryFieldValueLocalAccess(c *QueryFieldValueLocalAccessContext)

	// ExitQueryFieldValueRootAccess is called when exiting the QueryFieldValueRootAccess production.
	ExitQueryFieldValueRootAccess(c *QueryFieldValueRootAccessContext)

	// ExitQueryFieldDoubleValue is called when exiting the QueryFieldDoubleValue production.
	ExitQueryFieldDoubleValue(c *QueryFieldDoubleValueContext)

	// ExitQueryFieldStringValue is called when exiting the QueryFieldStringValue production.
	ExitQueryFieldStringValue(c *QueryFieldStringValueContext)

	// ExitFilterOperation is called when exiting the filterOperation production.
	ExitFilterOperation(c *FilterOperationContext)

	// ExitQueryWithLogicalOp is called when exiting the QueryWithLogicalOp production.
	ExitQueryWithLogicalOp(c *QueryWithLogicalOpContext)

	// ExitQueryExistence is called when exiting the QueryExistence production.
	ExitQueryExistence(c *QueryExistenceContext)

	// ExitQueryWithComparator is called when exiting the QueryWithComparator production.
	ExitQueryWithComparator(c *QueryWithComparatorContext)

	// ExitDbl is called when exiting the dbl production.
	ExitDbl(c *DblContext)
}

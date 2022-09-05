package protopath

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/caraml-dev/protopath/parser"
)

// JsonPathVisitor responsible to visit the tree parser
// and convert it into list of operations
type JsonPathVisitor struct {
	operations        []Operation
	fieldGetter       FieldGetter
	comparatorMapping map[int]operator
}

// NewJsonPathVisitor creates JsonPathVisitor to traverse all the parsed syntax into operation
func NewJsonPathVisitor(fieldGetter FieldGetter) *JsonPathVisitor {
	comparatorMapping := map[int]operator{
		parser.JsonPathLexerEQ:         eq,
		parser.JsonPathLexerNEQ:        nEq,
		parser.JsonPathLexerGREATER:    greater,
		parser.JsonPathLexerGREATER_EQ: greaterEq,
		parser.JsonPathLexerLESS:       less,
		parser.JsonPathLexerLESS_EQ:    lessEq,
		parser.JsonPathLexerAND:        and,
		parser.JsonPathLexerOR:         or,
	}
	return &JsonPathVisitor{
		operations:        make([]Operation, 0),
		fieldGetter:       fieldGetter,
		comparatorMapping: comparatorMapping,
	}
}

func (j *JsonPathVisitor) traverseTree(tree antlr.Tree, ops []Operation) ([]Operation, error) {
	switch tr := tree.(type) {
	case antlr.ErrorNode:
		return nil, fmt.Errorf("got syntax error")
	case antlr.TerminalNode:
		return ops, nil
	case *parser.RootContext:
		return j.visitRootContext(tr, ops)
	case *parser.JsonpathContext:
		return j.visitJsonPathContext(tr, ops)
	case *parser.FieldAccessContext:
		return j.visitFieldAccessContext(tr, false, ops)
	case *parser.ArrayQualifierContext:
		return j.visitArrayQualifierContext(tr, ops)
	case *parser.GetAllContext:
		return j.visitGetAllContext(tr, ops)
	case *parser.GetByIndicesContext:
		return j.visitGetByIndicesContext(tr, ops)
	case *parser.GetByIndexBackwardContext:
		return j.visitGetByIndexBackwardContext(tr, ops)
	case *parser.GetByRangeContext:
		return j.visitGetByRangeContext(tr, ops)
	case *parser.GetByQueryContext:
		return j.visitGetByQueryContext(tr, ops)
	case *parser.FilterOperationContext:
		return j.visitFilterOperationContext(tr, ops)
	}
	return ops, nil
}

func (j *JsonPathVisitor) visitJsonPathContext(jsonPathContext *parser.JsonpathContext, ops []Operation) ([]Operation, error) {
	for _, child := range jsonPathContext.GetChildren() {
		childOps, err := j.traverseTree(child, ops)
		if err != nil {
			return nil, err
		}
		ops = append(ops, childOps...)
	}
	return ops, nil
}

func (j *JsonPathVisitor) visitLocalFieldAccess(ctx *parser.LocalFieldAccessContext, ops []Operation) ([]Operation, error) {
	for _, child := range ctx.GetChildren() {
		switch childT := child.(type) {
		case *parser.FieldAccessContext:
			childOps, err := j.visitFieldAccessContext(childT, false, ops)
			if err != nil {
				return nil, err
			}
			ops = childOps
		case *parser.LocalFieldAccessContext:
			childOps, err := j.visitLocalFieldAccess(childT, ops)
			if err != nil {
				return nil, err
			}
			ops = childOps
		default:
			// most likely it goes to ErrorNode
			// traverseTree will handle the case when meet ErrorNode
			childOps, err := j.traverseTree(childT, ops)
			if err != nil {
				return nil, err
			}
			ops = childOps
		}
	}
	return ops, nil
}

func (j *JsonPathVisitor) visitRootContext(rootContext *parser.RootContext, ops []Operation) ([]Operation, error) {
	hasRootIdentifier := false
	for _, child := range rootContext.GetChildren() {
		switch childT := child.(type) {
		case antlr.TerminalNode:
			if childT.GetSymbol().GetTokenType() == parser.JsonPathLexerROOT_IDENTIFIER {
				hasRootIdentifier = true
			}
		case *parser.FieldAccessContext:
			childOps, err := j.visitFieldAccessContext(childT, hasRootIdentifier, ops)
			if err != nil {
				return nil, err
			}
			ops = childOps
		case *parser.RootContext:
			childOps, err := j.visitRootContext(childT, ops)
			if err != nil {
				return nil, err
			}
			ops = childOps
		default:
			// most likely it goes to ErrorNode
			// traverseTree will handle the case when meet ErrorNode
			childOps, err := j.traverseTree(childT, ops)
			if err != nil {
				return nil, err
			}
			ops = childOps
		}
	}
	return ops, nil
}

func (j *JsonPathVisitor) visitFieldAccessContext(fieldAccess *parser.FieldAccessContext, fromRoot bool, ops []Operation) ([]Operation, error) {
	var err error
	for _, child := range fieldAccess.GetChildren() {
		switch childT := child.(type) {
		case antlr.ErrorNode:
			return j.traverseTree(childT, ops)
		case antlr.TerminalNode:
			if childT.GetSymbol().GetTokenType() == parser.JsonPathLexerIDENTIFIER {
				ops = append(ops, &FieldAccessOperation{
					fieldName:   childT.GetText(),
					fromRoot:    fromRoot,
					fieldGetter: j.fieldGetter,
				})
			}
		default:
			ops, err = j.traverseTree(child, ops)
			if err != nil {
				return nil, err
			}
		}

	}
	return ops, nil
}

func (j *JsonPathVisitor) visitArrayQualifierContext(ctx *parser.ArrayQualifierContext, ops []Operation) ([]Operation, error) {
	for _, child := range ctx.GetChildren() {
		childOp, err := j.traverseTree(child, ops)
		if err != nil {
			return nil, err
		}
		ops = childOp
	}
	return ops, nil
}

func (j *JsonPathVisitor) visitGetAllContext(ctx *parser.GetAllContext, ops []Operation) ([]Operation, error) {
	ops = append(ops, &ArrayAccessOperation{method: &getByRange{}})
	return ops, nil
}

func (j *JsonPathVisitor) visitGetByIndicesContext(ctx *parser.GetByIndicesContext, ops []Operation) ([]Operation, error) {
	intSequence, ok := ctx.IntSequence().(*parser.IntSequenceContext)
	if !ok {
		return nil, fmt.Errorf("error accessing int sequence")
	}
	idxs := make([]int, len(intSequence.AllINT()))
	for i, idxVal := range intSequence.AllINT() {
		intVal, err := strconv.ParseInt(idxVal.GetText(), 10, 64)
		if err != nil {
			return nil, err
		}
		idxs[i] = int(intVal)
	}
	arrayAccessOp := &ArrayAccessOperation{
		method: &getByIndices{indices: idxs},
	}
	ops = append(ops, arrayAccessOp)
	return ops, nil
}

func (j *JsonPathVisitor) visitGetByIndexBackwardContext(ctx *parser.GetByIndexBackwardContext, ops []Operation) ([]Operation, error) {
	idx, err := strconv.ParseInt(ctx.INT().GetText(), 10, 64)
	if err != nil {
		return nil, err
	}

	ops = append(ops, &ArrayAccessOperation{
		method: &getByBackwardIndex{index: int(idx)},
	})
	return ops, nil
}

func (j *JsonPathVisitor) visitGetByRangeContext(ctx *parser.GetByRangeContext, ops []Operation) ([]Operation, error) {
	var rangeMethod *getByRange
	switch rangeType := ctx.ArrayRange().(type) {
	case *parser.GivenOnlyEndIdxRangeContext:
		endIdx, err := strconv.ParseInt(rangeType.INT(0).GetText(), 10, 64)
		if err != nil {
			return nil, err
		}
		endIdxInt := int(endIdx)
		rangeMethod = &getByRange{endIdx: &endIdxInt}
	case *parser.GivenOnlyStartIdxRangeContext:
		startIdx, err := strconv.ParseInt(rangeType.INT(0).GetText(), 10, 64)
		if err != nil {
			return nil, err
		}
		startIdxInt := int(startIdx)
		rangeMethod = &getByRange{startIdx: &startIdxInt}
	case *parser.GivenStartEndIdxRangeContext:
		startIdx, err := strconv.ParseInt(rangeType.INT(0).GetText(), 10, 64)
		if err != nil {
			return nil, err
		}
		startIdxInt := int(startIdx)
		endIdx, err := strconv.ParseInt(rangeType.INT(1).GetText(), 10, 64)
		if err != nil {
			return nil, err
		}
		endIdxInt := int(endIdx)
		rangeMethod = &getByRange{startIdx: &startIdxInt, endIdx: &endIdxInt}
	case *parser.AllRangeContext:
		rangeMethod = &getByRange{}
	}
	ops = append(ops, &ArrayAccessOperation{method: rangeMethod})
	return ops, nil
}

func (j *JsonPathVisitor) visitGetByQueryContext(ctx *parser.GetByQueryContext, ops []Operation) ([]Operation, error) {
	for _, child := range ctx.GetChildren() {
		childOp, err := j.traverseTree(child, ops)
		if err != nil {
			return nil, err
		}
		ops = childOp
	}
	return ops, nil
}

func (j *JsonPathVisitor) visitFilterOperationContext(ctx *parser.FilterOperationContext, ops []Operation) ([]Operation, error) {

	var filterOp *QueryFilterOp
	switch queryExpr := ctx.QueryExpr().(type) {
	case *parser.QueryExistenceContext:
		queryExistenceOps, err := j.visitQueryExistenceContext(queryExpr, make([]Operation, 0))
		if err != nil {
			return nil, err
		}
		filterOp = &QueryFilterOp{
			LeftOp: queryExistenceOps,
		}
	case *parser.QueryWithComparatorContext:
		res, err := j.visitQueryWithComparatorContext(queryExpr, make([]Operation, 0))
		if err != nil {
			return nil, err
		}
		filterOp = res
	case *parser.QueryWithLogicalOpContext:
		res, err := j.visitQueryWithLogicalOpContext(queryExpr, make([]Operation, 0))
		if err != nil {
			return nil, err
		}
		filterOp = res
	}

	ops = append(ops, filterOp)
	return ops, nil
}

func (j *JsonPathVisitor) visitQueryWithLogicalOpContext(ctx *parser.QueryWithLogicalOpContext, ops []Operation) (*QueryFilterOp, error) {
	leftOps, err := j.getQueryExprOperation(ctx.QueryExpr(0), ops)
	if err != nil {
		return nil, err
	}
	rightOps, err := j.getQueryExprOperation(ctx.QueryExpr(1), ops)
	if err != nil {
		return nil, err
	}

	operator, found := j.comparatorMapping[ctx.GetLogical_op().GetTokenType()]
	if !found {
		return nil, fmt.Errorf("operator is not found %v ", ctx.GetLogical_op().GetText())
	}

	return &QueryFilterOp{
		LeftOp:   leftOps,
		RightOp:  rightOps,
		Operator: operator,
	}, nil
}

func (j *JsonPathVisitor) getQueryExprOperation(ctx parser.IQueryExprContext, ops []Operation) (*QueryFilterOp, error) {
	switch queryExpr := ctx.(type) {
	case *parser.QueryWithComparatorContext:
		return j.visitQueryWithComparatorContext(queryExpr, ops)
	case *parser.QueryWithLogicalOpContext:
		return j.visitQueryWithLogicalOpContext(queryExpr, ops)
	}
	return nil, nil
}

func (j *JsonPathVisitor) visitQueryWithComparatorContext(ctx *parser.QueryWithComparatorContext, ops []Operation) (*QueryFilterOp, error) {
	leftOps, err := j.getQueryFieldValueOperations(ctx.QueryFieldValue(0), ops)
	if err != nil {
		return nil, err
	}
	rightOps, err := j.getQueryFieldValueOperations(ctx.QueryFieldValue(1), ops)
	if err != nil {
		return nil, err
	}

	filterOp := &QueryFilterOp{
		LeftOp:   leftOps,
		RightOp:  rightOps,
		Operator: j.comparatorMapping[ctx.GetComparator_op().GetTokenType()],
	}
	return filterOp, nil
}

func (j *JsonPathVisitor) getQueryFieldValueOperations(ctx antlr.Tree, ops []Operation) (any, error) {
	switch qFieldVal := ctx.(type) {
	case *parser.QueryFieldStringValueContext:
		// hack remove ' character
		return strings.Trim(qFieldVal.GetText(), "'"), nil
	case *parser.QueryFieldBoolValueContext:
		return toBool(qFieldVal.GetText())
	case *parser.QueryFieldValueRootAccessContext:
		rootCtx, ok := qFieldVal.Root().(*parser.RootContext)
		if !ok {
			return nil, fmt.Errorf("got unexpected type for root context")
		}
		childOps, err := j.visitRootContext(rootCtx, ops)
		if err != nil {
			return nil, err
		}
		return childOps, nil
	case *parser.QueryFieldValueLocalAccessContext:
		localCtx, ok := qFieldVal.LocalFieldAccess().(*parser.LocalFieldAccessContext)
		if !ok {
			return nil, fmt.Errorf("got unexpected type for local field context")
		}
		childOps, err := j.visitLocalFieldAccess(localCtx, ops)
		if err != nil {
			return nil, err
		}
		return childOps, nil
	case *parser.QueryFieldDoubleValueContext:
		doubleVal, err := toFloat64(qFieldVal.GetText())
		if err != nil {
			return nil, err
		}
		return doubleVal, nil
	default:
		return nil, fmt.Errorf("got syntax error")
	}
}

func (j *JsonPathVisitor) visitQueryExistenceContext(ctx *parser.QueryExistenceContext, ops []Operation) (any, error) {
	for _, child := range ctx.GetChildren() {
		return j.getQueryFieldValueOperations(child, ops)
	}

	return ops, nil
}

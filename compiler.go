package jsonpath

import (
	"context"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/caraml-dev/jsonpath/parser"
)

type Operation interface {
	Lookup(ctx context.Context, obj, rootObj interface{}) (interface{}, error)
}

// Compiler represents operations and way to fetch field based on jsonpath syntax
type Compiler struct {
	Operations  []Operation
	Path        string
	fieldGetter FieldGetter
}

// NewJsonPathCompiler compile all operations given the jsonpath
func NewJsonPathCompiler(jsonpath string) (*Compiler, error) {
	is := antlr.NewInputStream(jsonpath)
	lexer := parser.NewJsonPathLexer(is)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	fieldGetter := &ProtoFieldGetter{}
	p := parser.NewJsonPathParser(stream)
	visitor := NewJsonPathVisitor(fieldGetter)
	allOps, err := visitor.traverseTree(p.Jsonpath(), visitor.operations)
	if err != nil {
		return nil, err
	}
	return &Compiler{
		Operations:  allOps,
		Path:        jsonpath,
		fieldGetter: fieldGetter,
	}, nil
}

func (c *Compiler) Lookup(ctx context.Context, obj interface{}) (interface{}, error) {
	res := obj
	var err error
	for _, op := range c.Operations {
		res, err = op.Lookup(ctx, res, obj)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

package protopath

import (
	"context"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/caraml-dev/protopath/parser"
	"google.golang.org/protobuf/proto"
)

// Operation responsible to find the value based on compiled jsonpath syntax
type Operation interface {
	Lookup(ctx context.Context, obj, rootObj any) (any, error)
}

// Compiled represents operations and way to fetch field based on jsonpath syntax
type Compiled struct {
	Operations  []Operation
	Path        string
	fieldGetter FieldGetter
}

// NewJsonPathCompiler compile all operations given the jsonpath
func NewJsonPathCompiler(jsonpath string) (*Compiled, error) {
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
	return &Compiled{
		Operations:  allOps,
		Path:        jsonpath,
		fieldGetter: fieldGetter,
	}, nil
}

// Lookup run all the compiled operation based on the jsonpath syntax given proto message obj
// Return of this method not necessary is proto message it can be anything
func (c *Compiled) Lookup(ctx context.Context, obj proto.Message) (any, error) {
	var res any = obj
	var err error
	for _, op := range c.Operations {
		res, err = op.Lookup(ctx, res, obj)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

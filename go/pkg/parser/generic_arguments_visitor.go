package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type GenericArgumentsVisitor struct {
	*BaseMojoParserVisitor
}

func NewGenericArgumentsVisitor() *GenericArgumentsVisitor {
	return &GenericArgumentsVisitor{}
}

func (a *GenericArgumentsVisitor) VisitGenericArgumentClause(ctx *GenericArgumentClauseContext) interface{} {
	argumentList := ctx.GenericArguments()
	if argumentList != nil {
		return argumentList.Accept(a)
	}

	return []*lang.NominalType{}
}

func (a *GenericArgumentsVisitor) VisitGenericArguments(ctx *GenericArgumentsContext) interface{} {
	var arguments []*lang.NominalType
	argumentCtxes := ctx.AllGenericArgument()
	for _, argumentCtx := range argumentCtxes {
		visitor := NewGenericArgumentsVisitor()
		argument := argumentCtx.Accept(visitor).(*lang.NominalType)

		arguments = append(arguments, argument)
	}
	return arguments
}

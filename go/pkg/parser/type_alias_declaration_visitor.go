package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type TypeAliasDeclarationVisitor struct {
	*BaseMojoParserVisitor
}

func NewTypeAliasDeclarationVisitor() *TypeAliasDeclarationVisitor {
	visitor := &TypeAliasDeclarationVisitor{}
	return visitor
}

func (s *TypeAliasDeclarationVisitor) VisitTypeAliasDeclaration(ctx *TypeAliasDeclarationContext) interface{} {
	decl := &lang.TypeAliasDecl{}

	typeAliasName := ctx.TypealiasName()
	if typeAliasName != nil {
		var ok bool
		if decl.Name, ok = typeAliasName.Accept(s).(string); ok {
			decl.GenericParameters = GetGenericParameters(ctx.GenericParameterClause())

			assignmentCtx := ctx.TypealiasAssignment()
			if assignmentCtx != nil {
				if decl.Type, ok = assignmentCtx.Accept(s).(*lang.NominalType); ok {
					return decl
				}
			}
		}
	}
	return nil
}

func (s *TypeAliasDeclarationVisitor) VisitTypealiasAssignment(ctx *TypealiasAssignmentContext) interface{} {
	return GetType(ctx.Type_())
}

func (s *TypeAliasDeclarationVisitor) VisitTypealiasName(ctx *TypealiasNameContext) interface{} {
	return ctx.GetText()
}

package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type StatementsVisitor struct {
	*BaseMojoParserVisitor
}

func NewStatementsVisitor() *StatementsVisitor {
	visitor := &StatementsVisitor{}
	return visitor
}

func (s *StatementsVisitor) VisitStatements(ctx *StatementsContext) interface{} {
	statementCtxes := ctx.AllStatement()
	var statements []*lang.Statement
	for _, statementCtx := range statementCtxes {
		if statement, ok := statementCtx.Accept(s).(*lang.Statement); ok {
			if statement != nil {
				statements = append(statements, statement)
			}
		}
	}
	return statements
}

func (s *StatementsVisitor) VisitStatement(ctx *StatementContext) interface{}  {
	declCtx := ctx.Declaration()
	if declCtx != nil {
		return declCtx.Accept(s)
	}

	expressionCtx := ctx.Expression()
	if expressionCtx != nil {
		visitor := NewExpressionVisitor()
		expression := expressionCtx.Accept(visitor).(*lang.Expression)
		return lang.NewExpressionStatement(expression)
	}

	return nil
}

func (s *StatementsVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	packageCtx := ctx.PackageDeclaration()
	if packageCtx != nil {
		visitor := NewPackageDeclarationVisitor()
		if decl, ok := packageCtx.Accept(visitor).(*lang.PackageDecl); ok {
			return lang.NewPackageDeclStatement(decl)
		}
	}

	importCtx := ctx.ImportDeclaration()
	if importCtx != nil {
		visitor := NewImportDeclarationVisitor()
		if decl, ok := packageCtx.Accept(visitor).(*lang.ImportDecl); ok {
			return lang.NewImportDeclStatement(decl)
		}
	}

	enumCtx := ctx.EnumDeclaration()
	if enumCtx != nil {
		visitor := NewEnumDeclarationVisitor()
		if decl, ok := enumCtx.Accept(visitor).(*lang.EnumDecl); ok {
			return lang.NewEnumDeclStatement(decl)
		}
	}

	structCtx := ctx.StructDeclaration()
	if structCtx != nil {
		visitor := NewStructDeclarationVisitor()
		if decl, ok := structCtx.Accept(visitor).(*lang.StructDecl); ok {
			return lang.NewStructDeclStatement(decl)
		}
	}

	interfaceCtx := ctx.InterfaceDeclaration()
	if interfaceCtx != nil {
		visitor := NewInterfaceDeclarationVisitor()
		if decl, ok := interfaceCtx.Accept(visitor).(*lang.InterfaceDecl); ok {
			return lang.NewInterfaceDeclStatement(decl)
		}
	}

	typeAliasCtx := ctx.TypeAliasDeclaration()
	if typeAliasCtx != nil {
		visitor := NewTypeAliasDeclarationVisitor()
		if decl, ok := typeAliasCtx.Accept(visitor).(*lang.TypeAliasDecl); ok {
			return lang.NewTypeAliasDeclStatement(decl)
		}
	}

	functionCtx := ctx.FunctionDeclaration()
	if functionCtx != nil {
		visitor := NewFuncDeclarationVisitor()
		if decl, ok := functionCtx.Accept(visitor).(*lang.FunctionDecl); ok {
			return lang.NewFunctionDeclStatement(decl)
		}
	}

	//variableDeclaration := ctx.VariableDeclaration()
	//if variableDeclaration != nil {
	//	visitor := NewValueDeclarationVisitor()
	//	if decl, ok := variableDeclaration.Accept(visitor).(*lang.ValueDecl); ok {
	//		return lang.New
	//	}
	//}

	return nil
}

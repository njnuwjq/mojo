package protobuf

import (
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/mojo-lang/lang/go/pkg/lang"
	"github.com/mojo-lang/mojo/go/pkg/compiler/protobuf/genproto"
	"github.com/mojo-lang/mojo/go/pkg/compiler/protobuf/plugin"
	"strings"
)

type Compiler struct {
	Context *plugin.Context
	Files   []*genproto.FileDescriptor
}

func New() *Compiler {
	compiler := &Compiler{}
	compiler.Context = &plugin.Context{}
	compiler.Files = make([]*genproto.FileDescriptor, 0)
	return compiler
}

func (c *Compiler) CompilePackages(packages map[string]*lang.Package) error {
	for _, pkg := range packages {
		err := c.CompilePackage(pkg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Compiler) CompilePackage(pkg *lang.Package) error {
	for _, sourceFile := range pkg.SourceFiles {
		err := c.CompileFile(sourceFile)
		if err != nil {
			return err
		}
		c.Context.File.Package = &pkg.FullName
		c.Files = append(c.Files, c.Context.File)
	}
	return nil
}

func (c *Compiler) CompileFile(file *lang.SourceFile) error {
	if c.Context != nil {
		c.Context.Clear()
		c.Context.File = genproto.NewFileDescriptor()
	} else {
		c.Context = &plugin.Context{
			File: genproto.NewFileDescriptor(),
		}
	}

	c.Context.File.Proto3 = true

	name := strings.TrimSuffix(file.FullName, ".mojo") + ".proto"
	c.Context.File.Name = &name

	for _, statement := range file.Statements {
		switch statement.Statement.(type) {
		case *lang.Statement_Declaration:
			decl := statement.GetDeclaration()
			if decl == nil {
				return errors.New("declaration statement is nil")
			}

			switch decl.Declaration.(type) {
			case *lang.Declaration_StructDecl:
				err := c.compileStruct(&plugin.Context{
					File:    c.Context.File,
					Message: genproto.NewMessageDescriptor(c.Context.File),
				}, decl.GetStructDecl())
				if err != nil {
					return err
				}
			case *lang.Declaration_EnumDecl:
				err := c.compileEnum(&plugin.Context{
					File: c.Context.File,
					Enum: genproto.NewEnumDescriptor(c.Context.File),
				}, decl.GetEnumDecl())
				if err != nil {
					return err
				}
			case *lang.Declaration_InterfaceDecl:
				err := c.compileInterface(&plugin.Context{
					File:    c.Context.File,
					Service: genproto.NewServiceDescriptor(c.Context.File),
				}, decl.GetInterfaceDecl())
				if err != nil {
					return err
				}
			case *lang.Declaration_PackageDecl:
				err := c.compilePackageDecl(c.Context, decl.GetPackageDecl())
				if err != nil {
					return err
				}
			case *lang.Declaration_ImportDecl:
				err := c.compileImport(c.Context, decl.GetImportDecl())
				if err != nil {
					return err
				}
			}
		case *lang.Statement_Expression:
		default:
		}
	}
	return nil
}

func (c *Compiler) compilePackageDecl(ctx *plugin.Context, decl *lang.PackageDecl) error {
	return nil
}

func (c *Compiler) compileImport(ctx *plugin.Context, decl *lang.ImportDecl) error {
	return nil
}

func (c *Compiler) compileEnum(ctx *plugin.Context, decl *lang.EnumDecl) error {
	err := plugin.CompileEnum(ctx, decl)
	if err != nil {
		return err
	}

	ctx.File.Enums = append(ctx.File.Enums, ctx.Enum)
	return nil
}

func (c *Compiler) compileStruct(ctx *plugin.Context, decl *lang.StructDecl) error {
	err := plugin.CompileStruct(ctx, decl)
	if err != nil {
		return err
	}

	ctx.File.Messages = append(ctx.File.Messages, ctx.Message)
	return nil
}

func (c *Compiler) compileInterface(ctx *plugin.Context, i *lang.InterfaceDecl) error {
	ctx.Service.Name = &i.Name

	//if i.Document != nil {
	//	for _, l := range i.Document.Lines {
	//		service.Document = append(service.Document, l.Line)
	//	}
	//}

	for _, method := range i.Type.Methods {
		err := c.compileMethod(ctx, method)
		if err != nil {
			return err
		}
	}

	ctx.File.Services = append(ctx.File.Services, ctx.Service)
	return nil
}

func (c *Compiler) compileMethod(ctx *plugin.Context, method *lang.FunctionDecl) error {
	m := &genproto.MethodDescriptorProto{
		Name:       &method.Name,
		InputType:  nil,
		OutputType: nil,
		Options:    nil,
	}

	// generate the request message
	s := &lang.StructDecl{}
	s.Name = strcase.ToCamel(method.Name) + "Request"

	// add number attribute if there is no field has number attribute
	//for i, param := range method.Type.Parameters {
	//	param.Attributes
	//}

	s.Type = &lang.StructType{
		Fields: method.Signature.Parameters,
	}

	context := plugin.Context{
		File:    ctx.File,
		Message: genproto.NewMessageDescriptor(ctx.File),
	}
	err := plugin.CompileStruct(&context, s)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to compile the request in %s: %s", method.Name, err.Error()))
	}
	ctx.File.Messages = append(ctx.File.Messages, context.Message)

	m.InputType = &s.Name

	var nullTypeName = "mojo.core.Null"
	result := method.Signature.Result
	if result == nil {
		m.OutputType = &nullTypeName
	} else {
		if result.Name == "Tuple" {
			name := strcase.ToCamel(method.Name) + "Response"
			result.Attributes = lang.SetStringAttribute(result.Attributes, "alias", name)
		}

		_, name, err := plugin.CompileNominalType(ctx, result)
		if err != nil {
			return errors.New(
				fmt.Sprintf("failed to compile the type %s: %s", result.Name, err.Error()))
		}

		//TODO the type should be the struct
		m.OutputType = &name
	}

	// options

	ctx.Service.Method = append(ctx.Service.Method, m)
	return nil
}

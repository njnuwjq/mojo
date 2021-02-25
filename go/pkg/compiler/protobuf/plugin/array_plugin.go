package plugin

import (
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/mojo-lang/lang/go/pkg/lang"
	"github.com/mojo-lang/mojo/go/pkg/compiler/protobuf/genproto"
)

type ArrayPlugin struct {
}

func init() {
	p := plugins["Array"]
	if p == nil {
		p = make([]Plugin, 0)
	}
	p = append(p, &ArrayPlugin{})

	plugins["Array"] = p
}

func (p *ArrayPlugin) Compile(ctx *Context, t *lang.NominalType) (string, string, error) {
	if t.Name != "Array" {
		return "", "", errors.New("")
	}

	if len(t.GenericArguments) != 1 {
		return "", "", errors.New("")
	}

	t.Attributes = lang.SetIntegerAttribute(t.Attributes, "number", 1)

	s := &lang.StructDecl{}
	s.Type = &lang.StructType{
		Fields: []*lang.ValueDecl{{
			Name: "values",
			Type: t,
		}},
	}

	val := t.GenericArguments[0]
	s.Name = strcase.ToCamel(val.Name) + "Array"

	context := Context{
		File:    ctx.File,
		Message: genproto.NewMessageDescriptor(ctx.File),
	}
	err := CompileStruct(&context, s)
	if err != nil {
		return "", "", errors.New(fmt.Sprintf("failed to compile the map field entry in %s.%s: %s",
			ctx.Message.Name,
			ctx.FieldName,
			err.Error()))
	}

	//ctx.File.Message.AddInnerMessage(context.Message)
	return "struct", s.Name, nil
}
package plugin

import (
	"errors"
	"fmt"
	"github.com/mojo-lang/lang/go/pkg/lang"
	"github.com/mojo-lang/mojo/go/pkg/compiler/protobuf/genproto"
)

type TuplePlugin struct {
}

func init() {
	p := plugins["Tuple"]
	if p == nil {
		p = make([]Plugin, 0)
	}
	p = append(p, &TuplePlugin{})

	plugins["Tuple"] = p
}

func (p *TuplePlugin) Compile(ctx *Context, t *lang.NominalType) (string, string, error) {
	if t.Name != "Tuple" {
		return "", "", errors.New(fmt.Sprintf("type invalid, need Tuple, but is %s", t.Name))
	}

	if len(t.GenericArguments) == 1 {
		argument := t.GenericArguments[0]
		_, err := lang.GetStringAttribute(argument.Attributes, "label")
		if err != nil { // has no label
			argument.Attributes = append(argument.Attributes, t.Attributes...)
			return CompileNominalType(ctx, t.GenericArguments[0])
		}
	}

	// generate the message
	s := &lang.StructDecl{}
	alias, err := lang.GetStringAttribute(t.Attributes, "alias")
	if err == nil {
		s.Name = alias
	}
	s.Type = &lang.StructType{}
	for _, argument := range t.GenericArguments {
		label, err := lang.GetStringAttribute(argument.Attributes, "label")
		if err != nil {
			return "", "", errors.New(fmt.Sprintf("failed to get the label attribute from %s", argument.Name))
		}

		s.Type.Fields = append(s.Type.Fields, &lang.ValueDecl{
			Name: label,
			Type: argument,
		})
	}

	context := Context{
		File:    ctx.File,
		Message: genproto.NewMessageDescriptor(ctx.File),
	}
	err = CompileStruct(&context, s)
	if err != nil {
		return "", "", errors.New(fmt.Sprintf("failed to compile struct: %s", err.Error()))
	}

	ctx.File.Messages = append(ctx.File.Messages, context.Message)
	return "struct", s.Name, nil
}

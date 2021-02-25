package plugin

import (
	"errors"
	"fmt"
	"github.com/mojo-lang/lang/go/pkg/lang"
	"github.com/mojo-lang/mojo/go/pkg/compiler/protobuf/genproto"
)

type UnionPlugin struct {
}

func init() {
	p := plugins["Union"]
	if p == nil {
		p = make([]Plugin, 0)
	}
	p = append(p, &UnionPlugin{})

	plugins["Union"] = p
}

func (p *UnionPlugin) Compile(ctx *Context, t *lang.NominalType) (string, string, error) {
	_, err := lang.GetStringAttribute(t.Attributes, "type_alias_original_name")
	if err != nil { // directly union declaration

	} else {
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
				// auto generate
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

	return "", "", nil
}

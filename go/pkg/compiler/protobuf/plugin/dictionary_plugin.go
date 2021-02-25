package plugin

import (
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/mojo-lang/lang/go/pkg/lang"
	"github.com/mojo-lang/mojo/go/pkg/compiler/protobuf/genproto"
)

type DictionaryPlugin struct {
}

func init() {
	p := plugins["Dictionary"]
	if p == nil {
		p = make([]Plugin, 0)
	}
	p = append(p, &DictionaryPlugin{})

	plugins["Dictionary"] = p
}

func (p *DictionaryPlugin) Compile(ctx *Context, t *lang.NominalType) (string, string, error) {
	if t.Name != "Dictionary" {
		return "", "", errors.New("")
	}

	if len(t.GenericArguments) != 2 {
		return "", "", errors.New("")
	}

	s := &lang.StructDecl{}
	s.Name = strcase.ToCamel(ctx.FieldName) + "Entry"

	keyType := t.GenericArguments[0]
	keyType.Attributes = lang.SetIntegerAttribute(keyType.Attributes, "number", 1)

	valType := t.GenericArguments[1]
	valType.Attributes = lang.SetIntegerAttribute(valType.Attributes, "number", 2)

	s.Type = &lang.StructType{
		Fields: []*lang.ValueDecl{{
			Name: "key",
			Type: keyType,
		}, {
			Name: "value",
			Type: valType,
		}},
	}

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

	mapEntry := true
	context.Message.Options = &genproto.MessageOptions{
		MapEntry: &mapEntry,
	}

	ctx.Message.AddInnerMessage(context.Message)
	return "struct", s.Name, nil
}

package plugin

import (
	"errors"
	"fmt"
	"github.com/mojo-lang/lang/go/pkg/lang"
	"github.com/mojo-lang/mojo/go/pkg/compiler/protobuf/genproto"
)

type Plugin interface {
	Compile(ctx *Context, t *lang.NominalType) (string, string, error)
}

var plugins = make(map[string][]Plugin)

// type alias transform
// type: Scalar, Enum, Struct
func CompileNominalType(ctx *Context, t *lang.NominalType) (string, string, error) { // type, typeName, error
	switch t.Name {
	case "Bool", "Int32", "UInt32", "Int", "UInt", "Int64", "UInt64",
		 "Float", "Float32", "Double", "Float64",
		 "String", "Bytes":
		return "Scalar", t.Name, nil
	default:
		tp := "Struct"

		if decl := t.TypeDeclaration; decl != nil {
			if aliasDecl := decl.GetTypeAliasDecl(); aliasDecl != nil {
				if aliasDecl.Type != nil {
					// type_alias_original_name
					_, err := lang.GetStringAttribute(aliasDecl.Type.Attributes, "type_alias_original_name")
					if err != nil { // not found
						lang.SetStringAttribute(aliasDecl.Type.Attributes, "type_alias_original_name", t.Name)
					}

					return CompileNominalType(ctx, aliasDecl.Type)
				}
			}

			if decl.GetEnumDecl() != nil {
				tp = "Enum"
			}
		}

		ps := plugins[t.Name]
		for _, p := range ps {
			tp, typeName, err := p.Compile(ctx, t)
			if err != nil && len(typeName) > 0 {
				return tp, typeName, err
			}
		}

		return tp, t.Name, nil
	}

	return "Struct", t.Name, nil
}

func CompileEnum(ctx *Context, decl *lang.EnumDecl) error {
	ctx.Enum.Name = &decl.Name

	for i, e := range decl.Type.Enumerators {
		value := &genproto.EnumValueDescriptorProto{
			Name: &e.Name,
			//Number:               nil,
			//Options:              nil,
		}

		number, err := lang.GetIntegerAttribute(e.Attributes, "number")
		if err == nil {
			if number <= 0 {
				return errors.New("number attribute value must be positive")
			}
			n := int32(number)
			value.Number = &n
		} else {
			n := int32(i)
			value.Number = &n
		}

		ctx.Enum.Value = append(ctx.Enum.Value, value)
	}
	return nil
}

func CompileStruct(ctx *Context, decl *lang.StructDecl) error {
	ctx.Message.Name = &decl.Name

	for _, e := range decl.EnumTypeDecls {
		context := &Context{
			Parent: ctx,
			File:   ctx.File,
			Enum:   genproto.NewEnumDescriptor(ctx.File),
		}
		err := CompileEnum(context, e)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to parse the inner enum decl %s in %s: %s", e.Name, decl.Name, err.Error()))
		}

		ctx.Message.AddInnerEnum(context.Enum)
	}

	for _, s := range decl.StructTypeDecls {
		context := &Context{
			Parent:  ctx,
			File:    ctx.File,
			Message: genproto.NewMessageDescriptor(ctx.File),
		}
		err := CompileStruct(context, s)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to parse the inner struct decl %s in %s: %s", s.Name, decl.Name, err.Error()))
		}

		ctx.Message.AddInnerMessage(context.Message)
	}

	for _, field := range decl.Type.Fields {
		member := &genproto.FieldDescriptorProto{}
		member.Name = &field.Name

		switch field.Type.Name {
		case "Array":
			if len(field.Type.GenericArguments) > 0 {
				argument := field.Type.GenericArguments[0]
				repeated := genproto.FieldDescriptorProto_LABEL_REPEATED
				member.Label = &repeated
				t, name, err := CompileNominalType(ctx, argument)
				if err != nil {
					return errors.New(
						fmt.Sprintf("failed to compile the type %s: %s", argument.Name, err.Error()))
				}
				pType, pName := protoType(t, name)
				member.Type = &pType
				member.TypeName = &pName
			} else {
				return errors.New(fmt.Sprintf("unexpect array type of %s", field.Type.Name))
			}
		default:
			t, name, err := CompileNominalType(ctx, field.Type)
			if err != nil {
				return errors.New(
					fmt.Sprintf("failed to compile the type %s: %s", field.Type.Name, err.Error()))
			}

			pType, pName := protoType(t, name)
			member.Type = &pType
			member.TypeName = &pName
		}

		number, err := lang.GetIntegerAttribute(field.Type.Attributes, "number")
		if err != nil {
			return errors.New("has not set the number in field")
		}
		if number <= 0 {
			return errors.New("number attribute value must be positive")
		}
		n := int32(number)
		member.Number = &n

		ctx.Message.Field = append(ctx.Message.Field, member)
	}
	return nil
}

func protoType(t string, typeName string) (genproto.FieldDescriptorProto_Type, string) {
	switch typeName {
	case "Double", "Float64":
		return genproto.FieldDescriptorProto_TYPE_DOUBLE, ""
	case "Float", "Float32":
		return genproto.FieldDescriptorProto_TYPE_FLOAT, ""
	case "Int64", "Int":
		return genproto.FieldDescriptorProto_TYPE_INT64, ""
	case "UInt64", "UInt":
		return genproto.FieldDescriptorProto_TYPE_UINT64, ""
	case "Int32":
		return genproto.FieldDescriptorProto_TYPE_INT32, ""
	case "Bool":
		return genproto.FieldDescriptorProto_TYPE_BOOL, ""
	case "String":
		return genproto.FieldDescriptorProto_TYPE_STRING, ""
	case "Bytes":
		return genproto.FieldDescriptorProto_TYPE_BYTES, ""
	default:
		switch t {
		case "Enum":
			return genproto.FieldDescriptorProto_TYPE_ENUM, typeName
		default:
			return genproto.FieldDescriptorProto_TYPE_MESSAGE, typeName
		}
	}
}

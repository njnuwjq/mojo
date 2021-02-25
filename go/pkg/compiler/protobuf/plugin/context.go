package plugin

import (
	"context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/protobuf/genproto"
)

type Context struct {
	context.Context

	Parent *Context

	// the current file for parsing
	File *genproto.FileDescriptor

	// the current enum for parsing
	Enum *genproto.EnumDescriptor

	// the current message for parsing
	Message *genproto.MessageDescriptor

	FieldName string

	// the current service for parsing
	Service *genproto.ServiceDescriptor
}

func (c *Context) Clear()  {
	c.Parent = nil
	c.File = nil
	c.Enum = nil
	c.Message = nil
	c.Service = nil

	c.FieldName = ""
}

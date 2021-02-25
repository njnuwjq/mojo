package genproto

import (
	"strconv"
	"strings"
)

// Each type we import as a protocol buffer (other than FileDescriptorProto) needs
// a pointer to the FileDescriptorProto that represents it.  These types achieve that
// wrapping by placing each Proto inside a struct with the pointer to its File. The
// structs have the same names as their contents, with "Proto" removed.
// FileDescriptor is used to store the things that it points to.

// The File and package name method are Common to Messages and Enums.
type Common struct {
	File     *FileDescriptor // File this object comes from.
	Comments *SourceCodeInfo_Location
}

func (c *Common) LeadingComments() *string {
	if c.Comments != nil {
		return c.Comments.LeadingComments
	}
	return nil
}
func (c *Common) TrailingComments() *string {
	if c.Comments != nil {
		return c.Comments.TrailingComments
	}
	return nil
}

// MessageDescriptor represents a protocol buffer message.
type MessageDescriptor struct {
	Common
	*DescriptorProto

	Parent   *MessageDescriptor   // The containing message, if any.
	Messages []*MessageDescriptor // Inner Messages, if any.
	Enums    []*EnumDescriptor    // Inner Enums, if any.

	Index int    // The Index into the container, whether the File or another message.
	Path  string // The SourceCodeInfo Path as comma-separated integers.
}

func (m *MessageDescriptor) GetInnerMessage(name string) *MessageDescriptor {
	for _, msg := range m.Messages {
		if msg.Name != nil && *msg.Name == name {
			return msg
		}
	}
	return nil
}

func (m *MessageDescriptor) AddInnerMessage(msg *MessageDescriptor) {
	m.Messages = append(m.Messages, msg)
	m.NestedType = append(m.NestedType, msg.DescriptorProto)
}

func (m *MessageDescriptor) AddInnerEnum(enum *EnumDescriptor) {
	m.Enums = append(m.Enums, enum)
	m.EnumType = append(m.EnumType, enum.EnumDescriptorProto)
}

// EnumDescriptor describes an enum. If it's at top level, its Parent will be nil.
// Otherwise it will be the descriptor of the message in which it is defined.
type EnumDescriptor struct {
	Common
	*EnumDescriptorProto

	Parent *MessageDescriptor // The containing message, if any.
	Index  int                // The Index into the container, whether the File or a message.
	Path   string             // The SourceCodeInfo Path as comma-separated integers.
}

// ServiceDescriptor describes an service.
type ServiceDescriptor struct {
	Common
	*ServiceDescriptorProto

	Index int // The Index into the container, whether the File or a message.
}

// FileDescriptor describes an protocol buffer descriptor File (.proto).
// It includes slices of all the Messages and Enums defined within it.
// Those slices are constructed by WrapTypes.
type FileDescriptor struct {
	*FileDescriptorProto

	Messages []*MessageDescriptor // All the Messages defined in this File.
	Enums    []*EnumDescriptor    // All the Enums defined in this File.
	Services []*ServiceDescriptor // All the top-level extensions defined in this File.

	// Comments, stored as a map of Path (comma-separated integers) to the comment.
	Comments map[string]*SourceCodeInfo_Location

	Proto3 bool // whether to generate proto3 code for this File
}

func NewFileDescriptor() *FileDescriptor {
	return &FileDescriptor{
		FileDescriptorProto: &FileDescriptorProto{},

		Messages: make([]*MessageDescriptor, 0),
	}
}

func NewMessageDescriptor(file *FileDescriptor) *MessageDescriptor {
	return &MessageDescriptor{
		Common: Common{
			File: file,
		},
		DescriptorProto: &DescriptorProto {
		},
	}
}

func NewEnumDescriptor(file *FileDescriptor) *EnumDescriptor {
	return &EnumDescriptor{
		Common: Common{
			File: file,
		},
		EnumDescriptorProto: &EnumDescriptorProto {},
	}
}

func NewServiceDescriptor(file *FileDescriptor) *ServiceDescriptor {
	return &ServiceDescriptor{
		Common: Common{
			File: file,
		},
		ServiceDescriptorProto: &ServiceDescriptorProto {
		},
	}
}

func FileIsProto3(file *FileDescriptorProto) bool {
	return file.GetSyntax() == "Proto3"
}

func ExtractComments(file *FileDescriptor) {
	file.Comments = make(map[string]*SourceCodeInfo_Location)
	for _, loc := range file.GetSourceCodeInfo().GetLocation() {
		if loc.LeadingComments == nil {
			continue
		}
		var p []string
		for _, n := range loc.Path {
			p = append(p, strconv.Itoa(int(n)))
		}
		file.Comments[strings.Join(p, ",")] = loc
	}
}

// Return a slice of all the Descriptors defined within this File
func WrapMessageDescriptors(file *FileDescriptor) []*MessageDescriptor {
	sl := make([]*MessageDescriptor, 0, len(file.MessageType)+10)
	for i, desc := range file.MessageType {
		sl = wrapThisMessageDescriptor(sl, desc, nil, file, i)
	}
	return sl
}

// Wrap this Descriptor, recursively
func wrapThisMessageDescriptor(sl []*MessageDescriptor, desc *DescriptorProto, parent *MessageDescriptor, file *FileDescriptor, index int) []*MessageDescriptor {
	sl = append(sl, newMessageDescriptor(desc, parent, file, index))
	me := sl[len(sl)-1]
	for i, nested := range desc.NestedType {
		sl = wrapThisMessageDescriptor(sl, nested, me, file, i)
	}
	return sl
}

// Return a slice of all the EnumDescriptors defined within this File
func WrapEnumDescriptors(file *FileDescriptor, descs []*MessageDescriptor) []*EnumDescriptor {
	sl := make([]*EnumDescriptor, 0, len(file.EnumType)+10)
	// Top-level Enums.
	for i, enum := range file.EnumType {
		sl = append(sl, newEnumDescriptor(enum, nil, file, i))
	}
	// Enums within Messages. Enums within embedded Messages appear in the outer-most message.
	for _, nested := range descs {
		for i, enum := range nested.EnumType {
			sl = append(sl, newEnumDescriptor(enum, nested, file, i))
		}
	}
	return sl
}

// Return a slice of all the EnumDescriptors defined within this File
func WrapServiceDescriptors(file *FileDescriptor) []*ServiceDescriptor {
	sl := make([]*ServiceDescriptor, 0, len(file.Service)+10)
	// Top-level Enums.
	for i, service := range file.Service {
		sl = append(sl, newServiceDescriptor(service, nil, file, i))
	}
	return sl
}

// Construct the MessageDescriptor
func newMessageDescriptor(desc *DescriptorProto, parent *MessageDescriptor, file *FileDescriptor, index int) *MessageDescriptor {
	d := &MessageDescriptor{
		Common:          Common{file, nil},
		DescriptorProto: desc,
		Parent:          parent,
		Index:           index,
	}
	if parent == nil {
		//d.Path = fmt.Sprintf("%d,%d", messagePath, Index)
	} else {
		//d.Path = fmt.Sprintf("%s,%d,%d", Parent.Path, messageMessagePath, Index)
	}

	// The only way to distinguish a group from a message is whether
	// the containing message has a TYPE_GROUP field that matches.
	//if Parent != nil {
	//	parts := d.TypeName()
	//	if File.Package != nil {
	//		parts = append([]string{*File.Package}, parts...)
	//	}
	//	exp := "." + strings.Join(parts, ".")
	//	for _, field := range Parent.Field {
	//		if field.GetType() == FieldDescriptorProto_TYPE_GROUP && field.GetTypeName() == exp {
	//			d.group = true
	//			break
	//		}
	//	}
	//}

	return d
}

// Construct the EnumDescriptor
func newEnumDescriptor(desc *EnumDescriptorProto, parent *MessageDescriptor, file *FileDescriptor, index int) *EnumDescriptor {
	ed := &EnumDescriptor{
		Common:              Common{file, nil},
		EnumDescriptorProto: desc,
		Parent:              parent,
		Index:               index,
	}
	if parent == nil {
		//ed.Path = fmt.Sprintf("%d,%d", enumPath, Index)
	} else {
		//ed.Path = fmt.Sprintf("%s,%d,%d", Parent.Path, messageEnumPath, Index)
	}
	return ed
}

// Construct the ServiceDescriptor
func newServiceDescriptor(desc *ServiceDescriptorProto, parent *MessageDescriptor, file *FileDescriptor, index int) *ServiceDescriptor {
	sd := &ServiceDescriptor{
		Common:                 Common{file, nil},
		ServiceDescriptorProto: desc,
		Index:                  index,
	}
	if parent == nil {
		//ed.Path = fmt.Sprintf("%d,%d", enumPath, Index)
	} else {
		//ed.Path = fmt.Sprintf("%s,%d,%d", Parent.Path, messageEnumPath, Index)
	}
	return sd
}

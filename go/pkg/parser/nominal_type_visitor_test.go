package parser

import (
	"github.com/mojo-lang/lang/go/pkg/lang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNominalTypeVisitor_VisitArrayType(t *testing.T) {
	const arrayType = `type Val{ val: [Int] }`

	parser := &Parser{}
	file, err := parser.Parse(arrayType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Array", nominalType.Name)
	assert.Equal(t, 1, len(nominalType.GenericArguments))
	assert.Equal(t, "Int", nominalType.GenericArguments[0].Name)
}

func TestNominalTypeVisitor_VisitDictionaryType(t *testing.T) {
	const dictType = `type Val{ val: {String: Int} }`

	parser := &Parser{}
	file, err := parser.Parse(dictType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Dictionary", nominalType.Name)
	assert.Equal(t, 2, len(nominalType.GenericArguments))
	assert.Equal(t, "String", nominalType.GenericArguments[0].Name)
	assert.Equal(t, "Int", nominalType.GenericArguments[1].Name)
}

func TestNominalTypeVisitor_VisitUnion(t *testing.T) {
	const unionType = `type Val{ val: String | Int }`

	parser := &Parser{}
	file, err := parser.Parse(unionType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Union", nominalType.Name)
	assert.Equal(t, 2, len(nominalType.GenericArguments))
	assert.Equal(t, "String", nominalType.GenericArguments[0].Name)
	assert.Equal(t, "Int", nominalType.GenericArguments[1].Name)
}

func TestNominalTypeVisitor_VisitTupleType(t *testing.T) {
	const tupleType = `type Val{ val: (String, Int) }`

	parser := &Parser{}
	file, err := parser.Parse(tupleType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Tuple", nominalType.Name)
	assert.Equal(t, 2, len(nominalType.GenericArguments))
	assert.Equal(t, "String", nominalType.GenericArguments[0].Name)
	assert.Equal(t, "Int", nominalType.GenericArguments[1].Name)
}

func TestNominalTypeVisitor_VisitTupleTypeWithLabel(t *testing.T) {
	const tupleType = `type Val{ val: (str: String, integer: Int) }`

	parser := &Parser{}
	file, err := parser.Parse(tupleType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Tuple", nominalType.Name)
	assert.Equal(t, 2, len(nominalType.GenericArguments))
	assert.Equal(t, "String", nominalType.GenericArguments[0].Name)
	assert.Equal(t, "Int", nominalType.GenericArguments[1].Name)

	label, _ := lang.GetStringAttribute(nominalType.GenericArguments[0].Attributes, "label")
	assert.Equal(t, "str", label)

	label, _ = lang.GetStringAttribute(nominalType.GenericArguments[1].Attributes, "label")
	assert.Equal(t, "integer", label)
}

func getNominalType(file *lang.SourceFile) *lang.NominalType {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			if structDecl := decl.GetStructDecl(); structDecl != nil {
				if t := structDecl.Type; t != nil {
					if len(t.Fields) > 0 {
						if field := t.Fields[0]; field != nil {
							return field.Type
						}
					}
				}
			}
		}
	}
	return nil
}

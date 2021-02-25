package parser

import (
	"github.com/mojo-lang/lang/go/pkg/lang"
	"strings"
)

type DocumentVisitor struct {
	*BaseMojoParserVisitor
}

func NewDocumentVisitor() *DocumentVisitor {
	return &DocumentVisitor{}
}

func GetDocument(ctx IDocumentContext) *lang.Document {
	if ctx != nil {
		visitor := NewDocumentVisitor()
		return ctx.Accept(visitor).(*lang.Document)
	}
	return nil
}

func GetFollowingDocument(ctx IFollowingDocumentContext) *lang.Document {
	if ctx != nil {
		visitor := NewDocumentVisitor()
		return ctx.Accept(visitor).(*lang.Document)
	}
	return nil
}

func (d *DocumentVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	lines := ctx.AllLineDocument()
	document := &lang.Document{}
	for _, line := range lines {
		lineDocument := &lang.LineDocument{}

		lineDocument.StartPosition = GetPosition(line.GetSymbol())
		lineDocument.EndPosition = &lang.Position{
			Line:   lineDocument.StartPosition.Line,
			Column: lineDocument.StartPosition.Column + int32(len(line.GetText())),
		}

		lineDocument.Line = strings.TrimPrefix(line.GetText(), "///")

		document.Lines = append(document.Lines, lineDocument)
	}

	if len(lines) > 0 {
		document.StartPosition = document.Lines[0].StartPosition
		document.EndPosition = document.Lines[len(document.Lines) - 1].EndPosition
	}

	return document
}

func (d *DocumentVisitor) VisitFollowingDocument(ctx *FollowingDocumentContext) interface{} {
	document := &lang.Document{}

	line := &lang.LineDocument{}
	line.StartPosition = GetPosition(ctx.GetStart())
	line.EndPosition = GetPosition(ctx.GetStop())
	line.Line = strings.TrimPrefix(ctx.GetText(), "//<")

	document.Lines = append(document.Lines, line)
	return document
}

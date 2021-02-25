package parser

import (
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mojo-lang/lang/go/pkg/lang"
	"io/ioutil"
	"log"
	path2 "path"
	"sort"
	"strings"
)

type Parser struct {
}

func New() *Parser {
	return &Parser{}
}

func (p Parser) Parse(mojo string) (*lang.SourceFile, error) {
	input := antlr.NewInputStream(mojo)
	return p.ParseStream(input)
}

func (p Parser) ParseStream(input *antlr.InputStream) (*lang.SourceFile, error) {
	lexer := NewMojoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	parser := NewMojoParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true

	tree := parser.MojoFile()
	visitor := NewMojoFileVisitor()
	result := visitor.Visit(tree).(bool)
	if result {
		print(visitor.SourceFile.String())
		return visitor.SourceFile, nil
	} else {
		print("parser failed")
	}

	return nil, errors.New("parse failed")
}

func (p Parser) ParseFile(filename string) (*lang.SourceFile, error) {
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, err
	}

	sourceFile, err := p.ParseStream(input.InputStream)
	if err != nil {
		return nil, err
	}

	sourceFile.Name = path2.Base(filename)
	sourceFile.FullName = reassemblePath(filename, "/")
	return sourceFile, nil
}

func (p Parser) ParsePackage(path string) (map[string]*lang.Package, error) {
	packages, error := p.parsePackage(path)
	if error != nil {
		return nil, error
	}

	p.treePackages(packages)
	return packages, nil
}

func (p Parser) treePackages(packages map[string]*lang.Package) {
	var names []string
	for key, _ := range packages {
		names = append(names, key)
	}
	sort.Strings(names)

	for i, name := range names {
		for j := i; j < len(names); j++ {
			if strings.HasPrefix(names[j], name) && names[j] != name {
				parent := packages[name]
				child := packages[names[j]]
				parent.Children = append(parent.Children, child)
				child.Parent = parent
			}
		}
	}
}

func (p Parser) parsePackage(path string) (map[string]*lang.Package, error) {
	packages := make(map[string]*lang.Package)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	currentPkgName := reassemblePath(path, ".")
	for _, f := range files {
		if f.IsDir() {
			pkgs, err := p.parsePackage(path2.Join(path, f.Name()))
			if err != nil {
				return nil, err
			}

			for k, v := range pkgs {
				packages[k] = v
			}
		} else {
			sourceFile, err := p.ParseFile(path2.Join(path, f.Name()))
			if err != nil {
				return nil, err
			}

			if packages[currentPkgName] == nil {
				packages[currentPkgName] = &lang.Package{}
			}

			sourceFile.Package = currentPkgName
			packages[currentPkgName].SourceFiles = append(packages[currentPkgName].SourceFiles, sourceFile)
			packages[currentPkgName].Name = path2.Base(path)
			packages[currentPkgName].FullName = currentPkgName
		}
	}
	return packages, nil
}

func reassemblePath(path string, connector string) string {
	if len(path) > 0 {
		dir := path2.Dir(path)
		segments := []string{ path2.Base(path) }

		for {
			base := path2.Base(dir)
			if len(base) > 0 && base != "." && base != ".." {
				dir = path2.Dir(dir)
				segments = append(segments, base)
			} else {
				break
			}
		}

		for i, j := 0, len(segments)-1; i < j; i, j = i+1, j-1 {
			segments[i], segments[j] = segments[j], segments[i]
		}
		return strings.Join(segments, connector)
	}

	return ""
}

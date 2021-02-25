package main

import (
	"fmt"
	"github.com/mojo-lang/mojo/go/pkg/compiler/protobuf"
	"github.com/mojo-lang/mojo/go/pkg/compiler/protobuf/genproto"
	"github.com/mojo-lang/mojo/go/pkg/parser"
	flag "github.com/spf13/pflag"
	"os"
	"path/filepath"
	"strings"
)

// parse command
// mojo parse --file --package     --> out ast

// compile command
// mojo compile --file --package --type=protobuf --output=./

var (
	parseCmd = flag.NewFlagSet("parse", flag.ExitOnError)
	compileCmd = flag.NewFlagSet("compile", flag.ExitOnError)

	fileNameInParseCmd = parseCmd.StringP("file", "f", "", "the mojo file to parse")
	packageInParseCmd = parseCmd.StringP("package", "p", "", "the mojo package folder to parse")

	fileNameInCompileCmd = compileCmd.StringP("file", "f", "", "the mojo file to compile")
	packageInCompileCmd = compileCmd.StringP("package", "p", "", "the mojo package folder to compile")
	typeInCompileCmd = compileCmd.StringP("type", "t", "", "the type of mojo compiler")
	pathInCompileCmd = compileCmd.StringP("output", "o", "", "the output directory of the compiler")

	verboseFlag  = flag.BoolP("verbose", "v", false, "Verbose output")
	helpFlag     = flag.BoolP("help", "h", false, "Print usage")
)

var binName = filepath.Base(os.Args[0])
var workplace = ""

var (
	// Version is compiled into truss with the flag
	// go install -ldflags "-X main.Version=$SHA"
	Version string
	// BuildDate is compiled into truss with the flag
	// go install -ldflags "-X main.VersionDate=$VERSION_DATE"
	VersionDate string
)

func init() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	workplace = filepath.Dir(ex)
	fmt.Println(workplace)

	// If Version or VersionDate are not set, truss was not built with make
	if Version == "" || VersionDate == "" {
		//rebuild := promptNoMake()
		//if !rebuild {
		//	os.Exit(1)
		//}
		//err := makeAndRunTruss(os.Args)
		//exitIfError(errors.Wrap(err, "please install truss with make manually"))
		//os.Exit(0)
	}

	var buildInfo string
	buildInfo = fmt.Sprintf("version: %s", Version)
	buildInfo = fmt.Sprintf("%s version date: %s", buildInfo, VersionDate)

	flag.Usage = func() {
		if buildInfo != "" && (*verboseFlag || *helpFlag) {
			fmt.Fprintf(os.Stderr, "%s (%s)\n", binName, strings.TrimSpace(buildInfo))
		}
		fmt.Fprintf(os.Stderr, "\nUsage: %s [options] <mojo file>...\n", binName)
		fmt.Fprintf(os.Stderr, "\nGenerates protobuffer files using mojo.\n")
		fmt.Fprintln(os.Stderr, "\nOptions:")
		flag.PrintDefaults()
	}

	parseCmd.Usage = func() {
		if buildInfo != "" && (*verboseFlag || *helpFlag) {
			fmt.Fprintf(os.Stderr, "%s (%s)\n", binName, strings.TrimSpace(buildInfo))
		}
		fmt.Fprintf(os.Stderr, "\nUsage: %s parse [options] <mojo file>...\n", binName)
		fmt.Fprintf(os.Stderr, "\nGenerates protobuffer files using mojo.\n")
		fmt.Fprintln(os.Stderr, "\nOptions:")
		flag.PrintDefaults()
	}

	compileCmd.Usage = func() {
		if buildInfo != "" && (*verboseFlag || *helpFlag) {
			fmt.Fprintf(os.Stderr, "%s (%s)\n", binName, strings.TrimSpace(buildInfo))
		}
		fmt.Fprintf(os.Stderr, "\nUsage: %s compile [options] <mojo file>...\n", binName)
		fmt.Fprintf(os.Stderr, "\nGenerates protobuffer files using mojo.\n")
		fmt.Fprintln(os.Stderr, "\nOptions:")
		flag.PrintDefaults()
	}
}

func main() {
	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s: expected 'compile' or 'parse' subcommands\n", binName)
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "compile":
		compileCmd.Parse(os.Args[2:])
		switch *typeInCompileCmd {
		case "protobuf":
			compileProtobuf()
		case "openapi":
			compileOpenApi()
		case "markdown":
			compileMarkdown()
		case "go":
			compileGo()
		}

		//fmt.Println("subcommand 'foo'")
	case "parse":
		parseCmd.Parse(os.Args[2:])
		parse()
		os.Exit(0)
	default:
		fmt.Println("expected 'compile' or 'parse' subcommands")
		os.Exit(1)
	}
}

func parse() {
	parser := parser.New()

	if *fileNameInParseCmd != "" {
		file, err := parser.ParseFile(*fileNameInParseCmd)
		if err != nil {
			fmt.Printf("parser %s failed, because of %s", *fileNameInParseCmd, err.Error())
			os.Exit(1)
		}
		fmt.Print(file)
	} else if *packageInParseCmd != "" { // compile mojo package to protobuffer files
		pkgs, err := parser.ParsePackage(*packageInParseCmd)
		if err != nil {
			fmt.Printf("parser %s failed, because of %s", *packageInParseCmd, err.Error())
			os.Exit(1)
		}
		fmt.Print(pkgs)
	}
}

func compileProtobuf()  {
	parser := parser.New()

	if *fileNameInCompileCmd != "" {
		file, err := parser.ParseFile(*fileNameInCompileCmd)
		if err != nil {
			fmt.Printf("parse %s failed, because of %s", *fileNameInCompileCmd, err.Error())
			os.Exit(1)
		}
		compiler := protobuf.New()
		err = compiler.CompileFile(file)
		if err != nil {
			fmt.Printf("compile %s failed, because of %s", *fileNameInCompileCmd, err.Error())
			os.Exit(1)
		}

		generator := genproto.New(compiler.Files)
		generator.GenerateAllFiles().WriteAllFiles(*pathInCompileCmd)
	} else if *packageInCompileCmd != "" { // compile mojo package to protobuffer files
		pkgs, err := parser.ParsePackage(*packageInCompileCmd)
		if err != nil {
			fmt.Printf("parse %s failed, because of %s", *fileNameInCompileCmd, err.Error())
			os.Exit(1)
		}
		compiler := protobuf.New()
		err = compiler.CompilePackages(pkgs)
		if err != nil {
			fmt.Printf("compile %s failed, because of %s", *fileNameInCompileCmd, err.Error())
			os.Exit(1)
		}
		generator := genproto.New(compiler.Files)
		generator.GenerateAllFiles().WriteAllFiles(*pathInCompileCmd)
	}
}

func compileOpenApi()  {
}

func compileMarkdown()  {
}

func compileGo()  {
}
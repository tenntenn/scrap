package replacer

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/tools/go/ast/astutil"
)

// TODO: add line comment
const newFunc = `(func() *%[1]s {
	o := new(%[1]s)
	runtime.SetFinalizer(o, func(o any) {
		println(o)
	})
	return o
})()`

func Replace(name string) (string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, name, nil, parser.ParseComments)
	if err != nil {
		return "", err
	}

	var rerr error
	newFile := astutil.Apply(file, func(c *astutil.Cursor) bool {
		if rerr != nil {
			return false
		}

		call, _ := c.Node().(*ast.CallExpr)
		if call == nil {
			return true
		}

		fun, _ := call.Fun.(*ast.Ident)
		if fun == nil || fun.Name != "new" { // TODO: type checking
			return true
		}

		if len(call.Args) != 1 {
			return true
		}

		var arg bytes.Buffer
		if err := format.Node(&arg, fset, call.Args[0]); err != nil {
			rerr = err
			return false
		}

		expr, err := parser.ParseExpr(fmt.Sprintf(newFunc, arg.String()))
		if err != nil {
			rerr = err
			return false
		}

		c.Replace(expr)

		return true
	}, nil)

	if rerr != nil {
		return "", rerr
	}

	astutil.AddImport(fset, newFile.(*ast.File), "runtime")

	var src bytes.Buffer
	if err := format.Node(&src, fset, newFile); err != nil {
		return "", err
	}

	dst := filepath.Base(name)
	tmp, err := os.CreateTemp("", "*-"+dst)
	if err != nil {
		return "", err
	}

	if _, err := fmt.Fprintf(tmp, "//line: %s:1\n", name); err != nil {
		return "", err
	}

	if _, err := io.Copy(tmp, &src); err != nil {
		return "", err
	}

	return tmp.Name(), nil
}

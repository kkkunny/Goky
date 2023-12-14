package parse

import (
	"os"
	"path/filepath"
	"reflect"

	"github.com/kkkunny/stl/container/linkedlist"
	stlerror "github.com/kkkunny/stl/error"
	"github.com/samber/lo"

	"github.com/kkkunny/Sim/ast"
	errors "github.com/kkkunny/Sim/error"
	"github.com/kkkunny/Sim/lex"
	"github.com/kkkunny/Sim/reader"
	"github.com/kkkunny/Sim/token"
)

func loopParseWithUtil[T any](self *Parser, sem, end token.Kind, f func() T) (res []T) {
	for self.skipSEM(); !self.nextIs(end); self.skipSEM() {
		res = append(res, f())
		if !self.skipNextIs(sem) {
			break
		}
	}
	self.skipSEM()
	return res
}

func expectAttrIn(attrs []ast.Attr, expectAttr ...ast.Attr) {
	expectAttrTypes := lo.Map(expectAttr, func(item ast.Attr, _ int) reflect.Type {
		return reflect.ValueOf(item).Type()
	})
	for _, attr := range attrs {
		if len(expectAttrTypes) == 0 {
			errors.ThrowUnExpectAttr(attr.Position())
		}
		attrType := reflect.ValueOf(attr).Type()
		for _, expectAttrType := range expectAttrTypes {
			if !attrType.AssignableTo(expectAttrType) {
				errors.ThrowUnExpectAttr(attr.Position())
			}
		}
	}
}

// 语法解析目标文件
func parseFile(path string) (linkedlist.LinkedList[ast.Global], stlerror.Error) {
	_, r, err := reader.NewReaderFromFile(path)
	if err != nil {
		return linkedlist.LinkedList[ast.Global]{}, err
	}
	return New(lex.New(r)).Parse(), nil
}

// 语法解析目标目录
func parseDir(path string) (linkedlist.LinkedList[ast.Global], stlerror.Error) {
	entries, err := stlerror.ErrorWith(os.ReadDir(path))
	if err != nil {
		return linkedlist.LinkedList[ast.Global]{}, err
	}
	var asts linkedlist.LinkedList[ast.Global]
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".sim" {
			continue
		}
		fileAst, err := parseFile(filepath.Join(path, entry.Name()))
		if err != nil {
			return linkedlist.LinkedList[ast.Global]{}, err
		}
		asts.Append(fileAst)
	}
	return asts, nil
}

// Parse 语法解析
func Parse(path string) (linkedlist.LinkedList[ast.Global], stlerror.Error) {
	fs, err := stlerror.ErrorWith(os.Stat(path))
	if err != nil{
		return linkedlist.LinkedList[ast.Global]{}, err
	}
	if fs.IsDir(){
		return parseDir(path)
	}else{
		return parseFile(path)
	}
}

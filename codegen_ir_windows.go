//go:build codegenir && windows

package main

import (
	"fmt"
	"os"

	"github.com/kkkunny/go-llvm"
	stlerror "github.com/kkkunny/stl/error"
	stlos "github.com/kkkunny/stl/os"

	"github.com/kkkunny/Sim/compiler/codegen_ir"
)

func main() {
	llvm.EnablePrettyStackTrace()
	stlerror.Must(llvm.InitializeTargetInfo(llvm.X86))
	stlerror.Must(llvm.InitializeTarget(llvm.X86))
	stlerror.Must(llvm.InitializeTargetMC(llvm.X86))
	target := stlerror.MustWith(llvm.NewTargetFromTriple("x86_64-pc-windows-msvc"))
	path := stlerror.MustWith(stlos.NewFilePath(os.Args[1]).Abs())
	module := stlerror.MustWith(codegen_ir.CodegenIr(target, path))
	fmt.Println(module)
}

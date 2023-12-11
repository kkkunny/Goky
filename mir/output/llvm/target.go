package llvm

import (
	"sync"

	"github.com/kkkunny/go-llvm"
	stlerror "github.com/kkkunny/stl/error"

	"github.com/kkkunny/Sim/mir"
)

func getTarget(target mir.Target)*llvm.Target{
	switch target.Arch(){
	case mir.ArchX8664:
		sync.OnceFunc(func() {
			stlerror.Must(llvm.InitializeTargetInfo(llvm.X86))
			stlerror.Must(llvm.InitializeTarget(llvm.X86))
			stlerror.Must(llvm.InitializeTargetMC(llvm.X86))
		})
	default:
		panic("unreachable")
	}

	switch target.Name() {
	case mir.PackTargetName(mir.ArchX8664, mir.OSWindows):
		return stlerror.MustWith(llvm.NewTargetFromTriple("x86_64-w64-windows-gnu", "generic", ""))
	case mir.PackTargetName(mir.ArchX8664, mir.OSLinux):
		return stlerror.MustWith(llvm.NewTargetFromTriple("x86_64-unknown-linux-gnu", "generic", ""))
	default:
		panic("unreachable")
	}
}

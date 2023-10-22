package codegen

import (
	"github.com/kkkunny/go-llvm"

	"github.com/kkkunny/Sim/analyse"
	"github.com/kkkunny/Sim/mean"
)

// CodeGenerator 代码生成器
type CodeGenerator struct {
	analyser *analyse.Analyser // 语法分析器

	target  *llvm.Target
	ctx     llvm.Context
	module  llvm.Module
	builder llvm.Builder
}

func New(target *llvm.Target, analyser *analyse.Analyser) *CodeGenerator {
	ctx := llvm.NewContext()
	module := ctx.NewModule("main")
	module.SetTarget(target)
	return &CodeGenerator{
		analyser: analyser,
		target:   target,
		ctx:      ctx,
		module:   module,
		builder:  ctx.NewBuilder(),
	}
}

// Codegen 代码生成
func (self *CodeGenerator) Codegen() llvm.Module {
	nodes := self.analyser.Analyse()
	iter := nodes.Iterator()
	// 声明
	iter.Foreach(func(v mean.Global) bool {
		self.codegenGlobalDecl(v)
		return true
	})
	iter.Reset()
	// 定义
	nodes.Iterator().Foreach(func(v mean.Global) bool {
		self.codegenGlobalDef(v)
		return true
	})
	return self.module
}

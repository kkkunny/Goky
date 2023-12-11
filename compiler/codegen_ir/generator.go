package codegen_ir

import (
	"github.com/kkkunny/stl/container/hashmap"
	"github.com/kkkunny/stl/container/iterator"
	"github.com/kkkunny/stl/container/linkedlist"

	"github.com/kkkunny/Sim/mean"
	"github.com/kkkunny/Sim/mir"
)

// CodeGenerator 代码生成器
type CodeGenerator struct {
	means linkedlist.LinkedList[mean.Global]

	target  mir.Target
	ctx     *mir.Context
	module  *mir.Module
	builder *mir.Builder

	values  hashmap.HashMap[mean.Expr, mir.Value]
	loops   hashmap.HashMap[mean.Loop, loop]
	strings hashmap.HashMap[string, *mir.Constant]
	structs hashmap.HashMap[*mean.StructDef, mir.StructType]

	genericParams hashmap.HashMap[*mean.GenericParam, mean.Type]
}

func New(target mir.Target, means linkedlist.LinkedList[mean.Global]) *CodeGenerator {
	ctx := mir.NewContext(target)
	return &CodeGenerator{
		means: means,
		target:   target,
		ctx:      ctx,
		module:   ctx.NewModule(),
		builder:  ctx.NewBuilder(),
	}
}

// Codegen 代码生成
func (self *CodeGenerator) Codegen() *mir.Module {
	// 类型声明
	iterator.Foreach(self.means, func(v mean.Global) bool {
		st, ok := v.(*mean.StructDef)
		if ok {
			self.declStructDef(st)
		}
		return true
	})
	// 值声明
	iterator.Foreach(self.means, func(v mean.Global) bool {
		self.codegenGlobalDecl(v)
		return true
	})
	// 值定义
	iterator.Foreach(self.means, func(v mean.Global) bool {
		self.codegenGlobalDef(v)
		return true
	})
	// 初始化函数
	// FIXME: jit无法运行llvm.global_ctors
	self.builder.MoveTo(self.getInitFunction().Blocks().Front().Value)
	self.builder.BuildReturn()
	// 主函数
	var hasMain bool
	iterator.Foreach(self.means, func(v mean.Global) bool {
		if funcNode, ok := v.(*mean.FuncDef); ok && funcNode.Name == "main" {
			hasMain = true
			f := self.values.Get(funcNode).(*mir.Function)
			self.builder.MoveTo(self.getMainFunction().Blocks().Front().Value)
			self.builder.BuildReturn(self.builder.BuildCall(f))
			return false
		}
		return true
	})
	if !hasMain {
		self.builder.MoveTo(self.getMainFunction().Blocks().Front().Value)
		self.builder.BuildReturn(mir.NewUint(self.ctx.U8(), 0))
	}
	return self.module
}

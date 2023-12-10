package llvm

import (
	"github.com/kkkunny/go-llvm"
	"github.com/kkkunny/stl/container/hashmap"

	"github.com/kkkunny/Sim/mir"
)

type LLVMOutputer struct {
	target *llvm.Target
	ctx llvm.Context
	module llvm.Module
	builder llvm.Builder

	types hashmap.HashMap[mir.Type, llvm.Type]
	values hashmap.HashMap[mir.Value, llvm.Value]
	blocks hashmap.HashMap[*mir.Block, llvm.Block]
}

func NewLLVMOutputer()*LLVMOutputer{
	return &LLVMOutputer{}
}

func (self *LLVMOutputer) init(module *mir.Module){
	self.target = getTarget(module.Context().Target())
	self.ctx = llvm.NewContext()
	self.module = self.ctx.NewModule("")
	self.module.SetTarget(self.target)
	self.builder = self.ctx.NewBuilder()

	self.types = hashmap.NewHashMap[mir.Type, llvm.Type]()
	self.values = hashmap.NewHashMap[mir.Value, llvm.Value]()
	self.blocks = hashmap.NewHashMap[*mir.Block, llvm.Block]()
}

func (self *LLVMOutputer) Codegen(module *mir.Module){
	self.init(module)

	for iter:=module.Globals().Iterator(); iter.Next(); {
		self.codegenDeclType(iter.Value())
	}
	for iter:=module.Globals().Iterator(); iter.Next(); {
		self.codegenDefType(iter.Value())
	}

	for iter:=module.Globals().Iterator(); iter.Next(); {
		self.codegenDeclValue(iter.Value())
	}
	for iter:=module.Globals().Iterator(); iter.Next(); {
		self.codegenDefValue(iter.Value())
	}
}

func (self *LLVMOutputer) Module()llvm.Module{
	return self.module
}

func (self *LLVMOutputer) Target()*llvm.Target{
	return self.target
}

package ast

import (
	"github.com/kkkunny/Sim/reader"
	"github.com/kkkunny/Sim/token"
	"github.com/kkkunny/Sim/util"
)

// Global 全局ast
type Global interface {
	Ast
	global()
}

// FuncDef 函数定义
type FuncDef struct {
	Begin reader.Position
	Name  token.Token
	Ret   util.Option[Type]
	Body  *Block
}

func (self *FuncDef) Position() reader.Position {
	return reader.MixPosition(self.Begin, self.Body.Position())
}

func (self *FuncDef) global() {}

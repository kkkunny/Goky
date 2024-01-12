package ast

import (
	"github.com/kkkunny/stl/container/linkedlist"
	"github.com/kkkunny/stl/container/pair"

	"github.com/kkkunny/Sim/reader"
	"github.com/kkkunny/Sim/token"
	"github.com/kkkunny/Sim/util"
)

// Stmt 语句ast
type Stmt interface {
	Ast
	stmt()
}

// Block 代码块
type Block struct {
	Begin reader.Position
	Stmts linkedlist.LinkedList[Stmt]
	End   reader.Position
}

func (self *Block) Position() reader.Position {
	return reader.MixPosition(self.Begin, self.End)
}

func (*Block) stmt() {}

// Return 函数返回
type Return struct {
	Begin reader.Position
	Value util.Option[Expr]
}

func (self *Return) Position() reader.Position {
	v, ok := self.Value.Value()
	if !ok {
		return self.Begin
	}
	return reader.MixPosition(self.Begin, v.Position())
}

func (*Return) stmt() {}

// IfElse if else
type IfElse struct {
	Begin reader.Position
	Cond  util.Option[Expr]
	Body  *Block
	Next  util.Option[*IfElse]
}

func (self *IfElse) IsIf() bool {
	return self.Cond.IsSome()
}

func (self *IfElse) IsElse() bool {
	return self.Cond.IsNone()
}

func (self *IfElse) Position() reader.Position {
	if next, ok := self.Next.Value(); !ok {
		return reader.MixPosition(self.Begin, self.Body.Position())
	} else {
		return reader.MixPosition(self.Begin, next.Position())
	}
}

func (*IfElse) stmt() {}

// While 循环
type While struct {
	Begin reader.Position
	Cond Expr
	Body  *Block
}

func (self *While) Position() reader.Position {
	return reader.MixPosition(self.Begin, self.Body.Position())
}

func (*While) stmt() {}

// Break 跳出循环
type Break struct {
	Token token.Token
}

func (self *Break) Position() reader.Position {
	return self.Token.Position
}

func (*Break) stmt() {}

// Continue 下一次循环
type Continue struct {
	Token token.Token
}

func (self *Continue) Position() reader.Position {
	return self.Token.Position
}

func (*Continue) stmt() {}

// For 遍历
type For struct {
	Begin     reader.Position
	CursorMut bool
	Cursor    token.Token
	Iterator  Expr
	Body      *Block
}

func (self *For) Position() reader.Position {
	return reader.MixPosition(self.Begin, self.Body.Position())
}

func (*For) stmt() {}

// Match 匹配
type Match struct {
	Begin reader.Position
	Value Expr
	Cases []pair.Pair[Type, *Block]
	Other util.Option[*Block]
	End   reader.Position
}

func (self *Match) Position() reader.Position {
	return reader.MixPosition(self.Begin, self.End)
}

func (*Match) stmt() {}

package analyse

import (
	"github.com/kkkunny/stl/container/linkedlist"

	"github.com/kkkunny/Sim/ast"
	errors "github.com/kkkunny/Sim/error"
	. "github.com/kkkunny/Sim/mean"
	"github.com/kkkunny/Sim/util"
)

func (self *Analyser) analyseStmt(node ast.Stmt) (Stmt, BlockEof) {
	switch stmtNode := node.(type) {
	case *ast.Return:
		ret := self.analyseReturn(stmtNode)
		return ret, BlockEofReturn
	case *ast.Variable:
		return self.analyseLocalVariable(stmtNode), BlockEofNone
	case *ast.Block:
		return self.analyseBlock(stmtNode, nil)
	case *ast.IfElse:
		return self.analyseIfElse(stmtNode)
	case ast.Expr:
		return self.analyseExpr(nil, stmtNode), BlockEofNone
	case *ast.Loop:
		return self.analyseEndlessLoop(stmtNode)
	case *ast.Break:
		return self.analyseBreak(stmtNode), BlockEofBreakLoop
	case *ast.Continue:
		return self.analyseContinue(stmtNode), BlockEofNextLoop
	case *ast.For:
		return self.analyseFor(stmtNode)
	default:
		panic("unreachable")
	}
}

func (self *Analyser) analyseBlock(node *ast.Block, afterBlockCreate func(scope _LocalScope)) (*Block, BlockEof) {
	blockScope := _NewBlockScope(self.localScope)
	if afterBlockCreate != nil {
		afterBlockCreate(blockScope)
	}

	self.localScope = blockScope
	defer func() {
		self.localScope = self.localScope.GetParent().(_LocalScope)
	}()

	var jump BlockEof
	stmts := linkedlist.NewLinkedList[Stmt]()
	for iter := node.Stmts.Iterator(); iter.Next(); {
		stmt, stmtJump := self.analyseStmt(iter.Value())
		if b, ok := stmt.(*Block); ok {
			for iter := b.Stmts.Iterator(); iter.Next(); {
				stmts.PushBack(iter.Value())
			}
		} else {
			stmts.PushBack(stmt)
		}
		jump = max(jump, stmtJump)
	}
	return &Block{Stmts: stmts}, jump
}

func (self *Analyser) analyseReturn(node *ast.Return) *Return {
	expectRetType := self.localScope.GetRetType()
	if v, ok := node.Value.Value(); ok {
		value := self.expectExpr(expectRetType, v)
		return &Return{Value: util.Some[Expr](value)}
	} else {
		if !expectRetType.Equal(Empty) {
			errors.ThrowTypeMismatchError(node.Position(), expectRetType, Empty)
		}
		return &Return{Value: util.None[Expr]()}
	}
}

func (self *Analyser) analyseLocalVariable(node *ast.Variable) *Variable {
	v := &Variable{
		Mut:  node.Mutable,
		Name: node.Name.Source(),
	}
	if !self.localScope.SetValue(v.Name, v) {
		errors.ThrowIdentifierDuplicationError(node.Position(), node.Name)
	}

	v.Type = self.analyseType(node.Type)
	v.Value = self.expectExpr(v.Type, node.Value)
	return v
}

func (self *Analyser) analyseIfElse(node *ast.IfElse) (*IfElse, BlockEof) {
	if condNode, ok := node.Cond.Value(); ok {
		cond := self.expectExpr(Bool, condNode)
		body, jump := self.analyseBlock(node.Body, nil)

		var next util.Option[*IfElse]
		if nextNode, ok := node.Next.Value(); ok {
			nextIf, nextJump := self.analyseIfElse(nextNode)
			next = util.Some(nextIf)
			jump = max(jump, nextJump)
		} else {
			jump = BlockEofNone
		}

		return &IfElse{
			Cond: util.Some(cond),
			Body: body,
			Next: next,
		}, jump
	} else {
		body, jump := self.analyseBlock(node.Body, nil)
		return &IfElse{Body: body}, jump
	}
}

func (self *Analyser) analyseEndlessLoop(node *ast.Loop) (*EndlessLoop, BlockEof) {
	loop := &EndlessLoop{}
	body, eof := self.analyseBlock(node.Body, func(scope _LocalScope) {
		scope.SetLoop(loop)
	})
	loop.Body = body

	if eof == BlockEofNextLoop || eof == BlockEofBreakLoop {
		eof = BlockEofNone
	}
	return loop, eof
}

func (self *Analyser) analyseBreak(node *ast.Break) *Break {
	loop := self.localScope.GetLoop()
	if loop == nil {
		errors.ThrowLoopControlError(node.Position())
	}
	return &Break{Loop: loop}
}

func (self *Analyser) analyseContinue(node *ast.Continue) *Continue {
	loop := self.localScope.GetLoop()
	if loop == nil {
		errors.ThrowLoopControlError(node.Position())
	}
	return &Continue{Loop: loop}
}

func (self *Analyser) analyseFor(node *ast.For) (*For, BlockEof) {
	iterator := self.analyseExpr(nil, node.Iterator)
	iterType := iterator.GetType()
	if !TypeIs[*ArrayType](iterType) {
		errors.ThrowNotArrayError(node.Iterator.Position(), iterType)
	}

	et := iterType.(*ArrayType).Elem
	loop := &For{
		Iterator: iterator,
		Cursor: &Variable{
			Mut:   node.CursorMut,
			Type:  et,
			Name:  node.Cursor.Source(),
			Value: &Zero{Type: et},
		},
	}
	body, eof := self.analyseBlock(node.Body, func(scope _LocalScope) {
		if !scope.SetValue(loop.Cursor.Name, loop.Cursor) {
			errors.ThrowIdentifierDuplicationError(node.Cursor.Position, node.Cursor)
		}
		scope.SetLoop(loop)
	})
	loop.Body = body

	if eof == BlockEofNextLoop || eof == BlockEofBreakLoop {
		eof = BlockEofNone
	}
	return loop, eof
}

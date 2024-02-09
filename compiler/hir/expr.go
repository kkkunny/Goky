package hir

import (
	"math/big"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/either"
	stlslices "github.com/kkkunny/stl/slices"
)

// Expr 表达式
type Expr interface {
	Stmt
	GetType() Type
	Mutable() bool
}

// Ident 标识符
type Ident interface {
	Expr
	ident()
}

// Integer 整数
type Integer struct {
	Type  Type
	Value *big.Int
}

func (self *Integer) stmt() {}

func (self *Integer) GetType() Type {
	return self.Type
}

func (self *Integer) Mutable() bool {
	return false
}

// Float 浮点数
type Float struct {
	Type  Type
	Value *big.Float
}

func (self *Float) stmt() {}

func (self *Float) GetType() Type {
	return self.Type
}

func (self *Float) Mutable() bool {
	return false
}

// Binary 二元运算
type Binary interface {
	Expr
	GetLeft() Expr
	GetRight() Expr
}

// Assign 赋值
type Assign struct {
	Left, Right Expr
}

func (self *Assign) stmt() {}

func (self *Assign) GetType() Type {
	return NoThing
}

func (self *Assign) Mutable() bool {
	return false
}

func (self *Assign) GetLeft() Expr {
	return self.Left
}

func (self *Assign) GetRight() Expr {
	return self.Right
}

// IntAndInt 整数且整数
type IntAndInt struct {
	Left, Right Expr
}

func (self *IntAndInt) stmt() {}

func (self *IntAndInt) GetType() Type {
	return self.Left.GetType()
}

func (self *IntAndInt) Mutable() bool {
	return false
}

func (self *IntAndInt) GetLeft() Expr {
	return self.Left
}

func (self *IntAndInt) GetRight() Expr {
	return self.Right
}

// IntOrInt 整数或整数
type IntOrInt struct {
	Left, Right Expr
}

func (self *IntOrInt) stmt() {}

func (self *IntOrInt) GetType() Type {
	return self.Left.GetType()
}

func (self *IntOrInt) Mutable() bool {
	return false
}

func (self *IntOrInt) GetLeft() Expr {
	return self.Left
}

func (self *IntOrInt) GetRight() Expr {
	return self.Right
}

// IntXorInt 整数异或整数
type IntXorInt struct {
	Left, Right Expr
}

func (self *IntXorInt) stmt() {}

func (self *IntXorInt) GetType() Type {
	return self.Left.GetType()
}

func (self *IntXorInt) Mutable() bool {
	return false
}

func (self *IntXorInt) GetLeft() Expr {
	return self.Left
}

func (self *IntXorInt) GetRight() Expr {
	return self.Right
}

// IntShlInt 整数左移整数
type IntShlInt struct {
	Left, Right Expr
}

func (self *IntShlInt) stmt() {}

func (self *IntShlInt) GetType() Type {
	return self.Left.GetType()
}

func (self *IntShlInt) Mutable() bool {
	return false
}

func (self *IntShlInt) GetLeft() Expr {
	return self.Left
}

func (self *IntShlInt) GetRight() Expr {
	return self.Right
}

// IntShrInt 整数右移整数
type IntShrInt struct {
	Left, Right Expr
}

func (self *IntShrInt) stmt() {}

func (self *IntShrInt) GetType() Type {
	return self.Left.GetType()
}

func (self *IntShrInt) Mutable() bool {
	return false
}

func (self *IntShrInt) GetLeft() Expr {
	return self.Left
}

func (self *IntShrInt) GetRight() Expr {
	return self.Right
}

// NumAddNum 数字加数字
type NumAddNum struct {
	Left, Right Expr
}

func (self *NumAddNum) stmt() {}

func (self *NumAddNum) GetType() Type {
	return self.Left.GetType()
}

func (self *NumAddNum) Mutable() bool {
	return false
}

func (self *NumAddNum) GetLeft() Expr {
	return self.Left
}

func (self *NumAddNum) GetRight() Expr {
	return self.Right
}

// NumSubNum 数字减数字
type NumSubNum struct {
	Left, Right Expr
}

func (self *NumSubNum) stmt() {}

func (self *NumSubNum) GetType() Type {
	return self.Left.GetType()
}

func (self *NumSubNum) Mutable() bool {
	return false
}

func (self *NumSubNum) GetLeft() Expr {
	return self.Left
}

func (self *NumSubNum) GetRight() Expr {
	return self.Right
}

// NumMulNum 数字乘数字
type NumMulNum struct {
	Left, Right Expr
}

func (self *NumMulNum) stmt() {}

func (self *NumMulNum) GetType() Type {
	return self.Left.GetType()
}

func (self *NumMulNum) Mutable() bool {
	return false
}

func (self *NumMulNum) GetLeft() Expr {
	return self.Left
}

func (self *NumMulNum) GetRight() Expr {
	return self.Right
}

// NumDivNum 数字除数字
type NumDivNum struct {
	Left, Right Expr
}

func (self *NumDivNum) stmt() {}

func (self *NumDivNum) GetType() Type {
	return self.Left.GetType()
}

func (self *NumDivNum) Mutable() bool {
	return false
}

func (self *NumDivNum) GetLeft() Expr {
	return self.Left
}

func (self *NumDivNum) GetRight() Expr {
	return self.Right
}

// NumRemNum 数字取余数字
type NumRemNum struct {
	Left, Right Expr
}

func (self *NumRemNum) stmt() {}

func (self *NumRemNum) GetType() Type {
	return self.Left.GetType()
}

func (self *NumRemNum) Mutable() bool {
	return false
}

func (self *NumRemNum) GetLeft() Expr {
	return self.Left
}

func (self *NumRemNum) GetRight() Expr {
	return self.Right
}

// NumLtNum 数字小于数字
type NumLtNum struct {
	BoolType    Type
	Left, Right Expr
}

func (self *NumLtNum) stmt() {}

func (self *NumLtNum) GetType() Type {
	return self.BoolType
}

func (self *NumLtNum) Mutable() bool {
	return false
}

func (self *NumLtNum) GetLeft() Expr {
	return self.Left
}

func (self *NumLtNum) GetRight() Expr {
	return self.Right
}

// NumGtNum 数字大于数字
type NumGtNum struct {
	BoolType    Type
	Left, Right Expr
}

func (self *NumGtNum) stmt() {}

func (self *NumGtNum) GetType() Type {
	return self.BoolType
}

func (self *NumGtNum) Mutable() bool {
	return false
}

func (self *NumGtNum) GetLeft() Expr {
	return self.Left
}

func (self *NumGtNum) GetRight() Expr {
	return self.Right
}

// NumLeNum 数字小于等于数字
type NumLeNum struct {
	BoolType    Type
	Left, Right Expr
}

func (self *NumLeNum) stmt() {}

func (self *NumLeNum) GetType() Type {
	return self.BoolType
}

func (self *NumLeNum) Mutable() bool {
	return false
}

func (self *NumLeNum) GetLeft() Expr {
	return self.Left
}

func (self *NumLeNum) GetRight() Expr {
	return self.Right
}

// NumGeNum 数字大于等于数字
type NumGeNum struct {
	BoolType    Type
	Left, Right Expr
}

func (self *NumGeNum) stmt() {}

func (self *NumGeNum) GetType() Type {
	return self.BoolType
}

func (self *NumGeNum) Mutable() bool {
	return false
}

func (self *NumGeNum) GetLeft() Expr {
	return self.Left
}

func (self *NumGeNum) GetRight() Expr {
	return self.Right
}

// Equal 比较相等
type Equal struct {
	BoolType    Type
	Left, Right Expr
}

func (self *Equal) stmt() {}

func (self *Equal) GetType() Type {
	return self.BoolType
}

func (self *Equal) Mutable() bool {
	return false
}

func (self *Equal) GetLeft() Expr {
	return self.Left
}

func (self *Equal) GetRight() Expr {
	return self.Right
}

// NotEqual 比较不相等
type NotEqual struct {
	BoolType    Type
	Left, Right Expr
}

func (self *NotEqual) stmt() {}

func (self *NotEqual) GetType() Type {
	return self.BoolType
}

func (self *NotEqual) Mutable() bool {
	return false
}

func (self *NotEqual) GetLeft() Expr {
	return self.Left
}

func (self *NotEqual) GetRight() Expr {
	return self.Right
}

// BoolAndBool 布尔并且布尔
type BoolAndBool struct {
	Left, Right Expr
}

func (self *BoolAndBool) stmt() {}

func (self *BoolAndBool) GetType() Type {
	return self.Left.GetType()
}

func (self *BoolAndBool) Mutable() bool {
	return false
}

func (self *BoolAndBool) GetLeft() Expr {
	return self.Left
}

func (self *BoolAndBool) GetRight() Expr {
	return self.Right
}

// BoolOrBool 布尔或者布尔
type BoolOrBool struct {
	Left, Right Expr
}

func (self *BoolOrBool) stmt() {}

func (self *BoolOrBool) GetType() Type {
	return self.Left.GetType()
}

func (self *BoolOrBool) Mutable() bool {
	return false
}

func (self *BoolOrBool) GetLeft() Expr {
	return self.Left
}

func (self *BoolOrBool) GetRight() Expr {
	return self.Right
}

// Unary 一元运算
type Unary interface {
	Expr
	GetValue() Expr
}

// NumNegate 数字取相反数
type NumNegate struct {
	Value Expr
}

func (self *NumNegate) stmt() {}

func (self *NumNegate) GetType() Type {
	return self.Value.GetType()
}

func (self *NumNegate) Mutable() bool {
	return false
}

func (self *NumNegate) GetValue() Expr {
	return self.Value
}

// IntBitNegate 整数按位取反
type IntBitNegate struct {
	Value Expr
}

func (self *IntBitNegate) stmt() {}

func (self *IntBitNegate) GetType() Type {
	return self.Value.GetType()
}

func (self *IntBitNegate) Mutable() bool {
	return false
}

func (self *IntBitNegate) GetValue() Expr {
	return self.Value
}

// BoolNegate 布尔取反
type BoolNegate struct {
	Value Expr
}

func (self *BoolNegate) stmt() {}

func (self *BoolNegate) GetType() Type {
	return self.Value.GetType()
}

func (self *BoolNegate) Mutable() bool {
	return false
}

func (self *BoolNegate) GetValue() Expr {
	return self.Value
}

// Boolean 布尔值
type Boolean struct {
	BoolType Type
	Value    bool
}

func (self *Boolean) stmt() {}

func (self *Boolean) GetType() Type {
	return self.BoolType
}

func (self *Boolean) Mutable() bool {
	return false
}

// Call 调用
type Call struct {
	Func Expr
	Args []Expr
}

func (self *Call) stmt() {}

func (self *Call) GetType() Type {
	return AsType[*FuncType](self.Func.GetType()).Ret
}

func (self *Call) Mutable() bool {
	return false
}

// TypeCovert 类型转换
type TypeCovert interface {
	Expr
	GetFrom() Expr
}

// Num2Num 数字类型转数字类型
type Num2Num struct {
	From Expr
	To   Type
}

func (self *Num2Num) stmt() {}

func (self *Num2Num) GetType() Type {
	return self.To
}

func (*Num2Num) Mutable() bool {
	return false
}

func (self *Num2Num) GetFrom() Expr {
	return self.From
}

// DoNothingCovert 什么事都不做的转换
type DoNothingCovert struct {
	From Expr
	To   Type
}

func (self *DoNothingCovert) stmt() {}

func (self *DoNothingCovert) GetType() Type {
	return self.To
}

func (*DoNothingCovert) Mutable() bool {
	return false
}

func (self *DoNothingCovert) GetFrom() Expr {
	return self.From
}

// NoReturn2Any 无返回转任意类型
type NoReturn2Any struct {
	From Expr
	To   Type
}

func (self *NoReturn2Any) stmt() {}

func (self *NoReturn2Any) GetType() Type {
	return self.To
}

func (*NoReturn2Any) Mutable() bool {
	return false
}

func (self *NoReturn2Any) GetFrom() Expr {
	return self.From
}

// Array 数组
type Array struct {
	Type  Type
	Elems []Expr
}

func (self *Array) stmt() {}

func (self *Array) GetType() Type {
	return self.Type
}

func (self *Array) Mutable() bool {
	return false
}

// Index 索引
type Index struct {
	From  Expr
	Index Expr
}

func (self *Index) stmt() {}

func (self *Index) GetType() Type {
	return AsType[*ArrayType](self.From.GetType()).Elem
}

func (self *Index) Mutable() bool {
	return self.From.Mutable()
}

// Tuple 元组
type Tuple struct {
	Elems []Expr
}

func (self *Tuple) stmt() {}

func (self *Tuple) GetType() Type {
	ets := stlslices.Map(self.Elems, func(_ int, e Expr) Type {
		return e.GetType()
	})
	return NewTupleType(ets...)
}

func (self *Tuple) Mutable() bool {
	for _, elem := range self.Elems {
		if !elem.Mutable() {
			return false
		}
	}
	return true
}

// Extract 提取
type Extract struct {
	From  Expr
	Index uint
}

func (self *Extract) stmt() {}

func (self *Extract) GetType() Type {
	return AsType[*TupleType](self.From.GetType()).Elems[self.Index]
}

func (self *Extract) Mutable() bool {
	return self.From.Mutable()
}

// Struct 结构体
type Struct struct {
	Type   Type
	Fields []Expr
}

func (*Struct) stmt() {}

func (self *Struct) GetType() Type {
	return self.Type
}

func (self *Struct) Mutable() bool {
	return false
}

// Default 默认值
type Default struct {
	Type Type
}

func (*Default) stmt() {}

func (self *Default) GetType() Type {
	return self.Type
}

func (self *Default) Mutable() bool {
	return false
}

// GetField 取字段
type GetField struct {
	Internal bool // 是否在包内部，不受字段可变性影响
	From     Expr
	Index    uint
}

func (self *GetField) stmt() {}

func (self *GetField) GetType() Type {
	return AsType[*StructType](self.From.GetType()).Fields.Values().Get(self.Index).Type
}

func (self *GetField) Mutable() bool {
	if self.Internal {
		return true
	}
	return self.From.Mutable() && AsType[*StructType](self.From.GetType()).Fields.Values().Get(self.Index).Mutable
}

// String 字符串
type String struct {
	Type  Type
	Value string
}

func (self *String) stmt() {}

func (self *String) GetType() Type {
	return self.Type
}

func (self *String) Mutable() bool {
	return false
}

// Union 联合
type Union struct {
	Type  Type
	Value Expr
}

func (self *Union) stmt() {}

func (self *Union) GetType() Type {
	return self.Type
}

func (self *Union) Mutable() bool {
	return false
}

// TypeJudgment 类型判断
type TypeJudgment struct {
	BoolType Type
	Value    Expr
	Type     Type
}

func (self *TypeJudgment) stmt() {}

func (self *TypeJudgment) GetType() Type {
	return self.BoolType
}

func (self *TypeJudgment) Mutable() bool {
	return false
}

// UnUnion 解联合
type UnUnion struct {
	Type  Type
	Value Expr
}

func (self *UnUnion) stmt() {}

func (self *UnUnion) GetType() Type {
	return self.Type
}

func (self *UnUnion) Mutable() bool {
	return false
}

// GetRef 取引用
type GetRef struct {
	Mut   bool
	Value Expr
}

func (self *GetRef) stmt() {}

func (self *GetRef) GetType() Type {
	return NewRefType(self.Mut, self.Value.GetType())
}

func (self *GetRef) Mutable() bool {
	return self.Mut
}

func (self *GetRef) GetValue() Expr {
	return self.Value
}

// DeRef 解引用
type DeRef struct {
	Value Expr
}

func (self *DeRef) stmt() {}

func (self *DeRef) GetType() Type {
	return AsType[*RefType](self.Value.GetType()).Elem
}

func (self *DeRef) Mutable() bool {
	return AsType[*RefType](self.Value.GetType()).Mut && self.Value.Mutable()
}

func (self *DeRef) GetValue() Expr {
	return self.Value
}

// Method 方法
type Method struct {
	Self   either.Either[Expr, *CustomType]
	Define *MethodDef
}

func (self *Method) stmt() {}

func (self *Method) GetScope() *TypeDef {
	if left, ok := self.Self.Left(); ok {
		return AsCustomType(left.GetType())
	} else {
		return stlbasic.IgnoreWith(self.Self.Right())
	}
}

func (self *Method) GetDefine() GlobalMethod {
	return self.Define
}

func (self *Method) GetType() Type {
	if self.Define.IsStatic() {
		return self.Define.GetFuncType()
	} else {
		return self.Define.GetMethodType()
	}
}

func (self *Method) Mutable() bool {
	return false
}

func (*Method) ident() {}

// ShrinkUnion 缩小联合
type ShrinkUnion struct {
	Type  Type
	Value Expr
}

func (self *ShrinkUnion) stmt() {}

func (self *ShrinkUnion) GetType() Type {
	return self.Type
}

func (self *ShrinkUnion) Mutable() bool {
	return false
}

func (self *ShrinkUnion) GetFrom() Expr {
	return self.Value
}

// ExpandUnion 扩大联合
type ExpandUnion struct {
	Type  Type
	Value Expr
}

func (self *ExpandUnion) stmt() {}

func (self *ExpandUnion) GetType() Type {
	return self.Type
}

func (self *ExpandUnion) Mutable() bool {
	return false
}

func (self *ExpandUnion) GetFrom() Expr {
	return self.Value
}

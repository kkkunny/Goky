package types

import (
	"fmt"
)

var (
	U8    UintType = &_UintType_{kind: IntTypeKindByte}
	U16   UintType = &_UintType_{kind: IntTypeKindShort}
	U32   UintType = &_UintType_{kind: IntTypeKindInt}
	U64   UintType = &_UintType_{kind: IntTypeKindLong}
	Usize UintType = &_UintType_{kind: IntTypeKindSize}
)

// UintType 无符号整型
type UintType interface {
	IntType
	Unsigned()
}

type _UintType_ struct {
	kind IntTypeKind
}

func (self *_UintType_) String() string {
	if self.kind == IntTypeKindSize {
		return "usize"
	} else {
		return fmt.Sprintf("u%d", self.kind*8)
	}
}

func (self *_UintType_) Equal(dst Type) bool {
	t, ok := dst.(UintType)
	return ok && self.kind == t.Kind()
}

func (self *_UintType_) Kind() IntTypeKind {
	return self.kind
}

func (self *_UintType_) Number()   {}
func (self *_UintType_) Unsigned() {}

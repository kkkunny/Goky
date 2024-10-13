package types

import (
	"fmt"
)

// RefType 引用类型
type RefType interface {
	Type
	Mutable() bool
	Pointer() Type
}

func NewRefType(mut bool, p Type) RefType {
	return &_RefType_{
		mut: mut,
		ptr: p,
	}
}

type _RefType_ struct {
	mut bool
	ptr Type
}

func (self *_RefType_) String() string {
	if self.mut {
		return fmt.Sprintf("&mut %s", self.ptr.String())
	} else {
		return fmt.Sprintf("&%s", self.ptr.String())
	}
}

func (self *_RefType_) Equal(dst Type) bool {
	t, ok := dst.(RefType)
	return ok && self.ptr.Equal(t.Pointer())
}

func (self *_RefType_) Pointer() Type {
	return self.ptr
}

func (self *_RefType_) Mutable() bool {
	return self.mut
}

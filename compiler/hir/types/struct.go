package types

import (
	"fmt"
	"strings"

	"github.com/kkkunny/stl/container/linkedhashmap"
	stlslices "github.com/kkkunny/stl/container/slices"
)

type Field struct {
	pub  bool
	mut  bool
	name string
	typ  Type
}

func NewField(pub bool, mut bool, name string, typ Type) *Field {
	return &Field{
		pub:  pub,
		mut:  mut,
		name: name,
		typ:  typ,
	}
}

func (self *Field) Name() string {
	return self.name
}

func (self *Field) Type() Type {
	return self.typ
}

func (self *Field) Public() bool {
	return self.pub
}

func (self *Field) Mutable() bool {
	return self.mut
}

// StructType 结构体类型
type StructType interface {
	BuildInType
	Fields() linkedhashmap.LinkedHashMap[string, *Field]
}

func NewStructType(fs ...*Field) StructType {
	lhm := linkedhashmap.StdWithCap[string, *Field](uint(len(fs)))
	for _, f := range fs {
		lhm.Set(f.Name(), f)
	}
	return &_StructType_{
		fields: lhm,
	}
}

type _StructType_ struct {
	fields linkedhashmap.LinkedHashMap[string, *Field]
}

func (self *_StructType_) String() string {
	fields := stlslices.Map(self.fields.Values(), func(i int, f *Field) string {
		return fmt.Sprintf("%s:%s", f.name, f.typ.String())
	})
	return fmt.Sprintf("struct{%s}", strings.Join(fields, ";"))
}

func (self *_StructType_) Equal(dst Type, selfs ...Type) bool {
	if dst.Equal(Self) && len(selfs) > 0 {
		dst = stlslices.Last(selfs)
	}

	t, ok := dst.(StructType)
	if !ok || self.fields.Length() != t.Fields().Length() {
		return false
	}
	fields2 := t.Fields().Values()
	return stlslices.All(self.fields.Values(), func(i int, f1 *Field) bool {
		f2 := fields2[i]
		return f1.pub == f2.pub &&
			f1.mut == f2.mut &&
			f1.name == f2.name &&
			f1.typ.Equal(f2.typ, selfs...)
	})
}

func (self *_StructType_) Fields() linkedhashmap.LinkedHashMap[string, *Field] {
	return self.fields
}

func (self *_StructType_) BuildIn() {}

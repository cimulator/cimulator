package t

import "errors"

type TypeKey int

type (
	Char       int8
	UChar      uint8
	Short      int16
	UShort     uint16
	Int        int32
	UInt       uint32
	Long       int64
	ULong      uint64
	LongLong   int64
	ULongLong  uint64
	Float      float32
	Double     float64
	LongDouble float64 // Go doesn't have 'long double' type.
)

var TypeKeyStrings = [...]string{"Char", "UChar", "Short", "UShort", "Int",
	"UInt", "Long", "ULong", "LongLong", "ULongLong", "Float", "Double",
	"LongDouble"}
var StringsTypeKey = map[string]TypeKey{}
var FlagOverflow = false
var OverflowError = errors.New("Value overflow")

func init() {
	for i, v := range TypeKeyStrings {
		StringsTypeKey[v] = TypeKey(i)
	}
}

type Calculable interface {
	Add(Calculable) Calculable
	Sub(Calculable) Calculable
	Cast(TypeKey) Calculable

	Priority() TypeKey
}

func MaxPriority(a, b TypeKey) TypeKey {
	if a > b {
		return a
	}
	return b
}

func (t TypeKey) String() string {
	return TypeKeyStrings[t]
}

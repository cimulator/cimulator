package t

import (
	"math"
	"strconv"
)

const {{ToUpper .Generic}} TypeKey = {{.ValGeneric}}

func (i {{.Generic}}) Key() TypeKey {
	return {{ToUpper .Generic}}
}

func (i {{.Generic}}) Priority() TypeKey {
	return i.Key()
}

func (i {{.Generic}}) TypeString() string {
	return i.Key().String()
}

func (i {{.Generic}}) Add(v Calculable) Calculable {
	if i.Priority() != v.Priority() {
		max := MaxPriority(i.Priority(), v.Priority())
		return i.Cast(max).Add(v.Cast(max))
	}
	result := i + v.({{.Generic}})
	{{ $backup := .Generic }}
	{{ if lt .PriorityGeneric .FloatPriority }}
	if FlagOverflow {
		if ((i > 0 && v.({{$backup}}) > {{.MaxGeneric}} - i) || (i < 0 && v.({{$backup}}) < {{.MinGeneric}} - i)) {
			panic(OverflowError)
		}
	}
	{{ end }}
	return result
}

func (i {{.Generic}}) Sub(v Calculable) Calculable {
	if i.Priority() != v.Priority() {
		max := MaxPriority(i.Priority(), v.Priority())
		return i.Cast(max).Sub(v.Cast(max))
	}
	result := i - v.({{.Generic}})
	{{ $backup := .Generic }}
	{{ if lt .PriorityGeneric .FloatPriority }}
	if FlagOverflow {
		if ((i > 0 && -v.({{$backup}}) > {{.MaxGeneric}} - i) || (i < 0 && -v.({{$backup}}) < {{.MinGeneric}} - i)) {
			panic(OverflowError)
		}
	}
	{{ end }}
  return result
}

func (i {{.Generic}}) Cast(t TypeKey) Calculable {
	switch t {
	{{ $backup := .Generic }}
	{{ range .Types }}
	case {{ ToUpper . }}:
		ret := {{ . }}(i)
		if FlagOverflow && i != {{$backup}}(ret) {
			panic(OverflowError)
		}
		return ret
	{{end}}

	default:
		panic("Type not understood.")
	}
}

// vim: ft=go

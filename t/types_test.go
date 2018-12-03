package t

import "testing"

func TestString(t *testing.T) {
	m := map[TypeKey]string{
		CHAR:       "Char",
		UCHAR:      "UChar",
		SHORT:      "Short",
		USHORT:     "UShort",
		INT:        "Int",
		UINT:       "UInt",
		LONG:       "Long",
		ULONG:      "ULong",
		LONGLONG:   "LongLong",
		ULONGLONG:  "ULongLong",
		FLOAT:      "Float",
		DOUBLE:     "Double",
		LONGDOUBLE: "LongDouble",
	}
	for key, value := range m {
		if key.String() != value {
			t.Errorf("String value didn't match. Want: %s, Got: %s",
				value, key.String())
		}
	}
}

func TestAdd_IntInt(t *testing.T) {
	a := Int(90)
	b := Int(50)

	if want, got := a.Add(b), Int(140); want != got {
		t.Errorf("Add output incorrect. Want: %T(%v), Got: %T(%v)",
			want, want, got, got)
	}
	if want, got := b.Add(a), Int(140); want != got {
		t.Errorf("Add output incorrect. Want: %T(%v), Got: %T(%v)",
			want, want, got, got)
	}
}

func TestAdd_IntShort(t *testing.T) {
	a := Int(90)
	b := Short(50)

	if want, got := a.Add(b), Int(140); want != got {
		t.Errorf("Add output incorrect. Want: %T(%v), Got: %T(%v)",
			want, want, got, got)
	}
	if want, got := b.Add(a), Int(140); want != got {
		t.Errorf("Add output incorrect. Want: %T(%v), Got: %T(%v)",
			want, want, got, got)
	}
}

func TestAdd_IntDouble(t *testing.T) {
	a := Int(90)
	b := Double(50)

	if want, got := a.Add(b), Double(140.0); want != got {
		t.Errorf("Add output incorrect. Want: %T(%v), Got: %T(%v)",
			want, want, got, got)
	}
	if want, got := b.Add(a), Double(140.0); want != got {
		t.Errorf("Add output incorrect. Want: %T(%v), Got: %T(%v)",
			want, want, got, got)
	}
}

func TestAdd_ShortShort(t *testing.T) {
	a := Short(90)
	b := Short(50)

	if want, got := a.Add(b), Short(140); want != got {
		t.Errorf("Add output incorrect. Want: %T(%v), Got: %T(%v)",
			want, want, got, got)
	}
	if want, got := b.Add(a), Short(140); want != got {
		t.Errorf("Add output incorrect. Want: %T(%v), Got: %T(%v)",
			want, want, got, got)
	}
}

func TestSub_IntShort(t *testing.T) {
	a := Int(90)
	b := Short(50)

	if want, got := a.Sub(b), Int(40); want != got {
		t.Errorf("Sub output incorrect. Want: %T(%v), Got: %T(%v)",
			want, want, got, got)
	}
	if want, got := b.Sub(a), Int(-40); want != got {
		t.Errorf("Sub output incorrect. Want: %T(%v), Got: %T(%v)",
			want, want, got, got)
	}
}

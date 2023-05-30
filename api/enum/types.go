package enum


type Enum[T1 comparable] struct {
	Values []T1
	Invalid T1
}

func NewEnum[T1 comparable](invalid T1,values ...T1) Enum[T1] {
	enum := Enum[T1]{
		Values: values,
		Invalid: invalid,
	}

	return enum
}

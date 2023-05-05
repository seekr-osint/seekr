package newtypes

type ValidStruct interface {
	~struct{}
}
type Custom[T ValidStruct] struct { // the parsable type. It makes T parsable
	data T
}

func (c Custom[T]) Get() T { // make this a pars function
	return c.data
}

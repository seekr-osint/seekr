package tcmultiarg

type Args []interface{}
type Test struct {
	Input    Args
	Expect   Args
	Function interface{}
}
type Tests []Test

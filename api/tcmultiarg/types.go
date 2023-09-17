package tcmultiarg

// Used to define function Arguments and Function returns.
type Args []interface{}

// Type used to run the generated tests from.
type Test struct {
	Input    Args
	Expect   Args
	Function interface{}
}
type Tests []Test

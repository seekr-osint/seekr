package tc

type TestCase[T1 comparable,T2 comparable] struct {
	Input  T1
	Expect T2
}
type TestCases[T1 comparable, T2 comparable] []TestCase[T1,T2]
type Test[T1 comparable, T2 comparable] struct {
	Cases TestCases[T1, T2]
	Func  func(T1) T2
}


func NewTest[T1 comparable, T2 comparable](testCaseMap map[T1]T2, function func(T1) T2) Test[T1, T2] {
	var testCases TestCases[T1, T2]
	for input, expect := range testCaseMap {
		testCase := TestCase[T1, T2]{Input: input, Expect: expect}
		testCases = append(testCases, testCase)
	}

	return Test[T1, T2]{Cases: testCases, Func: function}
}

func NewEnumIsValidTest[T1 comparable, T2 comparable](function func(T1) T2,invalidExpect T2, invalidEnum T1, validExpect T2, validEnum ...T1) Test[T1, T2] {
	var testCases TestCases[T1, T2]

	for _,input := range validEnum {
		testCase := TestCase[T1, T2]{Input: input, Expect: validExpect}
		testCases = append(testCases, testCase)
	}

	testCase := TestCase[T1, T2]{Input: invalidEnum, Expect: invalidExpect}
	testCases = append(testCases, testCase)

	return Test[T1, T2]{Cases: testCases, Func: function}
}

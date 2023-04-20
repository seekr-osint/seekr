package errortypes

type APIError struct {
	Message string
	Status  int
}

func (e APIError) Error() string {
	return e.Message
}

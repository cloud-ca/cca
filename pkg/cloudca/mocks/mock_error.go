package mocks

type MockError struct {
	Message string
}

func (mockError MockError) Error() string {
	return mockError.Message
}

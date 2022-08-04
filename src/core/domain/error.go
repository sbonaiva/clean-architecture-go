package domain

type CoreError struct {
	Status  int
	Message string
	Fields  map[string]string
}

func NewCoreError(status int, message string) error {
	return &CoreError{
		Status:  status,
		Message: message,
		Fields:  make(map[string]string),
	}
}

func NewCoreErrorWithFields(status int, message string, fields map[string]string) error {
	return &CoreError{
		Status:  status,
		Message: message,
		Fields:  fields,
	}
}

func (e *CoreError) Error() string {
	return e.Message
}

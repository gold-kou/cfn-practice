package helper

type BadRequestError struct {
	Message string
}

type NotFoundError struct {
	Message string
}

type InternalServerError struct {
	Message string
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{
		Message: message,
	}
}

func (e *BadRequestError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Message: message,
	}
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{
		Message: message,
	}
}

func (e *InternalServerError) Error() string {
	return e.Message
}

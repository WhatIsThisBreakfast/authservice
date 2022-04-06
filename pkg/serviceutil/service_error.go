package serviceutil

type ServiceError struct {
	Httpcode int
	Messgae  string
}

func NewError(httpcode int, message string) *ServiceError {
	return &ServiceError{
		Httpcode: httpcode,
		Messgae:  message,
	}
}

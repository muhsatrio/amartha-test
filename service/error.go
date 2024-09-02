package service

type Error interface {
	error
}

type ServiceErr struct {
	Message string
}

var _ Error = ServiceErr{}

func (s ServiceErr) Error() string {
	return s.Message
}

var (
	ErrInvalidInput = ServiceErr{
		Message: "Invalid input",
	}
	ErrRequiredFieldEmpty = ServiceErr{
		Message: "Required field is empty",
	}
	ErrUnauthorized = ServiceErr{
		Message: "Unauthorized",
	}
	ErrForbiddenAccess = ServiceErr{
		Message: "Forbidden access",
	}
	ErrDataNotFound = ServiceErr{
		Message: "Data not found",
	}
	ErrDuplicateDataAdd = ServiceErr{
		Message: "Can not add data with duplicate identifier",
	}
)

func InternalErrorCustom(msg string) ServiceErr {
	return ServiceErr{
		Message: msg,
	}
}

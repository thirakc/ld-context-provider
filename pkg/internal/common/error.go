package common

const (
	LDContext = 1000
)

type Error interface {
	error
	Code() int32
}

type svcError struct {
	error
	code int32
}

func (c *svcError) Code() int32 {
	return c.code
}

func NewError(code int32, err error) *svcError {
	return &svcError{err, code}
}

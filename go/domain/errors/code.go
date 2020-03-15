package errors

type DDDGuysErrorCode uint

const (
	NoError = iota
	UserIdError
	UserNameError
	Unknown
)

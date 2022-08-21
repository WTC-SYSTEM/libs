package apperror

var (
	ErrNotFound               = NewAppError("not found", WTC000003, "")
	ErrEmailIsRegistered      = NewAppError("email is already registered", WTC000004, "")
	ErrInvalidEmailOrPassword = NewAppError("invalid email or password", WTC000005, "")
	ErrIncorrectOldPassword   = NewAppError("old password is wrong", WTC000006, "")
)

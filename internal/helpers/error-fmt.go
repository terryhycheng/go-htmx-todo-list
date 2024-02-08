package helpers

import "errors"

const (
	ErrUnknownError = "ERR_UNKNOWN_ERROR"
	ErrTitleRequired = "ERR_TITLE_REQUIRED"
	ErrDescriptionRequired = "ERR_DESCRIPTION_REQUIRED"
	ErrStatusRequired = "ERR_STATUS_REQUIRED"
	ErrTodoNotFound = "ERR_TODO_NOT_FOUND"
)

func Error(errorType string) error {
	switch errorType {
	case ErrTitleRequired:
		return errors.New(ErrTitleRequired)
	default:
		return errors.New(ErrUnknownError)
	}
}